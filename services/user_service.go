package services

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/ent/groups"
	"github.com/mrzack99s/mrz-identity-management/ent/users"
	"golang.org/x/text/encoding/unicode"
)

type MultiCreateSchema struct {
	PersonalID    string `json:"personal_id"`
	OrgID         string `json:"org_id"`
	GroupName     string `json:"group_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	EnableAccount string `json:"enable_account"`
	Years         int    `json:"years"`
}
type MultiCreateErrorMsg struct {
	PersonalID string `json:"personal_id"`
	Error      string `json:"error"`
}

func CreateUser(data ent.Users) (res *ent.Users, err error) {

	if data.UExpiredAt.Unix() == -62135596800 {
		res, err = DB_CLIENT.Users.Create().
			SetUPid(data.UPid).
			SetUOrgid(data.UOrgid).
			SetUFirstName(data.UFirstName).
			SetULastName(data.ULastName).
			SetUIsActive(data.UIsActive).
			SetInGroup(data.Edges.InGroup).
			Save(context.Background())

	} else {
		res, err = DB_CLIENT.Users.Create().
			SetUPid(data.UPid).
			SetUOrgid(data.UOrgid).
			SetUFirstName(data.UFirstName).
			SetULastName(data.ULastName).
			SetUIsActive(data.UIsActive).
			SetUExpiredAt(data.UExpiredAt).
			SetInGroup(data.Edges.InGroup).
			Save(context.Background())
	}

	if err != nil {
		return
	}

	user_dn := ""
	if data.Edges.InGroup.GIsSuperAdmin {
		user_dn = fmt.Sprintf("CN=%s,CN=RTNAdministrator,%s", res.UOrgid, BASE_DN)
	} else {
		if data.Edges.InGroup.GIsIntOrg {
			user_dn = fmt.Sprintf("CN=%s,OU=InternalPersonnel,%s", res.UOrgid, BASE_DN)
		} else {
			user_dn = fmt.Sprintf("CN=%s,OU=ExternalOrg,%s", res.UOrgid, BASE_DN)
		}
	}

	addReq := ldap.NewAddRequest(user_dn, []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "organizationalPerson", "user", "person"})
	addReq.Attribute("givenName", []string{res.UFirstName})
	addReq.Attribute("sn", []string{res.ULastName})
	addReq.Attribute("description", []string{fmt.Sprintf("%s %s %s", res.UPid, res.UFirstName, res.ULastName)})
	addReq.Attribute("sAMAccountName", []string{res.UOrgid})
	addReq.Attribute("displayName", []string{res.UOrgid})
	addReq.Attribute("uid", []string{res.UOrgid})
	addReq.Attribute("mail", []string{fmt.Sprintf("%s@%s", res.UOrgid, EMAIL_DOMAIN_NAME)})
	addReq.Attribute("userAccountControl", []string{fmt.Sprintf("%d", 0x0202)})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("userPrincipalName", []string{fmt.Sprintf("%s@%s", res.UOrgid, DOMAIN_NAME)})

	if res.UExpiredAt.Unix() == -62135596800 {
		addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", NEVER_EXPIRE)})
	} else {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		timeInThailand := res.UExpiredAt.In(loc)
		ldapTime := (timeInThailand.Unix() + 11644473600) * 10000000
		addReq.Attribute("accountExpires", []string{fmt.Sprintf("%d", ldapTime)})
	}

	if err = LDAP_CLIENT.Add(addReq); err != nil {
		DB_CLIENT.Users.DeleteOneID(res.ID).ExecX(context.Background())
		return
	}

	group_dn := fmt.Sprintf("CN=%s,CN=RTNGroups,%s", data.Edges.InGroup.GName, BASE_DN)
	modReq := ldap.NewModifyRequest(group_dn, []ldap.Control{})
	modReq.Add("member", []string{user_dn})
	if err = LDAP_CLIENT.Modify(modReq); err != nil {
		return
	}

	utf16 := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM)
	pwdEncoded, err := utf16.NewEncoder().String(fmt.Sprintf("%q", data.UPid))
	if err != nil {
		return
	}

	modReq = ldap.NewModifyRequest(user_dn, []ldap.Control{})
	modReq.Replace("unicodePwd", []string{pwdEncoded})
	err = LDAP_CLIENT.Modify(modReq)

	if res.UIsActive {
		modReq := ldap.NewModifyRequest(user_dn, []ldap.Control{})
		modReq.Replace("userAccountControl", []string{fmt.Sprintf("%d", 0x0200)})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return
		}
	}

	return
}

func UpsertUser(data ent.Users) (err error) {

	_, err = DB_CLIENT.Users.Query().Where(users.Or(
		users.UPidEQ(data.UPid),
		users.UOrgidEQ(data.UOrgid),
	)).WithInGroup().Only(context.Background())
	if err != nil {
		_, err = CreateUser(data)
	} else {
		_, err = UpdateUser(data)
	}

	return
}

func MultiCreateUser(data []MultiCreateSchema) (errmsg []MultiCreateErrorMsg) {
	for _, usr := range data {

		g, err := DB_CLIENT.Groups.Query().Where(groups.GNameEQ(usr.GroupName)).Only(context.Background())
		if err != nil {
			newErrMsg := MultiCreateErrorMsg{
				PersonalID: usr.PersonalID,
				Error:      fmt.Sprintf("%s", err),
			}
			errmsg = append(errmsg, newErrMsg)
			continue
		} else {

			newUser := ent.Users{}
			newUser.Edges.InGroup = g
			newUser.UPid = usr.PersonalID
			newUser.UOrgid = usr.OrgID
			newUser.UFirstName = usr.FirstName
			newUser.ULastName = usr.LastName

			isEnable, _ := strconv.ParseBool(usr.EnableAccount)
			newUser.UIsActive = isEnable

			if usr.Years != 0 {
				expired := time.Now().AddDate(usr.Years, 0, 0).Round(0)
				newUser.UExpiredAt = expired
			}

			err = UpsertUser(newUser)
			if err != nil {
				newErrMsg := MultiCreateErrorMsg{
					PersonalID: usr.PersonalID,
					Error:      fmt.Sprintf("%s", err),
				}
				errmsg = append(errmsg, newErrMsg)
			}
		}

	}

	return
}

func ReadAllUser() (res []*ent.Users, err error) {
	res, err = DB_CLIENT.Users.Query().WithInGroup().All(context.Background())
	return
}

func ReadUserWithPagination(page, p_page int) (res []*ent.Users, record_count int, err error) {
	if page > 1 {
		page -= 1
		page *= p_page
	} else {
		page = 0
	}

	res, err = DB_CLIENT.Users.Query().WithInGroup().Offset(page).Limit(p_page).All(context.Background())
	if err != nil {
		return
	}

	record_count, err = DB_CLIENT.Users.Query().Count(context.Background())
	if err != nil {
		return
	}

	return
}

func ReadUserWithSearch(search string, page, p_page int) (res []*ent.Users, record_count int, err error) {

	if page > 1 {
		page -= 1
		page *= p_page
	} else {
		page = 0
	}

	res, err = DB_CLIENT.Users.Query().Where(
		users.Or(
			users.UPidContains(search),
			users.UOrgidContains(search),
			users.UFirstNameContains(search),
			users.ULastNameContains(search),
			users.HasInGroupWith(groups.GNameContains(search)),
		)).WithInGroup().Offset(page).Limit(p_page).All(context.Background())
	if err != nil {
		return
	}

	record_count, err = DB_CLIENT.Users.Query().Where(
		users.Or(
			users.UPidContains(search),
			users.UOrgidContains(search),
			users.UFirstNameContains(search),
			users.ULastNameContains(search),
			users.HasInGroupWith(groups.GNameContains(search)),
		)).Count(context.Background())
	if err != nil {
		return
	}

	return
}

func CountAllUser() (count int, err error) {
	count, err = DB_CLIENT.Users.Query().Count(context.Background())
	return
}

func ReadUser(id int) (res *ent.Users, err error) {
	res, err = DB_CLIENT.Users.Query().WithInGroup().Where(users.IDEQ(id)).Only(context.Background())
	return
}

func UpdateUser(data ent.Users) (res *ent.Users, err error) {

	beforUpdate, err := DB_CLIENT.Users.Query().Where(users.UPidEQ(data.UPid)).WithInGroup().Only(context.Background())
	if err != nil {
		return
	}

	if data.UExpiredAt.Unix() == -62135596800 {
		res, err = DB_CLIENT.Users.UpdateOneID(data.ID).
			SetUFirstName(data.UFirstName).
			SetULastName(data.ULastName).
			SetUIsActive(data.UIsActive).
			SetInGroup(data.Edges.InGroup).
			ClearUExpiredAt().
			Save(context.Background())
	} else {
		res, err = DB_CLIENT.Users.UpdateOneID(data.ID).
			SetUFirstName(data.UFirstName).
			SetULastName(data.ULastName).
			SetUIsActive(data.UIsActive).
			SetInGroup(data.Edges.InGroup).
			SetUExpiredAt(data.UExpiredAt).
			Save(context.Background())
	}

	if err != nil {
		return
	}

	res.Edges.InGroup = res.QueryInGroup().OnlyX(context.Background())

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

	req := ldap.NewModifyRequest(user_dn, []ldap.Control{})
	req.Replace("givenName", []string{res.UFirstName})
	req.Replace("sn", []string{res.ULastName})
	req.Replace("description", []string{fmt.Sprintf("%s %s %s", res.UPid, res.UFirstName, res.ULastName)})

	if res.UIsActive {
		req.Replace("userAccountControl", []string{fmt.Sprintf("%d", 0x0200)})
	} else {
		req.Replace("userAccountControl", []string{fmt.Sprintf("%d", 0x0202)})
	}

	if res.UExpiredAt.Unix() == -62135596800 {
		req.Replace("accountExpires", []string{fmt.Sprintf("%d", NEVER_EXPIRE)})
	} else {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		timeInThailand := res.UExpiredAt.In(loc)
		ldapTime := (timeInThailand.Unix() + 11644473600) * 10000000
		req.Replace("accountExpires", []string{fmt.Sprintf("%d", ldapTime)})
	}

	if err = LDAP_CLIENT.Modify(req); err != nil {
		return
	}

	group_dn := fmt.Sprintf("CN=%s,CN=RTNGroups,%s", beforUpdate.Edges.InGroup.GName, BASE_DN)
	modReq := ldap.NewModifyRequest(group_dn, []ldap.Control{})
	modReq.Delete("member", []string{user_dn})
	if err = LDAP_CLIENT.Modify(modReq); err != nil {
		return nil, err
	}

	group_dn = fmt.Sprintf("CN=%s,CN=RTNGroups,%s", res.Edges.InGroup.GName, BASE_DN)
	modReq = ldap.NewModifyRequest(group_dn, []ldap.Control{})
	modReq.Add("member", []string{user_dn})
	if err = LDAP_CLIENT.Modify(modReq); err != nil {
		return nil, err
	}

	return
}

func DeleteUser(id int) (err error) {
	res, err := DB_CLIENT.Users.Query().Where(users.IDEQ(id)).WithInGroup().Only(context.Background())
	if err != nil {
		return err
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
	delReq := ldap.NewDelRequest(user_dn, []ldap.Control{})

	if err := LDAP_CLIENT.Del(delReq); err != nil {
		log.Fatalf("Error deleting service: %v", err)
	}

	err = DB_CLIENT.Users.DeleteOneID(id).Exec(context.Background())
	return
}
