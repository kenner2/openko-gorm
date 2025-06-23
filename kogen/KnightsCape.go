package kogen

import (
	"fmt"
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
	CapeIndex int16    `gorm:"column:sCapeIndex;type:smallint;not null" json:"sCapeIndex"`
	Name      [30]byte `gorm:"column:strName;type:varchar(30);not null" json:"strName"`
	BuyPrice  int      `gorm:"column:nBuyPrice;type:int;not null" json:"nBuyPrice"`
	Duration  int      `gorm:"column:nDuration;type:int;not null" json:"nDuration"`
	Grade     uint8    `gorm:"column:byGrade;type:tinyint;not null" json:"byGrade"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KnightsCape) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsCapeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KnightsCape) GetTableName() string {
	return _KnightsCapeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KnightsCape) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_CAPE] (sCapeIndex, strName, nBuyPrice, nDuration, byGrade) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CapeIndex),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(&this.BuyPrice),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Grade))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KnightsCape) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_CAPE] (\n\t[sCapeIndex] smallint NOT NULL,\n\t[strName] varchar(30) NOT NULL,\n\t[nBuyPrice] int NOT NULL,\n\t[nDuration] int NOT NULL,\n\t[byGrade] tinyint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
