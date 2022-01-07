package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ashish", "New Dehli", "110011", "2000-01-01", "1"},
		{"1002", "Maria", "New Dehli", "110011", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
