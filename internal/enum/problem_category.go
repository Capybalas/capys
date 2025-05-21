package enum

type ProblemCategory uint

const (
	ProblemCategoryHtml ProblemCategory = iota + 1 // html
	ProblemCategoryCss                             // css
	ProblemCategoryJs                              // js
)
