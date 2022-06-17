package dto



type AdminLoginInput struct {
	UserName string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"password" json:"user_name" binding:"required"`
}



