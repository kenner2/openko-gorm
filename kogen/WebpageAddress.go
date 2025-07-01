package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_WebpageAddressDatabaseNbr = "GAME"
	_WebpageAddressTableName   = "WEBPAGE_ADDRESS"
)

func init() {
	ModelList = append(ModelList, &WebpageAddress{})
}

// WebpageAddress Webpage URL list
type WebpageAddress struct {
	Index          int            `gorm:"column:nIndex;type:int;primaryKey;not null" json:"nIndex"`
	WebPageAddress *mssql.VarChar `gorm:"column:strWebPageAddress;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strWebPageAddress,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this WebpageAddress) GetDatabaseName() string {
	return GetDatabaseName(_WebpageAddressDatabaseNbr)
}

// TableName Returns the table name
func (this WebpageAddress) TableName() string {
	return _WebpageAddressTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this WebpageAddress) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WEBPAGE_ADDRESS] ([nIndex], [strWebPageAddress]) VALUES\n(%s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.WebPageAddress, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this WebpageAddress) GetInsertHeader() string {
	return "INSERT INTO [WEBPAGE_ADDRESS] ([nIndex], [strWebPageAddress]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this WebpageAddress) GetInsertData() string {
	return fmt.Sprintf("(%s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.WebPageAddress, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this WebpageAddress) GetCreateTableString() string {
	query := "CREATE TABLE [WEBPAGE_ADDRESS] (\n\t[nIndex] int NOT NULL,\n\t[strWebPageAddress] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS\n\tCONSTRAINT [PK_WEBPAGE_ADDRESS] PRIMARY KEY CLUSTERED ([nIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this WebpageAddress) SelectClause() (selectClause clause.Select) {
	return _WebpageAddressSelectClause
}

// GetAllTableData Returns a list of all table data
func (this WebpageAddress) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []WebpageAddress{}
	rawSql := "SELECT [nIndex], [strWebPageAddress] FROM [WEBPAGE_ADDRESS]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _WebpageAddressSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nIndex]",
		},
		clause.Column{
			Name: "[strWebPageAddress]",
		},
	},
}
