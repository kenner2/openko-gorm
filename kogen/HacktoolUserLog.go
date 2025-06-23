package kogen

import (
	"fmt"
)

const (
	_HacktoolUserLogDatabaseNbr = 0
	_HacktoolUserLogTableName   = "HACKTOOL_USERLOG"
)

func init() {
	ModelList = append(ModelList, &HacktoolUserLog{})
}

// HacktoolUserLog Tracks possible detections of a hack tool by a user
type HacktoolUserLog struct {
	AccountId    [21]byte  `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId       [21]byte  `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	HackToolName [512]byte `gorm:"column:strHackToolName;type:varchar(512)" json:"strHackToolName,omitempty"`
	WriteTime    int       `gorm:"column:tWriteTime;type:smalldatetime;not null;default:getdate()" json:"tWriteTime"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *HacktoolUserLog) GetDatabaseName() string {
	return GetDatabaseName(DbType(_HacktoolUserLogDatabaseNbr))
}

// GetTableName Returns the table name
func (this *HacktoolUserLog) GetTableName() string {
	return _HacktoolUserLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *HacktoolUserLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HACKTOOL_USERLOG] (strAccountID, strCharID, strHackToolName, tWriteTime) \nVALUES (%s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.CharId),
		GetOptionalBinaryVal(this.HackToolName),
		GetOptionalDecVal(&this.WriteTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *HacktoolUserLog) GetCreateTableString() string {
	query := "CREATE TABLE [HACKTOOL_USERLOG] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[strHackToolName] varchar(512),\n\t[tWriteTime] smalldatetime NOT NULL\n\n)\nGO\nALTER TABLE [HACKTOOL_USERLOG] ADD CONSTRAINT [DF_HACKTOOL_USERLOG_tWriteTime] DEFAULT getdate() FOR [tWriteTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
