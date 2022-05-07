package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: measurement <valores> <unidade>")
		os.Exit(1)
	}

	sourceUnit := os.Args[len(os.Args)-1]
	sourceValues := os.Args[1 : len(os.Args)-1]

	var targetUnid string

	if sourceUnit == "celsius" {
		targetUnid = "fahrenheit"
	} else if sourceUnit == "quilometros" {
		targetUnid = "milhas"
	} else {
		fmt.Printf("%s não é uma unidade conhecida! ", sourceUnit)
		os.Exit(1)
	}

	for index, value := range sourceValues {
		sourceValue, err := strconv.ParseFloat(value, 64)

		if err != nil {
			fmt.Printf("O valor %s na posição %d não é um número válido!\n", value, index)
			os.Exit(1)
		}
		var targetValue float64

		if sourceUnit == "celsius" {
			targetValue = sourceValue*1.8 + 32
		} else {
			targetValue = sourceValue / 1.60934
		}

		fmt.Printf("%.2f %s = %.2f %s\n", sourceValue, sourceUnit, targetValue, targetUnid)
	}
}
