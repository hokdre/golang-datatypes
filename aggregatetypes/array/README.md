# Array

> array is aggregate type. Means safe to copy
> 
> example :
```golang
arr1 := [1]int{1} //[1]
arr2 := arr1 //copy 

arr2[0] = 2 //not change arr1
// arr2 = [2] 
// arr1 = [1]
```

## Declare
```golang
// one dimensional
var arr [3]int // [ 0, 0, 0 ]
len(arr) // 3
cap(arr) // 3
var arr *[3]int = new([3]int) //pointer of array 
arr :=  [2]int {1} //len : 1, cap : 2
arr := [...]int {1, 2, 3} //len : 3, cap : 3

//multi dimensional
var multiArr [1][1]int // [ [0] ]
multiArr = [1][1]int{ [1]int{1} } // [ [1] ]

//not array
var notArr *[]int := new([]int) //pointer of slice
notArr := make([]int, 0, 0) //slice
```

## Operations
> just read and update 
<br>

> **cannot be adding new item**

## Common Miss
1. fun new()
   > fun new() it's return pointer
2. func make()
    > func make([]int, 0, 2) is return slice not array! make just for slices, maps, and channel 
3.  func append()
    > Array Cannot be Appended ! It's capacity is fixed!<br>
    > example : 
    ```golang
    arr1 := [1]int { 1 }
    append(arr1, 2) //error must be slice
    ```