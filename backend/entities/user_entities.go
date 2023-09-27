package entities

type AuthUser struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(255)"`
	Username  string `json:"username" gorm:"type:varchar(255)"`
	Email     string `json:"email" gorm:"type:varchar(255)"`
	Password  string `json:"password" gorm:"type:varchar(255)"`
	CreatedAt string `json:"created_at" gorm:"type:timestamp"`
}

func (AuthUser) TableName() string {
	return "users"
}
