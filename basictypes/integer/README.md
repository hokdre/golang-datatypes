# Numeric Value

## dicover convertion
1. ### convert bigger value to smaller
   ```golang
    /*
    if value is bigger, value starts counting from minimum value that's -128 for int8 and the rest value will acumulate increment (+) 
   */
   var i32 int16 = 129
   //int8 range : -128 to 127
   var i8 int8 = int8(i32) //-128 + 1

   //IMPORTANT : Make sure declare variabel first
   var i8 int8 := int8(129) //error constanta overflow

   ```
2. ### convert int(can negatif / positif) to uint(positif cannot negatif)
   ```golang
    /*
    if value is smaller, value starts counting from max value that's 255 for uint8 and the rest value will acumulate decrement (-)
   */
   var i int = -2
   //uint8 range : 0 to 255
   var ui uint8 = uint8(i) //254

   //IMPORTANT : Make sure declare variabel first
   var ui8 uint8 := uint8(-2) //error constanta overflow
   ```