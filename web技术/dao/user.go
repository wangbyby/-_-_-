package dao

//Login type
type Login struct {
	ID       int64  `form:"id" json:"id" xml:"id" `
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}
