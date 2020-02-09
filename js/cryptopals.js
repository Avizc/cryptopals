/*
    Convert hex to base64

    This: 49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

    Should produce: SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

    Note: Always operate on raw bytes, never on encoded strings. Only use hex and base64 for pretty
*/

const hexToByte = (hex) => {
  const key = '0123456789abcdef'
  let newBytes = []
  let currentChar = 0
  let currentByte = 0
  for (let i=0; i<hex.length; i++) {   // Go over two 4-bit hex chars to convert into one 8-bit byte
    currentChar = key.indexOf(hex[i])
    if (i%2===0) { // First hex char
      currentByte = (currentChar << 4) // Get 4-bits from first hex char
    }
    if (i%2===1) { // Second hex char
      currentByte += (currentChar)     // Concat 4-bits from second hex char
      newBytes.push(currentByte)       // Add byte
    }
  }
  return new Uint8Array(newBytes)
}
// hexToByte('68656c6c6f')
// => [104, 101, 108, 108, 111]

const byteToHex = (byte) => {
  const key = '0123456789abcdef'
  let bytes = new Uint8Array(byte)
  let newHex = ''
  let currentChar = 0
  for (let i=0; i<bytes.length; i++) { // Go over each 8-bit bytes
    currentChar = (bytes[i] >> 4)      // First 4-bits for first hex char
    newHex += key[currentChar]         // Add first hex char to string
    currentChar = (bytes[i] & 15)      // Erase first 4-bits, get last 4-bits for second hex char
    newHex += key[currentChar]         // Add second hex char to string
  }
  return newHex
}
// byteToHex([104,101,108,108,111])
// => '68656c6c6f'

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
  const key = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'
  let newBytes = []
  let currentChar = 0
  let currentByte = 0
  for (let i=0; i<base64.length; i++) {     // Go over four 6-bit base64 chars to decode into three 8-bit bytes
    currentChar = key.indexOf(base64[i])
    if (i%4===0) { // First base64 char
      currentByte = (currentChar << 2)      // Get the 6-bits from first base64 char
    }
    if (i%4===1) { // Second base64 char
      currentByte += (currentChar >> 6)     // Concat the first 2-bits from second base64 char
      newBytes.push(currentByte)            // Push the first byte
      currentByte = (currentChar & 15) << 4 // Erase bits before last 4-bits, add last 4-bits for second byte
    }
    if (i%4===2) { // Third base64 char
      currentByte += (currentChar >> 2)     // Concat first 4-bits from third base64 char for second byte
      newBytes.push(currentByte)            // Push second byte
      currentByte = (currentChar & 3) << 6  // Erase bits before last 2-bits, add last 2-bits for third byte
    }
    if (i%4===3) { // Fourth base64 char
      currentByte += (currentChar)          // Concat fourth base64 char for third byte
      newBytes.push(currentByte)            // Push third byte
    }
  }
  if (newBytes[newBytes.length-1] === -1) { // Remove single padding from the end (=)
    newBytes.pop()
  }
  if (newBytes[newBytes.length-2] === -1) { // Remove double padding from the end (==)
    newBytes.pop()
    newBytes.pop()
  }
  return new Uint8Array(newBytes)
}
// base64ToByte('aGVsbG8=')
// hell   => 104 101 108 108         => aGVsbA==
// hello  => 104 101 108 108 111     => aGVsbG8=
// hellos => 104 101 108 108 111 115 => aGVsbG9z

const hexToBase64 = (hex) => {
  return byteToBase64(hexToByte(hex))
}
// hexToBase64('49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d')
// => 'SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t'

const base64ToHex = (base64) => {
  return byteToHex(base64ToByte(base64))
}
// base64ToHex('SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t')
// => '49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'
