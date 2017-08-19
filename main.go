package main

import "fmt"

func main() {
	checkType("SubAlgo")
	checkType(10)
	checkType(true)

	checkType2("SubAlgo")
	checkType2(10)
	checkType2(true)
}

func checkType(v interface{}) {
	switch p := v.(type) {
	case string:
		fmt.Println("String: ", p)
	case int:
		fmt.Println("Int: ", p+10)
	case bool:
		fmt.Println("Bool: ", p)
	}
}

func checkType2(v interface{}) {
	param, ok := v.(string)

	if ok {
		fmt.Println("Param is string: ", param)
	} else {
		fmt.Println("Param is not string!!")
	}
}
