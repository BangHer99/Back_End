package delivery

import (
	"be12/mentutor/config"
	"be12/mentutor/features/mentor"
	"be12/mentutor/middlewares"
	"be12/mentutor/utils/helper"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MentorDelivery struct {
	mentorUsecase mentor.UsecaseInterface
}

func New(e *echo.Echo, usecase mentor.UsecaseInterface) {

	handler := MentorDelivery{
		mentorUsecase: usecase,
	}

	e.PUT("/users", handler.UpdateProfile(), middleware.JWT([]byte(config.SECRET_JWT))) // UPDATE USER
	e.POST("/mentors/tasks", handler.AddTask(), middleware.JWT([]byte(config.SECRET_JWT))) // UPDATE USER
}

func (md *MentorDelivery) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input UpdateUserFormat

		IdUser, IdClass, role := middlewares.ExtractToken(c)
		input.IdClass = uint(IdClass)

		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}

		// CEK GAMBAR
		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadFotoProfile(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		} 

		input.ID = uint(IdUser)
		cnv := ToDomainUpdateUser(input)
		res, err := md.mentorUsecase.UpdateProfile(cnv, role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated,helper.SuccessResponse("success update profile", ToResponseUpdateUser(res)))
	}
}

func (md *MentorDelivery) AddTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		IdUser, IdClass, role := middlewares.ExtractToken(c)

		var input TaskRequest

		if err := c.Bind(&input); err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		
		// CEK GAMBAR
		file, fileheader, err := c.Request().FormFile("images")
		if err == nil {
			res, err := helper.UploadGambarTugas(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		} 

		// CEK FILE
		file, fileheader, err = c.Request().FormFile("file")
		if err == nil {
			res, err := helper.UploadGambarTugas(file, fileheader)
			if err != nil {
				log.Print(err)
				return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
			}
			log.Print(res)
			input.Images = res
		} 

		input.IdClass = uint(IdClass)
		input.IdMentor = uint(IdUser)
		cnv := ToDomainTask(input)
		res, err := md.mentorUsecase.AddTask(cnv, role)
		if err != nil {
			log.Print(err)
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("Invalid Input From Client"))
		}
		return c.JSON(http.StatusCreated,helper.SuccessResponse("Success insert task", ToResponseAddTask(res)))
	}
}