package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"log"
	"time"
)
type User struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `'json: "email"`
	BloodType string `json:"blood_type"`
}
var Users []User
var dbClient *mongo.Client

func main ( ){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil{
		log.Fatalf("Could not connect to the db: %v\n", err)
	}

	dbClient = client
	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		// if there was an issue with the pin
		// print out the error and exit using log.Fatal()
		log.Fatalf("MOngo db not available: %v\n", err)
	}
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
	router.PATCH("/updateUser/:name", updateUserHandler)

	//delete
	router.DELETE("/deleteUser/:name", deleteUserHandler)


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
	
	err := c.ShouldBindJSON(&user)
	if err != nil{
			c.JSON(400, gin.H{
				"error": "invalid request data",
			})
			return

	}

	_, err = dbClient.Database("chuksDatabase").Collection("chuksCollection").InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println("error saving user", err)
	//	if saving ws not successful
		c.JSON(500, gin.H{
			"error": "Could not process request, could not save user",
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

	userAvailable := false
	
	for _, value := range Users {
		
		if value.Name == name {

			user = value

			userAvailable = true
		}
	}

	if !userAvailable  {
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

	name :=c.Param("name")

	var user User

	err :=c.ShouldBindJSON(&user)

	if err !=nil {
		c.JSON(400,gin.H{ 
			"error": "invalid request data",
			})
	return

	}

	filterQuery := bson.M{
		"name": name,
	}

	updateQurey := bson.M{
		"$set":bson.M{
			"name":user.Name,
			"age":user.Age,
			"email":user.Email,
		},
	}
	_,err = dbClient.Database("chuksDatabase").Collection("chuksCollection").UpdateOne(context.Background(),filterQuery,updateQurey)
	if err !=nil {
		c.JSON(500, gin.H{
			"error":"could not process request,database has not been updated",
		})
		return

	}

	c.JSON(200, gin.H{
		"message": "User update!",
	})
}

func deleteUserHandler(c *gin.Context)  {

	name := c.Param("name")

	query := bson.M{

		"name": name,
	}
	_, err :=dbClient.Database("chuksDatabase").Collection("chuksCollection").DeleteOne(context.Background(),query)
	if err != nil {
		c.JSON(500, gin.H{
			"error":"could not process request, could not bounce user",
		})
		return
	}

	c.JSON(200,gin.H{
		"message": "user has been bounced",
	})
}