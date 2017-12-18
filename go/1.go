package main
// Yay! Clean import of packages makes me happy.
// Packages: https://golang.org/pkg/
import(
    "fmt"
    // "encoding/hex"
    // "encoding/base64"
)
// This is my first function in Go huzzah!
func helloWorld(x string,y string,z string)string{
    return "I really love "+x+", "+y+", and "+z+"!"
}
// ------------------------------------- 
// Challenge 1 - Convert hex to base64
// The following string:
// "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d75"
// Should convert to:
// "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
// Cryptopals Rule
// Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
// -------------------------------------
// Put any string you want here!
const HexString="49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d75"
// Print out all the things here.
func main(){
    fmt.Println("I love rabbits, cheesecake, and cute things!")
    // fmt.Println(helloWorld("rabbits", "cheesecake", "cute things"))
}