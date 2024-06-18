package main

// Command to run -> go run main.go
// fmt ->  format basic strings, values, inputs, and outputs
// any return can be ignored as _
// no inheritance in GOLANG , NO Superclass
// Focus on whatever on screen is understandable
// 1st Upper letter -> Exported Thing that we are using
// No functions inside funtions

/*Variables*/
/*
import "fmt"
// user := 2993 NOT ALLOWED OUTSIDE , ONLY IN FUNCTION
const TokenLogin string = "sahilcr7"
// In GOLANG , publically declared  if var name start with Capital letter
func main() {
	fmt.Println("Hello from Sahil")

	var username string = "Sahil" //String
	var flag bool = true
	var smallUValue uint8 = 255     // > 255 out bound 2^8 - max
	var smallValue int8 = 23        // > 255 out bound 2^8 - max
	var floatingValue float64 = 2.5 // > 255 out bound 2^8 - max
	// unit -> unsigned , int -> signed , fload -> floating

	var something int // default value or some alise -> NO GARBAGE int->0

	// Implicit type
	var Implicit = "I haven't gave name"

	println(username)                        // print string
	fmt.Printf("My name is %T \n", username) // type
	println(flag)
	fmt.Printf(" %T \n", flag) // type
	println(smallUValue)
	fmt.Printf("%T \n", smallUValue) // type
	println(smallValue)
	fmt.Printf(" %T \n", smallValue) // type
	println(floatingValue)
	fmt.Printf("%T \n", floatingValue) // type
	println(something)
	fmt.Printf("%T \n", something) // type

	println(Implicit)
	fmt.Printf("%T \n", Implicit) // type
	// Implicit = 3  , error -> can not change type once assigned

	// ------------NO VAR STYLE---------------
	user := 30000 // Not outside
	// user := "sahil" error
	println(user)

	println(TokenLogin)
}
*/

/*ok Syntax*/
/*
func main() {
	reader := bufio.NewReader(os.Stdin)
	println("Enter your name :")

	// comma ok || err ok
	// input , err :  try catch
	// _, err

	input, _ := reader.ReadString('\n') // till end of string
	fmt.Println("Thanks you !", input)

	// Conversions in Go
	// String -> package common string syntax
	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(number + 10)
	}
}
*/

/*Time package
func main() {
	// package time in Go
	currentTime := time.Now()
	// This is format to use compulsari , 01 02 2006
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.September, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)

	// Build for window
	// GOOS = 'windows' go build
	// give exe file
	// current os , go build , build main file of directory
}
*/

/*Memory management( Automatic)
new() -> Allocate memory and no INITialize
zeroed storage (Can not put data), get memory added
make() -> Allocate memory and INITialize
non-zeroed storage (Can not put data), get memory added

Garbage collection happen automatically
Runtime Package (NumCpu(),NumCgoCall() .. etc core functions )
GOGC -> variable to set initial garbage collection target percentage
Collection is trigger when ratio of fresh data to live data after previous collection reach this percentage
default GOGC -> 100
GOGC -> off (Entirely off Garbage collection)


*/

/*Pointers in GO
Problem -> copy is created when pass to functions , classes lead to irregularities
This give guarantee that actual value is been passed
func main() {
	println("Pointers")
	var ptr *int // default is nill
	println("Value of ptr : ", ptr)

	number := 3
	fmt.Println(number)
	var pointer *int = &number // &reference to  current memory address
	println(pointer)
	println(*pointer)

	*pointer = *pointer * 2
	println(number) // actual is updated
	// pointers actually work on actual value
}
*/

/* Arrays and slices in GO
import (

	"fmt"
	"sort"

)

func main() {
	// Array (Statuc)
	var arr [3]int // static so have to give size, default 0,""
	arr[0] = 1
	arr[1] = 2
	// arr[2] = 3
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Printf("Type of Array is %T \n", arr)
	var arr2 = [2]int{1, 2}
	fmt.Println(arr2)

	// Slice (Dynamic array)
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("Type of Slices is %T \n", slice)
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)
	slice = append(slice[1:4]) //[start index,end index), default[:] ->[0,n]
	fmt.Println(slice)

	sliceusingmake := make([]int, 4)
	sliceusingmake[0] = 5 //default allocations
	sliceusingmake[1] = 6 //default allocations
	sliceusingmake[2] = 7 //default allocations
	sliceusingmake[3] = 0 //default allocations

	// append reallocate the memory and reaalocate the memory
	// tooks lot of time , memory
	sliceusingmake = append(sliceusingmake, 6, 7, 8)
	fmt.Println(sliceusingmake)

	// Sorting in slices
	sort.Ints(sliceusingmake) // void
	fmt.Println(sliceusingmake)
	fmt.Println(sort.IntsAreSorted(sliceusingmake))

	//Remove slices value based on index
	var chars = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(chars)
	var deleteIndex int = 3
	chars = append(chars[:deleteIndex], chars[deleteIndex+1:]...)
	fmt.Println(chars)
}
*/

/* Maps in GO
// import "fmt"

func main() {
	maps := make(map[string]int)
	maps["Sahil"] = 25
	maps["jyoti"] = 25
	fmt.Println(maps, maps["Sahil"])

	delete(maps, "Sahil") // can be use in slices ass well
	fmt.Println(maps)

	maps["Shile"] = 1
	maps["Sahil"] = 1
	// loops in Go
	for _, value := range maps {
		fmt.Println("Values is ", value)
	}
}
*/

