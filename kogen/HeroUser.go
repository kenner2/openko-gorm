package kogen

import (
	"fmt"
)

const (
	_HeroUserDatabaseNbr = 0
	_HeroUserTableName   = "HERO_USER"
)

func init() {
	ModelList = append(ModelList, &HeroUser{})
}

// HeroUser TODO Doc
type HeroUser struct {
	Index       int16    `gorm:"column:shIndex;type:smallint;not null" json:"shIndex"`
	UserId      [21]byte `gorm:"column:strUserID;type:varchar(21)" json:"strUserID,omitempty"`
	Nation      [20]byte `gorm:"column:strNation;type:varchar(20)" json:"strNation,omitempty"`
	ClassName   [30]byte `gorm:"column:strClass;type:varchar(30)" json:"strClass,omitempty"`
	Achievement [50]byte `gorm:"column:strAchievement;type:varchar(50)" json:"strAchievement,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *HeroUser) GetDatabaseName() string {
	return GetDatabaseName(DbType(_HeroUserDatabaseNbr))
}

// GetTableName Returns the table name
func (this *HeroUser) GetTableName() string {
	return _HeroUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *HeroUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HERO_USER] (shIndex, strUserID, strNation, strClass, strAchievement) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalBinaryVal(this.UserId),
		GetOptionalBinaryVal(this.Nation),
		GetOptionalBinaryVal(this.ClassName),
		GetOptionalBinaryVal(this.Achievement))
}

// GetCreateTableString Returns the create table statement for this object
func (this *HeroUser) GetCreateTableString() string {
	query := "CREATE TABLE [HERO_USER] (\n\t[shIndex] smallint NOT NULL,\n\t[strUserID] varchar(21),\n\t[strNation] varchar(20),\n\t[strClass] varchar(30),\n\t[strAchievement] varchar(50)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
