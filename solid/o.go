package solid

import "fmt"

// Интерфейс для различных типов оплаты
type PaymentMethod interface {
	Pay(amount float64) string
}

// Структура для оплаты кредитной картой
type CreditCard struct{}

func (c *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f with Credit Card", amount)
}

// Структура для оплаты наличными
type Cash struct{}

func (c *Cash) Pay(amount float64) string {
	return fmt.Sprintf("Paying %.2f with Cash", amount)
}

func ProcessPayment(pm PaymentMethod, amount float64) string {
	return pm.Pay(amount)
}

func O() {
	fmt.Print("[Open/Closed Principle]\n")

	creditCard := &CreditCard{}
	cash := &Cash{}

	fmt.Println(ProcessPayment(creditCard, 100.50))
	fmt.Println(ProcessPayment(cash, 50.25))
}
