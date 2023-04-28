package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tagirhamitov/services_practice_1/testing"
	"github.com/tagirhamitov/services_practice_1/types"
)

func main() {
	formatStr := os.Args[1]
	format, err := types.ParseFormat(formatStr)
	if err != nil {
		log.Fatal(err)
	}

	serializedSizeFunction := testing.GetSerializedSizeFunction(format)
	serializationFunction := testing.GetSerializationFunction(format)
	deserializationFunction := testing.GetDeserializationFunction(format)

	http.HandleFunc("/get_result", func(w http.ResponseWriter, r *http.Request) {
		object := types.NewStruct()

		serializedSize, err := serializedSizeFunction(object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, fmt.Errorf("failed to get serialized size: %w", err))
			return
		}

		serializationTime, err := serializationFunction(object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, fmt.Errorf("failed to test serialization: %w", err))
			return
		}

		deserializationTime, err := deserializationFunction(object)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, fmt.Errorf("failed to test deserialization: %w", err))
			return
		}

		fmt.Fprintf(
			w,
			"%v - %v - %v - %v\n",
			formatStr,
			serializedSize,
			serializationTime,
			deserializationTime,
		)
	})

	err = http.ListenAndServe("0.0.0.0:2000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
