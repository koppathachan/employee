package employee

import "go.mongodb.org/mongo-driver/bson/primitive"

//Employee model
type Employee struct {
	ID          *primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string              `json:"name"`
	Designation string              `json:"designation"`
	Exp         int                 `json:"exp"`
	Salary      float64             `json:"salary"`
}
