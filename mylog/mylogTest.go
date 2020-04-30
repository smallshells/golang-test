package mylog


func TestMyLog(){
	f1 := NewFileLogger(DEBUG,"./mylog","test.log")
	userId := 10
	f1.Debug("id是%d的用户一直尝试登陆",userId)
}