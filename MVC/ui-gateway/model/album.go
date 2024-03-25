package model

type Character struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Class string  `json:"class"`
	Level float64 `json:"level"`
	Mony  float64 `json:"mony"`
}

type deadCharacter struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Class string  `json:"cause_of_death"`
	Level float64 `json:"level"`
}
