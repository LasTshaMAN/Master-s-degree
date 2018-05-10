package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getContext(word string) Context {
	url := fmt.Sprintf("http://api.conceptnet.io/c/en/%s?offset=0&limit=1000", word)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Errorf("Couldn't create request: %s", err)
		return newContext(nil)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Errorf("Couldn't make request: %s", err)
		return newContext(nil)
	}
	defer resp.Body.Close()

	var ctxPage ContextPage
	if err := json.NewDecoder(resp.Body).Decode(&ctxPage); err != nil {
		fmt.Errorf("Couldn't decode response: %s", err)
		return newContext(nil)
	}

	relatedTerms := extractRelatedTerms(ctxPage, word)
	relatedTerms = removeDuplicates(relatedTerms)
	return newContext(relatedTerms)
}

func extractRelatedTerms(ctxPage ContextPage, word string) (result []string) {
	for _, edge := range ctxPage.Edges {
		if edge.Start.Label == word {
			result = append(result, edge.End.Label)
		}
	}
	return
}

var client = &http.Client{}

type ContextPage struct {
	Context []string `json:"@context"`
	ID      string   `json:"@id"`
	Edges   []struct {
		ID      string `json:"@id"`
		Type    string `json:"@type"`
		Dataset string `json:"dataset"`
		End     struct {
			ID       string `json:"@id"`
			Type     string `json:"@type"`
			Label    string `json:"label"`
			Language string `json:"language"`
			Term     string `json:"term"`
		} `json:"end"`
		License string `json:"license"`
		Rel     struct {
			ID    string `json:"@id"`
			Type  string `json:"@type"`
			Label string `json:"label"`
		} `json:"rel"`
		Sources []struct {
			ID          string `json:"@id"`
			Type        string `json:"@type"`
			Contributor string `json:"contributor"`
			Process     string `json:"process"`
		} `json:"sources"`
		Start struct {
			ID         string `json:"@id"`
			Type       string `json:"@type"`
			Label      string `json:"label"`
			Language   string `json:"language"`
			SenseLabel string `json:"sense_label"`
			Term       string `json:"term"`
		} `json:"start"`
		SurfaceText interface{} `json:"surfaceText"`
		Weight      float64     `json:"weight"`
	} `json:"edges"`
	View struct {
		ID                string `json:"@id"`
		Type              string `json:"@type"`
		Comment           string `json:"comment"`
		FirstPage         string `json:"firstPage"`
		NextPage          string `json:"nextPage"`
		PaginatedProperty string `json:"paginatedProperty"`
	} `json:"view"`
}
