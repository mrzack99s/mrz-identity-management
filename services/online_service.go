package services

import (
	"context"

	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/ent/onlinesession"
)

func CreateOnlineSession(data ent.OnlineSession) (res *ent.OnlineSession, err error) {

	res, err = DB_CLIENT.OnlineSession.Create().
		SetIPAddress(data.IPAddress).
		SetUPid(data.UPid).
		Save(context.Background())

	return
}

func ReadAllOnlineSession() (res []*ent.OnlineSession, err error) {
	res, err = DB_CLIENT.OnlineSession.Query().All(context.Background())
	return
}

func CountAllOnlineSession() (count int, err error) {
	count, err = DB_CLIENT.OnlineSession.Query().Count(context.Background())
	return
}

func ReadOnlineSession(id int) (res *ent.OnlineSession, err error) {
	res, err = DB_CLIENT.OnlineSession.Query().Where(onlinesession.IDEQ(id)).Only(context.Background())
	return
}

func DeleteOnlineSession(id int) (err error) {
	err = DB_CLIENT.OnlineSession.DeleteOneID(id).Exec(context.Background())
	return
}
