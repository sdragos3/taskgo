package models

import "fmt"

type Priority string

const (
	High   Priority = "high"
	Medium          = "medium"
	Low             = "low"
	None            = "none"
)

func (p *Priority) Set(s string) error {
	switch s {
	case "low", "medium", "high", "none":
		*p = Priority(s)
		return nil
	default:
		return fmt.Errorf("invalid Priority: %s", s)
	}
}

func (p *Priority) String() string {
	return string(*p)
}

func (p *Priority) Type() string {
	return "Priority"
}
