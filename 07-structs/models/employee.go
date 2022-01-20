package models

type Employee struct {
	Id   int
	Name string
	Org  Organization
}

func NewEmployee() *Employee {
	emp := new(Employee)
	return emp
}
