package exception

import "fmt"

type CustomerNotFound struct {
	Id int64
}

func (e CustomerNotFound) Error() string {
	return fmt.Sprintf("customer with id %v not found", e.Id)
}
