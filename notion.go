package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Notion struct {
	Parent     `json:"parent"`
	Properties `json:"properties"`
}

type Parent struct {
	DatabaseID string `json:"database_id"`
}

type Properties struct {
	Name   Name   `json:"Name"`
	Weight Weight `json:"Weight"`
}

type Name struct {
	Title []Title `json:"title"`
}

type Weight struct {
	Number float64 `json:"number"`
}

type Title struct {
	Text Text `json:"text"`
}

type Text struct {
	Content string `json:"content"`
}

func saveToNotion(date string, weight float64, dbID string, bearerToken string) {
	payload := &Notion{
		Parent: Parent{
			DatabaseID: dbID,
		},
		Properties: Properties{
			Name: Name{
				Title: []Title{
					{
						Text: Text{
							Content: date,
						},
					},
				},
			},
			Weight: Weight{
				Number: weight,
			},
		},
	}

	// Encode payload.
	pJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	// Make a request to Notion API.
	responseBody := bytes.NewBuffer(pJSON)
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.notion.com/v1/pages", responseBody)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Notion-Version", "2021-05-13")
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	//Handle Error.
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	defer req.Body.Close()
	log.Printf("API response %d", res.StatusCode)

	// Log response in case of non 200.
	if res.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}
