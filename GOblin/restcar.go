package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Car struct {
	ID                  int64   `json:"id"`
	Name                string  `json:"name"`
	Company             string  `json:"company"`
	YearOfManufacturing int     `json:"year"`
	Price               float64 `json:"price"`
}

var (
	carInventory = make(map[int64]*Car)
	mutex        sync.RWMutex
	idCounter    int64
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateID() int64 {
	idCounter++
	return idCounter
}

func addCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var car Car
	if err := json.NewDecoder(r.Body).Decode(&car); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if car.Name == "" || car.Company == "" {
		http.Error(w, "Name and Company are required", http.StatusBadRequest)
		return
	}

	// Lock the inventory before modifying it
	mutex.Lock()
	defer mutex.Unlock()

	// Auto-generate ID
	car.ID = generateID()
	carInventory[car.ID] = &car

	// Set content type and return the car details
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(car)
}

func getCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	if idStr == "" {
		getAllCarsHandler(w, r)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Use RLock for read operations
	mutex.RLock()
	car, exists := carInventory[id]
	mutex.RUnlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(car)
}

func getAllCarsHandler(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	cars := make([]*Car, 0, len(carInventory))
	for _, car := range carInventory {
		cars = append(cars, car)
	}
	mutex.RUnlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func deleteCarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/cars/")
	if idStr == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Lock the inventory before deleting
	mutex.Lock()
	_, exists := carInventory[id]
	if exists {
		delete(carInventory, id)
	}
	mutex.Unlock()

	if !exists {
		http.NotFound(w, r)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func carRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path == "/cars" {
			addCarHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	case http.MethodGet:
		if strings.HasPrefix(r.URL.Path, "/cars") {
			getCarHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	case http.MethodDelete:
		if strings.HasPrefix(r.URL.Path, "/cars/") {
			deleteCarHandler(w, r)
		} else {
			http.NotFound(w, r)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/cars", carRouter)
	http.HandleFunc("/cars/", carRouter)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
