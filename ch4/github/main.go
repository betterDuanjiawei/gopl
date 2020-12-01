package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"
// 即使对应的 json字段名称不是首字母大写,结构体的成员名称也必须首字母大写,由于在 unmarshal 阶段,json字段的名称关联到 go 结构体成员的名称是忽略大小写的.因此这里只需要在 json 中由下划线而 go里面没有下划线的时候使用一下成员变量的标签定义
type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items []*Issue
}

type Issue struct {
	Number int
	HTMLURL string 	`json:"html_url"`
	Title string
	State string
	User *User
	CreatedAt time.Time `json:"create_at"`
	Body string
}

type User struct {
	Login string
	HTMLURL string `json:"html_url"`
}

func SearchIssues(terms []string) (*IssueSearchResult, error)  {
	q := url.QueryEscape(strings.Join(terms, " "))
	//fmt.Println(q)
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { // 流式解码器 json.NewDecoder,可以利用它来依次从字节流里解码出多个 json实体.
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return &result, nil
}