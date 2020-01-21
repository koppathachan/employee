package employee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func addEmployee(e *Employee) {
	//do add here
	fmt.Printf("%#v\n", *e)
}

// Handler function adds an employee to the database
func Handler(w http.ResponseWriter, r *http.Request) {

	switch method := r.Method; method {
	case "POST":
		var e Employee
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			panic(err)
		}
		addEmployee(&e)
	}

}
