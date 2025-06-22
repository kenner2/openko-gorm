package kogen

import (
	"fmt"
)

const (
	_CopyTestDatabaseNbr = 0
	_CopyTestTableName   = "COPY_TEST"
)

func init() {
	ModelList = append(ModelList, &CopyTest{})
}

// CopyTest TODO: Doc
type CopyTest struct {
	ItemSerial int64 `gorm:"column:ITEMSERIAL;type:bigint;not null" json:"ITEMSERIAL"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *CopyTest) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CopyTestDatabaseNbr))
}

// GetTableName Returns the table name
func (this *CopyTest) GetTableName() string {
	return _CopyTestTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *CopyTest) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COPY_TEST] (ITEMSERIAL) \nVALUES (%s)", GetOptionalDecVal(&this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this *CopyTest) GetCreateTableString() string {
	query := "CREATE TABLE [COPY_TEST] (\n\t\"ITEMSERIAL\" bigint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
