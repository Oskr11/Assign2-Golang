package main

import "fmt"

func main() {
	factoryA := &ConcreteFactoryA{}
	factoryB := &ConcreteFactoryB{}

	productA := factoryA.CreateProduct()
	productB := factoryB.CreateProduct()

	fmt.Println("Product A:", productA.GetName())
	fmt.Println("Product B:", productB.GetName())
}
type Product interface {
	GetName() string
}

type ConcreteProductA struct {
	name string
}

func (p *ConcreteProductA) GetName() string {
	return p.name
}

type ConcreteProductB struct {
	name string
}

func (p *ConcreteProductB) GetName() string {
	return p.name
}

type Factory interface {
	CreateProduct() Product
}

type ConcreteFactoryA struct{}

func (f *ConcreteFactoryA) CreateProduct() Product {
	return &ConcreteProductA{name: "Product A"}
}
type ConcreteFactoryB struct{}

func (f *ConcreteFactoryB) CreateProduct() Product {
	return &ConcreteProductB{name: "Product B"}
}


