package model

type Character struct {
	ID      int64
	Name    string
	Status  string
	Species string
	Type    string
	Gender  string
	Image   string
}

func (Character) TableName() string {
	return "character"
}
