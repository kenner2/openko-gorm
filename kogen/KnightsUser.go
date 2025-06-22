package kogen

import (
	"fmt"
)

const (
	_KnightsUserDatabaseNbr = 0
	_KnightsUserTableName   = "KNIGHTS_USER"
)

func init() {
	ModelList = append(ModelList, &KnightsUser{})
}

// KnightsUser Knights to character relationships
type KnightsUser struct {
	KnightsId int16    `gorm:"column:sIDNum;type:smallint;not null" json:"sIDNum"`
	UserId    [21]byte `gorm:"column:strUserID;type:varchar(21)" json:"strUserID,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KnightsUser) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsUserDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KnightsUser) GetTableName() string {
	return _KnightsUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KnightsUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_USER] (sIDNum, strUserID) \nVALUES (%s, %s)", GetOptionalDecVal(&this.KnightsId),
		GetOptionalBinaryVal(this.UserId))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KnightsUser) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_USER] (\n\t\"sIDNum\" smallint NOT NULL,\n\t\"strUserID\" varchar(21)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
