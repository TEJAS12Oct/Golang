package Package2

import "fmt"

func Divide(a, b int) int {
	if b == 0 {
		panic("Division by zero is not allowed")
	}
	return a / b
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a / b, nil
}
