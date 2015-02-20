package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
)

var compact = flag.Bool("c", false, "Compact the json data")

func main() {
	flag.Parse()
	log.SetFlags(0)

	var data interface{}
	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&data)
	if err != nil {
		log.Fatalf("Invalid json format: %s", err)
	}

	var b []byte
	if *compact {
		b, err = json.Marshal(data)
	} else {
		b, err = json.MarshalIndent(data, "", "    ")
	}
	if err != nil {
		log.Fatalf("Failed to marshal: %s", err)
	}

	_, err = io.Copy(os.Stdout, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Failed to write: %s", err)
	}
}
