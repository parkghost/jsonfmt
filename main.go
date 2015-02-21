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
var input = flag.String("i", "stdin", "The input file")
var output = flag.String("o", "stdout", "The output file")

func main() {
	flag.Parse()
	log.SetFlags(0)

	r, err := file(*input, false)
	if err != nil {
		log.Fatalf("Unable to Open file: %s", err)
	}
	w, err := file(*output, true)
	if err != nil {
		log.Fatalf("Unable to Open file: %s", err)
	}

	var data interface{}
	dec := json.NewDecoder(r)
	err = dec.Decode(&data)
	if err != nil {
		log.Fatalf("Invalid json format: %s", err)
	}

	var b []byte
	if *compact {
		b, err = safeEncode(json.Marshal(data))
	} else {
		b, err = safeEncode(json.MarshalIndent(data, "", "    "))
	}
	if err != nil {
		log.Fatalf("Failed to marshal: %s", err)
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	if err != nil {
		log.Fatalf("Failed to write: %s", err)
	}
}

func file(name string, create bool) (*os.File, error) {
	switch name {
	case "stdin":
		return os.Stdin, nil
	case "stdout":
		return os.Stdout, nil
	default:
		if create {
			return os.Create(name)
		}
		return os.Open(name)
	}
}

func safeEncode(b []byte, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}

	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	return b, nil
}
