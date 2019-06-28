package model

type RenthubInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	LinkRoom    string `json:"link_room"`
	Project     string `json:"project"`
	LinkProject string `json:"link_project"`
}

type Config struct {
	Token   string `yaml:"token" json:"token"`
	Renthub struct {
		Link []string `yaml:"link" json:"link"`
		Max  int      `yaml:"max" json:"max"`
		Min  int      `yaml:"min" json:"min"`
	} `yaml:"renthub" json:"renthub"`
}
