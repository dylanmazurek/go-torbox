package constants

import (
	"errors"
	"strconv"
)

type PlanType int

func (p PlanType) String() string {
	switch p {
	case 0:
		return "Free"
	case 1:
		return "Essential"
	case 2:
		return "Pro"
	case 3:
		return "Standard"
	default:
		return "Unknown"
	}
}

func (p *PlanType) UnmarshalJSON(data []byte) error {
	var planStr = string(data)
	planInt, err := strconv.Atoi(planStr)
	if err != nil {
		return err
	}

	switch planInt {
	case 0:
		*p = PlanTypeFree
	case 1:
		*p = PlanTypeEssential
	case 2:
		*p = PlanTypePro
	case 3:
		*p = PlanTypeStandard
	default:
		return errors.New("unable to convert to plan type")
	}

	return nil
}

const (
	PlanTypeFree PlanType = iota
	PlanTypeEssential
	PlanTypePro
	PlanTypeStandard
)
