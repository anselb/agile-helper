package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// BoardLists is a struct that holds a single list object from a given board
type BoardLists struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IDBoard string `json:"idBoard"`
}

// ListCards is a struct that holds a single card object from a given list
type ListCards struct {
	ID        string   `json:"id"`
	IDBoard   string   `json:"idBoard"`
	Name      string   `json:"name"`
	IDMembers []string `json:"idMembers"`
}

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	trelloKey := os.Getenv("TRELLO_KEY")
	trelloToken := os.Getenv("TRELLO_TOKEN")

	boardID := "ilKTjffb"
	listName := "To Do Now"

	// Calling Trello API for lists inside of board
	boardLists := fmt.Sprintf("https://api.trello.com/1/boards/%s/lists?cards=none&card_fields=all&filter=open&fields=all&key=%s&token=%s", boardID, trelloKey, trelloToken)
	listResponse, err := http.Get(boardLists)
	if err != nil {
		log.Fatal(err)
	}
	// Close response body when done with it
	defer listResponse.Body.Close()

	// Get body of the response
	listData, err := ioutil.ReadAll(listResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	// New BoardLists object
	var lists []BoardLists
	// Unmarshal data into a pointer to the lists struct
	json.Unmarshal(listData, &lists)

	// Print data
	fmt.Println(lists)

	// Find list by name
	notFoundList := true
	var listCards string
	for _, list := range lists {
		if list.Name == listName {
			notFoundList = false
			listID := list.ID
			listCards = fmt.Sprintf("https://api.trello.com/1/lists/%s/cards?key=%s&token=%s", listID, trelloKey, trelloToken)
		}
	}
	if notFoundList {
		log.Fatal("Given board name not found")
	}

	// Calling Trello API for cards inside of list
	cardResponse, err := http.Get(listCards)
	if err != nil {
		log.Fatal(err)
	}
	// Close response body when done with it
	defer cardResponse.Body.Close()

	// Get body of the response
	cardData, err := ioutil.ReadAll(cardResponse.Body)
	if err != nil {
		log.Fatal(err)
	}

	// New ListCards object
	var cards []ListCards
	// Unmarshal data into a pointer to the lists struct
	json.Unmarshal(cardData, &cards)

	// Print data
	fmt.Println(cards)

}
