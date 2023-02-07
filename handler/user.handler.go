package handler

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/model/response"
	"go-fiber-gorm/utils"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	return ctx.JSON(users)

}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	//Validasi Req
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"data":    newUser,
		"message": "success",
	})
}

func UserHandlerById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	userResponse := response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Address:   user.Address,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    userResponse,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var user entity.User

	userId := ctx.Params("id")
	// Check Avail User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//Update User Data
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	// Check Avail User
	err := database.DB.Debug().First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
