package kogen

import (
	"fmt"
)

const (
	_MagicType3DatabaseNbr = 0
	_MagicType3TableName   = "MAGIC_TYPE3"
)

func init() {
	ModelList = append(ModelList, &MagicType3{})
}

// MagicType3 Type 3 supports Area of Effect and Damage over Time effects
type MagicType3 struct {
	MagicNumber int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name        [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Radius      uint8     `gorm:"column:Radius;type:tinyint;not null" json:"Radius"`
	Angle       int16     `gorm:"column:Angle;type:smallint;not null" json:"Angle"`
	DirectType  uint8     `gorm:"column:DirectType;type:tinyint;not null" json:"DirectType"`
	FirstDamage int16     `gorm:"column:FirstDamage;type:smallint;not null" json:"FirstDamage"`
	EndDamage   int16     `gorm:"column:EndDamage;type:smallint;not null" json:"EndDamage"`
	TimeDamage  int16     `gorm:"column:TimeDamage;type:smallint;not null" json:"TimeDamage"`
	Duration    uint8     `gorm:"column:Duration;type:tinyint;not null" json:"Duration"`
	Attribute   uint8     `gorm:"column:Attribute;type:tinyint;not null" json:"Attribute"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType3) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType3DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType3) GetTableName() string {
	return _MagicType3TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType3) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE3] (iNum, Name, Description, Radius, Angle, DirectType, FirstDamage, EndDamage, TimeDamage, Duration, Attribute) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.Angle),
		GetOptionalDecVal(&this.DirectType),
		GetOptionalDecVal(&this.FirstDamage),
		GetOptionalDecVal(&this.EndDamage),
		GetOptionalDecVal(&this.TimeDamage),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Attribute))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType3) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE3] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[Radius] tinyint NOT NULL,\n\t[Angle] smallint NOT NULL,\n\t[DirectType] tinyint NOT NULL,\n\t[FirstDamage] smallint NOT NULL,\n\t[EndDamage] smallint NOT NULL,\n\t[TimeDamage] smallint NOT NULL,\n\t[Duration] tinyint NOT NULL,\n\t[Attribute] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
