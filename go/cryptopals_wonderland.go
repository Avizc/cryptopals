package main

import (
	alice "github.com/llzes/cryptopals/go/rabbit_hole_01"
)

func main(){
	/* Challenge 1 - Convert hex to base64 */
	alice.Test()
	const hexCandy string = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	const baseCake string = "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	alice.HexToByte(hexCandy)
}