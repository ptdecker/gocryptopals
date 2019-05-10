package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

// Convert a slice of bytes containing hex encoded text into a base64 encoded string
func asciiHex2Base64(line []byte) string {

	// Decode hex encoded text into a byte slice trimming any trailing delimiter
	src := bytes.TrimRight(line, string([]byte{'\n'}))
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	// Output array of bytes in base64 encoding to console
	 return base64.StdEncoding.EncodeToString(dst[:n])

}

func main() {

	// Read in a line of text in from console
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	fmt.Println(asciiHex2Base64(line))
}
