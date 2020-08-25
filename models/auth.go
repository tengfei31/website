package models

type AUTH struct {
	ID int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username string, password string) bool {
	var auth AUTH
	db.Select("id").Where(AUTH{Username:username, Password: password}).First(&auth)
	if auth.ID > 0 {
		return true
	}
	return false
}