package group

import (
	"capys/internal/dao"
	"capys/internal/model"
	"capys/internal/model/do"
	"capys/internal/service"
	"capys/utility/casql"
	"context"
)

type sGroup struct{}

// GetList implements service.IGroup.
func (s *sGroup) GetList(ctx context.Context, in *casql.BaseListInput) (out *casql.BaseListOut, err error) {
	md := dao.Group.Ctx(ctx)
	out, err = casql.MakePage(md, in.Size)
	if err != nil {
		return
	}

	data := []model.Group{}

	err = md.Page(int(in.Page), int(in.Size)).Scan(&data)
	out.Items = data
	return
}

// Delete implements service.IGroup.
func (s *sGroup) Delete(ctx context.Context, id any) (err error) {
	_, err = dao.Group.Ctx(ctx).WherePri(id).Delete()
	return
}

// GetOne implements service.IGroup.
func (s *sGroup) GetOne(ctx context.Context, id any) (out *model.Group, err error) {
	err = dao.Group.Ctx(ctx).WherePri(id).Scan(&out)
	return
}

func (s *sGroup) Update(ctx context.Context, id any, in *model.CreateGroupInput) (err error) {
	data := &do.Group{
		GroupName: in.GroupName,
	}
	_, err = dao.Group.Ctx(ctx).WherePri(id).Update(data)
	return
}

// Create implements service.IGroup.
func (s *sGroup) Create(ctx context.Context, in *model.CreateGroupInput) (err error) {
	_, err = dao.Group.Ctx(ctx).Insert(in)
	return
}

func init() {
	service.RegisterGroup(&sGroup{})
}
