package entities

type NotificationAcceptment struct {
	Id                           int    `json:"id" gorm:"primaryKey"`
	UserId                       int    `json:"user_id" gorm:"type:int"`
	JobId                        int    `json:"job_id" gorm:"type:int"`
	StatusNotificationAcceptment int    `json:"status_notification_acceptment" gorm:"type:int"`
	CreatedAt                    string `json:"created_at" gorm:"type:timestamp"`
}

func (NotificationAcceptment) TableName() string {
	return "notification_acceptment"
}

type NotificationAcceptmentJob struct {
	Id           int    `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	CompanyName  string `json:"company_name" gorm:"column:company_name"`
	PositionName string `json:"position_name" gorm:"column:position_name"`
	LogoCompany  string `json:"logo_company" gorm:"column:logo_company"`
}
