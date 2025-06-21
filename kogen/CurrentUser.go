package kogen

import (
	"fmt"
)

const (
	_CurrentUserDatabaseNbr = 0
	_CurrentUserTableName   = "CURRENTUSER"
)

func init() {
	ModelList = append(ModelList, &CurrentUser{})
}

// CurrentUser: Keeps track of users currently connected to the server
type CurrentUser struct {
	ServerNumber int      `gorm:"column:nServerNo;type:int;not null" json:"nServerNo"`
	ServerIP     [20]byte `gorm:"column:strServerIP;type:varchar(20);not null" json:"strServerIP"`
	AccountId    [20]byte `gorm:"column:strAccountID;type:varchar(20);not null" json:"strAccountID"`
	CharId       [20]byte `gorm:"column:strCharID;type:varchar(20);not null" json:"strCharID"`
	ClientIP     [20]byte `gorm:"column:strClientIP;type:varchar(20);not null" json:"strClientIP"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *CurrentUser) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CurrentUserDatabaseNbr))
}

// GetTableName Returns the table name
func (this *CurrentUser) GetTableName() string {
	return _CurrentUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *CurrentUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [CURRENTUSER] (nServerNo, strServerIP, strAccountID, strCharID, strClientIP) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerNumber),
		GetOptionalBinaryVal(this.ServerIP),
		GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.CharId),
		GetOptionalBinaryVal(this.ClientIP))
}

// GetCreateTableString Returns the create table statement for this object
func (this *CurrentUser) GetCreateTableString() string {
	query := "CREATE TABLE \"CURRENTUSER\" (\n\t\"nServerNo\" int NOT NULL,\n\t\"strServerIP\" varchar(20) NOT NULL,\n\t\"strAccountID\" varchar(20) NOT NULL,\n\t\"strCharID\" varchar(20) NOT NULL,\n\t\"strClientIP\" varchar(20) NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
