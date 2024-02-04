package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"net/http"
)

type Repository struct {
	DB *gorm.DB
}

type Book struct {
	Author	string	`json:"author"`
	Title	string	`json:"title"`
	Publisher	string	`json:"publisher"`
}

func  (r *Repository) CreateBook(context *fiber.Ctx) error{
	book := Book{}

	err := context.BodyParser(&book)

	if err != nil  {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map("message":"request failed")
			return err
		)
	}
	err = r.DB.Create(&nook).Error
	if err!= nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map("message":"request failed")
			return err
		)
	}

	context.Status(http.StatusOk).JSON(&fiber.Map{
		"message:"book has been added
	})
	return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
	bookModels := &[]models.Books{}

	err := r.DB.Find(bookModels).Error
	if err != nil{
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map("message":"could not get books")
			return err
		)
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message":"books fetched succesfully",
		"data":bookModels
	})
	return nil

}

func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api")
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete_book/:id", r.DeleteBook)
	api.Get("/get_books/:id",r.GetBookByID)
	api.Get("/books",r.GetBooks)

}


func main(){
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal(err)
	}

	r := Repository{
		DB: db,

	}

	app := fiber.New() 
	r.SetupRoutes(app)
	app.Listen(":8080")


}