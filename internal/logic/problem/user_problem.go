package problem

import (
	"capys/internal/dao"
	"capys/internal/model/do"
	"capys/internal/model/entity"
	"capys/internal/service"
	"context"
)

type sUserProblem struct{}

func (s *sUserProblem) Create(ctx context.Context, in *entity.UserProblem) (err error) {
	_, err = dao.UserProblem.Ctx(ctx).Insert(in)
	return
}

func (s *sUserProblem) GetOne(ctx context.Context, id any) (out *entity.UserProblem, err error) {
	err = dao.UserProblem.Ctx(ctx).WherePri(id).Scan(&out)
	return
}

func (s *sUserProblem) Update(ctx context.Context, in *entity.UserProblem) (err error) {
	_, err = dao.UserProblem.Ctx(ctx).
		Where("user_id", in.UserId).
		Where("problem_id", in.ProblemId).
		Update(do.UserProblem{
			IsPass: in.IsPass,
		})
	return
}

func init() {
	service.RegisterUserProblem(&sUserProblem{})
}
