package main

import "fmt"

func main() {
	baseNotifier := NewNotifier("Base Notifier")
	notifierWithLogging := AddLoggingDecorator(baseNotifier)

	baseNotifier.SendNotification("SE-2201 THE BEST!")
	fmt.Println()
	notifierWithLogging.SendNotification("SE-2201 THE BEST!")
}

type Notifier interface {
	SendNotification(message string)
}
type ConcreteNotifier struct {
	name string
}

func (n *ConcreteNotifier) SendNotification(message string) {
	fmt.Printf("[%s] Sending Notification: %s\n", n.name, message)
}

type Decorator func(Notifier) Notifier

func NewNotifier(name string) Notifier {
	return &ConcreteNotifier{name}
}

func AddLoggingDecorator(notifier Notifier) Notifier {
	return &LoggingDecorator{
		Notifier: notifier,
	}
}

type LoggingDecorator struct {
	Notifier
}

func (d *LoggingDecorator) SendNotification(message string) {
	fmt.Println("Logging:before sending notification")
	d.Notifier.SendNotification(message)
	fmt.Println("Logging:after sending notification")
}
