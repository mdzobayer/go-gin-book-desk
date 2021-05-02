package model

type Filter struct {
	FieldName string      `json:"FieldName"`
	Value     interface{} `json:"Value"`
}
