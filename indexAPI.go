package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

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

func SendToAlgolia(repoList []*map[string]interface{}) error {

	client := algoliasearch.NewClient("EQDRH6QSH7", "6066c3e492d3a35cc0a425175afa89ff")

	//		ClearIndex(name string) (res UpdateTaskRes, err error)
	_, err := client.ClearIndex("githubRepo")
	if err == nil {
		log.Println("delete index ok")
	} else {
		log.Println("delete index fail")
	}

	index := client.InitIndex("githubRepo")

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
