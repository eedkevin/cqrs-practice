package util

import (
	"bytes"
	"encoding/gob"
)

func Encode(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (interface{}, error) {
	var data interface{}
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
