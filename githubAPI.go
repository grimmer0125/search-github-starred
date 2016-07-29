package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetReadme(token string, repoList []*map[string]interface{}, j int, channel chan int) {

	// log.Println("debug log:", j)
	repo := *repoList[j]
	readmeURL := repo["apiURL"].(string) + "/readme"

	log.Println("try to get readme:", readmeURL)
	req, err := http.NewRequest("GET", readmeURL, nil)
	if err != nil {
		log.Println("new request error :", err)
		channel <- j
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Accept", "application/vnd.github.raw")

	c := http.Client{}
	res, err := c.Do(req)
	if err != nil {
		log.Println("res error to readme:", err)

		channel <- j
		return
	}

	status := res.Header.Get("Status")
	if status == "404 Not Found" {
		log.Println("404 Not Found")
		// 	body: {"message":"Not Found","documentation_url":"https://developer.github.com/v3"}

	} else {
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("read body error:", err)
		} else {
			repo["readmd"] = string(b)
		}

		res.Body.Close()
	}

	channel <- j

}

func GetReposReadme(token string, repoList []*map[string]interface{}) error {

	lenList := len(repoList)

	log.Println("try getting all readme:", lenList)

	c := make(chan int, lenList)
	checkList := make([]int, lenList)

	for i := 0; i < lenList; i++ {
		go GetReadme(token, repoList, i, c)
	}

	for i := range c {

		checkList[i] = 1
		allGet := true

		for j := 0; j < len(checkList); j++ {
			if checkList[j] == 0 {
				allGet = false
				break
			}
		}

		if allGet == true {
			log.Println("check list:get all")
			close(c)
		} else {
			// log.Println("check list:not get all")
		}
	}

	log.Println("after getting all readme")

	return nil
}

// link := "Link: <https://api.github.com/user/5940941/starred?per_page=50&page=2>; rel=\"next\", <https://api.github.com/user/5940941/starred?per_page=50&page=5>; rel=\"last\""

//1 link:[<https://api.github.com/user/starred?page=2>; rel="next", <https://api.github.com/user/starred?page=8>; rel="last"]

//8 link:[<https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=7>; rel="prev"]
//4 <https://api.github.com/user/starred?per_page=30&page=5>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last"

// first, prev都寫了
// link:[<https://api.github.com/user/starred?per_page=30&page=8>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last", <https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=6>; rel="prev"]
func GetStarredInfo(tokenOwner, token string) ([]*map[string]interface{}, error) {

	log.Println("token:", token)

	var repoList []*map[string]interface{}

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

		var repoOrigList []map[string]interface{}

		if err := json.Unmarshal(b, &repoOrigList); err != nil {
			log.Println("can not parse repo list")
		}

		for _, repo := range repoOrigList {

			object := map[string]interface{}{
				"apiURL":        repo["url"],
				"repoURL":       repo["html_url"],
				"repoName":      repo["name"],
				"repofull_name": repo["full_name"], //no use now
				// "ownerName":     repo["owner"].(map[string]interface{})["login"],    //no use now
				// "ownerURL":      repo["owner"].(map[string]interface{})["html_url"], // no use now
				"starredBy":   tokenOwner,
				"description": repo["description"],
				"homepage":    repo["homepage"],
			}

			repoList = append(repoList, &object)
		}
	}

	GetReposReadme(token, repoList)

	return repoList, nil //map[string]interface{}{}, fmt.Errorf("unknown Content-Type: %v", ctype)
}

// ctype, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
// if err != nil {
// 	return nil, err
// }

// switch ctype {
// case "application/json", "text/javascript":
// 	var data map[string]interface{}
// 	json.Unmarshal(b, &data)
// 	return data, nil
// }
