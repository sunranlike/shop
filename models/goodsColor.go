package models

type GoodsColor struct {
	Id         int
	ColorName  string
	ColorValue string
	Status     int
	Checked    bool `gorm:"-"` // 忽略本字段,这个字段就是为了方便后端的渲染
}

func (GoodsColor) TableName() string {
	return "goods_color"
}
