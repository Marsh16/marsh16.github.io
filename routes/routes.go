package routes
// go mod init vp_week11_echo 
// GO111MODULE=on go get github.com/labstack/echo/v4

import (
	"github.com/labstack/echo/v4"
	"archcalculator.github.io/controllers"
)

func Init() *echo.Echo {
	e := echo.New()	

	//Book
	e.GET("/getAllBook", controllers.ReadAllBook)
	e.GET("/getBookByBookId", controllers.ReadBookByBookId)
	e.GET("/getBookByMemberId", controllers.ReadBookByMemberId)
	e.POST("/createBook", controllers.CreateBook)
	e.PATCH("/editBook", controllers.EditBook)
	e.DELETE("/deleteBook", controllers.DeleteBook)
    e.Static("/images", "./images")

	//Category
	e.GET("/getAllCategory", controllers.ReadAllCategory)
	e.POST("/createCategory", controllers.CreateCategory)
	e.PATCH("/editCategory", controllers.EditCategory)
	e.DELETE("/deleteCategory", controllers.DeleteCategory)

	//Member
	e.GET("/getAllMember", controllers.ReadAllMember)
	e.POST("/createMember", controllers.CreateMember)
	e.PATCH("/editMember", controllers.EditMember)
	e.DELETE("/deleteMember", controllers.DeleteMember)

	//Member
	e.GET("/getBookByCategoryId", controllers.ReadBookByCategoryId)
	e.GET("/getCategoryByBookId", controllers.ReadCategoryByBookId)
	e.POST("/createBookCategory", controllers.CreateBookCategory)
	e.DELETE("/deleteBookCategory", controllers.DeleteBookCategory)

	return e

}

