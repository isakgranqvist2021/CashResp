package models

type Offer struct {
	Href        string
	OfferID     string
	PublisherID string
	AppID       string
	CreatedAt   string
	UpdatedAt   string
}

/*
	https://www.offertoro.com/click_track/api
	?offer_id=19498565
	&pub_id=29892
	&pub_app_id=13859
	&USER_ID=[USER_ID]
*/
