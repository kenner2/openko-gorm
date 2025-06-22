package kogen

import (
	"fmt"
)

const (
	_UserEditorItemDatabaseNbr = 0
	_UserEditorItemTableName   = "USER_EDITOR_ITEM"
)

func init() {
	ModelList = append(ModelList, &UserEditorItem{})
}

// UserEditorItem User editor item
type UserEditorItem struct {
	CharId     [21]byte `gorm:"column:strCharID;type:char(21);not null" json:"strCharID"`
	AccountId  [21]byte `gorm:"column:strAccountID;type:char(21);not null" json:"strAccountID"`
	OpId       [21]byte `gorm:"column:strOpID;type:char(21);not null" json:"strOpID"`
	OpIP       [21]byte `gorm:"column:strOpIP;type:char(21);not null" json:"strOpIP"`
	DbIndex    int16    `gorm:"column:sDBIndex;type:smallint;not null" json:"sDBIndex"`
	Pos        int16    `gorm:"column:sPos;type:smallint;not null" json:"sPos"`
	Type       uint8    `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	ItemId1    int      `gorm:"column:nItemID1;type:int;not null" json:"nItemID1"`
	ItemId2    int      `gorm:"column:nItemID2;type:int;not null" json:"nItemID2"`
	UpdateTime *int     `gorm:"column:UpdateTime;type:smalldatetime" json:"UpdateTime,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserEditorItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserEditorItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserEditorItem) GetTableName() string {
	return _UserEditorItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserEditorItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_EDITOR_ITEM] (strCharID, strAccountID, strOpID, strOpIP, sDBIndex, sPos, byType, nItemID1, nItemID2, UpdateTime) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.CharId),
		GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.OpId),
		GetOptionalBinaryVal(this.OpIP),
		GetOptionalDecVal(&this.DbIndex),
		GetOptionalDecVal(&this.Pos),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ItemId1),
		GetOptionalDecVal(&this.ItemId2),
		GetOptionalDecVal(this.UpdateTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserEditorItem) GetCreateTableString() string {
	query := "CREATE TABLE [USER_EDITOR_ITEM] (\n\t[strCharID] char(21) NOT NULL,\n\t[strAccountID] char(21) NOT NULL,\n\t[strOpID] char(21) NOT NULL,\n\t[strOpIP] char(21) NOT NULL,\n\t[sDBIndex] smallint NOT NULL,\n\t[sPos] smallint NOT NULL,\n\t[byType] tinyint NOT NULL,\n\t[nItemID1] int NOT NULL,\n\t[nItemID2] int NOT NULL,\n\t[UpdateTime] smalldatetime\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
