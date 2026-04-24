package main

import (
	"fmt"
	"goDemo/color"
	"goDemo/mathops"
)

func main() {
	// Demo 1: stringer - auto-generated String() method for Color enum
	fmt.Println("=== Demo 1: stringer (enum -> String()) ===")
	colors := []color.Color{color.Red, color.Green, color.Blue, color.Yellow, color.Purple}
	for _, c := range colors {
		fmt.Printf("  color value %d -> %s\n", c, c)
	}

	// Demo 2: template-based code generator - math operations
	fmt.Println("\n=== Demo 2: template generator (math ops) ===")
	fmt.Printf("  Add(3, 4)      = %d\n", mathops.Add(3, 4))
	fmt.Printf("  Subtract(10, 3) = %d\n", mathops.Subtract(10, 3))
	fmt.Printf("  Multiply(6, 7)  = %d\n", mathops.Multiply(6, 7))

	fmt.Println("\n  All ops via map:")
	for name, fn := range mathops.All() {
		fmt.Printf("    %-10s -> %d\n", name, fn(10, 5))
	}
}
