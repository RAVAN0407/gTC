package main

import (
	"fmt"
	"testing"
)

func TestConvertGoToC(t *testing.T) {
	var testGo = Person{name: "John Doe", citNo: 12345, salary: 50000.0}
	var cStruct Person
	ConvertGoStructToCStruct(testGo, cStruct)
	fmt.Println(cStruct)

}
