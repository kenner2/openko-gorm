package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_CopyTestDatabaseNbr = 1
	_CopyTestTableName   = "COPY_TEST"
)

func init() {
	ModelList = append(ModelList, &CopyTest{})
}

// CopyTest TODO: Doc
type CopyTest struct {
	ItemSerial int64 `gorm:"column:ITEMSERIAL;type:bigint;not null" json:"ITEMSERIAL"`
}

// GetDatabaseName Returns the table's database name
func (this CopyTest) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CopyTestDatabaseNbr))
}

// TableName Returns the table name
func (this CopyTest) TableName() string {
	return _CopyTestTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this CopyTest) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COPY_TEST] ([ITEMSERIAL]) VALUES\n(%s)", GetOptionalDecVal(&this.ItemSerial))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this CopyTest) GetInsertHeader() string {
	return "INSERT INTO [COPY_TEST] (ITEMSERIAL) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this CopyTest) GetInsertData() string {
	return fmt.Sprintf("(%s)", GetOptionalDecVal(&this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this CopyTest) GetCreateTableString() string {
	query := "CREATE TABLE [COPY_TEST] (\n\t[ITEMSERIAL] bigint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this CopyTest) SelectClause() (selectClause clause.Select) {
	return _CopyTestSelectClause
}

// GetAllTableData Returns a list of all table data
func (this CopyTest) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []CopyTest{}
	rawSql := "SELECT [ITEMSERIAL] FROM [COPY_TEST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _CopyTestSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[ITEMSERIAL]",
		},
	},
}
