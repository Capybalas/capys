package enum

type PowerType uint

const (
	PowerStudent PowerType = iota + 1
	PowerTeacher
	PowerAdmin
)
