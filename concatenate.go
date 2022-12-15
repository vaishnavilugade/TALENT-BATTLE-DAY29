package main
import "fmt"

func main(){
    var a string
    fmt.Print("Enter first string:")
    fmt.Scan(&a)
    var b string
    fmt.Print("Enter second string:")
    fmt.Scan(&b)
    c:=a+" "+b
    fmt.Print(c)
}