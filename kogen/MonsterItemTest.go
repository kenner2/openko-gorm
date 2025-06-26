package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MonsterItemTestDatabaseNbr = 1
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

// GetDatabaseName Returns the table's database name
func (this MonsterItemTest) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MonsterItemTestDatabaseNbr))
}

// TableName Returns the table name
func (this MonsterItemTest) TableName() string {
	return _MonsterItemTestTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MonsterItemTest) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_ITEM_TEST] ([sIndex], [iItem01], [sPersent01], [iItem02], [sPersent02], [iItem03], [sPersent03], [iItem04], [sPersent04], [iItem05], [sPersent05]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
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

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MonsterItemTest) GetInsertHeader() string {
	return "INSERT INTO [MONSTER_ITEM_TEST] (sIndex, iItem01, sPersent01, iItem02, sPersent02, iItem03, sPersent03, iItem04, sPersent04, iItem05, sPersent05) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MonsterItemTest) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
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
func (this MonsterItemTest) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_ITEM_TEST] (\n\t[sIndex] smallint NOT NULL,\n\t[iItem01] int,\n\t[sPersent01] smallint,\n\t[iItem02] int,\n\t[sPersent02] smallint,\n\t[iItem03] int,\n\t[sPersent03] smallint,\n\t[iItem04] int,\n\t[sPersent04] smallint,\n\t[iItem05] int,\n\t[sPersent05] smallint\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MonsterItemTest) SelectClause() (selectClause clause.Select) {
	return _MonsterItemTestSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MonsterItemTest) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MonsterItemTest{}
	rawSql := "SELECT [sIndex], [iItem01], [sPersent01], [iItem02], [sPersent02], [iItem03], [sPersent03], [iItem04], [sPersent04], [iItem05], [sPersent05] FROM [MONSTER_ITEM_TEST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MonsterItemTestSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIndex]",
		},
		clause.Column{
			Name: "[iItem01]",
		},
		clause.Column{
			Name: "[sPersent01]",
		},
		clause.Column{
			Name: "[iItem02]",
		},
		clause.Column{
			Name: "[sPersent02]",
		},
		clause.Column{
			Name: "[iItem03]",
		},
		clause.Column{
			Name: "[sPersent03]",
		},
		clause.Column{
			Name: "[iItem04]",
		},
		clause.Column{
			Name: "[sPersent04]",
		},
		clause.Column{
			Name: "[iItem05]",
		},
		clause.Column{
			Name: "[sPersent05]",
		},
	},
}
