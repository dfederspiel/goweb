package models

type Animal struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Legs bool   `json:"legs"`
}
