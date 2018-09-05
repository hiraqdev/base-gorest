# base-gorest

Skeleton to create rest api using golang

# Stacks

- Package Manager : [glide](http://glide.sh/)
- Routing : [gorilla/mux](https://github.com/gorilla/mux)
- Middlewares : [gorilla/handlers](github.com/gorilla/handlers)
- Cmd Framework: [cobra](https://github.com/spf13/cobra)

# Standard

- Golang Code Standard
- REST API: [JSONAPI](http://jsonapi.org)

# Routes

All available routes should be registered at: `app/routes.go` and placed at `app/modules`

# Command Line

```go
go run *.go -h
```

```go
go run *.go --addr=":8080"
```

```go
go run *.go --writeTimeout=20 --readTimeout=30 --addr=":8080"
```