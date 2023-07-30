package customerUsecase

import (
	"bootcamp-api-hmsi/models"
	"bootcamp-api-hmsi/modules/customers"
)

type customerRepository struct {
	Repo customers.CustomerRepository
}

func NewCustomerUsecase(Repo customers.CustomerRepository) customers.CustomerUsecase {
	return &customerRepository{Repo}
}

func (r *customerRepository) FindAll() (*[]models.Customers, error) {
	results, err := r.Repo.GetAll()

	if err != nil {
		return nil, err
	}
	return results, nil
}

func (r *customerRepository) Insert(c *models.RequestInsertCustomer) error {
	err := r.Repo.Create(c)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Update(c *models.RequestUpdateCustomer) error {
	err := r.Repo.Edit(c)

	if err != nil {
		return err
	}

	return nil
}

func (r *customerRepository) Delete(c *models.RequestDeleteCustomer) error {
	err := r.Repo.Hapus(c)

	if err != nil {
		return err
	}

	return nil
}
