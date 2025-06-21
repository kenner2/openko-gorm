package kogen

import (
	"fmt"
)

const (
	_ItemUpProbabilityDatabaseNbr = 0
	_ItemUpProbabilityTableName   = "ITEMUP_PROBABILITY"
)

func init() {
	ModelList = append(ModelList, &ItemUpProbability{})
}

// ItemUpProbability: TODO: Doc usage
type ItemUpProbability struct {
	Type           uint8 `gorm:"column:bType;type:tinyint;primaryKey;not null" json:"bType"`
	MaxSuccess     int16 `gorm:"column:nMaxSuccess;type:smallint;not null" json:"nMaxSuccess"`
	MaxFail        int16 `gorm:"column:nMaxFail;type:smallint;not null" json:"nMaxFail"`
	CurrentSuccess int16 `gorm:"column:nCurSuccess;type:smallint;not null" json:"nCurSuccess"`
	CurrentFailure int16 `gorm:"column:nCurFail;type:smallint;not null" json:"nCurFail"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ItemUpProbability) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemUpProbabilityDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ItemUpProbability) GetTableName() string {
	return _ItemUpProbabilityTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ItemUpProbability) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEMUP_PROBABILITY] (bType, nMaxSuccess, nMaxFail, nCurSuccess, nCurFail) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.MaxSuccess),
		GetOptionalDecVal(&this.MaxFail),
		GetOptionalDecVal(&this.CurrentSuccess),
		GetOptionalDecVal(&this.CurrentFailure))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ItemUpProbability) GetCreateTableString() string {
	query := "CREATE TABLE \"ITEMUP_PROBABILITY\" (\n\t\"bType\" tinyint NOT NULL,\n\t\"nMaxSuccess\" smallint NOT NULL,\n\t\"nMaxFail\" smallint NOT NULL,\n\t\"nCurSuccess\" smallint NOT NULL,\n\t\"nCurFail\" smallint NOT NULL\n\tPRIMARY KEY (\"bType\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
