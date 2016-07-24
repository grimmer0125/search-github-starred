package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func getRepoReadme(token string, repoURLList []string) (map[string]interface{}, error) {

	log.Println("repo list len:", len(repoURLList))
	// log.Println("total:", repoURLList)
	// https://github.com/mhart/react-server-example

	// headers: self.authorizedRequestHeaders(with: ["Accept": "application/vnd.github.raw"]))
	//  "full_name": "sandstorm-io/sandstorm",
	//grimmer0125
	//repo name
	//get { return NSURL(string: "https://api.github.com/repos/\(self.ownerName)/\(self.name)/readme") }

	for i := 0; i < 5; i++ {
		readmeURL := repoURLList[i] + "/readme"

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

		log.Println("readmd body:", string(b))

		res.Body.Close()
	}

	return nil, nil
}

func getStarredInfo(token string) (map[string]interface{}, error) {

	// 1
	//   fmt.Println(strings.ContainsAny("Hello World", ",|"))
	//
	// 2
	//   if strings.IndexFunc("HelloWorld", f) != -1 {
	//
	// 3
	//     x := "chars@arefun"
	//
	//      i := strings.Index(x, "@")

	var repoURLList []string

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
			log.Println("i:", i)

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

		var repoList []map[string]interface{}

		if err := json.Unmarshal(b, &repoList); err != nil {
			log.Println("can not parse repo list")
		}
		// fmt.Println("repo", repoList)

		for _, repo := range repoList {
			repoURLList = append(repoURLList, repo["url"].(string))
		}

		// switch ctype {
		// case "application/json", "text/javascript":
		// 	var data map[string]interface{}
		// 	json.Unmarshal(b, &data)
		// 	return data, nil
		// }
	}

	getRepoReadme(token, repoURLList)

	// link := "Link: <https://api.github.com/user/5940941/starred?per_page=50&page=2>; rel=\"next\", <https://api.github.com/user/5940941/starred?per_page=50&page=5>; rel=\"last\""

	//1 link:[<https://api.github.com/user/starred?page=2>; rel="next", <https://api.github.com/user/starred?page=8>; rel="last"]

	//8 link:[<https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=7>; rel="prev"]
	//4 <https://api.github.com/user/starred?per_page=30&page=5>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last"

	// first, prev都寫了
	// link:[<https://api.github.com/user/starred?per_page=30&page=8>; rel="next", <https://api.github.com/user/starred?per_page=30&page=8>; rel="last", <https://api.github.com/user/starred?per_page=30&page=1>; rel="first", <https://api.github.com/user/starred?per_page=30&page=6>; rel="prev"]

	return nil, nil //map[string]interface{}{}, fmt.Errorf("unknown Content-Type: %v", ctype)
}
