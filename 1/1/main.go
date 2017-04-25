package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

// Decodes an hex encoded string and returns an base64 encoded string 
func hexToBase64(str string) {
	decoded, _ := hex.DecodeString(str)
	encoded := base64.StdEncoding.EncodeToString(decoded)
	if decoded == nil {
		fmt.Println("decode error")
		return
	}
	fmt.Println(string(encoded))
}

// Challenge 1 main function
func main() {
	hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
}
