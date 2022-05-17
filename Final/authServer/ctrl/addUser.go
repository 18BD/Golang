package ctrl

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"grpcAsan/server/models"
	"net/http"
)

type addrequest struct {
	UserID int `json:"user_id"`
	Login string `json:"login"`
	Password string `json:"password"`
}


func (ctrl *Ctrl) AddUser(c echo.Context) error {
	req := &addrequest{}

	err := c.Bind(req)
	if err != nil {
		return err
	}

	if req.Login == "" || req.Password == "" {
		return errors.New("no login or no pass")
	}

	// hashing
	hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), 8)

	// adding to DB
	userAdd := models.User{
		UserID:   req.UserID,
		Login:    req.Login,
		Password: string(hashed),
	}

	_, err = ctrl.db.Model(&userAdd).Insert()
	if err != nil {
		return fmt.Errorf(
			"trouble inserting new user, maybe hash is too long and doesnt fit in varchar(x)? %s",
			err)
	}

	return c.JSON(http.StatusOK, response{Response: "OK"})
}

