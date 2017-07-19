package process

import (
	"errors"

	"code.cloudfoundry.org/garden"
)

type Process struct{}

func (p *Process) ID() string {
	return ""
}

func (p *Process) Wait() (int, error) {
	return 0, errors.New("Not implemented")
}

func (p *Process) SetTTY(garden.TTYSpec) error {
	return errors.New("Not implemented")
}

func (p *Process) Signal(garden.Signal) error {
	return errors.New("Not implemented")
}
