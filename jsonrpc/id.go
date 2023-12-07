package jsonrpc

import "encoding/json"

type Id struct {
	strValue string
	intValue int32
	isNumber bool
}

func NewStringId(id string) *Id {
	return &Id{
		strValue: id,
		intValue: 0,
		isNumber: false,
	}
}

func NewIntId(id int32) *Id {
	return &Id{
		strValue: "",
		intValue: id,
		isNumber: true,
	}
}

func (id *Id) MarshalJSON() ([]byte, error) {
	if id.isNumber {
		return json.Marshal(id.intValue)
	}
	return json.Marshal(id.strValue)
}
