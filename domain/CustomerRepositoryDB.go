package domain

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findCustomerSql := "select customer_id,	name, city, zipcode, date_of_birth, status  from customers"
	rows, err := d.client.Query(findCustomerSql)
	if err != nil {
		log.Println("Error while quering customer table" + err.Error())
	}
	customers := make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			log.Println("Error while scaning customer " + err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:Wc4rt4y4.07@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client}
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, error) {
	customerSql := "select customer_id,	name, city, zipcode, date_of_birth, status  from customers where customer_id = ?"
	row := d.client.QueryRow(customerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	log.Println("Id:" + id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Customer not found")
		} else {
			log.Println("Error while scaning customer " + err.Error())
			return nil, err
		}
	}
	return &c, nil
}
