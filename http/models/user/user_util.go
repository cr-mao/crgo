package user

import "crgo/infra/db"

func IsEmailExist(email string) bool {
	var count int64
	db.GetDb("default").Model(User{}).Where("email = ? ", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	db.GetDb("default").Model(User{}).Where("phone = ? ", phone).Count(&count)
	return count > 0
}
