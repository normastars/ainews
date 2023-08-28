package main

import (
	"errors"
	"net/http"

	"github.com/normastars/frame"
)

var (
	DbName = "user"
)

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func (u *User) TableName() string {
	return "tb_user_x"
}

func init() {
	frame.RegisterTable(DbName, &User{})
}

func main() {
	app := frame.New()
	app.GET("/hello", HelloWorld)
	app.Run()

}

// HelloWorld hell world handler
func HelloWorld(c *frame.Context) {
	db := c.GetDB(DbName)
	// create user
	user := User{Name: "richardyu"}
	result := db.Create(&user)
	if result.Error != nil {
		c.Fatalf("failed to create user: %v", result.Error)
	}
	c.Infof("created user: %v\n", user)
	// send http
	c.DoHTTP().R().Get("https://httpbin.org/uuid")
	c.HTTPError2(http.StatusOK, "X0111", "normal error", errors.New("system panic"))
}
