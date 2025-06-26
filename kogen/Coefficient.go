package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_CoefficientDatabaseNbr = 1
	_CoefficientTableName   = "COEFFICIENT"
)

func init() {
	ModelList = append(ModelList, &Coefficient{})
}

// Coefficient Coefficient relationship between a character class, weapon types, and stats
type Coefficient struct {
	ClassId     int16    `gorm:"column:sClass;type:smallint;primaryKey;not null" json:"sClass"`
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

// GetDatabaseName Returns the table's database name
func (this Coefficient) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CoefficientDatabaseNbr))
}

// TableName Returns the table name
func (this Coefficient) TableName() string {
	return _CoefficientTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Coefficient) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COEFFICIENT] ([sClass], [ShortSword], [Sword], [Axe], [Club], [Spear], [Pole], [Staff], [Bow], [Hp], [Mp], [Sp], [Ac], [Hitrate], [Evasionrate]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ClassId),
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

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Coefficient) GetInsertHeader() string {
	return "INSERT INTO [COEFFICIENT] (sClass, ShortSword, Sword, Axe, Club, Spear, Pole, Staff, Bow, Hp, Mp, Sp, Ac, Hitrate, Evasionrate) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Coefficient) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ClassId),
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
func (this Coefficient) GetCreateTableString() string {
	query := "CREATE TABLE [COEFFICIENT] (\n\t[sClass] smallint NOT NULL,\n\t[ShortSword] float NOT NULL,\n\t[Sword] float NOT NULL,\n\t[Axe] float NOT NULL,\n\t[Club] float NOT NULL,\n\t[Spear] float NOT NULL,\n\t[Pole] float NOT NULL,\n\t[Staff] float NOT NULL,\n\t[Bow] float,\n\t[Hp] float NOT NULL,\n\t[Mp] float NOT NULL,\n\t[Sp] float NOT NULL,\n\t[Ac] float NOT NULL,\n\t[Hitrate] float NOT NULL,\n\t[Evasionrate] float NOT NULL\n\tCONSTRAINT [PK_COEFFICIENT] PRIMARY KEY ([sClass])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Coefficient) SelectClause() (selectClause clause.Select) {
	return _CoefficientSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Coefficient) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Coefficient{}
	rawSql := "SELECT [sClass], [ShortSword], [Sword], [Axe], [Club], [Spear], [Pole], [Staff], [Bow], [Hp], [Mp], [Sp], [Ac], [Hitrate], [Evasionrate] FROM [COEFFICIENT]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _CoefficientSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sClass]",
		},
		clause.Column{
			Name: "[ShortSword]",
		},
		clause.Column{
			Name: "[Sword]",
		},
		clause.Column{
			Name: "[Axe]",
		},
		clause.Column{
			Name: "[Club]",
		},
		clause.Column{
			Name: "[Spear]",
		},
		clause.Column{
			Name: "[Pole]",
		},
		clause.Column{
			Name: "[Staff]",
		},
		clause.Column{
			Name: "[Bow]",
		},
		clause.Column{
			Name: "[Hp]",
		},
		clause.Column{
			Name: "[Mp]",
		},
		clause.Column{
			Name: "[Sp]",
		},
		clause.Column{
			Name: "[Ac]",
		},
		clause.Column{
			Name: "[Hitrate]",
		},
		clause.Column{
			Name: "[Evasionrate]",
		},
	},
}
