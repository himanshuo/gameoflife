package models

type Goal struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Color string `json:"color"`
	PicturePath string `json:"picturePath"`
}
