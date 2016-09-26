package main

import (
	. "design_pattern-golang/design_pattern-golang/static-factory-method"
	"fmt"
	"log"
)

func main() {
	operationFactory := OperationFatory{}
	operationAdd, err := operationFactory.CreateOperationObj("+")
	if err != nil {
		log.Fatal(err)
	}
	resultAdd, _ := operationAdd.GetResult(1,2)
	fmt.Printf("1 + 2 = %f\n", resultAdd)
	operationSub := operationFactory.CreateOperationFunc("-")
	resultSub , _ := operationSub(1,2)
	fmt.Printf("1 - 2 = %f\n", resultSub)
}
