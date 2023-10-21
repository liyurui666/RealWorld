package models

type User struct {
	Id       uint
	Username string
	Email    string
	Password string
	Bio      string
	Image    string
}

func (table *User) TableName() string {
	return "user"
}

// UserRegistration 用户注册
func UserRegistration(user *User) {
	//添加新用户数据
	DB.Create(&user)
}

// QueryEmail 查询邮箱是否已存在
func QueryEmail(user User) User {
	DB.Where("email = ?", user.Email).First(&user)
	return user
}

// Login 用户登录 根据Email和Password查询
func Login(user *User) error {
	err := DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error
	return err
}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(user *User) {
	DB.Where("email = ?", user.Email).Model(&user).Updates(User{Bio: user.Bio, Image: user.Image})
}
