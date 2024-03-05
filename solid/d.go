package solid

import "fmt"

// Интерфейс для отправки уведомлений
type Notifier interface {
	Notify(message string) error
}

// Структура почтового уведомления
type EmailNotifier struct{}

func (e *EmailNotifier) Notify(message string) error {
	fmt.Println("Sending email notification:", message)
	return nil
}

// Структура SMS уведомления
type SMSNotifier struct{}

func (s *SMSNotifier) Notify(message string) error {
	fmt.Println("Sending SMS notification:", message)
	return nil
}

// Структура, использующая уведомления
type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) error {
	return n.notifier.Notify(message)
}

func D() {
	fmt.Print("[Dependency Inversion Principle]\n")

	emailNotifier := &EmailNotifier{}
	smsNotifier := &SMSNotifier{}

	emailService := &NotificationService{notifier: emailNotifier}
	smsService := &NotificationService{notifier: smsNotifier}

	emailService.SendNotification("Hello, this is an email notification")
	smsService.SendNotification("Hello, this is an SMS notification")
}
