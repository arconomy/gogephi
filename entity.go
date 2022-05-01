package gephi

type Attributes map[string]interface{}

type RGB struct {
	Red   string `json:"r"`
	Green string `json:"g"`
	Blue  string `json:"b"`
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
