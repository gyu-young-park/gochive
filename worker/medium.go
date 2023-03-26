package worker

import (
	"fmt"
	"github/gyu-young-park/go-archive/repository"

	"github.com/gocolly/colly/v2"
)

type mediumWork struct {
	crawler  *colly.Collector
	postList []repository.Post
}

func newMediumWork() *mediumWork {
	m := &mediumWork{
		crawler: colly.NewCollector(),
	}
	m.ready()
	return m
}

func (m *mediumWork) ready() {
	m.crawler.OnHTML("article", func(e *colly.HTMLElement) {
		post := repository.Post{}
		//postingTime := e.ChildTexts("p")[1]
		post.Author = e.ChildTexts("p")[0]
		post.Link = makeMediumLink(e.ChildAttrs("a", "href")[2])
		post.Title = e.ChildText("h2")
		m.postList = append(m.postList, post)
	})
}

func (m *mediumWork) do(store *repository.Storer) {
	m.crawler.Visit("https://medium.com/tag/go/latest")
	for _, post := range m.postList {
		id, err := store.CreatePostInDB(
			post.Author,
			MEDIUM,
			post.Title,
			post.Content,
			post.Link,
			post.PublishedAt,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Create Post: ", id)
	}
}

func makeMediumLink(link string) string {
	return fmt.Sprintf("%s%s", MEDIUM_ROOT_ADDRESS, link)
}
