package structuraldesignpattern

import "log"

// FuncObject type define function
type FuncObject func(int) int

// LogDecorate log
func LogDecorate(fn FuncObject) FuncObject {
	return func(n int) int {
		log.Println("Starting the execution with the integer", n)

		result := fn(n)

		log.Println("Execution is completed with the result", result)

		return result
	}
}
