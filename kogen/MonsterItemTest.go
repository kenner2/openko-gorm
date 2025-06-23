package kogen

import (
	"fmt"
)

const (
	_MonsterItemTestDatabaseNbr = 0
	_MonsterItemTestTableName   = "MONSTER_ITEM_TEST"
)

func init() {
	ModelList = append(ModelList, &MonsterItemTest{})
}

// MonsterItemTest Monster item test
type MonsterItemTest struct {
	Index    int16  `gorm:"column:sIndex;type:smallint;not null" json:"sIndex"`
	Item1    *int   `gorm:"column:iItem01;type:int" json:"iItem01,omitempty"`
	Percent1 *int16 `gorm:"column:sPersent01;type:smallint" json:"sPersent01,omitempty"`
	Item2    *int   `gorm:"column:iItem02;type:int" json:"iItem02,omitempty"`
	Percent2 *int16 `gorm:"column:sPersent02;type:smallint" json:"sPersent02,omitempty"`
	Item3    *int   `gorm:"column:iItem03;type:int" json:"iItem03,omitempty"`
	Percent3 *int16 `gorm:"column:sPersent03;type:smallint" json:"sPersent03,omitempty"`
	Item4    *int   `gorm:"column:iItem04;type:int" json:"iItem04,omitempty"`
	Percent4 *int16 `gorm:"column:sPersent04;type:smallint" json:"sPersent04,omitempty"`
	Item5    *int   `gorm:"column:iItem05;type:int" json:"iItem05,omitempty"`
	Percent5 *int16 `gorm:"column:sPersent05;type:smallint" json:"sPersent05,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MonsterItemTest) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MonsterItemTestDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MonsterItemTest) GetTableName() string {
	return _MonsterItemTestTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MonsterItemTest) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_ITEM_TEST] (sIndex, iItem01, sPersent01, iItem02, sPersent02, iItem03, sPersent03, iItem04, sPersent04, iItem05, sPersent05) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(this.Item1),
		GetOptionalDecVal(this.Percent1),
		GetOptionalDecVal(this.Item2),
		GetOptionalDecVal(this.Percent2),
		GetOptionalDecVal(this.Item3),
		GetOptionalDecVal(this.Percent3),
		GetOptionalDecVal(this.Item4),
		GetOptionalDecVal(this.Percent4),
		GetOptionalDecVal(this.Item5),
		GetOptionalDecVal(this.Percent5))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MonsterItemTest) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_ITEM_TEST] (\n\t[sIndex] smallint NOT NULL,\n\t[iItem01] int,\n\t[sPersent01] smallint,\n\t[iItem02] int,\n\t[sPersent02] smallint,\n\t[iItem03] int,\n\t[sPersent03] smallint,\n\t[iItem04] int,\n\t[sPersent04] smallint,\n\t[iItem05] int,\n\t[sPersent05] smallint\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
