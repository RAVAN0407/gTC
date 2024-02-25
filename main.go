package main

/*
	#include <stdio.h>
	#include <stdlib.h>

	struct person {
	char *name;
	int citNo;
	float salary;
	};
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

type Person struct {
	name   string
	citNo  int
	salary float64
}

// ConvertGoStructToCStruct converts a Go struct to a C struct
func ConvertGoStructToCStruct(goStruct interface{}, cStructInterface interface{}) interface{} {
	cStructType := reflect.TypeOf(cStructInterface)

	cStructValue := reflect.ValueOf(cStructInterface)

	t := reflect.TypeOf(goStruct)
	v := reflect.ValueOf(goStruct)

	for i := 0; i < t.NumField(); i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.Int:
			fieldValue := v.Field(i).Int()
			cStructField := cStructValue.FieldByName(cStructType.Field(i).Name)
			if cStructField.IsValid() && cStructField.CanSet() {
				cStructField.SetInt(fieldValue)
			}
		case reflect.Float64:
			fieldValue := v.Field(i).Float()
			cStructField := cStructValue.FieldByName(cStructType.Field(i).Name)
			if cStructField.IsValid() && cStructField.CanSet() {
				cStructField.SetFloat(fieldValue)
			}
		case reflect.String:
			fmt.Println("3")
			fieldValue := v.Field(i).String()
			name := C.CString(fieldValue)
			defer C.free(unsafe.Pointer(name))

			cStructField := cStructValue.FieldByName(cStructType.Field(i).Name)
			if cStructField.IsValid() && cStructField.CanSet() {
				cStructField.SetString(fieldValue)
			}
		}
	}
	return cStructInterface
}

func main() {
	var testGo = Person{name: "John Doe", citNo: 12345, salary: 50000.0}
	var cStruct C.struct_person
	res := ConvertGoStructToCStruct(testGo, cStruct)
	fmt.Println(res)
}
