package kogen

import (
	"fmt"
)

const (
	_WebpageAddressDatabaseNbr = 0
	_WebpageAddressTableName   = "WEBPAGE_ADDRESS"
)

func init() {
	ModelList = append(ModelList, &WebpageAddress{})
}

// WebpageAddress Webpage URL list
type WebpageAddress struct {
	Index          [10]byte `gorm:"column:nIndex;type:varchar(10)" json:"nIndex,omitempty"`
	WebPageAddress [10]byte `gorm:"column:strWebPageAddress;type:varchar(10)" json:"strWebPageAddress,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *WebpageAddress) GetDatabaseName() string {
	return GetDatabaseName(DbType(_WebpageAddressDatabaseNbr))
}

// GetTableName Returns the table name
func (this *WebpageAddress) GetTableName() string {
	return _WebpageAddressTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *WebpageAddress) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WEBPAGE_ADDRESS] (nIndex, strWebPageAddress) \nVALUES (%s, %s)", GetOptionalBinaryVal(this.Index),
		GetOptionalBinaryVal(this.WebPageAddress))
}

// GetCreateTableString Returns the create table statement for this object
func (this *WebpageAddress) GetCreateTableString() string {
	query := "CREATE TABLE [WEBPAGE_ADDRESS] (\n\t[nIndex] varchar(10),\n\t[strWebPageAddress] varchar(10)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
