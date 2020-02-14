package pet

type Pet struct {
	ID    string `json:"id"`
	Name  string `json:"name" binding:"required"`
	Age   int    `json:"age"`
	Legs  bool   `json:"legs"`
	Color string `json:"color"`
}

func init() {}
