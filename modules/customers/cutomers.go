package customers

import (
	"bootcamp-api-hmsi/models"
)

type (
	CustomerRepository interface {
		GetAll() (*[]models.Customers, error)
		Create(c *models.RequestInsertCustomer) error
		Edit(c *models.RequestUpdateCustomer) error
		Hapus(c *models.RequestDeleteCustomer) error
	}
	CustomerUsecase interface {
		FindAll() (*[]models.Customers, error)
		Insert(c *models.RequestInsertCustomer) error
		Update(c *models.RequestUpdateCustomer) error
		Delete(c *models.RequestDeleteCustomer) error
	}
)
