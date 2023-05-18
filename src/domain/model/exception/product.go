package exception

import "fmt"

type ProductNotFound struct {
	Id int64
}

func (e ProductNotFound) Error() string {
	return fmt.Sprintf("product with id %v not found", e.Id)
}

type ProductWithNoStock struct {
	Id int64
}

func (e ProductWithNoStock) Error() string {
	return fmt.Sprintf("product with id %v have no stock", e.Id)
}
