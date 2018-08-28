package common

import (
	"bytes"
	"encoding/gob"
)

func GetGobFromInterface(key interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(key)
	return buf.Bytes(), err
}

func GetInterfaceFromGob(byt []byte, obj interface{}) error {
	buf := bytes.NewBuffer(byt)
	enc := gob.NewDecoder(buf)
	err := enc.Decode(obj)
	return err
}
