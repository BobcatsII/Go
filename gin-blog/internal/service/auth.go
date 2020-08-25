package service

import "errors"

//AuthRequest 用于接口入参校验，required表示这是必填项
type AuthRequest struct {
	AppKey    string `form:"app_key" binding:"required"`
	AppSecret string `form:"app_secret" binding:"required"`
}

//对相应的基本逻辑进行处理
func (svc *Service) CheckAuth(param *AuthRequest) error {
	//根据是否取到认证信息ID判定认证信息ID是否存在。
	auth, err := svc.dao.GetAuth(
		param.AppKey,
		param.AppSecret,
	)
	if err != nil {
		return err
	}
	if auth.ID > 0 {
		return nil
	}
	return errors.New("auth info does not exist.")
}