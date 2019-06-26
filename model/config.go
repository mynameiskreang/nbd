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
	Token   string   `yaml:"token" json:"token"`
	Renthub []string `yaml:"renthub" json:"renthub"`
}
