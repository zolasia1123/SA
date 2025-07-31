package main

import (
	"github.com/SA/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"github.com/SA/controller/student"
	"github.com/SA/config"
	"github.com/gin-gonic/gin"
	"github.com/SA/middlewares"
)

const PORT = "8000"

func main() {
	db, err := gorm.Open(sqlite.Open("sa.db"), &gorm.Config{})
	if err != nil {
	panic("failed to connect database")
	}
	config.ConnectionDB()

	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	router := r.Group("/")
   {
       router.Use(middlewares.Authorizes())
       // User Route
       router.PUT("/user/:id", student.Update)
       router.GET("/users", student.GetAll)
       router.GET("/user/:id", student.Get)
       router.DELETE("/user/:id", student.Delete)
   }

   r.GET("/genders", student.GetAll)
   r.GET("/", func(c *gin.Context) {
       c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
   })
    // Run the server
    r.Run("localhost:" + PORT)

	// Migrate the schema
	db.AutoMigrate(&entity.Admin{}, &entity.Billing{}, &entity.Payment{}, &entity.Contract{}, &entity.Room{}, &entity.Student{})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
 }