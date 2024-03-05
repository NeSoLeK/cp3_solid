package solid

import "fmt"

// Работник
type Employee struct {
	Name     string
	Position string
	Salary   int
}

// Метод для вывода информации о сотруднике
func (e *Employee) PrintDetails() {
	fmt.Printf("Name: %s\nPosition: %s\nSalary: %d\n", e.Name, e.Position, e.Salary)
}

func S() {
	fmt.Print("[Single Responsibility Principle]\n")
	emp := Employee{Name: "John", Position: "Developer", Salary: 50000}
	emp.PrintDetails()
}
