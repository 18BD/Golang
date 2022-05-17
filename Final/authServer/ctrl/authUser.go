package ctrl

import (
	"errors"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"grpcAsan/server/models"
	"net/http"
)

type response struct {
	Response string `json:"response"`
}

type request struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func (ctrl *Ctrl) AuthUser(c echo.Context) error {
	req := &request{}

	err := c.Bind(req)
	if err != nil {
		return err
	}

	var users []models.User

	err = ctrl.db.Model(&users).Where("login = ?", req.Login).Select()
	if err != nil {
		return err
	}

	user := models.User{}
	if len(users) != 0 {
		user = users[0]
	}

	if user.Login == "" {
		return errors.New("no such user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return errors.New("wrong password")
	}

	return c.JSON(http.StatusOK, response{Response: "OK"})
}
