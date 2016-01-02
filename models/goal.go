package models

type Goal struct {
	Id   int    `json:"id"`
	Text string `json:"text"`
	Color string `json:"color"`
	PicturePath string `json:"picturePath"`
}

func (g * Goal) Equals(other * Goal) bool{
	if g.Id != other.Id {
		return false
	}
	if g.Text != other.Text {
		return false
	}
	if g.Color != other.Color {
		return false
	}
	if g.PicturePath != other.PicturePath {
		return false
	}
	return true

}
