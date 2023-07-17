package main

import (
	"fmt"

	"github.com/dwarvesf/go-threads"
	"github.com/dwarvesf/go-threads/model"
)

func main() {
	cfg, err := threads.InitConfig(
		threads.WithDoLogin("instagram_username", "instagram_password"),
	)

	if err != nil {
		fmt.Println("unable init config", err)
		return
	}

	client, err := threads.NewPrivateAPIClient(cfg)

	if err != nil {
		fmt.Println("unable init API client", err)
		return
	}

	p, err := client.CreatePost(model.CreatePostRequest{Caption: "hello threads"})
	if err != nil {
		fmt.Println("unable create a post", err)
		return
	}

	fmt.Println(p.Media.Pk)
}
