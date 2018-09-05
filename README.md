# base-gorest

Skeleton to create rest api using golang

# Stacks

- Package Manager : [glide](http://glide.sh/)
- Routing : [gorilla/mux](https://github.com/gorilla/mux)
- Middlewares : [gorilla/handlers](github.com/gorilla/handlers)
- Cmd Framework: [cobra](https://github.com/spf13/cobra)
- Configuration: [viper](https://github.com/spf13/viper)

# Configurations

We have two configurations:

- Server
- Application

For server side, we only have:

- addr : ip and port
- readTimeout
- writeTimeout

And all of these configured via command line when we run our api engine.

For the application side, configured via environment variables and managed
using Viper.  All key variables should prefixed with "GOREST_" , example :

```go
GOREST_TEST go run *.go
```

When using this engine with docker, you can use dotenv file and use parameter
`--env-file`, it will automatically inject all key and values to environment
variables inside container.

# Standard

- Golang Code Standard
- REST API: [JSONAPI](http://jsonapi.org)

# Setup

```
WARNING

You have to install glide!
```

```
glide up
```

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
