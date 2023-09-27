package entities

type Jobs struct {
	Id              int    `json:"id" gorm:"primaryKey"`
	CompanyName     string `json:"company_name" gorm:"type:varchar(255)"`
	PositionName    string `json:"position_name" gorm:"type:varchar(255)"`
	CompanyLocation string `json:"company_location" gorm:"type:varchar(255)"`
	LogoCompany     string `json:"logo_company" gorm:"type:varchar(255)"`
	Salary          uint   `json:"salary" gorm:"type:int"`
	TypeSalary      uint   `json:"type_salary" gorm:"type:int"`
	JobDescription  string `json:"job_description" gorm:"type:text"`
	JobRequirement  string `json:"job_requirement" gorm:"type:text"`
	JobBenefit      string `json:"job_benefit" gorm:"type:text"`
	JobType         string `json:"job_type" gorm:"type:varchar(255)"`
	JobCategory     string `json:"job_category" gorm:"type:varchar(255)"`
	JobStatus       string `json:"job_status" gorm:"type:varchar(255)"`
	JobSource       string `json:"job_source" gorm:"type:varchar(255)"`
	JobUrl          string `json:"job_url" gorm:"type:varchar(255)"`
	CreatedAt       string `json:"created_at" gorm:"type:timestamp"`
}

func (Jobs) TableName() string {
	return "jobs"
}
