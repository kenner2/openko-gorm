package kogen

import (
	"fmt"
)

const (
	_MagicType9DatabaseNbr = 0
	_MagicType9TableName   = "MAGIC_TYPE9"
)

func init() {
	ModelList = append(ModelList, &MagicType9{})
}

// MagicType9 Type 9 supports stealth and detection abilities
type MagicType9 struct {
	MagicNumber   int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name          [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description   [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	ValidGroup    uint8     `gorm:"column:ValidGroup;type:tinyint;not null" json:"ValidGroup"`
	NationChange  uint8     `gorm:"column:NationChange;type:tinyint;not null" json:"NationChange"`
	MonsterNumber int16     `gorm:"column:MonsterNum;type:smallint;not null" json:"MonsterNum"`
	TargetChange  uint8     `gorm:"column:TargetChange;type:tinyint;not null" json:"TargetChange"`
	StateChange   uint8     `gorm:"column:StateChange;type:tinyint;not null" json:"StateChange"`
	Radius        int16     `gorm:"column:Radius;type:smallint;not null" json:"Radius"`
	HitRate       int16     `gorm:"column:Hitrate;type:smallint;not null" json:"Hitrate"`
	Duration      int16     `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	AddDamage     int16     `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	Vision        int16     `gorm:"column:Vision;type:smallint;not null" json:"Vision"`
	NeedItem      int       `gorm:"column:NeedItem;type:int;not null" json:"NeedItem"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType9) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType9DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType9) GetTableName() string {
	return _MagicType9TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType9) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE9] (iNum, Name, Description, ValidGroup, NationChange, MonsterNum, TargetChange, StateChange, Radius, Hitrate, Duration, AddDamage, Vision, NeedItem) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType9) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE9] (\n\t\"iNum\" int NOT NULL,\n\t\"Name\" varchar(50),\n\t\"Description\" varchar(100),\n\t\"ValidGroup\" tinyint NOT NULL,\n\t\"NationChange\" tinyint NOT NULL,\n\t\"MonsterNum\" smallint NOT NULL,\n\t\"TargetChange\" tinyint NOT NULL,\n\t\"StateChange\" tinyint NOT NULL,\n\t\"Radius\" smallint NOT NULL,\n\t\"Hitrate\" smallint NOT NULL,\n\t\"Duration\" smallint NOT NULL,\n\t\"AddDamage\" smallint NOT NULL,\n\t\"Vision\" smallint NOT NULL,\n\t\"NeedItem\" int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
