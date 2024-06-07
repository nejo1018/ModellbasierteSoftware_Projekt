package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Payable interface {
	calculatePay(hours int) float64
}

type CEO struct {
	hourlyRate float64
}

type Developer struct {
	hourlyRate float64
}

type Janitor struct {
	hourlyRate float64
}

func (c CEO) calculatePay(hours int) float64 {
	return float64(hours) * c.hourlyRate
}

func (d Developer) calculatePay(hours int) float64 {
	return float64(hours) * d.hourlyRate
}

func (j Janitor) calculatePay(hours int) float64 {
	return float64(hours) * j.hourlyRate
}

func calculateSalary(p Payable, hours int) float64 {
	return p.calculatePay(hours)
}

type EmployeeFunction struct {
	calculate func(int) float64
}

func calculateCEOPay(hours int) float64 {
	return 250.0 * float64(hours) // Beispiel eines Gehalts
}

func calculateDeveloperPay(hours int) float64 {
	return 150.0 * float64(hours)
}

func calculateJanitorPay(hours int) float64 {
	return 25.0 * float64(hours)
}

func random1to3() int {
	return rand.Intn(3) + 1 // rand.Intn(3) gibt eine Zahl von 0 bis 2 zurück, daher + 1
}

func Berechnung() {
	// Zufallsgenerator initialisieren
	rand.Seed(time.Now().UnixNano())

	// Mitarbeiter-Map vorbereiten
	employeeMap := map[string]EmployeeFunction{
		"CEO":       {calculate: calculateCEOPay},
		"Developer": {calculate: calculateDeveloperPay},
		"Janitor":   {calculate: calculateJanitorPay},
	}

	// Eine Zufallszahl zwischen 1 und 3 generieren und entsprechend den Mitarbeitertyp setzen
	ran := random1to3()

	var employeeType string
	switch ran {
	case 1:
		employeeType = "CEO"
	case 2:
		employeeType = "Developer"
	case 3:
		employeeType = "Janitor"
	}

	hoursWorked := 40

	// Lohn berechnen
	if empFunc, exists := employeeMap[employeeType]; exists {
		fmt.Println(employeeType, ", Gesamtlohn: ", empFunc.calculate(hoursWorked))
	} else {
		fmt.Println("Unbekannter Angestellter")
	}

}

func main() {

	for i := 0; i < 10; i++ {
		Berechnung()
	}
}
