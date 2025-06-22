package kogen

import (
	"fmt"
)

const (
	_MagicType7DatabaseNbr = 0
	_MagicType7TableName   = "MAGIC_TYPE7"
)

func init() {
	ModelList = append(ModelList, &MagicType7{})
}

// MagicType7 Type 7 supports targeting modifications
type MagicType7 struct {
	MagicNumber   int       `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	Name          [50]byte  `gorm:"column:strName;type:varchar(50)" json:"strName,omitempty"`
	Note          [100]byte `gorm:"column:strNote;type:varchar(100)" json:"strNote,omitempty"`
	ValidGroup    uint8     `gorm:"column:byValidGroup;type:tinyint;not null" json:"byValidGroup"`
	NationChange  uint8     `gorm:"column:byNatoinChange;type:tinyint;not null" json:"byNatoinChange"`
	MonsterNumber int16     `gorm:"column:shMonsterNum;type:smallint;not null" json:"shMonsterNum"`
	TargetChange  uint8     `gorm:"column:byTargetChange;type:tinyint;not null" json:"byTargetChange"`
	StateChange   uint8     `gorm:"column:byStateChange;type:tinyint;not null" json:"byStateChange"`
	Radius        uint8     `gorm:"column:byRadius;type:tinyint;not null" json:"byRadius"`
	HitRate       int16     `gorm:"column:shHitrate;type:smallint;not null" json:"shHitrate"`
	Duration      int16     `gorm:"column:shDuration;type:smallint;not null" json:"shDuration"`
	Damage        int16     `gorm:"column:shDamage;type:smallint;not null" json:"shDamage"`
	Vision        uint8     `gorm:"column:byVisoin;type:tinyint;not null" json:"byVisoin"`
	NeedItem      int       `gorm:"column:nNeedItem;type:int;not null" json:"nNeedItem"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType7) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType7DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType7) GetTableName() string {
	return _MagicType7TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType7) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE7] (nIndex, strName, strNote, byValidGroup, byNatoinChange, shMonsterNum, byTargetChange, byStateChange, byRadius, shHitrate, shDuration, shDamage, byVisoin, nNeedItem) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Note),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType7) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE7] (\n\t[nIndex] int NOT NULL,\n\t[strName] varchar(50),\n\t[strNote] varchar(100),\n\t[byValidGroup] tinyint NOT NULL,\n\t[byNatoinChange] tinyint NOT NULL,\n\t[shMonsterNum] smallint NOT NULL,\n\t[byTargetChange] tinyint NOT NULL,\n\t[byStateChange] tinyint NOT NULL,\n\t[byRadius] tinyint NOT NULL,\n\t[shHitrate] smallint NOT NULL,\n\t[shDuration] smallint NOT NULL,\n\t[shDamage] smallint NOT NULL,\n\t[byVisoin] tinyint NOT NULL,\n\t[nNeedItem] int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
