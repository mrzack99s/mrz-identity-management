package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/go-ldap/ldap/v3"
	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/ent/groupbandwidth"
	"github.com/mrzack99s/mrz-identity-management/ent/groups"
)

func CreateGroup(data ent.Groups) (res *ent.Groups, err error) {

	_, err = DB_CLIENT.Groups.Query().Where(
		groups.And(
			groups.GName(data.GName),
			groups.HasUseBandwidthWith(groupbandwidth.IDEQ(data.Edges.UseBandwidth.ID)),
		),
	).Only(context.Background())
	if err == nil {
		err = errors.New("duplicate group profile")
		return
	}

	res, err = DB_CLIENT.Groups.Create().
		SetGName(data.GName).
		SetGIsIntOrg(data.GIsIntOrg).
		SetGIsSuperAdmin(data.GIsSuperAdmin).
		SetUseBandwidth(data.Edges.UseBandwidth).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	group_dn := fmt.Sprintf("CN=%s,CN=RTNGroups,%s", res.GName, BASE_DN)
	addReq := ldap.NewAddRequest(group_dn, []ldap.Control{})
	addReq.Attribute("objectClass", []string{"top", "group"})
	addReq.Attribute("name", []string{res.GName})
	addReq.Attribute("sAMAccountName", []string{res.GName})
	addReq.Attribute("instanceType", []string{fmt.Sprintf("%d", 0x00000004)})
	addReq.Attribute("groupType", []string{fmt.Sprintf("%d", 0x80000002)})

	if err = LDAP_CLIENT.Add(addReq); err != nil {
		DB_CLIENT.Groups.DeleteOneID(res.ID).ExecX(context.Background())
		return nil, err
	}

	modReq := ldap.NewModifyRequest(fmt.Sprintf("CN=AllowNetworkAccess,%s", BASE_DN), []ldap.Control{})
	modReq.Add("member", []string{group_dn})
	if err = LDAP_CLIENT.Modify(modReq); err != nil {
		return nil, err
	}

	if res.GIsSuperAdmin {
		modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=SuperAdmin,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
		modReq.Add("member", []string{group_dn})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return nil, err
		}
	} else {
		if res.GIsIntOrg {
			modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=InternalPersonnelGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Add("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}
		} else {
			modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=ExternalOrgGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Add("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}
		}
	}

	return
}

func ReadAllGroup() (res []*ent.Groups, err error) {
	res, err = DB_CLIENT.Groups.Query().WithUseBandwidth().All(context.Background())
	return
}

func ReadGroupWithSearch(search string, page, p_page int) (res []*ent.Groups, record_count int, err error) {

	if page > 1 {
		page -= 1
		page *= p_page
	} else {
		page = 0
	}

	res, err = DB_CLIENT.Groups.Query().Where(
		groups.Or(
			groups.GNameContains(search),
		)).WithUseBandwidth().Offset(page).Limit(p_page).All(context.Background())
	if err != nil {
		return
	}

	record_count, err = DB_CLIENT.Groups.Query().Where(
		groups.Or(
			groups.GNameContains(search),
		)).WithUseBandwidth().Count(context.Background())
	if err != nil {
		return
	}

	return
}

func ReadGroupWithPagination(page, p_page int) (res []*ent.Groups, record_count int, err error) {
	if page > 1 {
		page -= 1
		page *= p_page
	} else {
		page = 0
	}

	res, err = DB_CLIENT.Groups.Query().WithUseBandwidth().Offset(page).Limit(p_page).All(context.Background())
	if err != nil {
		return
	}

	record_count, err = DB_CLIENT.Groups.Query().Count(context.Background())
	if err != nil {
		return
	}

	return
}

func CountAllGroup() (count int, err error) {
	count, err = DB_CLIENT.Groups.Query().Count(context.Background())
	return
}

func ReadGroup(id int) (res *ent.Groups, err error) {
	res, err = DB_CLIENT.Groups.Query().WithUseBandwidth().Where(groups.IDEQ(id)).Only(context.Background())
	return
}

func UpdateGroup(data ent.Groups) (res *ent.Groups, err error) {
	res, err = DB_CLIENT.Groups.UpdateOneID(data.ID).
		SetGName(data.GName).
		SetGIsIntOrg(data.GIsIntOrg).
		SetGIsSuperAdmin(data.GIsSuperAdmin).
		SetUseBandwidth(data.Edges.UseBandwidth).Save(context.Background())
	if err != nil {
		return nil, err
	}

	group_dn := fmt.Sprintf("CN=%s,CN=RTNGroups,%s", res.GName, BASE_DN)
	if res.GIsSuperAdmin {

		modReq := ldap.NewModifyRequest(fmt.Sprintf("CN=ExternalOrgGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
		modReq.Delete("member", []string{group_dn})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return nil, err
		}

		modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=InternalPersonnelGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
		modReq.Delete("member", []string{group_dn})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return nil, err
		}

		modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=SuperAdmin,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
		modReq.Add("member", []string{group_dn})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return nil, err
		}
	} else {

		modReq := ldap.NewModifyRequest(fmt.Sprintf("CN=SuperAdmin,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
		modReq.Delete("member", []string{group_dn})
		if err = LDAP_CLIENT.Modify(modReq); err != nil {
			return nil, err
		}

		if res.GIsIntOrg {

			modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=ExternalOrgGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Delete("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}

			modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=InternalPersonnelGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Add("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}

		} else {
			modReq := ldap.NewModifyRequest(fmt.Sprintf("CN=InternalPersonnelGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Delete("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}

			modReq = ldap.NewModifyRequest(fmt.Sprintf("CN=ExternalOrgGroup,CN=RTNGroups,%s", BASE_DN), []ldap.Control{})
			modReq.Add("member", []string{group_dn})
			if err = LDAP_CLIENT.Modify(modReq); err != nil {
				return nil, err
			}
		}
	}
	return
}

func DeleteGroup(id int) (err error) {
	res, err := DB_CLIENT.Groups.Query().Where(groups.IDEQ(id)).WithUsers().Only(context.Background())
	if err != nil {
		return err
	}

	for _, usr := range res.Edges.Users {
		err = DeleteUser(usr.ID)
		if err != nil {
			return
		}
	}

	err = DB_CLIENT.Groups.DeleteOneID(id).Exec(context.Background())

	group_dn := fmt.Sprintf("CN=%s,CN=RTNGroups,%s", res.GName, BASE_DN)
	delReq := ldap.NewDelRequest(group_dn, []ldap.Control{})

	if err := LDAP_CLIENT.Del(delReq); err != nil {
		log.Fatalf("Error deleting service: %v", err)
	}

	return
}
