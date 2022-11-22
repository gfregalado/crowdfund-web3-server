package services

import "github.com/gfregalado/crowdfund-api/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByEmail(string) (*models.DBResponse, error)
	UpdateUserById(id string, field string, value string) (*models.DBResponse, error)
}
