package main

import (
	"fmt"

	"myproject/Package"
	"myproject/Package2"
	"myproject/Package3"
)

func main() {

	var ExportedVariable = "Hello, World!"
	fmt.Println(ExportedVariable)
	fmt.Println(Package.AnotherExportedVariable)
	fmt.Println(Package2.StringOperation("Tejas"))

	boyAge := 12
	if boyAge < 18 {
		fmt.Println("Age status: Non Adult")
	} else {
		fmt.Println("Age status: Adult")
	}

	a := 10
	b := 20

	if a > b {
		fmt.Println("a is greater than b")
	} else if a < b {
		fmt.Println("a is less than b")
	} else {
		fmt.Println("a is equal to b")
	}
	fmt.Printf("Sum: %d\n", Package.Sum(a, b))

	fmt.Printf("Subtract: %d\n", Package.Subtract(b, a))
	fmt.Printf("Multiply: %d\n", Package3.Multiply(3, 4))

	if quotient, err := Package2.SafeDivide(20, 5); err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Printf("Divide: %d\n", quotient)
	}

	if remainder, err := Package2.Remainder(20, 3); err != nil {
		fmt.Println("Remainder error:", err)
	} else {
		fmt.Printf("Remainder: %d\n", remainder)
	}

	x := 10
	y := 3.14
	z := 'A'
	s := "Hello"
	flag := true

	fmt.Printf("Integer: %d\n", x)
	fmt.Printf("Binary: %b\n", x)
	fmt.Printf("Hex: %x\n", x)

	fmt.Printf("Float: %f\n", y)
	fmt.Printf("Float (2 decimal): %.2f\n", y)

	fmt.Printf("Character: %c\n", z)
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Quoted string: %q\n", s)

	fmt.Printf("Boolean: %t\n", flag)

	fmt.Printf("Type of x: %T\n", x)
	fmt.Printf("Default format: %v\n", s)

}
