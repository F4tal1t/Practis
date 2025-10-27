package main

import "fmt"

func main() {
	// Initialize a map
	scores := make(map[string]int)

	// Add elements to the map
	scores["Alice"] = 90
	scores["Bob"] = 85

	// Access an element
	fmt.Println("Alice's score:", scores["Alice"]) // Output: Alice's score: 90

	// Check if a key exists
	score, exists := scores["Charlie"]
	if exists {
		fmt.Println("Charlie's score:", score)
	} else {
		fmt.Println("Charlie does not have a score.") // Output
	}

	// Delete an element
	delete(scores, "Bob")
	fmt.Println("Scores after deleting Bob:", scores) // Output: map[Alice:90]
}
