package main

import (
	"github.com/gin-gonic/gin"
	"jwt-gin/controllers"
	"jwt-gin/middlewares"
	"jwt-gin/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", controllers.Index)
	r.GET("/login", controllers.LoginGetHandler())
	r.POST("/login", controllers.LoginPostHandler())
	//r.POST("/register", controllers.RegisterPost)
	//r.GET("/logout", controllers.Logout)
	//r.GET("/dashboard", controllers.Dashboard)
	//r.GET("/profile", controllers.Profile)
	//r.GET("/edit", controllers.Edit)
	//r.POST("/edit", controllers.EditPost)
	//r.GET("/delete", controllers.Delete)
	//r.POST("/delete", controllers.DeletePost)
	//r.GET("/create", controllers.Create)
	//r.POST("/create", controllers.CreatePost)
	//r.GET("/read", controllers.Read)
	//r.POST("/read", controllers.ReadPost)
	//r.GET("/update", controllers.Update)
	//r.POST("/update", controllers.UpdatePost)

	public := r.Group("/api")
	public.POST("/register", controllers.Register)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	protected.GET("/logout", controllers.LogoutGetHandler())
	protected.GET("/dashboard", controllers.DashboardGetHandler())
	r.Run(":8080")

}
