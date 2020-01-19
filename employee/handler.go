package employee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AddEmployee function adds an employee to the database
func AddEmployee(w http.ResponseWriter, r *http.Request) {
	var e Employee
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", e)
}
