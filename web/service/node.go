package service

import (
	"x-ui/database"
	"x-ui/database/model"
)

// NodeService provides CRUD operations for managed nodes.
type NodeService struct{}

func (s *NodeService) List() ([]model.Node, error) {
	db := database.GetDB()
	var nodes []model.Node
	err := db.Find(&nodes).Error
	return nodes, err
}

func (s *NodeService) Create(name, apiURL, apiKey string) error {
	db := database.GetDB()
	node := &model.Node{Name: name, ApiURL: apiURL, ApiKey: apiKey}
	return db.Create(node).Error
}
