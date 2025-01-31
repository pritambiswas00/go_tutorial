package main // Entry point of the program
import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("sdsdsdsdsdsd")
	//In order to declare a variable you first need to type var followed by name of the variable then type of the variable below example

	var declared_variable int = 33
	//int8, int16, int32, int64 bit depending on your needs
	//float32, float64 bit depending on your needs
	fmt.Println(declared_variable)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	var result float32 = floatNum32 + float32(intNum32) //Before you int to a float addition, you have cast the int to float then you can add
	fmt.Println(result)

	//String

	var someString string = "Hello" //Double quotes are just a single you cant press enter and continue typing string
	var someString2 string = `Hello 
	
		World` // `` This we can use to extend the string
	var testString = "test" + "string" + "test" //Also can do concatinate
	fmt.Println(someString, someString2, testString)
	fmt.Println(utf8.RuneCountInString(testString)) //Here Rune is data type in Go that represents character

	var someRune rune = 'd'
	fmt.Println(someRune) //Give 100 that is ASCII value

	//Boolean

	var someBoolean bool = false
	fmt.Println(someBoolean)

}
