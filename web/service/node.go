package service

import (
	"errors"

	"x-ui/database"
	"x-ui/database/model"

	"gorm.io/gorm"
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

func (s *NodeService) Get(id int) (*model.Node, error) {
	db := database.GetDB()
	node := &model.Node{}
	err := db.First(node, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (s *NodeService) Update(id int, name, apiURL, apiKey string) error {
	db := database.GetDB()
	node := &model.Node{}
	if err := db.First(node, id).Error; err != nil {
		return err
	}
	node.Name = name
	node.ApiURL = apiURL
	node.ApiKey = apiKey
	return db.Save(node).Error
}

func (s *NodeService) Delete(id int) error {
	db := database.GetDB()
	res := db.Delete(&model.Node{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
