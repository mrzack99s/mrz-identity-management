package services

import (
	"context"
	"errors"

	"github.com/mrzack99s/mrz-identity-management/ent"
	"github.com/mrzack99s/mrz-identity-management/ent/groupbandwidth"
)

func CreateGroupBandwidth(data ent.GroupBandwidth) (res *ent.GroupBandwidth, err error) {
	_, err = DB_CLIENT.GroupBandwidth.Query().Where(
		groupbandwidth.And(
			groupbandwidth.GbwDownloadSpeedEQ(data.GbwDownloadSpeed),
			groupbandwidth.GbwUploadSpeedEQ(data.GbwUploadSpeed),
		),
	).Only(context.Background())
	if err == nil {
		err = errors.New("duplicate bandwidth profile")
		return
	}

	res, err = DB_CLIENT.GroupBandwidth.Create().
		SetGbwDownloadSpeed(data.GbwDownloadSpeed).
		SetGbwUploadSpeed(data.GbwUploadSpeed).
		Save(context.Background())

	return
}

func ReadAllAllGroupBandwidth() (res []*ent.GroupBandwidth, err error) {
	res, err = DB_CLIENT.GroupBandwidth.Query().All(context.Background())
	return
}

func ReadGroupBandwidthWithPagination(page, p_page int) (res []*ent.GroupBandwidth, record_count int, err error) {
	if page > 1 {
		page -= 1
		page *= p_page
	} else {
		page = 0
	}

	res, err = DB_CLIENT.GroupBandwidth.Query().Offset(page).Limit(p_page).All(context.Background())
	if err != nil {
		return
	}

	record_count, err = DB_CLIENT.GroupBandwidth.Query().Count(context.Background())
	if err != nil {
		return
	}

	return
}

func ReadGroupBandwidth(id int) (res *ent.GroupBandwidth, err error) {
	res, err = DB_CLIENT.GroupBandwidth.Query().Where(groupbandwidth.IDEQ(id)).Only(context.Background())
	return
}

func UpdateGroupBandwidth(data ent.GroupBandwidth) (res *ent.GroupBandwidth, err error) {
	res, err = DB_CLIENT.GroupBandwidth.UpdateOneID(data.ID).
		SetGbwDownloadSpeed(data.GbwDownloadSpeed).
		SetGbwUploadSpeed(data.GbwUploadSpeed).Save(context.Background())
	return
}

func DeleteGroupBandwidth(id int) (err error) {
	numberOfGroupAllocation := DB_CLIENT.GroupBandwidth.Query().Where(groupbandwidth.IDEQ(id)).QueryGroups().CountX(context.Background())
	if numberOfGroupAllocation > 0 {
		err = errors.New("having a group allocated, can't delete it")
		return
	}
	err = DB_CLIENT.GroupBandwidth.DeleteOneID(id).Exec(context.Background())
	return
}
