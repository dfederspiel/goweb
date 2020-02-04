package v2

type Pet struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Legs  bool   `json:"legs"`
	Owner `json:"owner"`
}
