package main

import "fmt"

type Appliance interface {
	TurnOn()
}


type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}


type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering Up")
}


func (c CoffeePot) Brew() {
	fmt.Println("Brewing, coffee machine made by", c)
}



func main() {
	var device Appliance
	device = Fan("Windco Breeze")
	device.TurnOn()

	device = CoffeePot("LuxBrew")
	device.TurnOn()

	//This code below works haha. but don't do it like this.
	//Instaed use Type conversion.
	//var coffeeMachine CoffeePot = "Brew Co"
	//coffeeMachine.Brew()

	//Below we used Type Assertion, which is saying that
	//we know device is of type Applianace. However,
	//we know it is actually concrete type CoffeePot
	var coffeeMaker CoffeePot = device.(CoffeePot)
	coffeeMaker.Brew()

	//notes:
	//if u remove func (c CoffeePot) TurnOn () {
	//.
	//Then it means CoffeePot type doesn't satisfy the interface
	//u can't assign value of that type to a variables that use that
	//interface as their type
	
}

