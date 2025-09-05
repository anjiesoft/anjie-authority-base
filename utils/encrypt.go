package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

// EncryptPassword 将密码加密，需要传入密码返回的是加密后的密码
func EncryptPassword(password string) (string, error) {
	// 加密密码，使用 bcrypt 包当中的 GenerateFromPassword 方法，bcrypt.DefaultCost 代表使用默认加密成本
	encryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		// 如果有错误则返回异常，加密后的空字符串返回为空字符串，因为加密失败
		return "", err
	} else {
		// 返回加密后的密码和空异常
		return string(encryptPassword), nil
	}
}

// EqualsPassword 对比密码是否正确
// password未加密的
// encryptPassword 加密的
func EqualsPassword(password, encryptPassword string) bool {
	// 使用 bcrypt 当中的 CompareHashAndPassword 对比密码是否正确，第一个参数为加密后的密码，第二个参数为未加密的密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptPassword), []byte(password))

	// 对比密码是否正确会返回一个异常，按照官方的说法是只要异常是 nil 就证明密码正确
	return err == nil
}

func GetSaltPassword(salt, password string) string {
	pwd := fmt.Sprintf("%s%s", salt, password)
	return pwd
}

// GetRandstring 取得随机字符串:使用字符串拼接
func GetRandstring(length int) string {
	if length < 1 {
		return ""
	}
	char := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charArr := strings.Split(char, "")
	charlen := len(charArr)
	ran := rand.New(rand.NewSource(time.Now().Unix()))

	var rchar string = ""
	for i := 1; i <= length; i++ {
		rchar = rchar + charArr[ran.Intn(charlen)]
	}
	return rchar
}

func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha256ToHex 将加密后的二进制转16进制字符串
func HmacSha256ToHex(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}

// HmacSha256ToBase64 将加密后的二进制转Base64字符串
func HmacSha256ToBase64(key string, data string) string {
	return base64.URLEncoding.EncodeToString(HmacSha256(key, data))
}
