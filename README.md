<h1 align="center">JsonHandlers</h1>

<div align="center" id="top"> 
  <img src="https://res.cloudinary.com/dahkenlmo/image/upload/v1678346700/jsonhandlers_hjca2i.png" alt="Json Handlers" />
  &#xa0;
</div>


<p align="center">JSON library to expose simple handlers that lets you easily read and write json from various sources.</p>


[![Build Status](https://github.com/abusomani/jsonhandlers/workflows/build/badge.svg)](https://github.com/abusomani/jsonhandlers/actions)
[![Github top language](https://img.shields.io/github/languages/top/abusomani/jsonhandlers)](https://img.shields.io/github/languages/top/abusomani/jsonhandlers)
[![Github language count](https://img.shields.io/github/languages/count/abusomani/jsonhandlers)](https://img.shields.io/github/languages/count/abusomani/jsonhandlers)
[![License](https://img.shields.io/badge/license-MIT-blue)](https://github.com/abusomani/jsonhandlers/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/abusomani/jsonhandlers)](https://goreportcard.com/report/github.com/abusomani/jsonhandlers)
[![Go Reference](https://pkg.go.dev/badge/github.com/abusomani/jsonhandlers.svg)](https://pkg.go.dev/github.com/abusomani/jsonhandlers)
[![Repo size](https://img.shields.io/github/repo-size/abusomani/jsonhandlers)](https://shields.io/github/repo-size/abusomani/jsonhandlers)
[![Coverage Status](https://coveralls.io/repos/github/abusomani/jsonhandlers/badge.svg?branch=main)](https://coveralls.io/github/abusomani/jsonhandlers?branch=main)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go#json)


## Prerequisites

A go module where you want to integrate jsonhandlers. To create one, follow this [guide](https://go.dev/doc/tutorial/create-module).

## Installation

```
go get github.com/abusomani/jsonhandlers
```


## Usage

A very useful feature of Go’s import statement are aliases. A common use case for import aliases is to provide a shorter alternative to a library’s package name.

In this example, we save ourselves having to type `jsonhandlers` everytime we want to call one of the library’s functions, we just use `jh` instead.

```
import (
    jh "github.com/abusomani/jsonhandlers"
)
```

## Options

Jsonhandlers package exposes multiple options while creating a new `jsonhandler` to be able to read/write json from sources like Files, Http Requests or Http responses. 


### WithFileHandler

You can use the `WithFileHandler` option to read/write Json from/to a file. For this, you need to create a new jsonhandler with the file handler option.

[Example](./example/operations/file_handling.go) to understand `WithFileHandler` in more detail.

**Sample Code**
```go
package operations

import (
	"fmt"

	"github.com/abusomani/jsonhandlers"
)


func handleFile() {
	jh := jsonhandlers.New(jsonhandlers.WithFileHandler(testFilePath))

	var sch school
	err := jh.Unmarshal(&sch)
}

```

### WithHTTPRequestHandler

You can use the `WithHTTPRequestHandler` option to read Json from a Http Request and to write Json to a Http ResponseWriter. For this, you need to create a new jsonhandler with the Http request handler option.

[Example](./example/operations/http_request_handling.go) to understand `WithHTTPRequestHandler` in more detail.

**Sample Code**
```go
package operations

import (
	"net/http"

	"github.com/abusomani/jsonhandlers"
)

type studentSearchRequest struct {
	Name string
}

type studentSearchResponse struct {
	Info student
}

func HandleHTTPRequest(students []student) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
    	jh := jsonhandlers.New(jsonhandlers.WithHTTPRequestHandler(w, r))

		var reqBody studentSearchRequest
		_ := jh.Unmarshal(&reqBody)
		

		for _, student := range students {
			// student found
			if student.Name == reqBody.Name {
				
				// write the response using jh.Marshal
				jh.Marshal(studentSearchResponse{
					Info: student,
				})
				return
			}
		}
	})
}

/*
  Sample request to be hit on the localhost server to test WithHTTPRequestHandler functionality.
  curl http://localhost:8080/search -d '{"Name": "Abhishek Somani"}'
*/
```

### WithHTTPResponseHandler

You can use the `WithHTTPResponseHandler` option to read/write Json from/to a Http Response. For this, you need to create a new jsonhandler with the Http response handler option.

[Example](./example/operations/http_response_handling.go) to understand `WithHTTPResponseHandler` in more detail.

**Sample Code**

```go
package operations

import (
	"fmt"
	"log"
	"net/http"

	"github.com/abusomani/jsonhandlers"
)

type user struct {
	Id        int
	FirstName string
	LastName  string
}

type getUsersResponse struct {
	Users []user
}

func HandleHTTPResponse() {
	resp, _ := http.Get("https://dummyjson.com/users")
	jh := jsonhandlers.New(jsonhandlers.WithHTTPResponseHandler(resp))

	var userResp getUsersResponse
	jh.Unmarshal(&userResp)
}

```

## Run examples
To run the examples present in the [example](./example/) folder you need to first checkout this package by doing a `git clone`. Once you have checked out this package, then you can run the [main.go](./example/main.go) using the following command to see all the examples in action:

```
go run example/main.go
```


## License
Licensed under [MIT](./LICENSE)