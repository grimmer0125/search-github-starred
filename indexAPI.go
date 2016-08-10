package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	// elastic "gopkg.in/olivere/elastic.v3"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/grimmer0125/elastic"
	// elastigo "github.com/mattbaird/elastigo/lib"
	// "gopkg.in/olivere/elastic.v3"
)

// no use now

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

// func SendToElastigo(repoList []*map[string]interface{}, account string) error {
// 	log.Println("Send to elastic Search  ")

// 	core := elastigo.NewConn()
// 	core.Domain = "search-searchgithub-7c4xubb6ne3t7keszcai7kqi3m.us-west-2.es.amazonaws.com"
// 	core.Port = "80"

// 	// 2nd
// 	// // Trace all requests
// 	// core.RequestTracer = func(method, url, body string) {
// 	// 	log.Printf("Requesting %s %s", method, url)
// 	// 	log.Printf("Request body: %s", body)
// 	// }
// 	//
// 	// // add single go struct entity
// 	// response, _ := core.Index("twitter", "tweet", "3", nil, Tweet{"kimchy", "Search is cool"})
// 	//
// 	// b, _ := json.Marshal(response)
// 	// fmt.Println("twitter1 json:", string(b))
// 	//
// 	// log.Println("twitter1:", response)
// 	//
// 	// // you have bytes
// 	// tw := Tweet{"kimchy", "Search is cool part 2"}
// 	// bytesLine, _ := json.Marshal(tw)
// 	//
// 	// response2, _ := core.Index("twitter", "tweet", "4", nil, bytesLine)
// 	//
// 	// log.Println("twitter2:", response2)
// 	// b, _ = json.Marshal(response2)
// 	// fmt.Println("twitter2 json:", string(b))
// 	// end

// 	// indexer := core.NewBulkIndexerErrors(10, 60) // NewBulkIndexer(maxConns) //IndexBulk("twitter", "tweet", "3", &t, Tweet{"kimchy", "Search is now cooler"})
// 	// indexer := core.NewBulkIndexer(10) // NewBulkIndexer(maxConns) //IndexBulk("twitter", "tweet", "3", &t, Tweet{"kimchy", "Search is now cooler"})
// 	indexer := core.NewBulkIndexerErrors(10, 5)

// 	indexer.Start()

// 	// done := make(chan bool)
// 	// indexer.Run(done)

// 	// go func() {
// 	// 	for errBuf := range indexer.ErrorChannel {
// 	// 		// just blissfully print errors forever
// 	// 		fmt.Println(errBuf.Err)
// 	// 	}
// 	// }()
// 	for i := 0; i < 800; i++ {
// 		fmt.Println("try bulk index:", i)

// 		//indexer.Index("twitter", "user", strconv.Itoa(i), "", nil, `{"name":"bob"}`, false)
// 		date := time.Unix(1257894000, 0)
// 		data := map[string]interface{}{"name": "smurfs", "age": 22, "date": date}
// 		err := indexer.Index("users5", "user5", strconv.Itoa(i), "", "", &date, data) //不代表真的送出
// 		fmt.Println("after bulk index:", err)
// 	}

// 	// errorCt := 0
// 	// var errBuf *ErrorBuffer
// 	// for errBuf := range indexer.ErrorChannel {
// 	// 	log.Println("get bulk error:", errBuf)
// 	// 	break
// 	// }

// 	log.Println("all after bulk")

// 	// done <- true
// 	// Indexing might take a while. So make sure the program runs
// 	// a little longer when trying this in main.
// 	indexer.Stop()
// 	log.Println("all after bulk2")

// 	return nil
// }
