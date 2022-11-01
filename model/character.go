package model

type Character struct {
	ID      int64  `gorm:"column:id"`
	Name    string `gorm:"column:name"`
	Status  string `gorm:"column:status"`
	Species string `gorm:"column:species"`
	Type    string `gorm:"column:type"`
	Gender  string `gorm:"column:gender"`
	Image   string `gorm:"column:image"`
}

func (Character) TableName() string {
	return "character"
}
