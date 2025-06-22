package kogen

import (
	"fmt"
)

const (
	_MakeItemGradeCodeDatabaseNbr = 0
	_MakeItemGradeCodeTableName   = "MAKE_ITEM_GRADECODE"
)

func init() {
	ModelList = append(ModelList, &MakeItemGradeCode{})
}

// MakeItemGradeCode Make item grade code
type MakeItemGradeCode struct {
	ItemIndex uint8  `gorm:"column:byItemIndex;type:tinyint;not null" json:"byItemIndex"`
	Grade1    int16  `gorm:"column:byGrade_1;type:smallint;not null" json:"byGrade_1"`
	Grade2    *int16 `gorm:"column:byGrade_2;type:smallint" json:"byGrade_2,omitempty"`
	Grade3    *int16 `gorm:"column:byGrade_3;type:smallint" json:"byGrade_3,omitempty"`
	Grade4    *int16 `gorm:"column:byGrade_4;type:smallint" json:"byGrade_4,omitempty"`
	Grade5    *int16 `gorm:"column:byGrade_5;type:smallint" json:"byGrade_5,omitempty"`
	Grade6    *int16 `gorm:"column:byGrade_6;type:smallint" json:"byGrade_6,omitempty"`
	Grade7    *int16 `gorm:"column:byGrade_7;type:smallint" json:"byGrade_7,omitempty"`
	Grade8    *int16 `gorm:"column:byGrade_8;type:smallint" json:"byGrade_8,omitempty"`
	Grade9    *int16 `gorm:"column:byGrade_9;type:smallint" json:"byGrade_9,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MakeItemGradeCode) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeItemGradeCodeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MakeItemGradeCode) GetTableName() string {
	return _MakeItemGradeCodeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MakeItemGradeCode) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM_GRADECODE] (byItemIndex, byGrade_1, byGrade_2, byGrade_3, byGrade_4, byGrade_5, byGrade_6, byGrade_7, byGrade_8, byGrade_9) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Grade1),
		GetOptionalDecVal(this.Grade2),
		GetOptionalDecVal(this.Grade3),
		GetOptionalDecVal(this.Grade4),
		GetOptionalDecVal(this.Grade5),
		GetOptionalDecVal(this.Grade6),
		GetOptionalDecVal(this.Grade7),
		GetOptionalDecVal(this.Grade8),
		GetOptionalDecVal(this.Grade9))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MakeItemGradeCode) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM_GRADECODE] (\n\t\"byItemIndex\" tinyint NOT NULL,\n\t\"byGrade_1\" smallint NOT NULL,\n\t\"byGrade_2\" smallint,\n\t\"byGrade_3\" smallint,\n\t\"byGrade_4\" smallint,\n\t\"byGrade_5\" smallint,\n\t\"byGrade_6\" smallint,\n\t\"byGrade_7\" smallint,\n\t\"byGrade_8\" smallint,\n\t\"byGrade_9\" smallint\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
