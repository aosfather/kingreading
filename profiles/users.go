package profiles

//用户
type User struct {
	ID    string   //唯一用户ID
	Nick  string   //用户显示名称
	Name  string   //用户真实姓名
	Email string   //用户联系用邮箱
	QQ    string   //用户QQ
	Tags  []string //用户标签

}

//标签
type Tag struct {
	Catalog string //类别
	Name    string //标签名称
	Label   string //标签显示名称
}
