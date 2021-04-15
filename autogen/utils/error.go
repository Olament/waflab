package utils

import "fmt"

type ErrNotSupported struct {
	Name string
	Type string
}

func (e *ErrNotSupported) Error() string {
	return fmt.Sprintf("autogen/factory: %s %s not supported", e.Type, e.Name)
}
