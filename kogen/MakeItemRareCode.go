package kogen

import (
	"fmt"
)

const (
	_MakeItemRareCodeDatabaseNbr = 0
	_MakeItemRareCodeTableName   = "MAKE_ITEM_LARECODE"
)

func init() {
	ModelList = append(ModelList, &MakeItemRareCode{})
}

// MakeItemRareCode Make item rarity codes
type MakeItemRareCode struct {
	LevelGrade  uint8 `gorm:"column:byLevelGrade;type:tinyint;not null" json:"byLevelGrade"`
	UpgradeItem int16 `gorm:"column:sUpgradeItem;type:smallint;not null;default:0" json:"sUpgradeItem"`
	RareItem    int16 `gorm:"column:sLareItem;type:smallint;not null;default:0" json:"sLareItem"`
	MagicItem   int16 `gorm:"column:sMagicItem;type:smallint;not null;default:0" json:"sMagicItem"`
	GeneralItem int16 `gorm:"column:sGereralItem;type:smallint;not null;default:0" json:"sGereralItem"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MakeItemRareCode) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeItemRareCodeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MakeItemRareCode) GetTableName() string {
	return _MakeItemRareCodeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MakeItemRareCode) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM_LARECODE] (byLevelGrade, sUpgradeItem, sLareItem, sMagicItem, sGereralItem) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.LevelGrade),
		GetOptionalDecVal(&this.UpgradeItem),
		GetOptionalDecVal(&this.RareItem),
		GetOptionalDecVal(&this.MagicItem),
		GetOptionalDecVal(&this.GeneralItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MakeItemRareCode) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM_LARECODE] (\n\t\"byLevelGrade\" tinyint NOT NULL,\n\t\"sUpgradeItem\" smallint NOT NULL,\n\t\"sLareItem\" smallint NOT NULL,\n\t\"sMagicItem\" smallint NOT NULL,\n\t\"sGereralItem\" smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
