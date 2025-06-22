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

// Coefficient Coefficient relationship between a character class, weapon types, and stats
type Coefficient struct {
	ClassId     int16    `gorm:"column:sClass;type:smallint;not null" json:"sClass"`
	ShortSword  float64  `gorm:"column:ShortSword;type:float;not null" json:"ShortSword"`
	Sword       float64  `gorm:"column:Sword;type:float;not null" json:"Sword"`
	Axe         float64  `gorm:"column:Axe;type:float;not null" json:"Axe"`
	Club        float64  `gorm:"column:Club;type:float;not null" json:"Club"`
	Spear       float64  `gorm:"column:Spear;type:float;not null" json:"Spear"`
	Pole        float64  `gorm:"column:Pole;type:float;not null" json:"Pole"`
	Staff       float64  `gorm:"column:Staff;type:float;not null" json:"Staff"`
	Bow         *float64 `gorm:"column:Bow;type:float" json:"Bow,omitempty"`
	HitPoint    float64  `gorm:"column:Hp;type:float;not null" json:"Hp"`
	MagicPower  float64  `gorm:"column:Mp;type:float;not null" json:"Mp"`
	Sp          float64  `gorm:"column:Sp;type:float;not null" json:"Sp"`
	Armor       float64  `gorm:"column:Ac;type:float;not null" json:"Ac"`
	HitRate     float64  `gorm:"column:Hitrate;type:float;not null" json:"Hitrate"`
	Evasionrate float64  `gorm:"column:Evasionrate;type:float;not null" json:"Evasionrate"`
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

// GetCreateTableString Returns the create table statement for this object
func (this *Coefficient) GetCreateTableString() string {
	query := "CREATE TABLE [COEFFICIENT] (\n\t\"sClass\" smallint NOT NULL,\n\t\"ShortSword\" float NOT NULL,\n\t\"Sword\" float NOT NULL,\n\t\"Axe\" float NOT NULL,\n\t\"Club\" float NOT NULL,\n\t\"Spear\" float NOT NULL,\n\t\"Pole\" float NOT NULL,\n\t\"Staff\" float NOT NULL,\n\t\"Bow\" float,\n\t\"Hp\" float NOT NULL,\n\t\"Mp\" float NOT NULL,\n\t\"Sp\" float NOT NULL,\n\t\"Ac\" float NOT NULL,\n\t\"Hitrate\" float NOT NULL,\n\t\"Evasionrate\" float NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
