package worker

import (
	"encoding/json"
	"fmt"
	"github/gyu-young-park/go-archive/repository"
	"io/ioutil"
	"net/http"
	"time"
	"unicode/utf8"
)

type redditWork struct {
}

func NewredditWork() *redditWork {
	r := &redditWork{}
	r.ready()
	return r
}

func (r *redditWork) ready() {
}

func (r *redditWork) do(store *repository.Storer) {
	url := makeTopNURL(GOLANG, 1, REDDIT_DAY)
	fmt.Println(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("User-Agent", "GOCHIVE_BOT")
	client := makeClientWithTimeout(5)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var postModel RedditTopPostListModel
	err = json.Unmarshal(data, &postModel)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(postModel.Data.Children) == 0 {
		fmt.Println(string(data))
		fmt.Println("error: There is no children")
		return
	}
	author := postModel.Data.Children[0].Data.Author
	title := postModel.Data.Children[0].Data.Title
	content := makeContentShort(postModel.Data.Children[0].Data.Selftext)
	seconds := postModel.Data.Children[0].Data.Created
	link := postModel.Data.Children[0].Data.URL
	t := time.Unix(int64(seconds), 0)
	publishedAt := t.UTC()

	fmt.Println(postModel.Data.Children[0].Data.Author)
	fmt.Println(postModel.Data.Children[0].Data.Title)
	fmt.Println(postModel.Data.Children[0].Data.Selftext)
	fmt.Println(publishedAt)

	_, err = store.CreatePostInDB(author, REDDIT, title, content, link, publishedAt.String())
	if err != nil {
		fmt.Println(err)
		return
	}
}

func makeTopNURL(subreddt string, limit int, t redditTime) string {
	return fmt.Sprintf(REDDIT_POST_TOP, subreddt, limit, t)
}

func makeContentShort(content string) string {
	if utf8.RuneCountInString(content) >= 200 {
		return content[:190] + "..."
	}
	return content
}

func makeClientWithTimeout(waitTime int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(waitTime) * time.Second,
	}
}
