package problem

import (
	v1 "capys/api/problem/v1"
	"capys/internal/model"
	"capys/internal/model/entity"
	"capys/internal/service"
	"context"
	"fmt"
	"strconv"
)

func (c *ControllerV1) ProblemSubmit(ctx context.Context, req *v1.ProblemSubmitReq) (res *v1.ProblemSubmitRes, err error) {
	// 运行沙盒
	// runInfo, err := service.Problem().Run(ctx, req.Id, req.UserSubmitProblem)
	// if err != nil {
	// 	return
	// }

	runInfo := &model.RunInfo{ // DEBUG用的
		IsPass: req.IsPass,
	}

	isFirstPass := false

	addNumber := &model.AddNumber{}

	userProblem, _ := service.UserProblem().GetOne(ctx, req.Id)
	if userProblem == nil {
		addNumber.IsPass = runInfo.IsPass
		addNumber.IsParticipation = true

		isFirstPass = runInfo.IsPass

		ProblemId, _ := strconv.Atoi(req.Id)
		UserId, _ := strconv.Atoi(service.BizCtx().GetUserId(ctx))

		in := &entity.UserProblem{
			UserId:    uint64(UserId),
			ProblemId: uint64(ProblemId),
			IsPass:    runInfo.IsPass,
		}
		service.UserProblem().Create(ctx, in)
	} else {
		if !userProblem.IsPass && runInfo.IsPass {
			addNumber.IsPass = true
			isFirstPass = true
		}
	}

	// 通关奖励
	if isFirstPass {
		// TODO: 获得奖励
		fmt.Println("获得奖励")
	}

	fmt.Println("---------------")
	err = service.Problem().AddNumber(ctx, req.Id, addNumber)
	fmt.Println("---------------")

	return
}
