# TaskManager-GoLang

A simple RESTful API for managing tasks, built with Go, Gorilla Mux, and MongoDB.

## Features

- List all tasks
- Create a new task
- Update an existing task
- Delete a task
- MongoDB integration for persistent storage

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) installed
- [MongoDB](https://www.mongodb.com/try/download/community) running locally or remotely

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/ffahriza/TaskManager-GoLang.git
    cd TaskManager-GoLang
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Configure MongoDB connection in `db/db.go` if needed.

### Running the Server

```sh
go run main.go
```

The server will start at [http://localhost:8000](http://localhost:8000).

## API Endpoints

| Method | Endpoint            | Description           |
|--------|---------------------|----------------------|
| GET    | `/tasks`            | List all tasks       |
| POST   | `/tasks`            | Create a new task    |
| PUT    | `/tasks/{id}`       | Update a task        |
| DELETE | `/tasks/{id}`       | Delete a task        |

## Example Task Object

```json
{
  "id": "1",
  "title": "Learn Go",
  "done": false
}
```

## Contributing

Feel free to fork the repo and submit pull requests!

## License

This project is licensed under the MIT License.

---

**Made by Andra**
