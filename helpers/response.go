package helpers

import (
	"github.com/gofiber/fiber/v2"
)

func ResponseSuccess(c *fiber.Ctx, message string, data interface{}) error {
    return c.JSON(fiber.Map{
        "success": true,
        "code":    "ALP-001",
        "message": message,
        "data":    data,
    })
}

func ResponseError(c *fiber.Ctx, code string, message string) error {
    return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
        "success": false,
        "code":    code,
        "message": message,
        "data":    make([]string, 0), // bisa ganti ke nil jika tidak ingin array kosong
    })
}
