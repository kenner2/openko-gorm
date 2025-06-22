package kogen

import (
	"fmt"
)

const (
	_KnightsRatingDatabaseNbr = 0
	_KnightsRatingTableName   = "KNIGHTS_RATING"
)

func init() {
	ModelList = append(ModelList, &KnightsRating{})
}

// KnightsRating Knights Ratings
type KnightsRating struct {
	Rank   int      `gorm:"column:nRank;type:int;not null" json:"nRank"`
	Index  *int16   `gorm:"column:shIndex;type:smallint" json:"shIndex,omitempty"`
	Name   [21]byte `gorm:"column:strName;type:varchar(21)" json:"strName,omitempty"`
	Points *int     `gorm:"column:nPoints;type:int" json:"nPoints,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KnightsRating) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsRatingDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KnightsRating) GetTableName() string {
	return _KnightsRatingTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KnightsRating) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_RATING] (nRank, shIndex, strName, nPoints) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(this.Index),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(this.Points))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KnightsRating) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_RATING] (\n\t[nRank] int NOT NULL,\n\t[shIndex] smallint,\n\t[strName] varchar(21),\n\t[nPoints] int\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
