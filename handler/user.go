package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/ruan-molinari/ecommerce/model"
	"github.com/ruan-molinari/ecommerce/view/user"
)

type UserHandler struct {}

func(h UserHandler) HandleUserShow(c echo.Context) error {
	usr := model.User{
		Email: "hello@world.com",
	}
	return render(c, user.Show(usr))
}
