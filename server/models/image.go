package models

import (
	"fmt"

	"github.com/isakgranqvist2021/surveys/utils"
)

type Image struct {
	Filename         string
	OriginalFilename string
	Type             string
	Alt              string
	Size             int64
	CreatedAt        string
	UpdatedAt        string
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
