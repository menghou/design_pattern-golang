package static_factory_method

import (
	"errors"
	"fmt"
)

type Operation interface {
	GetResult(numberA, numberB float64) (float64, error)
}

type OperationAdd struct {
}

func (a *OperationAdd) GetResult(numberA, numberB float64) (float64, error) {
	return numberA + numberB, nil
}

type OperationSub struct {
}

func (s *OperationSub) GetResult(numberA, numberB float64) (float64, error) {
	return numberA - numberB, nil
}

type OperationMul struct {
}

func (m *OperationMul) GetResult(numberA, numberB float64) (float64, error) {
	return numberA * numberB, nil
}

type OperationDiv struct {
}

func (d *OperationDiv) GetResult(numberA, numberB float64) (float64, error) {
	if numberA == 0 {
		return 0, errors.New("dividend are zero")
	} else {
		return numberA / numberB, nil
	}
}

type OperationFatory struct {
}

func (f *OperationFatory) CreateOperationObj(oper string) (Operation, error) {
	switch oper {
	case "+":
		return &OperationAdd{}, nil
	case "-":
		return &OperationSub{}, nil
	case "*":
		return &OperationMul{}, nil
	case "/":
		return &OperationDiv{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unknow operation %s", oper))
	}
}

func (f *OperationFatory) CreateOperationFunc(oper string) func(float64, float64) (float64, error) {
	operation, err := f.CreateOperationObj(oper)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}
	return operation.GetResult
}
