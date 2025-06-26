package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_LevelUpDatabaseNbr = 1
	_LevelUpTableName   = "LEVEL_UP"
)

func init() {
	ModelList = append(ModelList, &LevelUp{})
}

// LevelUp Level experience requirements
type LevelUp struct {
	Level       uint8 `gorm:"column:level;type:tinyint;primaryKey;not null" json:"level"`
	RequiredExp int   `gorm:"column:Exp;type:int;not null" json:"Exp"`
}

// GetDatabaseName Returns the table's database name
func (this LevelUp) GetDatabaseName() string {
	return GetDatabaseName(DbType(_LevelUpDatabaseNbr))
}

// TableName Returns the table name
func (this LevelUp) TableName() string {
	return _LevelUpTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this LevelUp) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [LEVEL_UP] ([level], [Exp]) VALUES\n(%s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.RequiredExp))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this LevelUp) GetInsertHeader() string {
	return "INSERT INTO [LEVEL_UP] (level, Exp) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this LevelUp) GetInsertData() string {
	return fmt.Sprintf("(%s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.RequiredExp))
}

// GetCreateTableString Returns the create table statement for this object
func (this LevelUp) GetCreateTableString() string {
	query := "CREATE TABLE [LEVEL_UP] (\n\t[level] tinyint NOT NULL,\n\t[Exp] int NOT NULL\n\tCONSTRAINT [PK_LEVEL_UP] PRIMARY KEY ([level])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this LevelUp) SelectClause() (selectClause clause.Select) {
	return _LevelUpSelectClause
}

// GetAllTableData Returns a list of all table data
func (this LevelUp) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []LevelUp{}
	rawSql := "SELECT [level], [Exp] FROM [LEVEL_UP]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _LevelUpSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[level]",
		},
		clause.Column{
			Name: "[Exp]",
		},
	},
}
