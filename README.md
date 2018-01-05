# 4d63.com/youtubelength
[![Build Status](https://img.shields.io/travis/leighmcculloch/go-youtubelength.svg)](https://travis-ci.org/leighmcculloch/go-youtubelength)
[![Codecov](https://img.shields.io/codecov/c/github/leighmcculloch/go-youtubelength.svg)](https://codecov.io/gh/leighmcculloch/go-youtubelength)
[![Go Report Card](https://goreportcard.com/badge/github.com/leighmcculloch/go-youtubelength)](https://goreportcard.com/report/github.com/leighmcculloch/go-youtubelength)
[![Go docs](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/4d63.com/youtubelength)

Go package that gets the length of Youtube videos.

## Usage

```go
import "4d63.com/youtubelength"
```

```go
length, err := youtubelength.Get(context.Background(), "G_OlRWGLdnw")
```

## Usage (App Engine)
```go
import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	"4d63.com/youtubelength"
)

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	httpClient := urlfetch.Client(c)
	ytlClient := youtubelength.Client{HTTPClient:httpClient}
	length, err := ytlClient.Get(c, "G_OlRWGLdnw")
	...
}
```
