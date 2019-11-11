package metadata

import (
	"bytes"
	"encoding/json"
)

type Enum interface {
	Name() string
	Ordinal() int
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(b []byte) error
}

func marshalJSON(enum Enum) ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(enum.Name())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func unmarshalJSON(b []byte) (string, error) {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return "", err
	}

	return j, nil
}

func toString(enum Enum, enumValues []string) string {

	ordinal := enum.Ordinal()
	length := len(enumValues)
	if ordinal < 1 || ordinal > length {
		return "Unknown"
	}

	return enumValues[ordinal-1]
}
