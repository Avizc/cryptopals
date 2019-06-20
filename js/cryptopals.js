/*
    Convert hex to base64

    This: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

    Should produce: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

    Note: Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty
*/

const hexToByte= (hex) => {

}

const byteToHex = (byte) => {

}

const byteToBase64 = (byte) => {
    const key = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'
    let bytes = new Uint8Array(byte)
    let newBase64 = ''
    let currentChar = 0
    for (let i=0; i<bytes.length; i++) {   // Go over three 8-bit bytes to encode four base64 6-bit chars
      if (i%3===0) { // First Byte
        currentChar = (bytes[i] >> 2)      // First 6-bits for first base64 char
        newBase64 += key[currentChar]      // Add the first base64 char to the string
        currentChar = (bytes[i] << 4) & 63 // Erase first 6-bits, add first 2 bits for second base64 char
      }
      if (i%3===1) { // Second Byte
        currentChar += (bytes[i] >> 4)     // Concat first 4-bits from second byte for second base64 char
        newBase64 += key[currentChar]      // Add the second base64 char to the string
        currentChar = (bytes[i] << 2) & 63 // Add two zeros, add 4-bits from second half of second byte
      }
      if (i%3===2) { // Third Byte
        currentChar += (bytes[i] >> 6)     // Concat first 2-bits of third byte for the third base64 char
        newBase64 += key[currentChar]      // Add the third base64 char to the string
        currentChar = bytes[i] & 63        // Add last 6-bits from third byte for the fourth base64 char
        newBase64 += key[currentChar]      // Add the fourth base64 char to the string
      }
    }
  if (bytes.length%3===1) { // Pad for two missing bytes
    newBase64 += `${key[currentChar]}==`
  }
  if (bytes.length%3===2) { // Pad one missing byte
    newBase64 += `${key[currentChar]}=`
  }
  return newBase64
}
// byteToBase64([104,101,108,108,111]) 
// => 'aGVsbG8='

const base64ToByte = (base64) => {
    const base64Table = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'

}