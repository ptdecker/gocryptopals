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

func main() {

	// Read in a line of text in from console
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	// Decode hex encoded text into a byte slice trimming any trailing delimiter
	src := bytes.TrimRight(line, string([]byte{'\n'}))
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	// Output array of bytes in base64 encoding to console
	fmt.Println(base64.StdEncoding.EncodeToString(dst[:n]))

}
