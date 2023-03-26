package repository

import (
	"fmt"
	"time"
)

type Post struct {
	Id          uint64
	Author      string
	Origin      string
	Title       string
	Content     string
	Link        string
	PublishedAt string
	CreatedAt   time.Time
}

func (s *Storer) duplicateCheckInDB(author, origin, title, content, link, publishedAt string) (*Post, error) {
	stmt := `SELECT id, author, origin, title, content, link, published_at, created_at
			 FROM post
			 WHERE origin=? AND title=? AND content=? AND link=? AND published_at=?;`

	row := s.db.QueryRow(stmt, origin, title, content, link, publishedAt)

	p := &Post{}
	err := row.Scan(&p.Id, &p.Author, &p.Origin, &p.Title, &p.Content, &p.Link, &p.PublishedAt, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *Storer) RetriveLatestPostInDB(origin string) (*Post, error) {
	stmt := `SELECT id, author, origin, title, content, link, published_at, created_at
			 FROM post
			 WHERE origin=?
			 ORDER BY id DESC LIMIT 1;`

	row := s.db.QueryRow(stmt, origin)

	p := &Post{}
	err := row.Scan(&p.Id, &p.Author, &p.Origin, &p.Title, &p.Content, &p.Link, &p.PublishedAt, &p.CreatedAt)
	if err != nil {
		return nil, err
	}

	return p, err
}

func (s *Storer) RetriveLatestPostsInDB(origin string, id, limit int) ([]*Post, error) {
	stmt := `SELECT id, author, origin, title, content, link, published_at, created_at
			 FROM post
			 WHERE origin=? AND id < ?
			 ORDER BY id DESC LIMIT ?;`

	rows, err := s.db.Query(stmt, origin, id, limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*Post{}

	for rows.Next() {
		p := &Post{}
		err := rows.Scan(&p.Id, &p.Author, &p.Origin, &p.Title, &p.Content, &p.Link, &p.PublishedAt, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *Storer) CreatePostInDB(author, origin, title, content, link, publishedAt string) (int, error) {
	post, err := s.duplicateCheckInDB(author, origin, title, content, link, publishedAt)
	if post != nil {
		fmt.Println("Duplicated data[Title:", post.Title, "]")
		return int(post.Id), nil
	}
	stmt := `INSERT INTO post (author, origin, title, content, link, published_at, created_at)
			VALUES(?, ?, ?, ?, ?, ?, UTC_TIMESTAMP())`

	result, err := s.db.Exec(stmt, author, origin, title, content, link, publishedAt, link)
	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}
