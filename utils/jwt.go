package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type Jwt struct {
	Secret string
	Ttl    int
	Issuer string
}

type JwtInfo struct {
	Id           int        `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Username     string     `gorm:"column:username;type:varchar(30);comment:账号;NOT NULL" json:"username"`
	Avatar       string     `gorm:"column:avatar;type:varchar(255);comment:头像" json:"avatar"`
	Name         string     `gorm:"column:name;type:varchar(30);comment:真实姓名;NOT NULL" json:"name"`
	Phone        string     `gorm:"column:phone;type:char(12);comment:电话" json:"phone"`
	Email        string     `gorm:"column:email;type:varchar(255);comment:邮箱" json:"email"`
	DepartmentId int        `gorm:"column:department_id;type:int(11);default:0;comment:部门ID" json:"department_id"`
	PostId       int        `gorm:"column:post_id;type:int(11);default:0;comment:岗位ID;NOT NULL" json:"post_id"`
	Roles        string     `gorm:"column:roles;type:varchar(255);comment:角色" json:"roles"`
	LastTime     CustomTime `gorm:"column:last_time;type:datetime;comment:最后一次登录时间" json:"last_time"`
	jwt.RegisteredClaims
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// GetToken 获取token
//
//	access is true 获取登录token
//	access is false 获取置换token
func (j *Jwt) GetToken(user JwtInfo, access bool) (string, error) {
	j.setJwtConfig(access)
	// 创建 Claims
	user.Issuer = j.Issuer
	user.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Duration(j.Ttl) * time.Second)) // 过期时间
	user.ID = strconv.Itoa(user.Id)
	// 生成token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, user)
	// 生成签名字符串
	return token.SignedString([]byte(j.Secret))
}

// ParseToken 校验token
//
//	access is true 校验登录token
//	access is false 校验置换token
func (j *Jwt) ParseToken(tokenString string, access bool) (*JwtInfo, error) {
	j.setJwtConfig(access)
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JwtInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil { // 解析token失败
		return nil, err
	}
	//对token对象中的Claim进行类型断言
	claims, ok := token.Claims.(*JwtInfo)
	if ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("is check false")
}

func (j *Jwt) setJwtConfig(access bool) {
	j.Ttl = GetConfigInt("jwt.access_ttl")
	j.Issuer = GetConfigString("jwt.issuer")
	j.Secret = GetConfigString("jwt.access_secret")

	if !access {
		j.Ttl = GetConfigInt("jwt.refresh_ttl")
		j.Secret = GetConfigString("jwt.refresh_secret")
	}
}
