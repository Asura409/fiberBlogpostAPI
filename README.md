
# 📝 Simple Blog Post API with GoFiber #

This is a basic RESTful API built using [Go](https://golang.org/) and the [Fiber](https://gofiber.io/) web framework. It provides CRUD operations for managing blog posts using an in-memory fake database (slice of structs).

---

## 🚀 Features ##

- Add a new post
- Edit an existing post
- Get all posts
- Get a post by ID
- Delete a post

---

## 📦 Tech Stack ##

- [Go (Golang)](https://golang.org/)
- [Fiber v2](https://gofiber.io/)

---

## 📁 Project Structure ##

```

.
├── main.go          # Main application file
├── go.mod           # Go module definition
└── README.md        # This file

````

---

## 🛠 Setup Instructions 

### 1. Clone the Repository
```bash
git clone https://github.com/yourusername/your-repo-name.git
cd your-repo-name
````

### 2. Initialize Go Module (if not already done) ###

```bash
go mod init github.com/yourusername/your-repo-name
```

### 3. Install Dependencies ###

```bash
go get github.com/gofiber/fiber/v2
```

### 4. Run the Server ###

```bash
go run main.go
```

Server will start on [http://localhost:8080](http://localhost:8080)

---

## 📮 API Endpoints ##

| Method | Endpoint     | Description         |
| ------ | ------------ | ------------------- |
| POST   | `/posts`     | Create a new post   |
| PUT    | `/posts/:id` | Update a post by ID |
| DELETE | `/posts/:id` | Delete a post by ID |
| GET    | `/posts`     | Get all posts       |
| GET    | `/posts/:id` | Get a specific post |

---

## 📦 Sample Post JSON Format  ##

```json
{
  "id": 1,
  "author": "Jane Doe",
  "title": "My First Post",
  "content": "This is the content of the post."
}
```

---

## ⚠️ Notes ##

* This app uses an in-memory slice to store posts — data is lost on server restart.
* This is intended for learning and prototyping. For production use, integrate a proper database (PostgreSQL, MongoDB, etc.)

---

## 🧠 Next Steps ##

* Add persistent storage (e.g., PostgreSQL)
* Add request validation
* Add Swagger/OpenAPI docs
* Add tests using Go's `testing` package

---

## 🤝 Contributing ##

Pull requests are welcome. For major changes, please open an issue first to discuss what you’d like to change.

---

## 📄 License ##

This project is open source and available under the [MIT License](LICENSE).
