package main

import (
	"fmt"

	"gopkg.in/Knetic/govaluate.v2"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("5*x*x + 1*x + 2")
	if err != nil {
		fmt.Println("Failed to parse equation")
		fmt.Println(err.Error())
	}
	parameters := make(map[string]interface{}, 8)
	parameters["x"] = 2
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println("Failed to calculate equation")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
}
