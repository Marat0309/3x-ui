package model

// Node represents a remote server that can be managed by the panel.
type Node struct {
	Id     int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	ApiURL string `json:"apiUrl"`
	ApiKey string `json:"apiKey"`
}
