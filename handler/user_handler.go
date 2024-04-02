package handler

import (
	"net/http"

	"github.com/HaAnhTu2/code01/entities"
	req "github.com/HaAnhTu2/code01/entities/req"
	"github.com/HaAnhTu2/code01/log"
	"github.com/HaAnhTu2/code01/reponsitory"

	validator "github.com/go-playground/validator/v10"
	//uuid "github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo reponsitory.UserRepo
}

func (u *UserHandler) HandleCreateUser(c echo.Context) error {
	// Nhận dữ liệu từ form
	req := req.ReqSignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, entities.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Kiểm tra tính hợp lệ của dữ liệu
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, entities.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Tạo một đối tượng người dùng mới
	newUser := entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	// Lưu người dùng vào cơ sở dữ liệu
	newUser, err := u.UserRepo.SaveUser(c.Request().Context(), newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, entities.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	// Trả về phản hồi cho người dùng
	newUser.Password = "" // Đảm bảo rằng mật khẩu không được trả về trong phản hồi
	return c.JSON(http.StatusOK, entities.Response{
		StatusCode: http.StatusOK,
		Message:    "Tài khoản đã được tạo thành công",
		Data:       newUser,
	})
}
