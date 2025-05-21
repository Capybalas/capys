package model

import (
	"capys/internal/enum"
	"capys/internal/model/entity"

)

type InProblem struct {
	Title      string                 `json:"title" v:"required|min-length:1|max-length:255#您必须提供简略题目|简略题目至少1个字符|简略题目最多255个字符"` // 简单的题目
	Topic      string                 `json:"topic" v:"required|min-length:1|max-length:255#您必须提供题目|题目至少1个字符|题目最多255个字符"`       // 详细题目
	Category   enum.ProblemCategory   `json:"category" v:"integer|min:1|max:5#题目分类必须符合规范|题目分类错误|题目分类错误" d:"1"`                  // 1 html
	Difficulty enum.ProblemDifficulty `json:"difficulty" v:"integer|min:1|max:5#难度必须符合规范|难度分类错误|难度分类错误" d:"1"`                  // 1 入门 2简单 3中等 4难 5困难
	Answer     string                 `json:"answer" v:"required|json#您必须提供正确答案|答案格式不正确"`                                       //
}

type OutProblem struct {
	*entity.Problem
}
