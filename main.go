package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)
type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `'json: "email"`
	BloodType string `json:"blood_type"`
}
var Users []User

func main ( ){
	// create a new gin router
	router:= gin.Default()

	//define a sinngle endpoint
	router.GET("/", helloWorldhandler)

	//CRUD endpoint for data

	//create
	router.POST("/createUser", createUserHandler)

	//retrieve
	router.GET("/getUsers", getAllUserHandler)

	// name is just a placeholder
	router.GET("/getUser/:name", getSingleUserHandler)

	// update
	router.PATCH("/updateUser", updateUserHandler)

	//delete
	router.DELETE("/deleteUser", deleteUserHandler)


	port :=os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	_= router.Run(":" + port)


}

func helloWorldhandler(c *gin.Context) {
	c.JSON(200,gin.H{
		"message":"hello world",
		"status": "we are live",
		"age": "25",
		"name": "father ruben",
	})
}

func createUserHandler(c *gin.Context) {
	//createUser
	var user User
	// user := User{}
	
	err :=c.ShouldBindJSON(&user)
	if err != nil{
			c.JSON(400, gin.H{
				"error": "invalid request data",
			})
			return

	}
	Users = append(Users, user)
	// save user somewhere
	c.JSON(200, gin.H{
		"message": "new user succeefully craedted",
		"data": user,
	})
}
func getSingleUserHandler( c *gin.Context)  {
	name := c.Param("name")

	fmt.Println("name", name)

	var user User
	
	for _, value := range Users {
		
		if value.Name == name {

			user = value
		}
	}
	if &user == nil {
		c.JSON( 404,gin.H{
			"error": "no user with name found:" + name,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data": user,
	})

}

func getAllUserHandler(c *gin.Context)  {
	
	
	c.JSON(200, gin.H{
		"message":"Hello world",
		"data": Users,
	})
	
}
func updateUserHandler(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "User update!",
	})
}

func deleteUserHandler(c *gin.Context)  {
	c.JSON(200,gin.H{
		"message": "user has been bounced",
	})
}