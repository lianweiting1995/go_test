package ch4

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestMovie(t *testing.T) {
	jsonStr := getMovieJsonStr()
	var Titles []struct {
		Title string
	}
	err := json.Unmarshal([]byte(jsonStr), &Titles)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Titles)
}
