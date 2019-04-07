// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

//!+
func main() {
	before := time.Now()
	for index, argument := range os.Args {
		fmt.Println(index, argument)
	}
	after := time.Now()
	fmt.Println(after.Sub(before))

	before = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	after = time.Now()
	fmt.Println(after.Sub(before))

}

//!-
