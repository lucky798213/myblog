package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"myblog/utils"
	"myblog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct { //用于存储一些必要的信息
	Username           string `json:"username"`
	jwt.StandardClaims        //声明中的预定义字段
}

var code int

// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(time.Hour * 10) //设置过期时间，Add方法用于在当前时间基础上增加指定的时间间隔
	SetClaims := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //其作用是将时间转换为 Unix 时间戳（即从 1970 年 1 月 1 日 00:00:00 UTC 起经过的秒数）
			Issuer:    "ginblog",
		},
	}

	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCSE
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) { //用于提供验证签名所需密钥的回调函数。
		return JwtKey, nil
	})
	if key, ok := setToken.Claims.(*MyClaims); ok && setToken.Valid {
		return key, errmsg.SUCCSE
	} else {
		return nil, errmsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")

		if tokenHeader == "" {
			code = errmsg.ERROR_TOKEN_EXIST
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)      //将token分割为n-1分，以sep进行分割，返回【】string
		if len(checkToken) != 2 || checkToken[0] != "Bearer" { //Bearer 是一个身份验证方案，这里看看是否符合格式
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1]) //将 Token 拆分为 Header、Payload、Signature 三部分，并储存在key中
		if tCode == errmsg.ERROR {
			code = errmsg.ERROR_TOKEN_WRONG
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ERROR_TOKEN_RUNTIME
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("username", key.Username)
		c.Next()
	}
}
