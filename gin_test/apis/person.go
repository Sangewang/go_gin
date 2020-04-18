package apis

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	. "gin_test/models"
	)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPersonApi()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(
		http.StatusOK, gin.H{
			"msg": msg,})
}