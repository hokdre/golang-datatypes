# String

1. ## Declare
   ```go
    var stringVal string // ""
    stringVal = "hello" // "hello"

    //using []byte
    var byteString []byte = []byte{'h','e','l','l','o'} // [ 104, 101, 108, 108, 111 ]
    
   //converting []byte to string
   stringConv1 := string(byteString[:]) // "hello"

   stringConv2 := bytes.NewBuffer(byteString).String() // "hello"   
   ```
2. ## Behind The Scene
   ```golang
   //string is using []byte in underlying data

   var byteString []byte = []byte{'h','e','l','l','o'} // [ 104, 101, 108, 108, 111 ]

   var stringVal string = "hello"
   for i := 0; i < len(stringVal); i++ {
      fmt.Print(stringVal[i]) // 104 101 108 108 111
   }

   //[]byte is mutable 
   byteString[0] = 108 
   //string is immutable
   stringVal[0] = "l" //not change
   ```
3. ## Operation
   1. ### Concat
      1. #### operator ( + , += )
         ```golang
         //concat
         fmt.Print("hello" + " world") // "hello world"

         //append
         str1 := "hello"
         str1 += " world"
         fmt.Print(str1) // "hello world"
         ``` 
      2. #### bytes.Buffer
         ```golang
         var stringVal string = "hello"
         
         var buffer *bytes.Buffer
         buffer = bytes.NewBuffer( []byte(stringVal) )
         buffer.WriteString(" world")

         var newString string = buffer.String() // "hello world"
         ``` 
      3. #### fmt.Sprintf
         ```golang
         str1 := "hello"
         fmt.Sprintf("%s world", str1) // "hello world"
         ``` 
   2. ### working with array
      1. #### string.Split(str string, separator string) []string
         ```golang
         str1 := "hello"
         var arrString []string = strings.Split(str1, "") // ["h","e","l","l","o"]
         ```  
      2. ### string.Join(arr []string, separator string ) string
         ```golang
         str1 := "helloworld"
         var arrString []string = strings.Split(str1, "") // ["h","e","l","l","o","w","o","r","l","d"]

         newStr := strings.Join(arrString, ",") // "h,e.l,l,o,w,o,r,l,d"
         ``` 
   3. ### Replace
      1. #### Replace(str, old, new string, n int ) string
         ```golang
         str := "hello world"

         //n > 0 limit
         newStr1 := strings.Replace(str, "l","i", 2) // "heii0 world"

         //n < 0 no limit
         newStr2 := strings.Replace(str, "l","i", -1) // "heii0 worid"
         ``` 
      2. ### ReplaceAll(str,old, new string ) string
         Behaviour same like replace no limit  
   4. ### Contains
      ```golang
      str1 := "hello world"

      //contains sub string
      strings.Contains(str1, "hello") //true

      //contains a char
      strings.ContainsAny(str1, "ce") //true
      ```
   5. ### Find
      ```golang
      //1. find word
      strings.Index("hello world", "world") //6

      //2. find char
      strings.IndexAny("hello world", "le") //1
      ``` 
   6. ### Remove White Spaces
      ```golang
      strings.Replace("h e l l o"," ","",-1) // "hello"
      ``` 
   
4. ## Convertion
   1. ### to Numeric
      ```golang
      //int
      i, err := strconv.Atoi("42")
      i, err := strconv.ParseInt("97",10,64) //97
      i, err := strconv.ParseInt("128",10, 8) //will throw error out of range int8 max number 127

      //float64
      f, err := strconv.ParseFloat("3.1415",64)
      
      //uint
      u, err  := strconv.ParseUint("0",10,64) //0
      u, err  := strconv.ParseUint("-1",10,64) //throw error because uint cannot be negatif value
      ``` 
   2. ### to Bool
      ```golang
      //truly value
      b, err := strconv.ParseBool("true")
      b, err := strconv.ParseBool("1")

      //falsy value
      b, err := strconv.ParseBool("false")
      b, err := strconv.ParseBool("0")
      ``` 
      > important : Empty String Will Throw Error 
      ```golang
      b, err := strconv.ParseBool("") // b will false but err will occur
      b, err := strconv.ParseBool("a") // b will false but err will occur
      ```
   3. ### to Date
      ```golang
      layout := "Mon, 01/02/06, 03:04PM"
      value := "Sat, 08/18/2020, 18:00PM"
      t, _ := time.Parse(layout, value) //type time.Time
      ``` 
   4. ### reverse
      ```golang
      s := strconv.FormatBool(true) // "true"
      s := strconv.FormatInt(1) // "1"
      s := strconv.FormatFloat(3.1415, 'E', -1, 64) // "3.1415E+00"
      s := strconv.FormatUint(97, 10) //"97"
      s := strconv.FormatInt(97, 10) //"97"
      ``` 
      > important : Don't use string( int ) to format. use strconv.FormatInt(int, 10) instead.
      ```golang
      //unexpected result
      s := string(97) // "a"
      ```
5. ## Printing
   ```golang
   fmt.Printf("%s", "hello") // "hello"
   ```