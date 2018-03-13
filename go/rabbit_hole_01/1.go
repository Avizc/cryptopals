/*
	Challenge 1 - Convert hex to base64
		The following string:
			"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
		Should convert to:
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
		Cryptopals Rule
			Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty-printing.
*/
package rabbit_hole_01

/*
	Re-doing this challenge, avoid imports as much as possible especially anything from Reflect! ✨
	Previous first attempt at this challenge can now be found in rabbit_hole_01_old/0.go
*/

import (
	"fmt"
)

func Test() {
	// fmt.Println("I love rabbits, cheesecake, and cute things!")
	var firstChar uint8 = 0x42
	fmt.Println("Following should be 01000010")
	fmt.Println("%08b\n", firstChar)
}

func HexToByte(hexCookies string) []byte {
	newCookies := []byte(hexCookies)
	fmt.Println("This is the original hex string: ", hexCookies)
	fmt.Println("String should now be byte encoded as ASCII: ", newCookies)
	return newCookies
}

// func ByteToHex(byteCookies []byte) string {

// }

// func BaseSixtyFourToByte(baseMacarons string) []byte {

// }

// func ByteToBaseSixtyFour(byteMacarons []byte) string {

// }

// func HexToBaseSixtyFour(hexCake string) string {

// }
