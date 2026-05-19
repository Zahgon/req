package main

import (
	"fmt"

	"github.com/imroc/req/v3"
)

// Change the name if you want
var username = "imroc"

func main() {
	repo, star, err := findTheMostPopularRepo(username)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The most popular repo of %s is %s, which have %d stars\n", username, repo, star)
}

func init() {
	req.EnableDebugLog().
		EnableTraceAll().
		EnableDumpEachRequest().
		SetCommonErrorResult(&ErrorMessage{}).
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil {
				return nil
			}
			if errMsg, ok := resp.ErrorResult().(*ErrorMessage); ok {
				resp.Err = errMsg
				return nil
			}
			if !resp.IsSuccessState() {
				resp.Err = fmt.Errorf("bad status: %s\nraw content:\n%s", resp.Status, resp.Dump())
			}
			return nil
		})
}

type Repo struct {
	Name string `json:"name"`
	Star int    `json:"stargazers_count"`
}
type ErrorMessage struct {
	Message string `json:"message"`
}

func (msg *ErrorMessage) Error() string { _ = "STUB: not implemented"; return "" }

func findTheMostPopularRepo(username string) (repo string, star int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

//  HTTP status `code >= 200 and <= 299` is considered as success by default

// Try Next page

// All repos have been traversed, return the final result
