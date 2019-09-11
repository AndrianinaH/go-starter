package main

import (
	"fmt"
)

func main() {
	fmt.Printf("hello, world\n")

	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var firstVar = "fomby"
	fmt.Println(firstVar)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// The := syntax is shorthand for declaring and initializing a variable,
	// e.g. for var wanyKamo string = "fomby foana" in this case.
	wanyKamo := "fomby foana"
	fmt.Println("fomby ve ? " + wanyKamo)

	const s string = "constant"

	// loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	// i = i + 1
	// 	i++
	// }
	fmt.Println("----------------")
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	fmt.Println("----------------")
	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("----------------")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// if else
	fmt.Println("----------------")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}
	fmt.Println("----------------")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}
	fmt.Println("----------------")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// switch case
	fmt.Println("----------------")
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("----------------")
	// comparer les types au lieu des valeurs
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	fmt.Println("----------------")
	// array
	firstArray := []int{1, 2, 5}
	fmt.Println("emp:", firstArray)

	fmt.Println("----------------")
	firstArray = append(firstArray, 3)
	fmt.Println("set:", firstArray)

	// slice
	dynamicArray := make([]int, 1)
	dynamicArray[0] = 40
	dynamicArray = append(dynamicArray, 5)
	dynamicArray = append(dynamicArray, 45, 55)
	fmt.Println("dynamicArray ", dynamicArray)
	fmt.Println("lenght", len(dynamicArray))

	newDynamicArray := make([]int, len(dynamicArray))
	copy(newDynamicArray, dynamicArray)
	fmt.Println("----------------")
	fmt.Println("newDynamicArray ", newDynamicArray)

	// map (cle, valeur)
	fmt.Println("----------------")
	myMap := make(map[string]int)
	myMap["k1"] = 7
	myMap["k2"] = 13
	fmt.Println("myMap ", myMap)
	fmt.Println("myMap lenght ", len(myMap))
	// access value of Map
	v1 := myMap["k1"]
	fmt.Println("v1: ", v1)

	delete(myMap, "k2")
	fmt.Println("myMap:", myMap)

	// The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". Here we didn’t need the value itself, so we ignored it with the blank identifier _.
	_, prs := myMap["k1"]
	fmt.Println("prs:", prs)

	newMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", newMap)

}
