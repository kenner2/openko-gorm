package kogen

import (
	"fmt"
)

const (
	_MonsterSummonListDatabaseNbr = 0
	_MonsterSummonListTableName   = "MONSTER_SUMMON_LIST"
)

func init() {
	ModelList = append(ModelList, &MonsterSummonList{})
}

// MonsterSummonList Monster summon list
type MonsterSummonList struct {
	MonsterId   int16    `gorm:"column:sSid;type:smallint;not null" json:"sSid"`
	Name        [31]byte `gorm:"column:strName;type:varchar(31)" json:"strName,omitempty"`
	Level       int16    `gorm:"column:sLevel;type:smallint;not null" json:"sLevel"`
	Probability int16    `gorm:"column:sProbability;type:smallint;not null" json:"sProbability"`
	Type        uint8    `gorm:"column:bType;type:tinyint;not null;default:0" json:"bType"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MonsterSummonList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MonsterSummonListDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MonsterSummonList) GetTableName() string {
	return _MonsterSummonListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MonsterSummonList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_SUMMON_LIST] (sSid, strName, sLevel, sProbability, bType) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MonsterId),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Probability),
		GetOptionalDecVal(&this.Type))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MonsterSummonList) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_SUMMON_LIST] (\n\t[sSid] smallint NOT NULL,\n\t[strName] varchar(31),\n\t[sLevel] smallint NOT NULL,\n\t[sProbability] smallint NOT NULL,\n\t[bType] tinyint NOT NULL\n\n)\nGO\nALTER TABLE [MONSTER_SUMMON_LIST] ADD CONSTRAINT [DF_MONSTER_SUMMON_LIST_bType] DEFAULT 0 FOR [bType]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
