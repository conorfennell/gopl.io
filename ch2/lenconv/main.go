package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/conorfennell/gopl.io/ch2/lenconv/conv"
)

func main() {

	for _, arg := range os.Args[1:] {
		value, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(lenconv.FeetToMeters(lenconv.Feet(value)).String())
		fmt.Println(lenconv.PoundsToKilograms(lenconv.Pounds(value)).String())
	}

}
