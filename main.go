package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"unicode"
)

func main() {
	log.SetFlags(0)
	var data interface{}
	br := bufio.NewReader(os.Stdin)

loop:
	for {
		r, _, err := br.ReadRune()
		if err != nil {
			log.Fatalf("Failed to read: %s", err)
		}

		switch {
		case r == '[': // json array
			data = make([]map[string]interface{}, 0)
			br.UnreadRune()
			break loop
		case r == '{': // json object
			data = make(map[string]interface{})
			br.UnreadRune()
			break loop
		case unicode.IsSpace(r): // ignore space
			continue
		default:
			log.Fatal("Invalid json format")
		}
	}

	dec := json.NewDecoder(br)
	err := dec.Decode(&data)
	if err != nil {
		log.Fatalf("Invalid json format: %s", err)
	}

	b, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal: %s", err)
	}

	_, err = io.Copy(os.Stdout, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Failed to write: %s", err)
	}
}
