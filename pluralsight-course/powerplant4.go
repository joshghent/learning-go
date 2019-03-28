package main

import (
	"fmt"
	"strings"
)

func main() {

	plants := []PowerPlant{
		PowerPlant{hydro, 100, active},
		PowerPlant{wind, 200, active},
		PowerPlant{solar, 50, active},
		PowerPlant{solar, 50, inactive},
		PowerPlant{wind, 25, unavailable},
		PowerPlant{wind, 35, inactive},
	}

	grid := PowerGrid{300, plants}

	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose an option: ")

	var option string

	// & means it passes the actual option memory address rather than passing a copy
	fmt.Scanln(&option)

	switch option {
	case "1":
		grid.generatePlantReport()
	case "2":
		grid.generateGridReport()
	default:
		print("Unknown option. Please try again.")
	}
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for i, p := range pg.plants {
		label := fmt.Sprintf("%s%d", "Plant #", i)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("%-20s%s\n", "Type: ", p.plantType)
		fmt.Printf("%-20s%.0f\n", "Capacity: ", p.capacity)
		fmt.Printf("%-20s%s\n", "Status: ", p.status)
	}
}

func (pg *PowerGrid) generateGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}

	label := "Power Grid"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("%-20s%.0f\n", "Capacity: ", capacity)
	fmt.Printf("%-20s%.0f\n", "Load: ", pg.load)
	fmt.Printf("%-20s%.0f\n", "Utilization: ", pg.load/capacity*100)
}
