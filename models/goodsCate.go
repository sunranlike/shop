package models

type GoodsCate struct {
	Id          int
	Title       string
	CateImg     string
	Link        string
	Template    string
	Pid         int
	SubTitle    string
	Keywords    string
	Description string
	Sort        int
	Status      int
	AddTime     int
	//这里是怎么配置的?其实就是自关联,外键关联的表示goodCate(其实就是自己),自己的主键就是ID(如果不写reference,默认也是主键)
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}
