package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_HacktoolUserLogDatabaseNbr = "GAME"
	_HacktoolUserLogTableName   = "HACKTOOL_USERLOG"
)

func init() {
	ModelList = append(ModelList, &HacktoolUserLog{})
}

// HacktoolUserLog Tracks possible detections of a hack tool by a user
type HacktoolUserLog struct {
	AccountId    mssql.VarChar `gorm:"column:strAccountID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strAccountID"`
	CharId       mssql.VarChar `gorm:"column:strCharID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strCharID"`
	HackToolName mssql.VarChar `gorm:"column:strHackToolName;type:varchar(512) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strHackToolName"`
	WriteTime    time.Time     `gorm:"column:tWriteTime;type:smalldatetime;not null;default:getdate()" json:"tWriteTime"`
}

// GetDatabaseName Returns the table's database name
func (this HacktoolUserLog) GetDatabaseName() string {
	return GetDatabaseName(_HacktoolUserLogDatabaseNbr)
}

// TableName Returns the table name
func (this HacktoolUserLog) TableName() string {
	return _HacktoolUserLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this HacktoolUserLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HACKTOOL_USERLOG] ([strAccountID], [strCharID], [strHackToolName], [tWriteTime]) VALUES\n(%s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.HackToolName, false),
		GetDateTimeExportFmt(&this.WriteTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this HacktoolUserLog) GetInsertHeader() string {
	return "INSERT INTO [HACKTOOL_USERLOG] ([strAccountID], [strCharID], [strHackToolName], [tWriteTime]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this HacktoolUserLog) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.HackToolName, false),
		GetDateTimeExportFmt(&this.WriteTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this HacktoolUserLog) GetCreateTableString() string {
	query := "CREATE TABLE [HACKTOOL_USERLOG] (\n\t[strAccountID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strCharID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strHackToolName] varchar(512) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[tWriteTime] smalldatetime NOT NULL\n)\nGO\nALTER TABLE [HACKTOOL_USERLOG] ADD CONSTRAINT [DF_HACKTOOL_USERLOG_tWriteTime] DEFAULT getdate() FOR [tWriteTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this HacktoolUserLog) SelectClause() (selectClause clause.Select) {
	return _HacktoolUserLogSelectClause
}

// GetAllTableData Returns a list of all table data
func (this HacktoolUserLog) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []HacktoolUserLog{}
	rawSql := "SELECT [strAccountID], [strCharID], [strHackToolName], [tWriteTime] FROM [HACKTOOL_USERLOG]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _HacktoolUserLogSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[strHackToolName]",
		},
		clause.Column{
			Name: "[tWriteTime]",
		},
	},
}
