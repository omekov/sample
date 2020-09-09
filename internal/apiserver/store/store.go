package store

import "github.com/omekov/sample/internal/apiserver/store/customers"

type Store struct {
	Customers customers.Customers
}

func InitMongodb() {

}
