package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ianodad/bookstore_users_microservice/domain/users"
	"github.com/ianodad/bookstore_users_microservice/services"
	"github.com/ianodad/bookstore_users_microservice/utils/errors"
)



func SearchUser(c *gin.Context){
	c.String(http.StatusNotImplemented, "Not Implemented")
}

func CreateUser(c *gin.Context){
	var user users.User

	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println(err)
	// 	//TODO handle error
	// 	return
	// }

	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// 	//TODO: handle error
	// 	return
	// }
	
	if err := c.ShouldBindJSON(&user); err != nil {
		// TODO: handle error
		resTErr := errors.RestErr{
			Message: "Invalid json body",
			Status: http.StatusBadRequest,
			Error: "bad request",
		}
		c.JSON(resTErr.Status, resTErr)
		return
	}
	result, saveErr := services.CreateUser(user)

	if saveErr != nil{
		//TODO: handle error
		c.JSON(saveErr.Status, saveErr)
		return
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil{
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}