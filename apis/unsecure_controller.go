package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/mrzack99s/mrz-identity-management/services"
)

type UnsecureController struct {
	router gin.IRouter
}

func (ctl *UnsecureController) Login(c *gin.Context) {

	obj := services.AuthSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	usr, err := services.GetAdminAuthentication(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}
}

func (ctl *UnsecureController) CheckSession(c *gin.Context) {

	obj := services.AuthSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	usr, err := services.CheckSessionWithUsernameAndAPISecret(obj)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, usr)

	}
}

func (ctl *UnsecureController) Logout(c *gin.Context) {

	obj := services.AuthSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	session, err := services.FindLoggedSession(obj.Username)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		services.Logout(session)
		c.JSON(200, gin.H{
			"success": true,
		})
	}
}

// AuthenticationController creates and register handles
func NewUnsecureController(router gin.IRouter) *UnsecureController {
	pc := &UnsecureController{
		router: router,
	}

	pc.register()

	return pc

}

func (ctl *UnsecureController) ChangePassword(c *gin.Context) {

	type ChangePasswordSchema struct {
		UPid        string `json:"u_pid"`
		OldPassword string `json:"old_password"`
		Password    string `json:"password"`
	}

	obj := ChangePasswordSchema{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	CheckLdapConnection()

	err := services.ChangePassword(obj.UPid, obj.OldPassword, obj.Password)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return

	} else {
		c.JSON(200, gin.H{
			"success": true,
		})
	}
}

func (ctl *UnsecureController) register() {
	auth := ctl.router.Group("/authentication")
	auth.POST("", ctl.Login)
	auth.POST("logout", ctl.Logout)
	auth.POST("check", ctl.CheckSession)
	user := ctl.router.Group("/authentication")
	user.POST("change-password", ctl.ChangePassword)
}
