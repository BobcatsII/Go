package timer

import "time"

//获取时间
func GetNowTime() time.Time {
	return time.Now()
}

//推算时间
func GetCalculateTime(currentTimer time.Time, d string) (time.Time, error) {
	//调用ParseDuration方法，用于从字符串解析出duration(持续时间)，其支持的有效单位有ns\us\ms\s\m\h
	duration, err := time.ParseDuration(d)
	if err != nil{
		return time.Time{}, err
	}
	//调用Add方法,可以传入其返回的duration，可以得到当前Timer时间加上duration后所得到的的最终时间。
	return currentTimer.Add(duration), nil
}