package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// Define some data to send as JSON
	// response := map[string]interface{}{
	// 	"name":  "John Doe",
	// 	"email": "johndoe@example.com",
	// }

	// // Set the data as JSON
	// c.Data["json"] = user

	// // Serve the JSON response
	// c.ServeJSON()

	// retrive data from db
	o := orm.NewOrm()
	var users []models.Users
	total, err := o.QueryTable("users").All(&users)
	if err != nil {
		errorResponse := map[string]string{
			"status": "failed",
			"error":  "Resource not found",
		}

		c.Ctx.Output.SetStatus(404) // Set HTTP status code to 404
		c.Data["json"] = errorResponse

	} else {
		c.Ctx.Output.SetStatus(200)
		successResponse := map[string]interface{}{
			"status":  "success",
			"message": "Welcome to Beego!",
			"total":   total,
			"data":    users,
		}

		c.Data["json"] = successResponse

	}
	c.ServeJSON()
}

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// adding user
func (c *UserController) Post() {
	var input UserInput
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid data"}
		c.ServeJSON()
		return
	}

	// response with
	response := map[string]string{
		"message": "Data received successfully",
		"name":    input.Name,
		"email":   input.Email,
	}
	c.Data["json"] = response
	c.ServeJSON()
}
