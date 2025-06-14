package main
import ( 
	"strconv"
	"github.com/gofiber/fiber/v2"
)

type posts struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"content"`
}
// fake db
var Posts []posts

// create a new post
func addPosts(c *fiber.Ctx) error {
	var newPost posts
	if err := c.BodyParser(&newPost);err != nil {
		return c.Status(400).SendString(err.Error())
	}
	Posts = append(Posts, newPost)
	return c.Status(201).JSON(fiber.Map{
		"message": "Post added successfully",
		"post":    newPost,
	})
}
// edit a post
func editPost(c *fiber.Ctx) error {
id,error := strconv.Atoi(c.Params("id"))
	if error != nil {
		return c.Status(400).SendString("Invalid post ID")
	}
	found := false
	for index, post := range Posts {
		if post.ID == id {
			found = true
			Posts[index].Author = c.FormValue(("author"))
			Posts[index].Title = c.FormValue("title")
			Posts[index].Content = c.FormValue("content")
			return c.Status(200).JSON(fiber.Map{
				"message" : "Post successfully updated",
				"post":    Posts[index],
			})
		}
	}
	if !found {
		return c.Status(404).SendString("Post not found")
	}
	return c.Status(500).SendString("Internal Server Error")
}
// get all posts
func getPosts (c *fiber.Ctx) error {
	if len(Posts) == 0 {
		return c.Status(404).SendString("No posts found")
	}
	return c.Status(200).JSON(Posts)
}
// get a post by id
func get_post_by_id(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid post ID")
	}
	found := false
	for _, post := range Posts {
		if post.ID == id {
			found = true
			return c.Status(200).JSON(post)
		}
	}
	if !found {
		return c.Status(404).SendString("Post not found")	
	}
	return c.Status(500).SendString("Internal Server Error")
}
// delete a post
func deletePost(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("Invalid post ID")
	}
	found := false
	for index, post := range Posts {
		if post.ID == id {
			found = true
			Posts = append(Posts[:index], Posts[index+1:]...)
			return c.Status(200).JSON(fiber.Map{
				"message": "Post deleted successfully",
				"posts":   Posts,
			})
		}
	}
	if !found {
		return c.Status(404).SendString("Post not found")
	}
	return c.Status(500).SendString("Internal Server Error")
}

func main() {
	app := fiber.New()
	app.Post("/posts", addPosts)
	app.Put("/posts/:id", editPost)
	app.Delete("/posts/:id", deletePost)
	app.Get("/posts", getPosts)	
	app.Get("/posts/:id", get_post_by_id)
	app.Listen(":8080")
}
