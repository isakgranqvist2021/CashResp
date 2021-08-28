package admin

import (
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/isakgranqvist2021/surveys/models"
	"github.com/isakgranqvist2021/surveys/utils"
)

func GetImages(c *fiber.Ctx) error {
	images, err := models.GetImages()

	if err != nil {
		images = []models.Image{}
	}

	return c.Render("pages/admin/images", fiber.Map{
		"Title":       "Images",
		"User":        c.Locals("User"),
		"Alert":       c.Locals("Alert"),
		"Stylesheets": []string{"images.min.css"},
		"Data": map[string]interface{}{
			"Images": images,
		},
	})
}

func PostImages(c *fiber.Ctx) error {
	fileHeader, err := c.FormFile("image")

	if err != nil {
		return c.Redirect(c.OriginalURL())
	}

	Filename := fmt.Sprintf("%s%s", utils.RandKey(50, false), filepath.Ext(fileHeader.Filename))

	if err = c.SaveFile(fileHeader, fmt.Sprintf("./uploads/%s", Filename)); err != nil {
		return c.Redirect(c.OriginalURL())
	}

	image := models.Image{
		Filename:         Filename,
		OriginalFilename: fileHeader.Filename,
		Type:             filepath.Ext(fileHeader.Filename),
		Size:             fileHeader.Size,
	}

	if err := c.BodyParser(&image); err != nil {
		return c.Redirect(c.OriginalURL())
	}

	if err := image.SaveImage(); err != nil {
		return c.Redirect(c.OriginalURL())
	}

	return c.Redirect(c.OriginalURL())
}
