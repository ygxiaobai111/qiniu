package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName        string `gorm:"unique"`
	PasswordDigest  string
	FollowCount     int64
	FanCount        int64
	CollectionCount int64
	Avatar          string  `gorm:"size:1000"`
	Follows         []*User `gorm:"many2many:follows;"`                         // 关注列表
	Fans            []*User `gorm:"many2many:follows;joinForeignKey:follow_id"` // 粉丝列表
}

const (
	PassWordCost = 12                                 // 密码加密难度
	Avatar       = "http://www.xzkckj.cn/image/T.jpg" //默认用户头像

)

// SetPassword 设置密码加密
func (u *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	u.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err == nil
}
