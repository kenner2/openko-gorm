package kogen

import (
	"fmt"
)

const (
	_CoefficientDatabaseNbr = 0
	_CoefficientTableName   = "COEFFICIENT"
)

func init() {
	ModelList = append(ModelList, &Coefficient{})
}

// Coefficient: Coefficient relationship between a character class, weapon types, and stats
type Coefficient struct {
	ClassId     int16    `gorm:"column:sClass;not null" json:"sClass"`
	ShortSword  float32  `gorm:"column:ShortSword;not null" json:"ShortSword"`
	Sword       float32  `gorm:"column:Sword;not null" json:"Sword"`
	Axe         float32  `gorm:"column:Axe;not null" json:"Axe"`
	Club        float32  `gorm:"column:Club;not null" json:"Club"`
	Spear       float32  `gorm:"column:Spear;not null" json:"Spear"`
	Pole        float32  `gorm:"column:Pole;not null" json:"Pole"`
	Staff       float32  `gorm:"column:Staff;not null" json:"Staff"`
	Bow         *float32 `gorm:"column:Bow" json:"Bow,omitempty"`
	HitPoint    float32  `gorm:"column:Hp;not null" json:"Hp"`
	MagicPower  float32  `gorm:"column:Mp;not null" json:"Mp"`
	Sp          float32  `gorm:"column:Sp;not null" json:"Sp"`
	Armor       float32  `gorm:"column:Ac;not null" json:"Ac"`
	HitRate     float32  `gorm:"column:Hitrate;not null" json:"Hitrate"`
	Evasionrate float32  `gorm:"column:Evasionrate;not null" json:"Evasionrate"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Coefficient) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CoefficientDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Coefficient) GetTableName() string {
	return _CoefficientTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Coefficient) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COEFFICIENT] (sClass, ShortSword, Sword, Axe, Club, Spear, Pole, Staff, Bow, Hp, Mp, Sp, Ac, Hitrate, Evasionrate) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ClassId),
		GetOptionalDecVal(&this.ShortSword),
		GetOptionalDecVal(&this.Sword),
		GetOptionalDecVal(&this.Axe),
		GetOptionalDecVal(&this.Club),
		GetOptionalDecVal(&this.Spear),
		GetOptionalDecVal(&this.Pole),
		GetOptionalDecVal(&this.Staff),
		GetOptionalDecVal(this.Bow),
		GetOptionalDecVal(&this.HitPoint),
		GetOptionalDecVal(&this.MagicPower),
		GetOptionalDecVal(&this.Sp),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Evasionrate))
}
