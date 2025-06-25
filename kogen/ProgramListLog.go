package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
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
	Id           int           `gorm:"column:id;type:int;primaryKey;not null" json:"id"`
	AccountId    mssql.VarChar `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId       mssql.VarChar `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	HackToolName mssql.VarChar `gorm:"column:strHackToolName;type:varchar(1024);not null" json:"strHackToolName"`
	WriteTime    time.Time     `gorm:"column:tWriteTime;type:smalldatetime;not null;default:getdate()" json:"tWriteTime"`
}

// GetDatabaseName Returns the table's database name
func (this ProgramListLog) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ProgramListLogDatabaseNbr))
}

// TableName Returns the table name
func (this ProgramListLog) TableName() string {
	return _ProgramListLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ProgramListLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [PROGRAMLIST_LOG] ([id], [strAccountID], [strCharID], [strHackToolName], [tWriteTime]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.HackToolName, false),
		GetDateTimeExportFmt(&this.WriteTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ProgramListLog) GetInsertHeader() string {
	return "INSERT INTO [PROGRAMLIST_LOG] (id, strAccountID, strCharID, strHackToolName, tWriteTime) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ProgramListLog) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.HackToolName, false),
		GetDateTimeExportFmt(&this.WriteTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this ProgramListLog) GetCreateTableString() string {
	query := "CREATE TABLE [PROGRAMLIST_LOG] (\n\t[id] int NOT NULL,\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[strHackToolName] varchar(1024) NOT NULL,\n\t[tWriteTime] smalldatetime NOT NULL\n\tCONSTRAINT [PK_PROGRAMLIST_LOG] PRIMARY KEY ([id])\n)\nGO\nALTER TABLE [PROGRAMLIST_LOG] ADD CONSTRAINT [DF_PROGRAMLIST_LOG_tWriteTime] DEFAULT getdate() FOR [tWriteTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ProgramListLog) SelectClause() (selectClause clause.Select) {
	return _ProgramListLogSelectClause
}

// GetAllTableData Returns a list of all table data
func (this ProgramListLog) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ProgramListLog{}
	rawSql := "SELECT [id], [strAccountID], [strCharID], [strHackToolName], [tWriteTime] FROM [PROGRAMLIST_LOG]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ProgramListLogSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[id]",
		},
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
