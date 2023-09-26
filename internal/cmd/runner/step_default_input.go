package runner

import (
	"github.com/gontainer/gontainer/internal/pkg/consts"
	"github.com/gontainer/gontainer/internal/pkg/input"
	"github.com/gontainer/gontainer/internal/pkg/output"
)

type StepDefaultInput struct {
}

func (StepDefaultInput) Name() string {
	return "Default input"
}

func (StepDefaultInput) Run(i *input.Input, _ *output.Output) error {
	i.Meta.Functions = map[string]string{
		consts.FuncEnv:    consts.BuiltInGetEnv,
		consts.FuncEnvInt: consts.BuiltinGetEnvInt,
		consts.FuncTodo:   consts.BuiltInParamTodo,
	}
	return nil
}
