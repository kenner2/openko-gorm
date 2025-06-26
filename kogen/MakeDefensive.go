package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MakeDefensiveDatabaseNbr = 1
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

// GetDatabaseName Returns the table's database name
func (this MakeDefensive) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeDefensiveDatabaseNbr))
}

// TableName Returns the table name
func (this MakeDefensive) TableName() string {
	return _MakeDefensiveTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MakeDefensive) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_DEFENSIVE] ([byLevel], [sClass_1], [sClass_2], [sClass_3], [sClass_4], [sClass_5], [sClass_6], [sClass_7]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(this.Class1),
		GetOptionalDecVal(this.Class2),
		GetOptionalDecVal(this.Class3),
		GetOptionalDecVal(this.Class4),
		GetOptionalDecVal(this.Class5),
		GetOptionalDecVal(this.Class6),
		GetOptionalDecVal(this.Class7))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MakeDefensive) GetInsertHeader() string {
	return "INSERT INTO [MAKE_DEFENSIVE] (byLevel, sClass_1, sClass_2, sClass_3, sClass_4, sClass_5, sClass_6, sClass_7) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MakeDefensive) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(this.Class1),
		GetOptionalDecVal(this.Class2),
		GetOptionalDecVal(this.Class3),
		GetOptionalDecVal(this.Class4),
		GetOptionalDecVal(this.Class5),
		GetOptionalDecVal(this.Class6),
		GetOptionalDecVal(this.Class7))
}

// GetCreateTableString Returns the create table statement for this object
func (this MakeDefensive) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_DEFENSIVE] (\n\t[byLevel] tinyint NOT NULL,\n\t[sClass_1] smallint,\n\t[sClass_2] smallint,\n\t[sClass_3] smallint,\n\t[sClass_4] smallint,\n\t[sClass_5] smallint,\n\t[sClass_6] smallint,\n\t[sClass_7] smallint\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MakeDefensive) SelectClause() (selectClause clause.Select) {
	return _MakeDefensiveSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MakeDefensive) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MakeDefensive{}
	rawSql := "SELECT [byLevel], [sClass_1], [sClass_2], [sClass_3], [sClass_4], [sClass_5], [sClass_6], [sClass_7] FROM [MAKE_DEFENSIVE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MakeDefensiveSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[byLevel]",
		},
		clause.Column{
			Name: "[sClass_1]",
		},
		clause.Column{
			Name: "[sClass_2]",
		},
		clause.Column{
			Name: "[sClass_3]",
		},
		clause.Column{
			Name: "[sClass_4]",
		},
		clause.Column{
			Name: "[sClass_5]",
		},
		clause.Column{
			Name: "[sClass_6]",
		},
		clause.Column{
			Name: "[sClass_7]",
		},
	},
}
