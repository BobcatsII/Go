package app

import (
	"gin-blog/global"
	"gin-blog/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//结构体，Appkey和AppSecret用于自定义的认证信息;
//jwt.StandardClaims结构体，在jwt-go中预定义的，详细可以ctrl+左键查看。
type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

//用于获取项目的JWT Secret，当前使用的是默认配置的Secret
func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//主要功能：生成JWT token
//依据客户端传入的Appkey和AppSecret在项目中设置的Issuer和ExpiresAt,生成token。
func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret()) 	//生成签名字串，根据传入得secret进行签名返回token
	return token, err
}

//解析、校验token：解析传入的Token，然后根据claims的属性要求进行校验。
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) { //用于解析权鉴声明
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims) //验证基于时间的声明。
		if ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}