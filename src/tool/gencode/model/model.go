package model

type CustomerModel struct {
	Name      string   `json:"name"`
	ValueType string   `json:"type"`
	Fields    []*field `json:"fields"`
}

type field struct {
	Name      string `json:"name"`
	TypeValue string `json:"type"`
	Desc      string `json:"desc"`
}
