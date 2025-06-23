package kogen

import (
	"fmt"
)

const (
	_UserEditorDatabaseNbr = 0
	_UserEditorTableName   = "USER_EDITOR"
)

func init() {
	ModelList = append(ModelList, &UserEditor{})
}

// UserEditor User editor
type UserEditor struct {
	CharId            [21]byte   `gorm:"column:strCharID;type:char(21);not null" json:"strCharID"`
	AccountId         [21]byte   `gorm:"column:strAccountID;type:char(21);not null" json:"strAccountID"`
	OpId              [21]byte   `gorm:"column:strOpID;type:char(21);not null" json:"strOpID"`
	OpIP              [21]byte   `gorm:"column:strOpIP;type:char(21);not null" json:"strOpIP"`
	OldUserValue      [600]byte  `gorm:"column:strOldUserValue;type:char(600);not null" json:"strOldUserValue"`
	NewUserValue      [600]byte  `gorm:"column:strNewUserValue;type:char(600);not null" json:"strNewUserValue"`
	OldUserSkill      [10]byte   `gorm:"column:strOldUserSkill;type:char(10);not null" json:"strOldUserSkill"`
	NewUserSkill      [10]byte   `gorm:"column:strNewUserSkill;type:char(10);not null" json:"strNewUserSkill"`
	OldUserItem       [400]byte  `gorm:"column:strOldUserItem;type:char(400);not null" json:"strOldUserItem"`
	NewUserItem       [400]byte  `gorm:"column:strNewUserItem;type:char(400);not null" json:"strNewUserItem"`
	OldWarehouseValue [100]byte  `gorm:"column:strOldWHValue;type:char(100);not null" json:"strOldWHValue"`
	NewWarehouseValue [100]byte  `gorm:"column:strNewWHValue;type:char(100);not null" json:"strNewWHValue"`
	OldWarehouseItem  [1600]byte `gorm:"column:strOldWHItem;type:char(1600);not null" json:"strOldWHItem"`
	NewWarehouseItem  [1600]byte `gorm:"column:strNewWHItem;type:char(1600);not null" json:"strNewWHItem"`
	EditorTime        int        `gorm:"column:EditorTime;type:smalldatetime;not null;default:getdate()" json:"EditorTime"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserEditor) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserEditorDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserEditor) GetTableName() string {
	return _UserEditorTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserEditor) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_EDITOR] (strCharID, strAccountID, strOpID, strOpIP, strOldUserValue, strNewUserValue, strOldUserSkill, strNewUserSkill, strOldUserItem, strNewUserItem, strOldWHValue, strNewWHValue, strOldWHItem, strNewWHItem, EditorTime) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.CharId),
		GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.OpId),
		GetOptionalBinaryVal(this.OpIP),
		GetOptionalBinaryVal(this.OldUserValue),
		GetOptionalBinaryVal(this.NewUserValue),
		GetOptionalBinaryVal(this.OldUserSkill),
		GetOptionalBinaryVal(this.NewUserSkill),
		GetOptionalBinaryVal(this.OldUserItem),
		GetOptionalBinaryVal(this.NewUserItem),
		GetOptionalBinaryVal(this.OldWarehouseValue),
		GetOptionalBinaryVal(this.NewWarehouseValue),
		GetOptionalBinaryVal(this.OldWarehouseItem),
		GetOptionalBinaryVal(this.NewWarehouseItem),
		GetOptionalDecVal(&this.EditorTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserEditor) GetCreateTableString() string {
	query := "CREATE TABLE [USER_EDITOR] (\n\t[strCharID] char(21) NOT NULL,\n\t[strAccountID] char(21) NOT NULL,\n\t[strOpID] char(21) NOT NULL,\n\t[strOpIP] char(21) NOT NULL,\n\t[strOldUserValue] char(600) NOT NULL,\n\t[strNewUserValue] char(600) NOT NULL,\n\t[strOldUserSkill] char(10) NOT NULL,\n\t[strNewUserSkill] char(10) NOT NULL,\n\t[strOldUserItem] char(400) NOT NULL,\n\t[strNewUserItem] char(400) NOT NULL,\n\t[strOldWHValue] char(100) NOT NULL,\n\t[strNewWHValue] char(100) NOT NULL,\n\t[strOldWHItem] char(1600) NOT NULL,\n\t[strNewWHItem] char(1600) NOT NULL,\n\t[EditorTime] smalldatetime NOT NULL\n\n)\nGO\nALTER TABLE [USER_EDITOR] ADD CONSTRAINT [DF_USER_EDITOR_EditorTime] DEFAULT getdate() FOR [EditorTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
