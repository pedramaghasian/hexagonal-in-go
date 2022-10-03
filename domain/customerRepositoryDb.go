package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pedramaghasian/hexagonal-in-go/errs"
	"github.com/pedramaghasian/hexagonal-in-go/logger"
)

// Adaptor
type customerRepositoryDb struct {
	client *sqlx.DB
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSql := "SELECT * from customers"
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "SELECT * from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while querying customer table" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected DB error")
	}
	return customers, nil
}

func (d customerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT * from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, customerSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) customerRepositoryDb {
	return customerRepositoryDb{client: dbClient}
}
