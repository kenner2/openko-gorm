package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KnightsUserDatabaseNbr = 1
	_KnightsUserTableName   = "KNIGHTS_USER"
)

func init() {
	ModelList = append(ModelList, &KnightsUser{})
}

// KnightsUser Knights to character relationships
type KnightsUser struct {
	KnightsId int16         `gorm:"column:sIDNum;type:smallint;primaryKey;not null" json:"sIDNum"`
	UserId    mssql.VarChar `gorm:"column:strUserID;type:varchar(21);primaryKey;not null" json:"strUserID"`
}

// GetDatabaseName Returns the table's database name
func (this KnightsUser) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsUserDatabaseNbr))
}

// TableName Returns the table name
func (this KnightsUser) TableName() string {
	return _KnightsUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KnightsUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_USER] ([sIDNum], [strUserID]) VALUES\n(%s, %s)", GetOptionalDecVal(&this.KnightsId),
		GetOptionalVarCharVal(&this.UserId, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KnightsUser) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS_USER] (sIDNum, strUserID) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KnightsUser) GetInsertData() string {
	return fmt.Sprintf("(%s, %s)", GetOptionalDecVal(&this.KnightsId),
		GetOptionalVarCharVal(&this.UserId, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this KnightsUser) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_USER] (\n\t[sIDNum] smallint NOT NULL,\n\t[strUserID] varchar(21) NOT NULL\n\tCONSTRAINT [PK_KNIGHTS_USER] PRIMARY KEY ([sIDNum], [strUserID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KnightsUser) SelectClause() (selectClause clause.Select) {
	return _KnightsUserSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KnightsUser) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KnightsUser{}
	rawSql := "SELECT [sIDNum], [strUserID] FROM [KNIGHTS_USER]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsUserSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIDNum]",
		},
		clause.Column{
			Name: "[strUserID]",
		},
	},
}
