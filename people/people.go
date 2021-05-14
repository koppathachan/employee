package people

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Server struct {
	UnimplementedPeopleServer
}

type People struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

func (s Server) GetPeople(_ *Empty, src People_GetPeopleServer) error {
	pathToFile := os.Getenv("PATH_TO_FILE")
	content, rferr := ioutil.ReadFile(pathToFile)
	if rferr != nil {
		return rferr
	}
	var people []People
	err2 := json.Unmarshal(content, &people)
	if err2 != nil {
		return err2
	}
	for _, x := range people {
		err := src.Send(&Person{
			Id:     x.Id,
			Name:   x.Name,
			Gender: x.Gender,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
