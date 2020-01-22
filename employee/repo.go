package employee

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository interface
type Repository interface {
	Add(emp *Employee) error
	Get(id []string) ([]Employee, error)
}

type repo struct{}

//NewRepository is the constructor for repo class.
func NewRepository() Repository {
	return &repo{}
}

var once sync.Once
var client *mongo.Client

func getCLient() *mongo.Client {
	once.Do(func() {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		cli, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			log.Panic("Error connecting mongoDB", err)
		}
		client = cli
	})
	return client
}

//Add an employee into the database.
func (*repo) Add(emp *Employee) error {
	client := getCLient()
	collection := client.Database("akhil").Collection("employee")
	_, err := collection.InsertOne(context.TODO(), emp)
	return err
}

func (*repo) Get(id []string) ([]Employee, error) {
	if id == nil || len(id) == 0 {
		var empList []Employee
		client := getCLient()
		collection := client.Database("akhil").Collection("employee")
		cur, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Panic("Could not find from database")
			return nil, err
		}
		for cur.Next(context.TODO()) {
			var emp Employee
			if err := cur.Decode(&emp); err != nil {
				log.Panic("Decode error")
			}
			empList = append(empList, emp)
		}
		return empList, nil
	}

	log.Panic("Not implemented.")
	return nil, nil
}
