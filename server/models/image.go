package models

import (
	"fmt"

	"github.com/isakgranqvist2021/surveys/utils"
)

type Image struct {
	ID               int
	Filename         string
	OriginalFilename string
	Type             string
	Alt              string
	CreatedAt        string
	UpdatedAt        string
	Size             int64
}

func GetImages() ([]Image, error) {
	db := utils.Connect()
	defer db.Close()

	query := "SELECT * FROM images"

	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	var images []Image
	for rows.Next() {
		var image Image
		if err := rows.Scan(
			&image.ID,
			&image.Filename,
			&image.OriginalFilename,
			&image.Type,
			&image.Alt,
			&image.CreatedAt,
			&image.UpdatedAt,
			&image.Size,
		); err != nil {
			fmt.Println(err)
		} else {
			images = append(images, image)
		}
	}

	return images, nil
}

func (i *Image) SaveImage() error {
	db := utils.Connect()
	defer db.Close()

	query := fmt.Sprintf(`
		INSERT INTO images (Filename, OriginalFilename, Type, Alt, Size) 
		VALUES ('%s', '%s', '%s', '%s', '%d')`,
		i.Filename, i.OriginalFilename, i.Type, i.Alt, i.Size)

	_, err := db.Exec(query)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
