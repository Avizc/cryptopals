//
// Challenge 2 - Fixed XOR
// Write a function that takes two equal-length buffers and produces their XOR combination.
// If your function works properly, then when you feed it the string:
// "1c0111001f010100061a024b53535009181c"
// ...after hex decoding, and when XOR'd against:
// "686974207468652062756c6c277320657965"
// ...should produce:
// "746865206b696420646f6e277420706c6179"
//
package rabbit_hole_02
import(
	"fmt"
	// "log"
	// "encoding/hex"
)
// Test string alice put any string here!
const alice="1c0111001f010100061a024b53535009181c"
// Test string whiteRabbit put any string of the same length as alice here!
const whiteRabbit="686974207468652062756c6c277320657965"
// Test string wonderland to check if alice and whiteRabbit have been XOR'd properly
const wonderland="746865206b696420646f6e277420706c6179"
// Print out all the things here!
func Test2(){
	fmt.Println("This should print alice: ",alice)
	fmt.Println("This should print whiteRabbit: ",whiteRabbit)
}