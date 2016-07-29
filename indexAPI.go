package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

// no use now
func QueryAlgolia(queryStr, starredBy string) {
	client := algoliasearch.NewClient("EQDRH6QSH7", "6066c3e492d3a35cc0a425175afa89ff")

	index := client.InitIndex("githubRepo")

	params := algoliasearch.Map{
		"attributesToSnippet": []string{"description:40"},
		"facetFilters":        "starredBy:" + starredBy,
	}

	res, err := index.Search(queryStr, params)

	if err != nil {
		log.Println("error:", err)
	}

	b, err := json.Marshal(res)
	fmt.Println("search result:", string(b))
}

func SendToAlgolia(repoList []*map[string]interface{}, account string) error {

	client := algoliasearch.NewClient("EQDRH6QSH7", "6066c3e492d3a35cc0a425175afa89ff")
	index := client.InitIndex("githubRepo")

	// delete first
	// _, err := client.ClearIndex("githubRepo")
	// if err == nil {
	// 	log.Println("delete index ok")
	// } else {
	// 	log.Println("delete index fail")
	// }

	filters := "starredBy:" + account
	// facet := "starredBy:" + account
	// facetFilters := []string{facet}

	params := algoliasearch.Map{
		// Set your query parameters here
		"filters": filters,
	}
	err := index.DeleteByQuery("*", params)
	if err == nil {
		log.Println("delete ok !!!!", account)
	} else {
		log.Println("delete fail !!!!", account)
	}

	setting := make(map[string]interface{})
	setting["attributesForFaceting"] = []string{"starredBy"}
	index.SetSettings(setting)

	var objects []algoliasearch.Object

	for _, object := range repoList {

		b, _ := json.Marshal(*object)
		// log.Println("repo size:", len(b)) //110589

		// repoList := algoliasearch.Object (normalMap)

		if len(b) < 100000 {
			// temp disable
			// object2 = algoliasearch.Object(*object)
			// objects = append(objects, object2)

			objects = append(objects, *object)
		} else {
			log.Println("record size is larger thant 100K, may ~ 10K after minified json")
			repoURL := (*object)["repoURL"]
			sendTwilioAlert(repoURL.(string))
		}
	}

	_, err = index.AddObjects(objects)
	if err != nil {
		log.Println("add to algolia error:", err)
		return err
	} else {
		log.Println("add to algolia ok")
		return nil
	}
}
