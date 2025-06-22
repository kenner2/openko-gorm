package kogen

import (
	"fmt"
)

const (
	_LevelUpDatabaseNbr = 0
	_LevelUpTableName   = "LEVEL_UP"
)

func init() {
	ModelList = append(ModelList, &LevelUp{})
}

// LevelUp Level experience requirements
type LevelUp struct {
	Level       uint8 `gorm:"column:level;type:tinyint;not null" json:"level"`
	RequiredExp int   `gorm:"column:Exp;type:int;not null" json:"Exp"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *LevelUp) GetDatabaseName() string {
	return GetDatabaseName(DbType(_LevelUpDatabaseNbr))
}

// GetTableName Returns the table name
func (this *LevelUp) GetTableName() string {
	return _LevelUpTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *LevelUp) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [LEVEL_UP] (level, Exp) \nVALUES (%s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.RequiredExp))
}

// GetCreateTableString Returns the create table statement for this object
func (this *LevelUp) GetCreateTableString() string {
	query := "CREATE TABLE [LEVEL_UP] (\n\t[level] tinyint NOT NULL,\n\t[Exp] int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
