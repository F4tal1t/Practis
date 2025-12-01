package main

import "fmt"

// Coffee interface defines the contract for all coffee types
type Coffee interface {
	Cost() float64
	Description() string
}

// SimpleCoffee is the base implementation
type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 {
	return 2.0
}

func (c *SimpleCoffee) Description() string {
	return "Simple coffee"
}

// Espresso is another base implementation
type Espresso struct{}

func (e *Espresso) Cost() float64 {
	return 3.0
}

func (e *Espresso) Description() string {
	return "Espresso"
}

// MilkDecorator adds milk to coffee
type MilkDecorator struct {
	coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.5
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", milk"
}

// SugarDecorator adds sugar to coffee
type SugarDecorator struct {
	coffee Coffee
}

func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.2
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", sugar"
}

// WhipDecorator adds whipped cream to coffee
type WhipDecorator struct {
	coffee Coffee
}

func (w *WhipDecorator) Cost() float64 {
	return w.coffee.Cost() + 0.7
}

func (w *WhipDecorator) Description() string {
	return w.coffee.Description() + ", whipped cream"
}

// VanillaDecorator adds vanilla flavor to coffee
type VanillaDecorator struct {
	coffee Coffee
}

func (v *VanillaDecorator) Cost() float64 {
	return v.coffee.Cost() + 0.3
}

func (v *VanillaDecorator) Description() string {
	return v.coffee.Description() + ", vanilla"
}

// ChocolateDecorator adds chocolate to coffee
type ChocolateDecorator struct {
	coffee Coffee
}

func (ch *ChocolateDecorator) Cost() float64 {
	return ch.coffee.Cost() + 0.6
}

func (ch *ChocolateDecorator) Description() string {
	return ch.coffee.Description() + ", chocolate"
}

// CoffeeBuilder helps build complex coffee combinations
type CoffeeBuilder struct {
	coffee Coffee
}

func NewCoffeeBuilder(baseType string) *CoffeeBuilder {
	var baseCoffee Coffee
	switch baseType {
	case "simple":
		baseCoffee = &SimpleCoffee{}
	case "espresso":
		baseCoffee = &Espresso{}
	default:
		baseCoffee = &SimpleCoffee{}
	}
	
	return &CoffeeBuilder{coffee: baseCoffee}
}

func (cb *CoffeeBuilder) AddMilk() *CoffeeBuilder {
	cb.coffee = &MilkDecorator{coffee: cb.coffee}
	return cb
}

func (cb *CoffeeBuilder) AddSugar() *CoffeeBuilder {
	cb.coffee = &SugarDecorator{coffee: cb.coffee}
	return cb
}

func (cb *CoffeeBuilder) AddWhip() *CoffeeBuilder {
	cb.coffee = &WhipDecorator{coffee: cb.coffee}
	return cb
}

func (cb *CoffeeBuilder) AddVanilla() *CoffeeBuilder {
	cb.coffee = &VanillaDecorator{coffee: cb.coffee}
	return cb
}

func (cb *CoffeeBuilder) AddChocolate() *CoffeeBuilder {
	cb.coffee = &ChocolateDecorator{coffee: cb.coffee}
	return cb
}

func (cb *CoffeeBuilder) Build() Coffee {
	return cb.coffee
}

// CoffeeShop manages coffee orders
type CoffeeShop struct {
	name string
}

func NewCoffeeShop(name string) *CoffeeShop {
	return &CoffeeShop{name: name}
}

func (cs *CoffeeShop) ProcessOrder(coffee Coffee) {
	fmt.Printf("☕ %s: Preparing %s\n", cs.name, coffee.Description())
	fmt.Printf("💰 Total cost: $%.2f\n", coffee.Cost())
	fmt.Println("✅ Order ready!")
}
