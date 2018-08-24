package model

type Category struct {
	Model
	CategoryId int `gorm:"primary_key" column:"category_id" json:"category_id"`
	Name string `json:"name"`
	ParentId int `json:"parent_id"`
	Status int `json:"status"`
	ImageId int `json:"image_id"`

	//外键
	ImageFile ImageFile `gorm:"ForeignKey:ImageId"`

	Good []Good `gorm:"ForeignKey:CategoryId"`
}

func (Category) TableName() string  {
	return "category"
}

func CategoryList() (category []Category) {
	db.Preload("ImageFile").Preload("Good").Where("status=?",1).Find(&category)
	return
}

//根据分类编号返回商品列表
func CategoryGoods(categoryId int) (category Category){
	//必须显示的设置表名 不知道什么鬼
	db.Table("rw_category").Preload("ImageFile").Preload("Good").Where("status=?",1).Where("category_id=?",categoryId).Find(&category)
	return
}