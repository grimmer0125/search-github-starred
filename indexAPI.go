package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/grimmer0125/elastic"
	// elastigo "github.com/mattbaird/elastigo/lib"
)

// no use now, for testing

// type Tweet struct {
// 	User    string `json:"user"`
// 	Message string `json:"message"`
// }

const (
	awsURL      = "https://search-searchgithub-7c4xubb6ne3t7keszcai7kqi3m.us-west-2.es.amazonaws.com"
	githubIndex = "githubrepos" //can not be githubRepos, should lower case !!
)

func initElastic() (client *elastic.Client) {
	//Create a client
	client, err := elastic.NewClient(
		elastic.SetURL(awsURL),
		// elastic.SetHttpClient(http.DefaultClient),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false))

	if err != nil {
		// Handle error
		log.Println("new elastic Search fail: ", err)
		return nil
	}

	return client
}

func useScrollToDelete(client *elastic.Client, account string) {

	q := elastic.NewMatchQuery("starredBy", account)

	searchResult, err := client.Scroll(githubIndex).Size(10000).Query(q).Do()

	bulkRequest := client.Bulk()

	if err != nil {
		// Handle error
		log.Println("new elastic Scroll fail: ", err)
	} else {
		log.Println("get scroll result:")
		if searchResult.Hits.TotalHits > 0 {
			fmt.Printf("Found a total of %d tweets\n", searchResult.Hits.TotalHits)

			// Iterate through results
			for _, hit := range searchResult.Hits.Hits {

				// fmt.Println("got id:", hit.Id)

				request := elastic.NewBulkDeleteRequest().Index(githubIndex).Type(account).Id(hit.Id)

				bulkRequest = bulkRequest.Add(request)
			}
		} else {
			// No hits
			fmt.Print("Found no hits in %s\n", account)
		}
	}

	log.Println("numbe of delete requests:", bulkRequest.NumberOfActions())
	bulkResponse, err := bulkRequest.Do()
	if err != nil {

		//		get bulk error: elastic: No bulk actions to commit !!!!

		log.Println("get bulk error:", err)
		// t.Fatal(err)
		return
	}
	log.Println("after, numbe of delete requests:", bulkRequest.NumberOfActions())

	if bulkResponse == nil {
		log.Println("expected bulkResponse to be != nil; got nil")
	} else {
		log.Println("buld resp ok")
	}

	log.Println("use elasticsearch done")
	// use this result to delete !!!!!!!!!!!!!!

	//ref 1. https://github.com/olivere/elastic/blob/e4233aab21b455148600a8da2172576082ad8125/scroll_test.go
	//    2. https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-scroll.html

}

func SendToGitHubElasticsearch(repoList []*GitHubRepo, account string) error {

	client := initElastic()

	if client == nil {
		// Handle error
		log.Println("new elastic Search fail")
		return nil
	}

	// Delete the content under the account/type
	useScrollToDelete(client, account)

	// Create an index
	// _, err = client.CreateIndex(githubIndex).Do()
	// if err != nil {
	// 	// Handle error
	// 	// panic(err)
	// 	log.Println("create index fail :", err)
	// }

	// Add a document to the index
	// tweet := Tweet{User: "olivere", Message: "Take Five"}
	// _, err = client.Index().
	// 	Index("twitter789").
	// 	Type("tweet").
	// 	Id("1").
	// 	BodyJson(tweet).
	// 	Refresh(true).
	// 	Do()
	// if err != nil {
	// 	// Handle error
	// 	// panic(err)
	// 	log.Println("index fail,", err)
	// }

	// Delete the index again
	// log.Println("try to delete first")
	// _, err := client.DeleteIndex(githubIndex).Do()
	// if err != nil {
	// 	// Handle error
	// 	log.Println("delete first fail:", err)

	// 	if err.Error() == "elastic: Error 404 (Not Found): no such index [type=index_not_found_exception]" {
	// 		log.Println("just not found index")
	// 	} else {
	// 		return err
	// 	}
	// }

	// bulk - add
	var err error
	if len(repoList) > 0 {
		bulkRequest := client.Bulk()
		for i, repo := range repoList {

			objectID := strconv.Itoa(i)
			// log.Println("repo:", *repo)
			index := elastic.NewBulkIndexRequest().Index(githubIndex).Type(account).Id(objectID).Doc(*repo)
			bulkRequest = bulkRequest.Add(index)
		}
		// tweet3 := make(map[string]interface{})
		// tweet3["user2"] = "x21"
		// tweet3["message2"] = "x22"

		// tweet3 := map[string]interface{}
		//tweet1 := GitHubRepo{"1", "1", "1", "1", "1", "1", "1", "1"}
		// tweet1 := Tweet{User: "olivere22", Message: "Welcome to Golang and Elasticsearch."}
		// tweet2 := Tweet{User: "sandrae22", Message: "Dancing all night long. Yeah."}

		// index1Req := elastic.NewBulkIndexRequest().Index("repo").Type("tweet").Id("91").Doc(tweet1)
		// index2Req := elastic.NewBulkIndexRequest().Index("twitter1271").Type("tweet").Id("3").Doc(tweet3)

		// // // bulkRequest := client.Bulk()
		// bulkRequest = bulkRequest.Add(index1Req)
		// bulkRequest = bulkRequest.Add(index2Req)

		log.Println("numbe of add requests:", bulkRequest.NumberOfActions())
		bulkResponse, err := bulkRequest.Do()
		if err != nil {
			log.Println("get bulk error:", err)
			// t.Fatal(err)
			return err
		}
		log.Println("after, number of add requests:", bulkRequest.NumberOfActions())

		if bulkResponse == nil {
			log.Println("expected bulkResponse to be != nil; got nil")
		} else {
			log.Println("buld resp ok")
		}
	} else {
		log.Println("repo list is empty, so directly return error nil")
		err = nil
	}

	log.Println("use elasticsearch done")

	return err
}
