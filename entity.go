package gephi

type Attributes map[string]interface{}

type RGB struct {
	Red   int `json:"r"`
	Green int `json:"g"`
	Blue  int `json:"b"`
}

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}

type Entity struct {
	ID string `json:"-"`
	Attributes
}
