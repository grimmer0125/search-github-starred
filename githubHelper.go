package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
)

func queryAlgolia(queryStr, starredBy string) {
	client := algoliasearch.NewClient("EQDRH6QSH7", "6066c3e492d3a35cc0a425175afa89ff")

	index := client.InitIndex("githubRepo")

	params := algoliasearch.Map{
		"attributesToSnippet": []string{"description:40"},
		"facetFilters":        "starredBy" + starredBy, // "firstname:Jimmie",
	}

	res, err := index.Search(queryStr, params)

	if err != nil {
		log.Println("error:", err)
	}

	b, err := json.Marshal(res)
	fmt.Println("search result:", string(b))
}

func sendToAlgolia(repoList []*algoliasearch.Object) {

	client := algoliasearch.NewClient("EQDRH6QSH7", "6066c3e492d3a35cc0a425175afa89ff")
	index := client.InitIndex("githubRepo")

	setting := make(map[string]interface{})
	setting["attributesForFaceting"] = []string{"starredBy"}
	index.SetSettings(setting)

	// object1 := algoliasearch.Object{
	// 	"firstname": "Jimmie apple tree",
	// 	"lastname":  "Barninger",
	// }
	//
	// object2 := algoliasearch.Object{
	// 	"firstname": "Jimmie",
	// 	"lastname":  "Barninger",
	// }

	// content, _ := ioutil.ReadFile("contacts.json")
	var objects []algoliasearch.Object
	// // if err := json.Unmarshal(content, &objects); err != nil {
	// // 	return
	// // }
	// objects = append(objects, object1)
	// objects = append(objects, object2)
	//

	for _, object := range repoList {
		objects = append(objects, *object)
	}

	_, err := index.AddObjects(objects)
	if err != nil {
		log.Println("add to algolia error", err)
	}
	log.Println("add to algolia ok")
}

//  string list
//->

func getRepoReadme(token string, repoList []*algoliasearch.Object) (map[string]interface{}, error) {

	log.Println("repo list len:", len(repoList))
	// log.Println("total:", repoURLList)
	// https://github.com/mhart/react-server-example

	// headers: self.authorizedRequestHeaders(with: ["Accept": "application/vnd.github.raw"]))
	//  "full_name": "sandstorm-io/sandstorm",
	//grimmer0125
	//repo name
	//get { return NSURL(string: "https://api.github.com/repos/\(self.ownerName)/\(self.name)/readme") }

	for i := 0; i < 5; i++ {
		repo := *repoList[i]
		readmeURL := repo["repoURL"].(string) + "/readme"

		req, err := http.NewRequest("GET", readmeURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Accept", "application/vnd.github.raw")

		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			return nil, err
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		// log.Println("readmd body:", string(b))
		repo["readmd"] = string(b)

		res.Body.Close()
	}

	sendToAlgolia(repoList)
	return nil, nil
}

func getStarredInfo(tokenOwner, token string) (map[string]interface{}, error) {

	log.Println("token:", token)

	var repoList []*algoliasearch.Object

	for ifRun, pageIndex := true, 1; ifRun == true; pageIndex++ {
		pageStr := strconv.Itoa(pageIndex)

		//?per_page=50
		starredURL := "https://api.github.com/user/starred?per_page=100&page=" + pageStr

		log.Println("start to query:", starredURL)
		req, err := http.NewRequest("GET", starredURL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		c := http.Client{}
		res, err := c.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		// ctype, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
		// if err != nil {
		// 	return nil, err
		// }

		// limit := res.Header["X-RateLimit-Limit"]
		// log.Println("hearders:", res.Header)

		linkList := res.Header["Link"]
		log.Println("link:", linkList)

		ifRun = false
		for _, link := range linkList {
			i := strings.Index(link, "last")

			if i > -1 {
				ifRun = true
				break
			}
		}

		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		//json array

		// too many
		// log.Println("response body:", string(b))
		// type mytype []map[string]string

		var repoOrigList []algoliasearch.Object

		if err := json.Unmarshal(b, &repoOrigList); err != nil {
			log.Println("can not parse repo list")
		}
		// fmt.Println("repo", repoList)

		for _, repo := range repoOrigList {

			object := algoliasearch.Object{
				"repoURL":     repo["url"],
				"repoName":    repo["name"],
				"ownerName":   repo["owner"].(map[string]interface{})["login"],
				"ownerURL":    repo["html_url"],
				"starredBy":   tokenOwner,
				"description": repo["description"],
			}
			// repo["url"].(string)
			repoList = append(repoList, &object)
		}
		// "description": "Golang标准库。对于程序员而言，标准库与语言本身同样重要，它好比一个百宝箱，能为各种常见的任务提供完美的解决方案。以示例驱动的方式讲解Golang的标准库。",
		// "homepage"

		// switch ctype {
		// case "application/json", "text/javascript":
		// 	var data map[string]interface{}
		// 	json.Unmarshal(b, &data)
		// 	return data, nil
		// }
	}

	getRepoReadme(token, repoList)

	// link := "Link: <https://api.github.com/user/5940941/starred?per_page=50&page=2>; rel=\"next\", <https://api.github.com/user/5940941/starred?per_page=50&page=5>; rel=\"last\""

	//1 link:[<https://api.github.com/user/starred?page=2>; rel="next", <https://api.github.com/user/starred?page=8>; rel="last"]

	//8 link:[<https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=7>; rel="prev"]
	//4 <https://api.github.com/user/starred?per_page=30&page=5>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last"

	// first, prev都寫了
	// link:[<https://api.github.com/user/starred?per_page=30&page=8>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last", <https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=6>; rel="prev"]

	return nil, nil //map[string]interface{}{}, fmt.Errorf("unknown Content-Type: %v", ctype)
}

// url
// description ?
// readme ?
// owner - html_url
// name (repo-name)
// owner -login (owner-login)
// starredBy - (原始的那個人)

// type GitHubRepo struct {
// 	URL       string `json:"url"`
// 	Name      string `json:"name"`
// 	OwnerName string `json:"ownerName"`
// 	OwnerURL  string `json:"ownerURL"`
// 	StarredBy string `json:"starredBy"`
// }
