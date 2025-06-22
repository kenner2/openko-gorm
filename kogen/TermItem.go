package kogen

import (
	"fmt"
)

const (
	_TermItemDatabaseNbr = 0
	_TermItemTableName   = "TERM_ITEM"
)

func init() {
	ModelList = append(ModelList, &TermItem{})
}

// TermItem Term item
type TermItem struct {
	UserId     [21]byte `gorm:"column:strUserId;type:char(21)" json:"strUserId,omitempty"`
	Type       *uint8   `gorm:"column:byType;type:tinyint" json:"byType,omitempty"`
	Pos        *int16   `gorm:"column:nPos;type:smallint" json:"nPos,omitempty"`
	ItemNumber [4]byte  `gorm:"column:ItemNum;type:binary(4)" json:"ItemNum,omitempty"`
	ItemSerial [8]byte  `gorm:"column:ItemSerial;type:binary(8)" json:"ItemSerial,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *TermItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_TermItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *TermItem) GetTableName() string {
	return _TermItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *TermItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [TERM_ITEM] (strUserId, byType, nPos, ItemNum, ItemSerial) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNumber),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this *TermItem) GetCreateTableString() string {
	query := "CREATE TABLE [TERM_ITEM] (\n\t\"strUserId\" char(21),\n\t\"byType\" tinyint,\n\t\"nPos\" smallint,\n\t\"ItemNum\" binary(4),\n\t\"ItemSerial\" binary(8)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
