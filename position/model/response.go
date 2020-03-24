package model

type Country struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Towns []Town `json:"towns"`
}

type Town struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Districts []District `json:"districts"`
}

type District struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Localisation struct {
	ID       string `json:"id" db:"id"`
	Position string `json:"position" db:"position"`
	Country  string `json:"country" db:"code"`
}
