
# Unit Converter

A simple web-based unit converter that allows users to convert between different units of length, weight, and temperature. The backend is built in Go, while the frontend is a static HTML, CSS, and JavaScript application served by the Go server.

Project URL: https://roadmap.sh/projects/unit-converter


## Deployment

To deploy this project run

```bash
  go build -o unit-converter cmd/main.go
```



## Features

- Convert between various units of length, weight, and temperature.
- RESTful API for handling unit conversion requests.
- Simple, user-friendly interface.
- Go and Frontend integration


## 🛠 Skills
Programming Languages: Go, Python, JavaScript

Frontend: HTML, CSS

Backend: REST APIs, Go net/http


## Installation

1.Clone the repository

```bash
  git clone https://github.com/r3iwan/unit-converter.git
  cd unit-converter
```

2.Install dependencies

```bash
  go mod tidy
```

3.Run the server

```bash
  go run cmd/main.go
```
    