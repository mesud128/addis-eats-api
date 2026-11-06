package main

import (
	"fmt"
	"unicode/utf8"
)

// Assignment 1: Menu Report

type Dish struct {
	Name      string
	Price     int // Price in santim
	IsFasting bool
}

func menuReport() {
	fmt.Println("\n========== MENU REPORT ==========")

	dishes := []Dish{
		{"Doro Wat", 18500, false},
		{"Shiro", 12000, true},
		{"Tibs", 24000, false},
		{"Injera Basket", 4500, true},
		{"Firfir", 10000, true},
		{"Kitfo", 28000, false},
		{"Beyaynet", 15000, true},
		{"Gomen Besiga", 16000, false},
	}

	total := 0
	cheapest := dishes[0]
	mostExpensive := dishes[0]
	fastingCount := 0
	nonFastingCount := 0

	fmt.Println("\nFood Menu:")

	for _, dish := range dishes {
		fmt.Printf("%-20s ETB %.2f\n", dish.Name, float64(dish.Price)/100)

		total += dish.Price

		if dish.Price < cheapest.Price {
			cheapest = dish
		}

		if dish.Price > mostExpensive.Price {
			mostExpensive = dish
		}

		if dish.IsFasting {
			fastingCount++
		} else {
			nonFastingCount++
		}
	}

	average := float64(total) / float64(len(dishes))

	fmt.Printf("\nCheapest dish: %s (ETB %.2f)\n",
		cheapest.Name, float64(cheapest.Price)/100)

	fmt.Printf("Most expensive dish: %s (ETB %.2f)\n",
		mostExpensive.Name, float64(mostExpensive.Price)/100)

	fmt.Printf("Average price: ETB %.2f\n", average/100)
	fmt.Printf("Fasting dishes: %d\n", fastingCount)
	fmt.Printf("Non-fasting dishes: %d\n", nonFastingCount)
}

// Assignment 2: Delivery Fee Table

func deliveryFee(area string, subtotal int) int {
	// Free delivery for orders of 50,000 santim or more.
	if subtotal >= 50000 {
		return 0
	}

	switch area {
	case "Bole":
		return 3000
	case "Kazanchis":
		return 3500
	case "Piassa":
		return 4000
	case "Megenagna":
		return 4500
	default:
		return 5000
	}
}

func feeTable() {
	fmt.Println("\n========== DELIVERY FEE TABLE ==========")

	areas := []string{
		"Bole",
		"Kazanchis",
		"Piassa",
		"Megenagna",
	}

	subtotals := []int{
		10000,
		30000,
		60000,
	}

	for _, area := range areas {
		fmt.Printf("\nArea: %s\n", area)

		for _, subtotal := range subtotals {
			fee := deliveryFee(area, subtotal)
			grandTotal := subtotal + fee

			fmt.Printf(
				"Subtotal: ETB %.2f | Delivery: ETB %.2f | Grand Total: ETB %.2f\n",
				float64(subtotal)/100,
				float64(fee)/100,
				float64(grandTotal)/100,
			)
		}
	}
}

// Assignment 3: Amharic Character Count

func amharicCount() {
	fmt.Println("\n========== AMHARIC CHARACTER COUNT ==========")

	dishes := []string{
		"Doro Wat",
		"ሽሮ",
		"ቲብስ",
		"ፍርፍር",
		"በያይነት",
	}

	for _, dish := range dishes {
		byteLength := len(dish)
		runeCount := utf8.RuneCountInString(dish)

		fits := runeCount <= 20

		fmt.Printf("\nDish: %s\n", dish)
		fmt.Printf("Byte length: %d\n", byteLength)
		fmt.Printf("Character count: %d\n", runeCount)
		fmt.Printf("Fits 20-character field: %t\n", fits)
	}
}

// Main function

func main() {
	menuReport()
	feeTable()
	amharicCount()
}
