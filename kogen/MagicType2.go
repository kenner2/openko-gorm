package kogen

import (
	"fmt"
)

const (
	_MagicType2DatabaseNbr = 0
	_MagicType2TableName   = "MAGIC_TYPE2"
)

func init() {
	ModelList = append(ModelList, &MagicType2{})
}

// MagicType2 Type 2 covers bow abilities
type MagicType2 struct {
	MagicNumber   int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name          [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description   [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	HitType       uint8     `gorm:"column:HitType;type:tinyint;not null" json:"HitType"`
	HitRateMod    int16     `gorm:"column:HitRate;type:smallint;not null" json:"HitRate"`
	DamageMod     int16     `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	RangeMod      int16     `gorm:"column:AddRange;type:smallint;not null" json:"AddRange"`
	NeedArrow     uint8     `gorm:"column:NeedArrow;type:tinyint;not null" json:"NeedArrow"`
	AddDamagePlus *int16    `gorm:"column:AddDamagePlus;type:smallint" json:"AddDamagePlus,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType2) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType2DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType2) GetTableName() string {
	return _MagicType2TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType2) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE2] (iNum, Name, Description, HitType, HitRate, AddDamage, AddRange, NeedArrow, AddDamagePlus) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.HitType),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.RangeMod),
		GetOptionalDecVal(&this.NeedArrow),
		GetOptionalDecVal(this.AddDamagePlus))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType2) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE2] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[HitType] tinyint NOT NULL,\n\t[HitRate] smallint NOT NULL,\n\t[AddDamage] smallint NOT NULL,\n\t[AddRange] smallint NOT NULL,\n\t[NeedArrow] tinyint NOT NULL,\n\t[AddDamagePlus] smallint\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
