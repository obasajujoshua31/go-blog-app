package auto

import "github.com/obasajujoshua31/blogos/api/models"

var users = []models.User{
	models.User{
		Name:     "John Doe",
		Email:    "johndoe@email.com",
		Password: "12345678",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "This is my first title",
		Content: "This is my first Content",
	},
}
