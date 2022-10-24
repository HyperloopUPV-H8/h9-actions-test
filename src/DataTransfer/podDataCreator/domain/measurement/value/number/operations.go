package number

import (
	"regexp"
	"strconv"

	"github.com/HyperloopUPV-H8/Backend-H8/DataTransfer/utils"
)

type Operation struct {
	operator string
	operand  float32
}

func doOperation(number float32, operation Operation) float32 {
	switch operation.operator {
	case "+":
		return number + operation.operand
	case "-":
		return number - operation.operand
	case "*":
		return number * operation.operand
	case "/":
		return number / operation.operand
	default:
		panic("Invalid operation")
	}
}

func getOperations(ops string) []Operation {
	opExp, err := regexp.Compile(`([+\-\/*]{1})(\d+)`)

	if err != nil {
		utils.PrintRegexErr("opExp", err)
	}

	matches := opExp.FindAllStringSubmatch(ops, -1)
	operations := make([]Operation, 0)
	for _, match := range matches {
		operation := getOperation(match[1], match[2])
		operations = append(operations, operation)
	}
	return operations
}

func getOperation(operator string, operand string) Operation {
	num, err := strconv.ParseFloat(operand, 32)

	if err != nil {
		utils.PrintParseNumberErr(err)
	}

	return Operation{operator: operator, operand: float32(num)}
}

func getOpositeAndReversedOperations(operations []Operation) []Operation {
	newOperations := make([]Operation, len(operations))
	for index, operation := range operations {
		newOperations[index] = getOpositeOperation(operation)
	}
	newOperations = getReversedOperations(newOperations)
	return newOperations
}

func getOpositeOperation(operation Operation) Operation {
	opositeOperation := Operation{operand: operation.operand}
	switch operation.operator {
	case "+":
		opositeOperation.operator = "-"
	case "-":
		opositeOperation.operator = "+"

	case "*":
		opositeOperation.operator = "/"

	case "/":
		opositeOperation.operator = "*"
	default:
		panic("Invalid operator")
	}

	return opositeOperation
}

func getReversedOperations(operations []Operation) []Operation {
	length := len(operations)
	reversedOperations := make([]Operation, length)

	for index, operation := range operations {
		reversedOperations[length-1-index] = operation
	}

	return reversedOperations
}