/* Structure in GO NO INHERITANCE
 First letter is capital because these things are exported out to another functions
 type User struct {
	// no var in struct
	Name   string
	Email  string
	Status bool
	Age    int }
  func main() {
	sahil := User{"Sahil", "sahil@gmail.com", true, 21}
	fmt.Println(sahil)
	fmt.Printf("Deep details are %v \n", sahil)  // without keys
	fmt.Printf("Deep details are %+v \n", sahil) // with keys
	// properties can access by .
	fmt.Println(sahil.Name)
 }
*/

/* Control Flow Statements
 import (
	"fmt"
	"math/rand"
	"time")
	func main() {
	// Simple if else
	if 2 > 4 {
		fmt.Println("2 is greater than 4")
	} else if 2 == 2 {
		fmt.Println("2 is equal to 2")
	} else {
		fmt.Println("2 is smaller than 4")
	}

	//SUPER IMPORTANT
	if num := 3; num < 5 {
		fmt.Println("num is less than 5")
	}

	// Switch Case
	id := 5
	switch id {
	case 1:
		fmt.Println("User id is 1")
	case 2:
		fmt.Println("User id is 1")
	case 3:
		fmt.Println("User id is 1")
	case 4:
		fmt.Println("User id is 1")
	default:
		fmt.Println("User id is invalid")
	}

	// Random number
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1 // range is 0 to 5
	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1")
	case 2:
		fmt.Println("Dice number is 2")
		fallthrough // fall below and take below conditions
	case 3:
		fmt.Println("Dice number is 3")
	case 4:
		fmt.Println("Dice number is 4")
	case 5:
		fmt.Println("Dice number is 5")
	case 6:
		fmt.Println("Dice number is 6")
	default:
		fmt.Println("Dice number is invalid")
	}}
*/

/* Loops in GO

 import "fmt"

 func main() {
	days := []int{6, 5, 4, 3, 2, 1}
	fmt.Println(days)

	// no ++d  only d++
	for i := 0; i < len(days); i++ {
		fmt.Print(days[i])
	}
	println()
	// i is index not value , not like js
	for i := range days {
		fmt.Print(days[i])
	}
	println()
	for i, value := range days {
		fmt.Println(i, value)
	}

	i := 1
	for i < 10 {
		if i == 1 {
			i = i + 1
			continue
		}
		if i == 3 {
			goto jumptome
		}
		fmt.Println(i)
		if i == 4 {
			break
		}
		i = i + 1
	}
 jumptome:
	fmt.Println("Someone is Jump on me") }
*/

/* Functions in GO

 import "fmt"

 func print() {
	fmt.Println("hello Sahil")
} // () functions to execute on the go

// Return value called as signatures
func add(a int, b int) int {
	return a + b
}

// pro functions
func addMultiple(values ...int) int {
	total := 0
	// now values are slice
	for i := range values {
		total += values[i]
	}
	return total
}

// return multiple
func returnBoth(a int, b int) (int, int) {
	return a, b
}
func main() {
	fmt.Println("Jay Shree Ram")
	print()
	fmt.Println(add(5, 6))
	fmt.Println(addMultiple(1, 2, 3, 4, 5))
	a, b := returnBoth(1, 2)
	fmt.Println(a, b)
}

*/

/* Methods in GO
import "fmt"

type Result struct {
	Pass  bool
	Marks int
	Cgpa  int
}

func (res Result) getResult() {
	fmt.Println("Is user pass", res.Pass)
}
func (res Result) getMarks() {
	fmt.Println("Is user pass", res.Marks)
}
func (res Result) getCgpa() {
	fmt.Println("Is user pass", res.Cgpa)
}

// mainupulate -> Create copy of res object
// NO Manipulations
func (res Result) updateMarks() {
	res.Marks = 100
}

func main() {
	r := Result{true, 172, 8}
	r.getResult()
	r.getMarks()
	r.getCgpa()
	r.updateMarks() // Not changes anything
	r.getMarks()
}
*/

/*Defer in GO
When function is executing is line by line
When defer , function value and params to call are evaluated as usual
ans saved a now but actual function is not invoked
defered functions is invoked immediately before surrounding functions returns
in reverse order they defered
LAST IN FIRST OUT

import "fmt"

func myDefer() {
	for i := 0; i < 5; i++ {
		fmt.Println("Printing :", i)
	}
}
func main() {
	// without defer -> defered (bottom to top statements)
	// defer statements executed as stack
	defer myDefer()
	defer fmt.Println("World") // defered for last
	fmt.Println("Hello")
	defer fmt.Println("Sahil") // execute first
}
*/

/* Files in GO
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
func main() {
	fmt.Println("Welcome to heaven")
	content := "This is precautions you need to take"
	file, err := os.Create("./file.txt")
	handleError(err)
	len, err := io.WriteString(file, content) // length
	handleError(err)
	fmt.Println("File length :", len)

	fileName := "./file.txt"
	dataByte, err := ioutil.ReadFile(fileName)
	handleError(err)
	fmt.Println(dataByte)
	fmt.Println(string(dataByte))
	file.Close()
}

*/
