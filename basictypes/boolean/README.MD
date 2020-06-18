# Boolean

1. ## Declare
   ```go
    var boolVal bool // false
   ```
2. ## Operation
   1. And ( && )
        ```go
        (true && true ) // true
        (true && false ) // false
        (false && false) // false
        ```
   2. Or ( || )
         ```go
        (true && true ) // true
        (true && false ) // true
        (false && false) // false
        ```
   3. Not ( ! )
       ```go
        !(true) //false
        !(false) //true
        ``` 
3. ## Convertion
   1. To String
      1. strconv.FormatBool(bool)
            ```go
            strconv.FormatBool(true) // "true"
            ```  
      2. fmt.Sprintf("%t or %v", bool)
            ```go
            fmt.Sprintf("%t", true) // "true"

            fmt.Sprintf("%v", true) // "true"            
            ``` 
    2. To Integer
       ```go
       var integerVal int // initial 0
       if true {
           intergerVal = 1
       }
       ``` 
4. ## Printing
   ```golang
   fmt.Printf("%t", true) //"true"
   ```