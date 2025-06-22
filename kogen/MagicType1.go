package kogen

import (
	"fmt"
)

const (
	_MagicType1DatabaseNbr = 0
	_MagicType1TableName   = "MAGIC_TYPE1"
)

func init() {
	ModelList = append(ModelList, &MagicType1{})
}

// MagicType1 Type 1 covers melee abilities
type MagicType1 struct {
	MagicNumber int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name        [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Type        uint8     `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	HitRateMod  int16     `gorm:"column:HitRate;type:smallint;not null" json:"HitRate"`
	DamageMod   int16     `gorm:"column:Hit;type:smallint;not null" json:"Hit"`
	AddDamage   int16     `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	Delay       uint8     `gorm:"column:Delay;type:tinyint;not null" json:"Delay"`
	ComboType   uint8     `gorm:"column:ComboType;type:tinyint;not null" json:"ComboType"`
	ComboCount  uint8     `gorm:"column:ComboCount;type:tinyint;not null" json:"ComboCount"`
	ComboDamage int16     `gorm:"column:ComboDamage;type:smallint;not null" json:"ComboDamage"`
	Range       int16     `gorm:"column:Range;type:smallint;not null" json:"Range"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType1) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType1DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType1) GetTableName() string {
	return _MagicType1TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType1) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE1] (iNum, Name, Description, Type, HitRate, Hit, AddDamage, Delay, ComboType, ComboCount, ComboDamage, Range) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Delay),
		GetOptionalDecVal(&this.ComboType),
		GetOptionalDecVal(&this.ComboCount),
		GetOptionalDecVal(&this.ComboDamage),
		GetOptionalDecVal(&this.Range))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType1) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE1] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[Type] tinyint NOT NULL,\n\t[HitRate] smallint NOT NULL,\n\t[Hit] smallint NOT NULL,\n\t[AddDamage] smallint NOT NULL,\n\t[Delay] tinyint NOT NULL,\n\t[ComboType] tinyint NOT NULL,\n\t[ComboCount] tinyint NOT NULL,\n\t[ComboDamage] smallint NOT NULL,\n\t[Range] smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
