package main

import (
	"fmt"
	"log"
)

// package singleton

const (
	NOTSTART = "NotStart"
	FETCHING = "Fetching"
	INDEXING = "Indexing"
	INDEXED  = "Indexed"
	ERROR    = "Error"
)

type GitHubUser struct {
	Account string `json:"account"`
	// AccessToken       string `json:"accessToken"`
	Tokens            []string `json:"tokens"`
	Status            string   `json:"status"`
	NumOfStarred      int
	IndicesOfStarrerd int
}

// go getStarredInfo(account, token)

// _, err = getStarredInfo(profile.Nickname(), profile.Token().AccessToken)

// if err != nil {
// 	log.Println("cant not get starred info.")
// }

// func GetStarredInfo(tokenOwner, token string) (map[string]interface{}, error) {

func (user *GitHubUser) GetStarredInfo(token string) {

	user.Status = FETCHING
	SetUser(user.Account, *user)

	repoList, _ := GetStarredInfo(user.Account, token)

	user.Status = INDEXING
	len := len(repoList)
	user.NumOfStarred = len
	SetUser(user.Account, *user)

	log.Println("get number of starred:", user.NumOfStarred)

	// b, _ := json.Marshal(&repoList)
	// fmt.Println("total repo:", string(b))

	// try to indexing
	err := SendToGitHubElasticsearch(repoList, user.Account)
	fmt.Println("after send")

	if err == nil {
		user.Status = INDEXED
		SetUser(user.Account, *user)
	}
}

// 一個表
// account  status  numOfStarred  indexOfStarred
