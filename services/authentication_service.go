package services

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/go-ldap/ldap/v3"
	"golang.org/x/text/encoding/unicode"

	"github.com/mrzack99s/mrz-identity-management/config"
	"github.com/mrzack99s/mrz-identity-management/ent/users"
)

type AuthResponse struct {
	Status  string `json:"status"`
	Role    string `json:"role"`
	Message string `json:"message"`
}

func GetAuthentication(username, password string, only_super_admin bool) (res AuthResponse) {
	r, e := DB_CLIENT.Users.Query().Where(users.Or(users.UOrgidEQ(username), users.UPidEQ(username))).WithInGroup().Only(context.Background())
	if e != nil {
		res.Status = "error"
		res.Message = "cannot find your personnel identity in organize"
		return
	}

	user_dn := ""
	if r.Edges.InGroup.GIsSuperAdmin {
		user_dn = fmt.Sprintf("CN=%s,CN=RTNAdministrator,%s", r.UOrgid, BASE_DN)
		res.Role = "super_admin"
	} else {
		res.Role = "user"
		if r.Edges.InGroup.GIsIntOrg {
			user_dn = fmt.Sprintf("CN=%s,OU=InternalPersonnel,%s", r.UOrgid, BASE_DN)
		} else {
			user_dn = fmt.Sprintf("CN=%s,OU=ExternalOrg,%s", r.UOrgid, BASE_DN)
		}
	}

	ldapURL := fmt.Sprintf("ldaps://%s", config.SystemConfigVar.LDAP.Hostname)
	l, err := ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		res.Status = "error"
		res.Message = "cannot connect to ldap server"
		return
	}
	defer l.Close()

	err = l.Bind(user_dn, password)
	if err != nil {
		res.Status = "error"
		res.Message = "password is not correct"
	} else {
		res.Status = "success"
	}

	l.Close()

	return
}

func ChangePassword(pid, old_password, password string) (err error) {

	res, e := DB_CLIENT.Users.Query().Where(users.UOrgid(pid)).WithInGroup().Only(context.Background())
	if e != nil {

		err = fmt.Errorf("cannot find your personnel identity in organize")
		return
	}

	user_dn := ""
	if res.Edges.InGroup.GIsSuperAdmin {
		user_dn = fmt.Sprintf("CN=%s,CN=RTNAdministrator,%s", res.UOrgid, BASE_DN)
	} else {
		if res.Edges.InGroup.GIsIntOrg {
			user_dn = fmt.Sprintf("CN=%s,OU=InternalPersonnel,%s", res.UOrgid, BASE_DN)
		} else {
			user_dn = fmt.Sprintf("CN=%s,OU=ExternalOrg,%s", res.UOrgid, BASE_DN)
		}
	}

	ldapURL := fmt.Sprintf("ldaps://%s", config.SystemConfigVar.LDAP.Hostname)
	l, err := ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {

		err = fmt.Errorf("cannot connect to ldap server")
		return
	}
	defer l.Close()

	err = l.Bind(user_dn, old_password)
	if err != nil {

		err = fmt.Errorf("password is not correct")
	}
	l.Close()

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String(fmt.Sprintf("%q", password))
	if err != nil {
		return
	}

	modReq := ldap.NewModifyRequest(user_dn, []ldap.Control{})
	modReq.Replace("unicodePwd", []string{pwdEncoded})
	err = LDAP_CLIENT.Modify(modReq)

	nowTime := time.Now()

	DB_CLIENT.Users.UpdateOneID(res.ID).SetUPasswordUpdatedAt(nowTime).SaveX(context.Background())

	return
}

func ResetPassword(pid string) (err error) {

	res, e := DB_CLIENT.Users.Query().Where(users.UPidEQ(pid)).WithInGroup().Only(context.Background())
	if e != nil {

		err = fmt.Errorf("cannot find your personnel identity in organize")
		return
	}

	user_dn := ""
	if res.Edges.InGroup.GIsSuperAdmin {
		user_dn = fmt.Sprintf("CN=%s,CN=RTNAdministrator,%s", res.UOrgid, BASE_DN)
	} else {
		if res.Edges.InGroup.GIsIntOrg {
			user_dn = fmt.Sprintf("CN=%s,OU=InternalPersonnel,%s", res.UOrgid, BASE_DN)
		} else {
			user_dn = fmt.Sprintf("CN=%s,OU=ExternalOrg,%s", res.UOrgid, BASE_DN)
		}
	}

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String(fmt.Sprintf("%q", res.UPid))
	if err != nil {
		return
	}

	modReq := ldap.NewModifyRequest(user_dn, []ldap.Control{})
	modReq.Replace("unicodePwd", []string{pwdEncoded})
	err = LDAP_CLIENT.Modify(modReq)

	nowTime := time.Now()

	DB_CLIENT.Users.UpdateOneID(res.ID).SetUPasswordUpdatedAt(nowTime).SaveX(context.Background())
	return
}
