package model

type Country struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	ISOCode string `json:"iso_code"`
	Towns   []Town `json:"towns"`
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

type Localization struct {
	ID       string `json:"id" db:"id"`
	Position string `json:"position" db:"position"`
	Country  string `json:"country" db:"code"`
}
