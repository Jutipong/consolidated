package entity

type User struct {
	ID       string `gorm:"column:ID" json:"id" binding:"required"`
	Code     string `gorm:"column:Code" json:"code" binding:"required" `
	Name     string `gorm:"column:Name" json:"name"`
	Age      int8   `gorm:"column:Age" json:"age" binding:"gte=1,lte=130"`
	Email    string `gorm:"column:Email" json:"email" binding:"required,email"`
	IsActive bool   `gorm:"column:IsActive" json:"isactive"`
}
