package kogen

import (
	"fmt"
)

const (
	_AccountCharDatabaseNbr = 0
	_AccountCharTableName   = "ACCOUNT_CHAR"
)

func init() {
	ModelList = append(ModelList, &AccountChar{})
}

// AccountChar Represents the relationship between accounts and characters
type AccountChar struct {
	AccountId [21]byte `gorm:"column:strAccountID;type:varchar(21);primaryKey;not null" json:"strAccountID"`
	Nation    uint8    `gorm:"column:bNation;type:tinyint;not null" json:"bNation"`
	CharNum   uint8    `gorm:"column:bCharNum;type:tinyint;not null;default:0" json:"bCharNum"`
	CharId1   [21]byte `gorm:"column:strCharID1;type:varchar(21)" json:"strCharID1,omitempty"`
	CharId2   [21]byte `gorm:"column:strCharID2;type:varchar(21)" json:"strCharID2,omitempty"`
	CharId3   [21]byte `gorm:"column:strCharID3;type:varchar(21)" json:"strCharID3,omitempty"`
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
	return fmt.Sprintf("INSERT INTO [ACCOUNT_CHAR] (strAccountID, bNation, bCharNum, strCharID1, strCharID2, strCharID3) \nVALUES (%s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.CharNum),
		GetOptionalBinaryVal(this.CharId1),
		GetOptionalBinaryVal(this.CharId2),
		GetOptionalBinaryVal(this.CharId3))
}

// GetCreateTableString Returns the create table statement for this object
func (this *AccountChar) GetCreateTableString() string {
	query := "CREATE TABLE [ACCOUNT_CHAR] (\n\t\"strAccountID\" varchar(21) NOT NULL,\n\t\"bNation\" tinyint NOT NULL,\n\t\"bCharNum\" tinyint NOT NULL,\n\t\"strCharID1\" varchar(21),\n\t\"strCharID2\" varchar(21),\n\t\"strCharID3\" varchar(21)\n\tPRIMARY KEY (\"strAccountID\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
