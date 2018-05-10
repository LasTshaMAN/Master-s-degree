package main

import (
	"fmt"
	"time"
)

func profile(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("====> %s took %s\n", name, elapsed)
}
