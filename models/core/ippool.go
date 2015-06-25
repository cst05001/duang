package core

/*
	此功能仅限用bridge独立IP方式运行container使用
*/
type IpPool struct {
	Id     int64
	IP     string `orm:"unique"`
	Status uint8  `orm:"default(1)"`
}

//获取可用IP
func (this *IpPool) GetFreeIP(n int) []string {
	return nil
}

//获取所有IP
func (this *IpPool) GetAllIP() []string {
	return nil
}

//获取被占用的IP
func (this *IpPool) GetUsedIP() []string {
	return nil
}

//把新IP加入池
func (this *IpPool) AddIP(ip string) error {
	return nil
}

//从IP池删除IP
func (this *IpPool) DelIP(ip string) error {
	return nil
}
