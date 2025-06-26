package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_CurrentUserDatabaseNbr = "GAME"
	_CurrentUserTableName   = "CURRENTUSER"
)

func init() {
	ModelList = append(ModelList, &CurrentUser{})
}

// CurrentUser Keeps track of users currently connected to the server
type CurrentUser struct {
	ServerNumber int           `gorm:"column:nServerNo;type:int;not null" json:"nServerNo"`
	ServerIP     mssql.VarChar `gorm:"column:strServerIP;type:varchar(20);not null" json:"strServerIP"`
	AccountId    mssql.VarChar `gorm:"column:strAccountID;type:varchar(20);not null" json:"strAccountID"`
	CharId       mssql.VarChar `gorm:"column:strCharID;type:varchar(20);not null" json:"strCharID"`
	ClientIP     mssql.VarChar `gorm:"column:strClientIP;type:varchar(20);not null" json:"strClientIP"`
}

// GetDatabaseName Returns the table's database name
func (this CurrentUser) GetDatabaseName() string {
	return GetDatabaseName(_CurrentUserDatabaseNbr)
}

// TableName Returns the table name
func (this CurrentUser) TableName() string {
	return _CurrentUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this CurrentUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [CURRENTUSER] ([nServerNo], [strServerIP], [strAccountID], [strCharID], [strClientIP]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerNumber),
		GetOptionalVarCharVal(&this.ServerIP, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.ClientIP, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this CurrentUser) GetInsertHeader() string {
	return "INSERT INTO [CURRENTUSER] ([nServerNo], [strServerIP], [strAccountID], [strCharID], [strClientIP]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this CurrentUser) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerNumber),
		GetOptionalVarCharVal(&this.ServerIP, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.ClientIP, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this CurrentUser) GetCreateTableString() string {
	query := "CREATE TABLE [CURRENTUSER] (\n\t[nServerNo] int NOT NULL,\n\t[strServerIP] varchar(20) NOT NULL,\n\t[strAccountID] varchar(20) NOT NULL,\n\t[strCharID] varchar(20) NOT NULL,\n\t[strClientIP] varchar(20) NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this CurrentUser) SelectClause() (selectClause clause.Select) {
	return _CurrentUserSelectClause
}

// GetAllTableData Returns a list of all table data
func (this CurrentUser) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []CurrentUser{}
	rawSql := "SELECT [nServerNo], [strServerIP], [strAccountID], [strCharID], [strClientIP] FROM [CURRENTUSER]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _CurrentUserSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nServerNo]",
		},
		clause.Column{
			Name: "[strServerIP]",
		},
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[strClientIP]",
		},
	},
}
