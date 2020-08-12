package global

import "gin-blog/pkg/setting"

//将配置信息和应用程序关联
var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
)
