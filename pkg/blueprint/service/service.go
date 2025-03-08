package service

var governor ComputeGovernor = nil

// var _ ComputeGovernor = &computeGovernor{}(nil)

type ComputeGovernor interface {
	Start(limit int)
}

type computeGovernor struct {
	limit int
}

func (g *computeGovernor) Start(limit int) {
	governor = g
}

type Runnable[T, U any] func(T) U

type ComputeService[T, U any] interface {
	Run(Runnable[T, U]) U
}

type computeService[T, U any] struct {
	runnable Runnable[T, U]
}

func NewComputeService[T, U any](runnable Runnable[T, U]) U {

}
