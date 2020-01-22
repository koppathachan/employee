package employee

import (
	"encoding/json"
	"log"
	"net/http"
)

func addEmployee(e *Employee) error {
	db := NewRepository()
	if err := db.Add(e); err != nil {
		return err
	}
	log.Println("Added new employee")
	return nil
}

func getEmployee() ([]Employee, error) {
	db := NewRepository()
	return db.Get(nil)
}

// Handler function adds an employee to the database
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch method := r.Method; method {
	case "POST":
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			panic(err)
		}
		if err := addEmployee(&e); err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		json.NewEncoder(w).Encode("Added new employee")
	case "GET":
		if el, err := getEmployee(); err != nil {
			json.NewEncoder(w).Encode(err)
		} else {
			json.NewEncoder(w).Encode(el)
		}
	}

}
