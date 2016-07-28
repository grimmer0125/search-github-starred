package main

import "log"

// package singleton

const (
	NOTSTART = "NotStart"
	FETCHING = "Fetching"
	INDEXING = "Indexing"
	INDEXED  = "Indexed"
	ERROR    = "Error"
)

type GitHubUser struct {
	account           string
	accessToken       string
	status            string
	numOfStarred      int
	indicesOfStarrerd int
}

// go getStarredInfo(account, token)

// _, err = getStarredInfo(profile.Nickname(), profile.Token().AccessToken)

// if err != nil {
// 	log.Println("cant not get starred info.")
// }

// func GetStarredInfo(tokenOwner, token string) (map[string]interface{}, error) {

func (user *GitHubUser) GetStarredInfo() {

	user.status = FETCHING
	repoList, _ := GetStarredInfo(user.account, user.accessToken)
	user.status = INDEXING
	len := len(repoList)
	user.numOfStarred = len
	log.Println("get number of starred:", user.numOfStarred)

	// b, _ := json.Marshal(&repoList)
	// fmt.Println("total repo:", string(b))

	//try to indexing
	err := SendToAlgolia(repoList, user.account)
	if err == nil {
		user.status = INDEXED
	}
}

// 一個表
// account  status  numOfStarred  indexOfStarred
