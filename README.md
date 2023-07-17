<h1 align="center">
    Dwarves Golang Threads API
</h1>
<p align="center">
    <a href="https://github.com/dwarvesf">
        <img src="https://img.shields.io/badge/-make%20by%20dwarves-%23e13f5e?style=for-the-badge&logo=data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACwAAAAsBAMAAADsqkcyAAAAD1BMVEUAAAD///////////////+PQt5oAAAABXRSTlMAQL//gOnhmfMAAAAJcEhZcwAAHsIAAB7CAW7QdT4AAACYSURBVHicndLRDYJAEIThMbGAI1qAYAO6bAGXYP81uSGBk+O/h3Mev4dhWJCkYZqreOi1xoh0eSIvoCaBRjc1B9+I31g9Z2aJ5jkOsYScBW8zDerO/fObnY/FiTl3caOEH2nMzpyZhezIlgqXr2OlOX617Up/nHnPUg0+LHl18YO50d3ghOy1ioeIq1ceTypsjpvYeJohfQEE5WtH+OEYkwAAAABJRU5ErkJggg==&&logoColor=white" alt="Dwarves Foundation" />
    </a>
    <a href="https://discord.gg/dwarvesv">
        <img src="https://img.shields.io/badge/-join%20the%20community-%235865F2?style=for-the-badge&logo=discord&&logoColor=white" alt="Dwarves Foundation Discord" />
    </a>
</p>

Unofficial, Reverse-Engineered Golang client for Meta's Threads. Supports Read and Write.

## Getting started
How to install
Install the library with the following command using go module:

```
$ go get github.com/dwarvesf/go-threads
```

Examples
Find examples of how to use the library in the examples folder:

```
ls examples
├── create_post
│   └── main.go
...
```

## API

### Disclaimer

The Threads API is a public API in the library, requiring no authorization. In contrast, the Instagram API, referred to as the private API, requires the Instagram username and password for interaction.

The public API offers read-only endpoints, while the private API provides both read and write endpoints. The private API is generally more stable as Instagram is a reliable product.

Using the public API reduces the risk of rate limits or account suspension. However, there is a trade-off between stability, bugs, rate limits, and suspension. The library allows for combining approaches, such as using the public API for read-only tasks and the private API for write operations. A retry mechanism can also be implemented, attempting the public API first and then falling back to the private API if necessary.

### Initialization

To start using the `GoThreads` package, import the relevant class for communication with the Threads API and create an instance of the object.

For utilizing only the public API, use the following code snippet:

```go
import (
    "github.com/dwarvesf/go-threads"
)

func main() {
    th := threads.NewThreads()
    th.GetThreadLikers(<thread_id>)

    // Using global instance
    threads.GetThreadLikers(<thread_id>)
}
```

If you intend to use the private API exclusively or both the private and public APIs, utilize the following code snippet:

```go
package main

import (
	"fmt"

	"github.com/dwarvesf/go-threads"
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
}
```

Or the shorter syntax

```go
package main

import (
	"fmt"

	"github.com/dwarvesf/go-threads"
	"github.com/dwarvesf/go-threads/model"
)

func main() {
	client, err := threads.InitAPIClient(
		threads.WithDoLogin("instagram_username", "instagram_password"),
	)
	if err != nil {
		fmt.Println("unable init API client", err)
		return
	}

	p, err := client.CreatePost(model.CreatePostRequest{Caption: "new post"})
	if err != nil {
		fmt.Println("unable create a post", err)
		return
	}
}
```

To mitigate the risk of blocking our users, an alternative initialization method can be implemented for the client. This method entails storing the API token and device token, which are subsequently utilized for initializing the API client.

```go
package main

import (
	"fmt"

	"github.com/dwarvesf/go-threads"
	"github.com/dwarvesf/go-threads/model"
)

func main() {
	client, err := threads.InitAPIClient(
		threads.WithCridential("instagram_username", "instagram_password"),
		threads.WithAPIToken("device_id", "api_token"),
	)
	if err != nil {
		fmt.Println("unable init API client", err)
		return
	}

	p, err := client.CreatePost(model.CreatePostRequest{Caption: "new post"})
	if err != nil {
		fmt.Println("unable create a post", err)
		return
	}
}
```

### Public API

`Coming soon`

### Private API

`Coming soon`

## Road map

- [ ] Improve the perfomance
- [ ] Listing API
- [ ] Mutation API

