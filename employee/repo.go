package employee

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Repository interface
type Repository interface {
	Add(emp *Employee) error
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
		clientOptions := options.Client().ApplyURI("mongodb://192.168.5.223:27017")
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
