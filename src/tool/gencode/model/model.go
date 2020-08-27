package model

type CustomerModel struct {
	Name      string  `json:"name"`
	ValueType string  `json:"value_type"`
	Fields    []field `json:"fields"`
}

type field struct {
	Name      string `json:"name"`
	TypeValue string `json:"type"`
	Desc      string `json:"desc"`
	Required  bool   `json:"required"`
}
