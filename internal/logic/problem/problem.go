package problem

import (
	"capys/internal/dao"
	"capys/internal/model"
	"capys/internal/model/do"
	"capys/internal/service"
	"capys/utility/casql"
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sProblem struct{}

// AddNumber implements service.IProblem.
func (s *sProblem) AddNumber(ctx context.Context, id any, in *model.AddNumber) (err error) {

	data := g.Map{}

	if in.IsPass {
		data["passing_number"] = gdb.Raw("passing_number+1")
	}

	if in.IsParticipation {
		data["participation_number"] = gdb.Raw("participation_number+1")
	}
	g.Model("problem").Data(data).WherePri(id).Update()
	return
}

// Run implements service.IProblem.
func (s *sProblem) Run(ctx context.Context, id any, userCode *model.UserSubmitProblem) (out *model.RunInfo, err error) {
	out = &model.RunInfo{
		IsPass: true,
	}
	return
}

// Delete implements service.IProblem.
func (s *sProblem) Delete(ctx context.Context, id any) (err error) {
	_, err = dao.Problem.Ctx(ctx).WherePri(id).Delete()
	return
}

// GetList implements service.IProblem.
func (s *sProblem) GetList(ctx context.Context, in *casql.BaseListInput) (out *casql.BaseListOut, err error) {
	md := dao.Problem.Ctx(ctx)
	out, err = casql.MakePage(md, in.Size)
	if err != nil {
		return
	}

	data := []model.OutProblem{}

	err = md.Page(int(in.Page), int(in.Size)).Scan(&data)
	out.Items = data
	return
}

// GetOne implements service.IProblem.
func (s *sProblem) GetOne(ctx context.Context, id any) (out *model.OutProblem, err error) {
	err = dao.Problem.Ctx(ctx).WherePri(id).Scan(&out)
	return
}

func (s *sProblem) Create(ctx context.Context, in *model.InProblem) (err error) {
	data := &do.Problem{}
	gconv.Scan(in, &data)
	_, err = dao.Problem.Ctx(ctx).Insert(data)
	return
}

func (s *sProblem) Update(ctx context.Context, id any, in *model.InProblem) (err error) {
	data := &do.Problem{}
	gconv.Scan(in, &data)
	_, err = dao.Problem.Ctx(ctx).WherePri(id).Update(data)
	return
}

func init() {
	service.RegisterProblem(&sProblem{})
}
