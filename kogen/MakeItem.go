package kogen

import (
	"fmt"
)

const (
	_MakeItemDatabaseNbr = 0
	_MakeItemTableName   = "MAKE_ITEM"
)

func init() {
	ModelList = append(ModelList, &MakeItem{})
}

// MakeItem Make item
type MakeItem struct {
	Index     int16    `gorm:"column:sIndex;type:smallint;not null" json:"sIndex"`
	ItemInfo  [20]byte `gorm:"column:strItemInfo;type:varchar(20)" json:"strItemInfo,omitempty"`
	ItemCode  int      `gorm:"column:iItemCode;type:int;not null" json:"iItemCode"`
	ItemLevel uint8    `gorm:"column:byItemLevel;type:tinyint;not null;default:0" json:"byItemLevel"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MakeItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MakeItem) GetTableName() string {
	return _MakeItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MakeItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM] (sIndex, strItemInfo, iItemCode, byItemLevel) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalBinaryVal(this.ItemInfo),
		GetOptionalDecVal(&this.ItemCode),
		GetOptionalDecVal(&this.ItemLevel))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MakeItem) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM] (\n\t[sIndex] smallint NOT NULL,\n\t[strItemInfo] varchar(20),\n\t[iItemCode] int NOT NULL,\n\t[byItemLevel] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
