package enum

type ProblemDifficulty uint

const (
	ProblemDifficultyVeryEasy ProblemDifficulty = iota + 1 // 简单
	ProblemDifficultyEasy                                  // 容易
	ProblemDifficultyMedium                                // 中等
	ProblemDifficultyHard                                  // 困难
	ProblemDifficultyVeryHard                              // 非常困难
)
