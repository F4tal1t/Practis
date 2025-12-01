package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Design Patterns in Go ===\n")

	// Singleton Pattern Demo
	fmt.Println("1. SINGLETON PATTERN")
	SingletonDemo()

	// Factory Pattern Demo
	fmt.Println("\n2. FACTORY PATTERN")
	FactoryDemo()

	// Observer Pattern Demo
	fmt.Println("\n3. OBSERVER PATTERN")
	ObserverDemo()

	// Decorator Pattern Demo
	fmt.Println("\n4. DECORATOR PATTERN")
	DecoratorDemo()
}

// Singleton Pattern Demo
func SingletonDemo() {
	db1 := GetDatabaseInstance()
	db2 := GetDatabaseInstance()

	fmt.Printf("Database instance 1: %p\n", db1)
	fmt.Printf("Database instance 2: %p\n", db2)
	fmt.Printf("Are they the same instance? %t\n", db1 == db2)

	db1.Connect()
	db2.Query("SELECT * FROM users")
}

// Factory Pattern Demo
func FactoryDemo() {
	// Create different types of notifications
	email := CreateNotification("email")
	sms := CreateNotification("sms")
	push := CreateNotification("push")

	email.Send("Hello via Email!")
	sms.Send("Hello via SMS!")
	push.Send("Hello via Push!")
}

// Observer Pattern Demo
func ObserverDemo() {
	// Create subject (news agency)
	newsAgency := NewNewsAgency()

	// Create observers (subscribers)
	subscriber1 := &NewsSubscriber{name: "John"}
	subscriber2 := &NewsSubscriber{name: "Jane"}
	
	// Register observers
	newsAgency.Register(subscriber1)
	newsAgency.Register(subscriber2)

	// Publish news
	newsAgency.SetNews("Breaking: Go 1.22 Released!")
}

// Decorator Pattern Demo
func DecoratorDemo() {
	// Base coffee
	coffee := &SimpleCoffee{}
	fmt.Printf("Base coffee: %s, Cost: $%.2f\n", coffee.Description(), coffee.Cost())

	// Add milk
	coffeeWithMilk := &MilkDecorator{coffee: coffee}
	fmt.Printf("With milk: %s, Cost: $%.2f\n", coffeeWithMilk.Description(), coffeeWithMilk.Cost())

	// Add sugar
	coffeeWithMilkAndSugar := &SugarDecorator{coffee: coffeeWithMilk}
	fmt.Printf("With milk and sugar: %s, Cost: $%.2f\n", coffeeWithMilkAndSugar.Description(), coffeeWithMilkAndSugar.Cost())
}
