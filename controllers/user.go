package controllers

import (

	// "fmt"

	"encoding/json"
	"fmt"
	"strconv"

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

// adding user
func (c *UserController) Post() {

	var input models.Users
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid data"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, error := o.Insert(&input)

	if error != nil {
		c.Data["json"] = map[string]string{
			"error": "Faled to add",
		}
	} else {
		c.Data["json"] = map[string]string{
			"message": "user saved",
		}
	}

	c.ServeJSON()
}

// update user data
func (c *UserController) Patch() {
	userId := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(userId)
	user := models.Users{Id: id}
	o := orm.NewOrm()

	var input models.Users
	bodyerr := json.Unmarshal(c.Ctx.Input.RequestBody, &input)

	if bodyerr != nil {
		fmt.Println(input)
	}

	if err := o.Read(&user); err == nil {
		user.Name = input.Name
		user.Email = input.Email

		if num, err := o.Update(&user); err == nil {
			successResponse := map[string]interface{}{
				"message": "updated",
				"num":     num,
			}
			c.Data["json"] = successResponse
		} else {
			errorResponse := map[string]interface{}{
				"message": "Faled to update",
			}
			c.Data["json"] = errorResponse
		}

		c.ServeJSON()
	}

}

// remove user
func (c *UserController) Delete() {
	userId := c.Ctx.Input.Param(":id")
	num, err := strconv.Atoi(userId)

	o := orm.NewOrm()
	user := models.Users{Id: num}

	o.Delete(&user)

	if err == nil {
		reponse := map[string]string{
			"message": "Deleted",
		}

		c.Data["json"] = reponse
	}

	c.ServeJSON()
}
