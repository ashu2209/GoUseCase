package main

import (
	"fmt"
)

func main() {

	// initialie the products map
	// key as String and value as Slice of strings
	products := make(map[string][]string)

	for {
		var choice int
		fmt.Println("Enter your choice:")
		fmt.Println("1. Adding a new product")
		fmt.Println("2. Delete a new product")
		fmt.Println("3. Print the products list")
		fmt.Println("4. Exit")
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var key string
			var value string
			fmt.Print("Enter key: ")
			fmt.Scan(&key)
			fmt.Print("Enter value: ")
			fmt.Scan(&value)
			// adding products directly to the map without checking if the key already exists
			products[key] = append(products[key], value)
			// Check if the key exists in the map
			if _, exists := products[key]; !exists {
				// Add the key-value pair to the map
				products[key] = append(products[key], value)
				fmt.Println("Element added successfully!")
			} else {
				fmt.Println("Element with key", key, "already exists in the map.")
			}

			fmt.Println("product added successfully!")
		case 2:
			var productKey string
			var valueToDelete string
			fmt.Print("Enter key : ")
			fmt.Scan(&productKey)
			fmt.Print("Enter value: ")
			fmt.Scan(&valueToDelete)
			index := findIndex(products[productKey], valueToDelete)
			if index != -1 {
				products[productKey] = append(products[productKey][:index], products[productKey][index+1:]...)
				fmt.Println("product", valueToDelete, "deleted from product list successfully!")
			} else {
				fmt.Println("product", valueToDelete, "not found in slice")
			}
			//delete(products, keyToDelete)
			fmt.Println("product deleted successfully!")
		case 3:
			fmt.Println("Printing a products and It's family Type:")
			for key, value := range products {
				fmt.Println("Key:", key, "Value:", value)
			}
		case 4:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// Displaying the map after each operation
		fmt.Println("Updated Map:", products)
	}

}

// Function to find the index of a value in a slice
func findIndex(slice []string, value string) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}
