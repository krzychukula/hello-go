package main

import (
	"fmt"
	"github.com/krzychukula/newmath"
	"math"
)

func main() {
	fmt.Printf("Sqrt(2) = %v. math.Sqrt(2) = %v. Diff:%v\n", newmath.Sqrt(2), math.Sqrt(2), newmath.Sqrt(2)-math.Sqrt(2))
}
