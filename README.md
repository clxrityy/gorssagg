# gorssagg

Go RSS feed aggregator authenticated with API keys

---

```zsh
git clone https://github.com/clxrityy/gorssagg.git
```
```zsh
go mod vendor && go mod tidy
```
```zsh
go build && ./gorssagg
```

## Requirements

- **PostgreSQL** server & client (view [here](/sql/README.md) for the guide)
    - [`sqlc`](https://sqlc.dev/)
    - [`goose`](https://github.com/c9s/goose)

---

## [`.env`](/.env.example)

```.env
PORT=8080
DB_URL=postgres://postgres:@localhost:5432/gorssagg?sslmode=disable
```
---

## Routes

- `/v1` - base route

```go
import (
	//...
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	//...
)
//...
func main() {
    //...
    router := chi.NewRouter();
    router.Use(cors.Handler(cors.Options{
        //...
    }))

    v1Router := chi.NewRouter();
    router.Mount("v1", v1Router);
}
```

- ##### `/v1/healthz` - test route (responds with `200` if working)
- ##### `/v1/error` - error route (responds with `400`)
- ##### `/v1/users` - users route
    - **POST** - creates a user, takes a `name` parameter
        - responds with a JSON object that should contain the user's `api_key`
    - **GET** - get a user by their API key
        - pass in the `Authorization: ApiKey <api_key>` header
- ##### `/v1/feeds` - RSS feeds route
    - **POST** - creates a feed, takes a `name` & `url` parameter (the URL should point to some sort of `/index.xml`)
        - responds with `201` if successful
        - pass in the `Authorization: ApiKey <api_key>` header
    - **GET** - returns an array of feeds (no parameters)
- ##### `/v1/feed_follows` - follow a feed
    - **POST** - follows a feed, takes a `feed_id` parameter
        - pass in the `Authorization: ApiKey <api_key>` header
        - responds with `201` if successful
    - **GET** - gets the feed(s) a user is following
        - pass in the `Authorization: ApiKey <api_key>` header
    - **DELETE** - `/v1/feed_follows/{feed_follow_id}` - unfollow a feed
        - pass in the `Authorization: ApiKey <api_key>` header

---