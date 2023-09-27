package entities

type Review struct {
	Id      int    `json:"id" gorm:"primaryKey"`
	UserID  int    `json:"user_id" gorm:"type:int"`
	Comment string `json:"comment" gorm:"type:text"`
	Rating  int    `json:"rating" gorm:"type:int"`
}
