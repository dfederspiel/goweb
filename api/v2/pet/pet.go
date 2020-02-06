package pet

type Pet struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Legs  bool   `json:"legs"`
	Color string `json:"color"`
}
