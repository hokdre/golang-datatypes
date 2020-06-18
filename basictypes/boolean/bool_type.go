package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1. Declare
	fmt.Println("1. DECLARE")
	var boolVal bool // false
	fmt.Printf("zero val bool varibel : %t \n \n", boolVal)

	//2. Operation
	fmt.Println("2. OPERATION")
	fmt.Println("2.1 Operation AND (&&)")
	fmt.Printf("true && true : %t \n", (true && true)) // true

	fmt.Printf("true && false : %t \n", (true && false)) // false

	fmt.Printf("false && false : %t \n", (false && false)) // false

	fmt.Println("2.2 Operation OR ( || )")
	fmt.Printf("true || true : %t \n", (true || true)) // true

	fmt.Printf("true || false : %t \n", (true || false)) // true

	fmt.Printf("false || false : %t \n", (false || false)) // false

	fmt.Println("2.3 Operation Not ( ! )")

	fmt.Printf("!(true) : %t \n", !(true)) // false

	fmt.Printf("!(false) : %t \n \n", !(false)) // true

	//3. Convertion
	fmt.Println("3. CONVERTION")

	fmt.Printf("Convert using strconv.FormatBool : %s \n", strconv.FormatBool(true)) // "true"

	fmt.Printf("Convert using fmt.Sprintf : %s \n", fmt.Sprintf("%t", true)) // "true"

	fmt.Printf("Convert using fmt.Sprintf : %s \n \n", fmt.Sprintf("%v", true)) // "true"

	//4. Printing
	fmt.Println("4. PRINTING")
	fmt.Printf("use t code for printing : %t \n", true) //"true"
}
