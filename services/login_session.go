package services

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/security"
)

type LoggedSession struct {
	Username   string    `json:"username"`
	APISecret  string    `json:"api_secret"`
	Authorized bool      `json:"authorized"`
	AtDatetime time.Time `json:"at_datetime"`
	Reason     string    `json:"reason"`
}

type AuthSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var LoggedSessionVar = make(map[string]*LoggedSession)

func FindLoggedSession(username string) (session *LoggedSession, err error) {
	if v, found := LoggedSessionVar[username]; found {
		session = v
		return
	}
	err = errors.New("not found session")
	return
}

func CheckTimeout(session *LoggedSession) {
	nowTime := time.Now()
	diffTime := nowTime.Sub(session.AtDatetime)
	if int(diffTime.Hours()) > 0 && int(diffTime.Minutes()) > 45 {
		session.Authorized = false
		session.Reason = "time-out"
	}

}

func Logout(session *LoggedSession) {
	delete(LoggedSessionVar, session.Username)
}

func CheckSessionWithUsernameAndAPISecret(obj AuthSchema) (auth *LoggedSession, e error) {
	session, err := FindLoggedSession(obj.Username)
	if err != nil {
		e = errors.New("not found this session ")
		return
	} else {
		if session.APISecret == obj.Password {
			auth = session
		} else {
			e = errors.New("api secret not match")
		}
	}
	return
}

func GetAdminAuthentication(obj AuthSchema) (auth *LoggedSession, err error) {

	res := GetAuthentication(obj.Username, obj.Password, true)
	if res.Status == "success" {

		if res.Role != "super_admin" {
			err = errors.New("cannot access to user management from this account")
			return
		}

		session, err := FindLoggedSession(obj.Username)
		if err == nil && session.Authorized {
			auth = session
		} else {
			auth = &LoggedSession{
				Username:   obj.Username,
				APISecret:  security.GeneratePasswordString(128),
				Authorized: true,
				AtDatetime: time.Now(),
			}
			LoggedSessionVar[obj.Username] = auth
		}
	} else {
		err = errors.New(res.Message)
		return
	}
	return
}

func APIAuthentication(c *gin.Context) {
	username, apiSecret, hasAuth := c.Request.BasicAuth()
	authorized := false

	if hasAuth {

		session, err := FindLoggedSession(username)
		CheckTimeout(session)
		if err == nil {
			if session.APISecret == apiSecret {
				authorized = true
			}
		}

	}

	if authorized {
		c.Next()
	} else {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

}
