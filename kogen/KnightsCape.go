package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KnightsCapeDatabaseNbr = 0
	_KnightsCapeTableName   = "KNIGHTS_CAPE"
)

func init() {
	ModelList = append(ModelList, &KnightsCape{})
}

// KnightsCape Knights cape information
type KnightsCape struct {
	CapeIndex int16         `gorm:"column:sCapeIndex;type:smallint;primaryKey;not null" json:"sCapeIndex"`
	Name      mssql.VarChar `gorm:"column:strName;type:varchar(30);not null" json:"strName"`
	BuyPrice  int           `gorm:"column:nBuyPrice;type:int;not null" json:"nBuyPrice"`
	Duration  int           `gorm:"column:nDuration;type:int;not null" json:"nDuration"`
	Grade     uint8         `gorm:"column:byGrade;type:tinyint;not null" json:"byGrade"`
}

// GetDatabaseName Returns the table's database name
func (this KnightsCape) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsCapeDatabaseNbr))
}

// TableName Returns the table name
func (this KnightsCape) TableName() string {
	return _KnightsCapeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KnightsCape) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_CAPE] ([sCapeIndex], [strName], [nBuyPrice], [nDuration], [byGrade]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CapeIndex),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.BuyPrice),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Grade))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KnightsCape) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS_CAPE] (sCapeIndex, strName, nBuyPrice, nDuration, byGrade) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KnightsCape) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CapeIndex),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.BuyPrice),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Grade))
}

// GetCreateTableString Returns the create table statement for this object
func (this KnightsCape) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_CAPE] (\n\t[sCapeIndex] smallint NOT NULL,\n\t[strName] varchar(30) NOT NULL,\n\t[nBuyPrice] int NOT NULL,\n\t[nDuration] int NOT NULL,\n\t[byGrade] tinyint NOT NULL\n\tCONSTRAINT [PK_KNIGHTS_CAPE] PRIMARY KEY ([sCapeIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KnightsCape) SelectClause() (selectClause clause.Select) {
	return _KnightsCapeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KnightsCape) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KnightsCape{}
	rawSql := "SELECT [sCapeIndex], [strName], [nBuyPrice], [nDuration], [byGrade] FROM [KNIGHTS_CAPE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsCapeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sCapeIndex]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[nBuyPrice]",
		},
		clause.Column{
			Name: "[nDuration]",
		},
		clause.Column{
			Name: "[byGrade]",
		},
	},
}
