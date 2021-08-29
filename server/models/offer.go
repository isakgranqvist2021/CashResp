package models

import (
	"fmt"

	"github.com/isakgranqvist2021/surveys/utils"
)

type Offer struct {
	ID          int
	OfferID     string
	PublisherID string
	AppID       string
	Description string
	CreatedAt   string
	UpdatedAt   string
	Href        string
	ImageID     int
	Draft       bool
	Provider    string
}

/*
	https://www.offertoro.com/click_track/api
	?offer_id=19498565
	&pub_id=29892
	&pub_app_id=13859
	&USER_ID=[USER_ID]
*/

func (o *Offer) CreateNew() error {
	db := utils.Connect()
	defer db.Close()

	draft := 0
	if o.Draft {
		draft = 1
	}

	query := fmt.Sprintf(`
		INSERT INTO offers (OfferID, PublisherID, AppID, Description, Href, ImageID, Provider, Draft)
		VALUES ('%s', '%s', '%s', '%s', '%s', '%d', '%s', '%v')`,
		o.OfferID, o.PublisherID, o.AppID, o.Description, o.Href, o.ImageID, o.Provider, draft,
	)

	_, err := db.Exec(query)

	return err
}
