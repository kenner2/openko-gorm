package kogen

import (
	"fmt"
)

const (
	_AccountCharDatabaseNbr = 0
	_AccountCharTableName   = "ACCOUNT_CHAR"
)

// AccountChar: Represents the relationship between accounts and characters
type AccountChar struct {
	AccountId string  `gorm:"column:strAccountID;primaryKey;not null" json:"strAccountID"`
	Nation    uint8   `gorm:"column:bNation;not null" json:"bNation"`
	CharNum   uint8   `gorm:"column:bCharNum;not null" json:"bCharNum"`
	CharId1   *string `gorm:"column:strCharID1" json:"strCharID1,omitempty"`
	CharId2   *string `gorm:"column:strCharID2" json:"strCharID2,omitempty"`
	CharId3   *string `gorm:"column:strCharID3" json:"strCharID3,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *AccountChar) GetDatabaseName() string {
	return GetDatabaseName(DbType(_AccountCharDatabaseNbr))
}

// GetTableName Returns the table name
func (this *AccountChar) GetTableName() string {
	return _AccountCharTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *AccountChar) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ACCOUNT_CHAR] (strAccountID, bNation, bCharNum, strCharID1, strCharID2, strCharID3) \nVALUES (%s, %s, %s, %s, %s, %s)", GetOptionalStringVal(&this.AccountId),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.CharNum),
		GetOptionalStringVal(this.CharId1),
		GetOptionalStringVal(this.CharId2),
		GetOptionalStringVal(this.CharId3))
}
