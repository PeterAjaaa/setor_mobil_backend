package dto

type CarResponseRequest struct {
	ID              uint   `json:"id"`
	RegistrationNum string `json:"registration_num"`
	Brand           string `json:"brand"`
	Model           string `json:"model"`
	Year            uint16 `json:"year"`
	PricePerDay     uint32 `json:"price_per_day"`
	Status          string `json:"status"`
	Description     string `json:"description"`
	ImageURL        string `json:"image_url"`
}
