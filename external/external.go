package external

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Post struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func GetPosts() []Post {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		log.Fatalln(err)
	}

	return posts
}
