package ch4

import (
	"encoding/json"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "movie1", Year: 1987, Color: true, Actors: []string{"Jhon", "Jake"}},
	{Title: "movie2", Year: 1986, Color: false, Actors: []string{"Acor", "cc"}},
}

func getMovieJsonStr() string {
	// data, err := json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}
