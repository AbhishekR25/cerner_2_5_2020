//Cerner_2^5_202
//Program to find Square of a given number
package main
 
import "fmt"
 
var square int   // Variable declaration outside the main function
 
func main(){
    var input int 
    var result int 
    fmt.Println("Please enter a number:")
    fmt.Scanf("%d", &input)
    result = input * input
    fmt.Println("Squre of " , input , " is: " ,result)
}