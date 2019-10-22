package main

import "fmt"
import "strings"

func printif(src interface{}) {
	if value, ok := src.(int); ok {
		fmt.Printf("parameter is int [value: %d].\n", value)
		return
	}

	if value, ok := src.(string); ok {
		value = strings.ToUpper(value)
		fmt.Printf("paramter is string [value: %s].\n", value)
		return
	}

	if value, ok := src.([]string); ok {
		value = append(value, "unknown")
		fmt.Printf("paramter is slice string. [value: %s\n", value)
		return
	}

	if value, ok := src.([3]string); ok {
		fmt.Printf("paramter is array3 string. [value: %s\n", value)
		return
	}

	if value, ok := src.([2]string); ok {
		fmt.Printf("paramter is array2 string. [value: %s\n", value)
		return
	}

	
	fmt.Printf("paramter is unknown type valueType: %T\n", src)
}

func main() {
	printif(12)
	printif("hello")
	printif([]string{"cat", "dog"})
	printif([2]string{"hello", "world"})
}
