package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	//1. DECLARE
	fmt.Println("1. DECLARE")

	var stringVal string // ""
	stringVal = "hello"  // "hello"
	fmt.Printf("%s \n", stringVal)
	//using []byte
	var byteString []byte = []byte{'h', 'e', 'l', 'l', 'o'} // [ 104, 101, 108, 108, 111 ]

	//converting []byte to string
	stringConv1 := string(byteString[:]) // "hello"
	fmt.Printf("string from byte using string() : %s \n", stringConv1)
	stringConv2 := bytes.NewBuffer(byteString).String() // "hello"
	fmt.Printf("string from byte using buffer : %s \n", stringConv2)

	//2. Behind The Scene
	fmt.Println("2. Behind The Scene")
	fmt.Println("string is using []byte in underlying data")

	var byteStr []byte = []byte{'h', 'e', 'l', 'l', 'o'} // [ 104, 101, 108, 108, 111 ]
	fmt.Printf("Byte str : %v \n", byteStr)
	var str string = "hello"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str caracter : %v \n", str[i]) // 104 101 108 108 111
	}

	//3. Operation
	fmt.Println("3. Operation")
	fmt.Println("1. Concat")
	fmt.Println("1. Operator ( + , += )")
	//concat
	fmt.Printf("Concat hello + world : %s \n", "hello "+" world") // "hello world"

	//append
	str1 := "hello"
	str1 += " world"
	fmt.Printf("Append hello += world : %s \n", str1) // "hello world"

	fmt.Println("2. bytes.Buffer")
	var strVal string = "hello"

	var buffer *bytes.Buffer = bytes.NewBuffer([]byte(strVal))
	buffer.WriteString(" World")
	var newString string = buffer.String() // "hello world"
	fmt.Printf("String concation using buffer : %s \n", newString)

	fmt.Println("3. fmt.Sprintf")
	str2 := "hello"
	resultFormat := fmt.Sprintf("%s world", str2) // "hello world"
	fmt.Printf("String concation using fmt.Sprintf : %s \n\n", resultFormat)

	fmt.Println("2. Working with array")
	str3 := "hello"
	var arrString []string = strings.Split(str3, "") // ["h","e","l","l","o"]
	fmt.Printf("Split string hello : %v \n", arrString)

	strJoin := strings.Join(arrString, ",") // "hello world"
	fmt.Printf("Join ['h','e','l','l','o'] with ',' : %s \n\n", strJoin)

	fmt.Println("3. Replace")
	str4 := "hello world"
	//n > 0 limit
	newStrReplace1 := strings.Replace(str4, "l", "i", 2) // "heii0 world"
	fmt.Printf("hello world replace l with i 2 times : %s \n", newStrReplace1)
	//n < 0 no limit
	newStrReplace2 := strings.Replace(str4, "l", "i", -1) // "heii0 worid"
	fmt.Printf("hello world replace l with i all : %s \n\n", newStrReplace2)

	fmt.Println("4. Contains")
	str5 := "hello world"
	//contains sub string
	isContainWord := strings.Contains(str5, "hello") //true
	fmt.Printf("hello world contains hello : %t \n", isContainWord)
	//contains a char
	isContainChar := strings.ContainsAny(str5, "ce") //true
	fmt.Printf("hello world contains c or e : %t \n\n", isContainChar)

	fmt.Println("5. Find")
	indexWorld := strings.Index("hello world", "world") //6
	fmt.Printf("index first character world in hello world : %d \n", indexWorld)
	//2. find char
	indexChar := strings.IndexAny("hello world", "le") //1
	fmt.Printf("index first character l or e in hello world : %d \n \n", indexChar)

	fmt.Println("6. Remove White Spaces")
	fmt.Printf("remove white space h e l l o : %s \n\n", strings.Replace("h e l l o", " ", "", -1)) //hello

	fmt.Println("4. Convertion")
	fmt.Println("1. to Numeric")
	i, _ := strconv.Atoi("42")
	fmt.Printf("convert string(42) to int : %d \n", i)
	i64, _ := strconv.ParseInt("97", 10, 64) //97
	fmt.Printf("convert string(97) to int : %d \n", i64)
	_, err := strconv.ParseInt("128", 10, 8)
	fmt.Printf("Error : %s \n", err)
	//will throw error out of range int8 max number 127
	//float
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("convert string(3.1415) to float : %f \n", f)
	//uint
	u, _ := strconv.ParseUint("0", 10, 64) //0
	fmt.Printf("convert string(0) to uint : %d \n", u)
	_, err = strconv.ParseUint("-1", 10, 64) //throw error because uint cannot be negatif value
	fmt.Printf("Erro convert string(-1) to uint : %s \n", err)

	fmt.Println("2. Boolean")
	//truly value
	b, _ := strconv.ParseBool("true")
	fmt.Printf("convert string(true) to boolean : %t \n", b)
	b, _ = strconv.ParseBool("1")
	fmt.Printf("convert string(1) to boolean : %t \n", b)
	//falsy value
	b, _ = strconv.ParseBool("false")
	fmt.Printf("convert string(false) to boolean : %t \n", b)
	b, _ = strconv.ParseBool("0")
	fmt.Printf("convert string(0) to boolean : %t \n", b)

	_, err = strconv.ParseBool("") // b will false but err will occur
	fmt.Printf("Error convert empty tring to boolean : %s \n\n", err)

	fmt.Println("3. to Date")
	layout := "Mon, 01/02/06, 03:04PM"
	value := "Sat, 08/18/2020, 18:00PM"
	t, _ := time.Parse(layout, value)
	fmt.Printf("value date : %v \n", t)
	fmt.Printf("Type date : %T \n", t)
}
