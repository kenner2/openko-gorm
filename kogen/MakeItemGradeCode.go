package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MakeItemGradeCodeDatabaseNbr = "GAME"
	_MakeItemGradeCodeTableName   = "MAKE_ITEM_GRADECODE"
)

func init() {
	ModelList = append(ModelList, &MakeItemGradeCode{})
}

// MakeItemGradeCode Make item grade code
type MakeItemGradeCode struct {
	ItemIndex uint8 `gorm:"column:byItemIndex;type:tinyint;primaryKey;not null" json:"byItemIndex"`
	Grade1    int16 `gorm:"column:byGrade_1;type:smallint;not null" json:"byGrade_1"`
	Grade2    int16 `gorm:"column:byGrade_2;type:smallint;not null" json:"byGrade_2"`
	Grade3    int16 `gorm:"column:byGrade_3;type:smallint;not null" json:"byGrade_3"`
	Grade4    int16 `gorm:"column:byGrade_4;type:smallint;not null" json:"byGrade_4"`
	Grade5    int16 `gorm:"column:byGrade_5;type:smallint;not null" json:"byGrade_5"`
	Grade6    int16 `gorm:"column:byGrade_6;type:smallint;not null" json:"byGrade_6"`
	Grade7    int16 `gorm:"column:byGrade_7;type:smallint;not null" json:"byGrade_7"`
	Grade8    int16 `gorm:"column:byGrade_8;type:smallint;not null" json:"byGrade_8"`
	Grade9    int16 `gorm:"column:byGrade_9;type:smallint;not null" json:"byGrade_9"`
}

// GetDatabaseName Returns the table's database name
func (this MakeItemGradeCode) GetDatabaseName() string {
	return GetDatabaseName(_MakeItemGradeCodeDatabaseNbr)
}

// TableName Returns the table name
func (this MakeItemGradeCode) TableName() string {
	return _MakeItemGradeCodeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MakeItemGradeCode) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM_GRADECODE] ([byItemIndex], [byGrade_1], [byGrade_2], [byGrade_3], [byGrade_4], [byGrade_5], [byGrade_6], [byGrade_7], [byGrade_8], [byGrade_9]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Grade1),
		GetOptionalDecVal(&this.Grade2),
		GetOptionalDecVal(&this.Grade3),
		GetOptionalDecVal(&this.Grade4),
		GetOptionalDecVal(&this.Grade5),
		GetOptionalDecVal(&this.Grade6),
		GetOptionalDecVal(&this.Grade7),
		GetOptionalDecVal(&this.Grade8),
		GetOptionalDecVal(&this.Grade9))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MakeItemGradeCode) GetInsertHeader() string {
	return "INSERT INTO [MAKE_ITEM_GRADECODE] ([byItemIndex], [byGrade_1], [byGrade_2], [byGrade_3], [byGrade_4], [byGrade_5], [byGrade_6], [byGrade_7], [byGrade_8], [byGrade_9]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MakeItemGradeCode) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Grade1),
		GetOptionalDecVal(&this.Grade2),
		GetOptionalDecVal(&this.Grade3),
		GetOptionalDecVal(&this.Grade4),
		GetOptionalDecVal(&this.Grade5),
		GetOptionalDecVal(&this.Grade6),
		GetOptionalDecVal(&this.Grade7),
		GetOptionalDecVal(&this.Grade8),
		GetOptionalDecVal(&this.Grade9))
}

// GetCreateTableString Returns the create table statement for this object
func (this MakeItemGradeCode) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM_GRADECODE] (\n\t[byItemIndex] tinyint NOT NULL,\n\t[byGrade_1] smallint NOT NULL,\n\t[byGrade_2] smallint NOT NULL,\n\t[byGrade_3] smallint NOT NULL,\n\t[byGrade_4] smallint NOT NULL,\n\t[byGrade_5] smallint NOT NULL,\n\t[byGrade_6] smallint NOT NULL,\n\t[byGrade_7] smallint NOT NULL,\n\t[byGrade_8] smallint NOT NULL,\n\t[byGrade_9] smallint NOT NULL\n\tCONSTRAINT [PK_MAKE_ITEM_GRADECODE] PRIMARY KEY CLUSTERED ([byItemIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MakeItemGradeCode) SelectClause() (selectClause clause.Select) {
	return _MakeItemGradeCodeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MakeItemGradeCode) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MakeItemGradeCode{}
	rawSql := "SELECT [byItemIndex], [byGrade_1], [byGrade_2], [byGrade_3], [byGrade_4], [byGrade_5], [byGrade_6], [byGrade_7], [byGrade_8], [byGrade_9] FROM [MAKE_ITEM_GRADECODE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MakeItemGradeCodeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[byItemIndex]",
		},
		clause.Column{
			Name: "[byGrade_1]",
		},
		clause.Column{
			Name: "[byGrade_2]",
		},
		clause.Column{
			Name: "[byGrade_3]",
		},
		clause.Column{
			Name: "[byGrade_4]",
		},
		clause.Column{
			Name: "[byGrade_5]",
		},
		clause.Column{
			Name: "[byGrade_6]",
		},
		clause.Column{
			Name: "[byGrade_7]",
		},
		clause.Column{
			Name: "[byGrade_8]",
		},
		clause.Column{
			Name: "[byGrade_9]",
		},
	},
}
