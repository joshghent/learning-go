package main

import "fmt"

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")

	var option string

	// & means it passes the actual option memory address rather than passing a copy
	fmt.Scanln(&option)

	switch option {
	case "1":
		generatePlantCapacityReport(plantCapacities)
	case "2":
		generatePowerGridReport(activePlants, plantCapacities, gridLoad)

	default:
		print("Unknown option. Please try again.")
	}
}

func generatePlantCapacityReport(plantCapacities []float64) {
	for index, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", index, cap)
	}
}

func generatePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantID := range activePlants {
		capacity += plantCapacities[plantID]
	}

	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", gridLoad)
	fmt.Printf("%-20s%.0f\n", "Utilization: ", gridLoad/capacity*100)
}
