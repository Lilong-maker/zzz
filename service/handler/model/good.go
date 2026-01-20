package model

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name string `grom:"type:varchar(30)"`
	Num  int    `grom:"type:int(11)"`
}

func (g *Goods) Findgood(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&g).Error
}

func (g *Goods) GoodAdd(db *gorm.DB) error {
	return db.Debug().Create(&g).Error
}

func (g *Goods) FindGoods(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Find(&g).Error
}

func (g *Goods) UpdateGoods(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Updates(&g).Error
}

func (g *Goods) FindGoodsList(db *gorm.DB, id int32) error {
	return db.Debug().Find(&g).Error
}

func (g *Goods) DeleteGoods(db *gorm.DB, id int32) error {
	return db.Debug().Where("id = ?", id).Delete(&g).Error
}
