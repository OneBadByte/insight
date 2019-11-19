package database

import (
	"context"
	"log"
	"time"
)

type Post struct {
	Id        int64     `json:"id"`
	Username  string    `json:"username"`
	Mood      string    `json:"mood"`
	Genre     string    `json:"genre"`
	Post      string    `json:"post"`
	Timestamp time.Time `json:"timestamp"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

const (
	GET_ALL_POSTS = "SELECT posts.id, users.username, mood, genre, post, time_stamp FROM posts INNER JOIN users ON posts.user_id = users.id ORDER BY posts.id DESC;"
	ADD_POST      = "INSERT INTO posts VALUES(DEFAULT, $1, $2, $3, $4);"
)

func (dbConn *DatabaseConnection) GetAllPosts() Posts {
	var posts Posts
	rows, err := dbConn.conn.Query(context.Background(), GET_ALL_POSTS)
	if err != nil {
		log.Println(err.Error())
		return posts
	}

	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.Id, &post.Username,
			&post.Mood, &post.Genre,
			&post.Post, &post.Timestamp)
		if err != nil {
			log.Println(err.Error())
			return posts
		}
		posts.Posts = append(posts.Posts, post)
	}
	return posts
}

func (dbConn *DatabaseConnection) AddPost(post Post) error {
	_, err := dbConn.conn.Exec(context.Background(), ADD_POST,
		dbConn.GetIdFromUsername(post.Username), post.Mood, post.Genre, post.Post)
	return err
}
