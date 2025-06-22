package kogen

import (
	"fmt"
)

const (
	_ProgramListLogDatabaseNbr = 0
	_ProgramListLogTableName   = "PROGRAMLIST_LOG"
)

func init() {
	ModelList = append(ModelList, &ProgramListLog{})
}

// ProgramListLog Program list log
type ProgramListLog struct {
	Id           int        `gorm:"column:id;type:int;primaryKey;not null" json:"id"`
	AccountId    [21]byte   `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId       [21]byte   `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	HackToolName [1024]byte `gorm:"column:strHackToolName;type:varchar(1024);not null" json:"strHackToolName"`
	WriteTime    int        `gorm:"column:tWriteTime;type:smalldatetime;not null;default:getdate()" json:"tWriteTime"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ProgramListLog) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ProgramListLogDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ProgramListLog) GetTableName() string {
	return _ProgramListLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ProgramListLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [PROGRAMLIST_LOG] (id, strAccountID, strCharID, strHackToolName, tWriteTime) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.CharId),
		GetOptionalBinaryVal(this.HackToolName),
		GetOptionalDecVal(&this.WriteTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ProgramListLog) GetCreateTableString() string {
	query := "CREATE TABLE [PROGRAMLIST_LOG] (\n\t[id] int NOT NULL,\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[strHackToolName] varchar(1024) NOT NULL,\n\t[tWriteTime] smalldatetime NOT NULL\n\tPRIMARY KEY (\"id\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
