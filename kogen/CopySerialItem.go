package kogen

import (
	"fmt"
)

const (
	_CopySerialItemDatabaseNbr = 0
	_CopySerialItemTableName   = "COPY_SERIAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &CopySerialItem{})
}

// CopySerialItem TODO: Doc
type CopySerialItem struct {
	UserId     [21]byte `gorm:"column:strUserId;type:char(21)" json:"strUserId,omitempty"`
	Type       *uint8   `gorm:"column:byType;type:tinyint" json:"byType,omitempty"`
	Pos        *int16   `gorm:"column:nPos;type:smallint" json:"nPos,omitempty"`
	ItemNum    [4]byte  `gorm:"column:ItemNum;type:binary(4)" json:"ItemNum,omitempty"`
	ItemSerial [8]byte  `gorm:"column:ItemSerial;type:binary(8)" json:"ItemSerial,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *CopySerialItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CopySerialItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *CopySerialItem) GetTableName() string {
	return _CopySerialItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *CopySerialItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COPY_SERIAL_ITEM] (strUserId, byType, nPos, ItemNum, ItemSerial) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNum),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this *CopySerialItem) GetCreateTableString() string {
	query := "CREATE TABLE [COPY_SERIAL_ITEM] (\n\t\"strUserId\" char(21),\n\t\"byType\" tinyint,\n\t\"nPos\" smallint,\n\t\"ItemNum\" binary(4),\n\t\"ItemSerial\" binary(8)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
