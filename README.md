<p align="center">
  <img width="250" height="200" src="readme-icon.png">
</p>

[![Go Report Card](https://goreportcard.com/badge/github.com/kushalshit27/go-rest-api)](https://goreportcard.com/report/github.com/kushalshit27/go-rest-api)
<!-- GETTING STARTED -->
# REST API with Golang
A RESTful API example for simple manage post application with Go


## Getting Started

To get a local copy up and running follow these simple steps.

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/kushalshit27/go-rest-api.git
   ```
2. Configuration
   ```sh
   configure .env
   ```

### Build and Run
```bash
# Build and Run
cd go-rest-api
go mod download
go build
./go-rest-api

# API Endpoint : API running on: http://127.0.0.1:8080
```
## API

#### /api/posts
* `GET` : Get all posts
* `POST` : Create a new posts

#### /api/posts/:id
* `GET` : Get a post
* `PUT` : Update a post
* `DELETE` : Delete a post

## Roadmap

See the [open issues](https://github.com/kushalshit27/go-rest-api/issues) for a list of proposed features (and known issues).


## License

Distributed under the MIT License. See `LICENSE` for more information.
