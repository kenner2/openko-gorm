package kogen

import (
	"fmt"
)

const (
	_MakeDefensiveDatabaseNbr = 0
	_MakeDefensiveTableName   = "MAKE_DEFENSIVE"
)

func init() {
	ModelList = append(ModelList, &MakeDefensive{})
}

// MakeDefensive Make defensive
type MakeDefensive struct {
	Level  uint8  `gorm:"column:byLevel;type:tinyint;not null" json:"byLevel"`
	Class1 *int16 `gorm:"column:sClass_1;type:smallint" json:"sClass_1,omitempty"`
	Class2 *int16 `gorm:"column:sClass_2;type:smallint" json:"sClass_2,omitempty"`
	Class3 *int16 `gorm:"column:sClass_3;type:smallint" json:"sClass_3,omitempty"`
	Class4 *int16 `gorm:"column:sClass_4;type:smallint" json:"sClass_4,omitempty"`
	Class5 *int16 `gorm:"column:sClass_5;type:smallint" json:"sClass_5,omitempty"`
	Class6 *int16 `gorm:"column:sClass_6;type:smallint" json:"sClass_6,omitempty"`
	Class7 *int16 `gorm:"column:sClass_7;type:smallint" json:"sClass_7,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MakeDefensive) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeDefensiveDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MakeDefensive) GetTableName() string {
	return _MakeDefensiveTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MakeDefensive) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_DEFENSIVE] (byLevel, sClass_1, sClass_2, sClass_3, sClass_4, sClass_5, sClass_6, sClass_7) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(this.Class1),
		GetOptionalDecVal(this.Class2),
		GetOptionalDecVal(this.Class3),
		GetOptionalDecVal(this.Class4),
		GetOptionalDecVal(this.Class5),
		GetOptionalDecVal(this.Class6),
		GetOptionalDecVal(this.Class7))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MakeDefensive) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_DEFENSIVE] (\n\t[byLevel] tinyint NOT NULL,\n\t[sClass_1] smallint,\n\t[sClass_2] smallint,\n\t[sClass_3] smallint,\n\t[sClass_4] smallint,\n\t[sClass_5] smallint,\n\t[sClass_6] smallint,\n\t[sClass_7] smallint\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
