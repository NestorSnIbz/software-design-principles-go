package main

import "fmt"

func calculateDiscount(orderType string, amount float64) float64 {
	switch orderType {
	case "regular":
		return amount * 0.9
	case "premium":
		return amount * 0.8
	default:
		return amount
	}
}

func sendNotification() {
	fmt.Println("Sending email notification...")
}

func processOrder(orderType string, amount float64) {
	finalAmount := calculateDiscount(orderType, amount)
	fmt.Println("Order processed. Total:", finalAmount)
	sendNotification()
}

func main() {
	processOrder("premium", 200)
}