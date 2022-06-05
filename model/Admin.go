package model

type Admin struct {
	ID        string `gorm:"primaryKey" param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
	Password  string `gorm:"not null;" param:"password" query:"password" header:"password" form:"password" json:"password" xml:"password"`
	Name      string `gorm:"not null" param:"name" query:"name" header:"name" form:"name" json:"name" xml:"name"`
	Email     string `gorm:"not null" param:"email" query:"email" header:"email" form:"email" json:"email" xml:"email"`
	Role      string `gorm:"not null" param:"role" query:"role" header:"role" form:"role" json:"role" xml:"role"`
	LastLogin uint64 `gorm:"not null" json:"lastLogin"`
	Img       string `gorm:"not null" param:"img" query:"img" header:"img" form:"img" json:"img" xml:"img"`
	Type      string `gorm:"not null" param:"type" query:"type" header:"type" form:"type" json:"type" xml:"type"`
}
