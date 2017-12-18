package main
import(
    "fmt"
    // "encoding/hex"
    // "encoding/base64"
)
func helloWorld(x string,y string,z string)string{
    return "I really love "+x+", "+y+", and "+z+"!"
}
func main(){
    fmt.Println("I love rabbits, cheesecake, and cute things!")
    fmt.Println(helloWorld("rabbits", "cheesecake", "cute things"))
}