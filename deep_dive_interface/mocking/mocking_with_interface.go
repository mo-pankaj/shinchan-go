package mocking

import (
	"database/sql"
	"fmt"
	"time"
)

type ShopDB struct {
	*sql.DB
}

// ShopModel making it interface not a concrete type
// this allows any db not only Shop db if it implements CountCustomers and CountSales we can use interface on that struct
type ShopModel interface {
	CountCustomers(time.Time) (int, error)
	CountSales(time.Time) (int, error)
}

func (sdb *ShopDB) CountCustomers(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM customers WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

func (sdb *ShopDB) CountSales(since time.Time) (int, error) {
	var count int
	err := sdb.QueryRow("SELECT count(*) FROM sales WHERE timestamp > $1", since).Scan(&count)
	return count, err
}

// calculateSalesRatio accepts an interface not a concrete type
// this allow this function to be used by any of the structs that implement thet interface
func calculateSalesRatio(shopModel ShopModel) (string, error) {
	since := time.Now().Add(-24 * time.Hour)
	customer, err := shopModel.CountCustomers(since)
	if err != nil {
		return "", nil
	}
	sales, err := shopModel.CountSales(since)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%.2f", float64(customer/sales)), nil
}


// mongoDb lets say we moved to another db mongodb
// this allow seemless transaction between them
type mongoDb struct{}

func (mdb *mongoDb) CountCustomers(_ time.Time) (int, error) {
	return 1000, nil
}

func (mdb *mongoDb) CountSales(_ time.Time) (int, error) {
	return 300, nil
}

func main() {
	// sdb := ShopDB{&sql.DB{}}
	// r,_ := calculateSalesRatio(&sdb)
	// fmt.Printf("rate %s", r)

	mdb := mongoDb{}
	s, _ := calculateSalesRatio(&mdb)
	fmt.Printf("rate %s", s)

}
