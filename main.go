package main

import "fmt"

func main() {
	//map
	//Create map type1 use "make"
	a := make(map[string]string) //make([index]value)
	a["name"] = "Gola"
	a["surname"] = "Su"
	a["tell"] = "091-8651234"
	//----------------------

	//Create map type2
	b := map[string]string{
		"sport": "Tennis",
		"name":  "SubAlgo",
	}

	x, check := a["x"] //check a["x"] if a have a["x"] will return true if not will return false
	y, check2 := a["tell"]

	fmt.Println(x, check)  // false
	fmt.Println(y, check2) //091-8651234 true

	delete(a, "tell") //delete map a["tell"]
	y, check2 = a["tell"]

	fmt.Println("After delete ", y, check2)
	fmt.Println("")

	fmt.Println("Check map value by if statement")
	//Check map value by if statement
	if z, ck := a["nam"]; ck {
		fmt.Println(z, ck)
	} else {
		fmt.Println("Not init")
	}
	fmt.Println("")

	fmt.Println("show map(a) value by for range")
	for key, value := range a {
		fmt.Println(key, ":", value)
	}
	fmt.Println("")

	fmt.Println("show map(b) value by for range")
	for key, value := range b {
		fmt.Println(key, ":", value)
	}

}

//
//false
//091-8651234 true
//After delete   false

//Check map value by if statement
//Not init

//show map(a) value by for range
//name : Gola
//surname : Su

//show map(b) value by for range
//sport : Tennis
//name : SubAlgo
