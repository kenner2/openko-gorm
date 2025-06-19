package kogen

import (
	"fmt"
)

const (
	_CopySerialItemDatabaseNbr = 0
	_CopySerialItemTableName   = "COPY_SERIAL_ITEM"
)

// CopySerialItem: Represents the relationship between accounts and characters
type CopySerialItem struct {
	AccountId  *string `gorm:"column:strAccountID" json:"strAccountID,omitempty"`
	Type       *uint8  `gorm:"column:byType" json:"byType,omitempty"`
	Pos        *int16  `gorm:"column:nPos" json:"nPos,omitempty"`
	ItemNum    *Binary `gorm:"column:ItemNum;type:binary(4)" json:"ItemNum,omitempty"`
	ItemSerial *Binary `gorm:"column:ItemSerial;type:binary(8)" json:"ItemSerial,omitempty"`
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
	return fmt.Sprintf("INSERT INTO [COPY_SERIAL_ITEM] (strAccountID, byType, nPos, ItemNum, ItemSerial) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalStringVal(this.AccountId),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNum),
		GetOptionalBinaryVal(this.ItemSerial))
}
