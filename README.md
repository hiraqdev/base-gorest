# base-gorest

Skeleton to create rest api using golang

# Stacks

- Package Manager : [dep](https://github.com/golang/dep)
- Routing : [gorilla/mux](https://github.com/gorilla/mux)
- Middlewares : [gorilla/handlers](github.com/gorilla/handlers)
- Cmd Framework: [cobra](https://github.com/spf13/cobra)
- Configuration: [viper](https://github.com/spf13/viper)

# How To Use

If you want to use this skeleton, you just need to clone this repo
and dont forget to change all required path `github.com/hiraqdev/base-gorest`
to your local development path.

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

You have to install dep!
```

```
dep ensure
```

# Routes

All available routes should be registered at: `app/routes.go` and placed at `app/modules`

You can list all of your routers from command line:

```
go run *.go routes
```

And an output will look like:

```
+-------+-------------+
| PATH  | HTTP METHOD |
+-------+-------------+
| /ping | GET         |
+-------+-------------+
```

# Command Line

```go
go run *.go -h
```

```go
go run *.go server --addr=":8080"
```

```go
go run *.go server --writeTimeout=20 --readTimeout=30 --addr=":8080"
```

# Build

```
go build -o gorest
```
