package main

import (
	"fmt"
	"os"
	"testing"
)

var testCounter = 0
var testSuccess = 0

// -----------------------------------------------------------------
func TestMain(m *testing.M) {
	// call flag.Parse() here if TestMain uses flags
	fmt.Println("========================================")
	fmt.Println(appName, " Test suite ---")
	fmt.Println("========================================")

	// THIS IS THE CORE OF THE TEST EXECUTION
	res := m.Run()

	fmt.Println("========================================")
	fmt.Println(appName, " Test suite res")
	fmt.Println("========================================")
	fmt.Println("Errors reported", res)
	fmt.Println("========================================")
	fmt.Printf("Executed %3d tests", testCounter)
	fmt.Println()
	fmt.Printf("Success  %3d tests", testSuccess)
	fmt.Println()
	fmt.Printf("FAILED   %3d tests", testCounter-testSuccess)
	fmt.Println()

	os.Exit(res)
}
