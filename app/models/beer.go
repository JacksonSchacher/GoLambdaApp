package models

type beer struct {
	ID          string  `json:"id"`
	NAME        string  `json:"name"`
	TAGLINE     string  `json:"tagline"`
	DESCRIPTION string  `json:"description"`
	IMAGE_URL   string  `json:"image_url"`
	ABV         float32 `json:"abv"`
	IBU         int     `json:"ibu"`
	EBC         int     `json:"ebc"`
	SRM         int     `json:"srm"`
	PH          float32 `json:"ph"`
}