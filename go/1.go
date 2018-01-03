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
	"reflect" // https://golang.org/pkg/reflect/#TypeOf
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
	// Result: I'm killing your brain like a poisonous mushroom
	fmt.Println("\nThis is decodedHex in decimal: ", decodedHex)
	// Result: [73 39 109 32 107 105 108 108 105 110 103 32 121 111 117 114 32 98 114 97 105 110 32 108 105 107 101 32 97 32 112 111 105 115 111 110 111 117 115 32 109 117 115 104 114 111 111 109]
	// fmt.Println("\nThis is the decodedHex typeOf here: ", reflect.TypeOf(decodedHex))
	// Result: []uint8

	encodedBase64 := base64.StdEncoding.EncodeToString([]byte(decodedHex))
	fmt.Println("\nThis is the encodedBase64 of decodedHex: ", encodedBase64)
	// Result: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t
	// fmt.Println("\nThis is the encodedBase64 typeOf here: ", reflect.TypeOf(encodedBase64))
	// Result: string

	// I'm not sure if this entire process was necessary, though I'm unsure why it wouldn't convert back cleanly.
	byteEncodedBase64 := []byte(encodedBase64)
	// fmt.Println("\nSee if byteEncodedBase64 works? Here: ", byteEncodedBase64)
	// Result: [83 83 100 116 73 71 116 112 98 71 120 112 98 109 99 103 101 87 57 49 99 105 66 105 99 109 70 112 98 105 66 115 97 87 116 108 73 71 69 103 99 71 57 112 99 50 57 117 98 51 86 122 73 71 49 49 99 50 104 121 98 50 57 116]
	// fmt.Println("\nCheck out the typeOf for byteEncodedBase64: ", reflect.TypeOf(byteEncodedBase64))
	// Result: []uint8

	decodedBase64, err := base64.StdEncoding.DecodeString(checkAgainstMe)
	if err != nil {
		fmt.Println("Decode error: ", err)
		return
	}
	fmt.Printf("\nThis is checkAgainstMe base64 decoded in symbol: %s", decodedBase64)
	// Result: I'm killing your brain like a poisonous mushroom
	fmt.Println("\nThis is checkAgainstMe base64 decoded in decimal: ", decodedBase64)
	// Result: [73 39 109 32 107 105 108 108 105 110 103 32 121 111 117 114 32 98 114 97 105 110 32 108 105 107 101 32 97 32 112 111 105 115 111 110 111 117 115 32 109 117 115 104 114 111 111 109]
	// fmt.Println("\nThis is the decodedBase64 typeOf here: ", reflect.TypeOf(decodedBase64))
	// Result: []uint8

	// I need to figure out why this is broken!
	// if byteEncodedBase64 == decodedBase64 {
	// 	fmt.Println("\nSuccess! Your variable encodedBase64 is equal to decodedBase64.")
	// } else { // Apparently else if/else has to be on the same line as the closing brace? Weird.
	// 	fmt.Println("\nUh oh something has gone wrong, try this again?")
	// }
	// Currenty result of this conditional is:
	// ./1.go:71: invalid operation: byteEncodedBase64 == decodedBase64 (slice can only be compared to nil)

	// Checking if two slices are equal is a package apparently?
	fmt.Println("\nIf this returns 'false' huzzah yay success (until I figure out why byteEncodedBase64 is wrong)!", reflect.DeepEqual(byteEncodedBase64, decodedBase64))
	fmt.Println("\nWait I think I only need to see if decodedHex and decodedBase64 are equal... if this is 'true' whoo success!", reflect.DeepEqual(decodedHex, decodedBase64))
}
