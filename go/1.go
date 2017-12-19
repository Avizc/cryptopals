//
// Challenge 1 - Convert hex to base64
// The following string:
// "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
// Should convert to:
// "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
// Cryptopals Rule
// Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
//
package main

// Yay! Clean import of packages makes me happy.
// Packages: https://golang.org/pkg/
import (
	"encoding/base64" // https://golang.org/pkg/encoding/base64/
	"encoding/hex"    // https://golang.org/pkg/encoding/hex/#DecodeString
	"fmt"
	"log"
	// "reflect" // https://golang.org/pkg/reflect/#TypeOf
)

// This is my first function in Go huzzah!
// func helloWorld(x string, y string, z string) string {
// 	return "I really love " + x + ", " + y + ", and " + z + "!"
// }

// Put any string you want here!
const s = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
const checkAgainstMe = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

// Print out all the things here.
func main() {
	// fmt.Println("I love rabbits, cheesecake, and cute things!")
	// fmt.Println(helloWorld("rabbits", "cheesecake", "cute things"))
	decodedHex, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("This is the decodedHex in symbol: %s", decodedHex)
	fmt.Println("\nThis is decodedHex in decimal: ", decodedHex)
	encodedBase64 := base64.StdEncoding.EncodeToString([]byte(decodedHex))
	fmt.Println("\nThis is the encodedBase64 of decodedHex: ", encodedBase64)
	decodedBase64, err := base64.StdEncoding.DecodeString(checkAgainstMe)
	if err != nil {
		fmt.Println("Decode error: ", err)
		return
	}
	fmt.Printf("\nThis is checkAgainstMe base64 decoded in symbol: %s", decodedBase64)
	fmt.Println("\nThis is checkAgainstMe base64 decoded in decimal: ", decodedBase64)
	// Uh oh I really broke this thing now sleep deprived oops. Check out the typeof in the morning!
	// if encodedBase64 == decodedBase64 {
	// 	fmt.Println("\nSuccess! Your variable encodedBase64 is equal to decodedBase64.")
	// } else { // Apparently else if/else has to be on the same line as the closing brace? Weird.
	// 	fmt.Println("\nUh oh something has gone wrong, try this again?")
	// }
}
