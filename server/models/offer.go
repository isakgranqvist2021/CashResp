package models

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

	return nil
}
