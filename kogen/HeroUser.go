package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_HeroUserDatabaseNbr = "GAME"
	_HeroUserTableName   = "HERO_USER"
)

func init() {
	ModelList = append(ModelList, &HeroUser{})
}

// HeroUser TODO Doc
type HeroUser struct {
	Index       int16          `gorm:"column:shIndex;type:smallint;not null" json:"shIndex"`
	UserId      *mssql.VarChar `gorm:"column:strUserID;type:varchar(21)" json:"strUserID,omitempty"`
	Nation      *mssql.VarChar `gorm:"column:strNation;type:varchar(20)" json:"strNation,omitempty"`
	ClassName   *mssql.VarChar `gorm:"column:strClass;type:varchar(30)" json:"strClass,omitempty"`
	Achievement *mssql.VarChar `gorm:"column:strAchievement;type:varchar(50)" json:"strAchievement,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this HeroUser) GetDatabaseName() string {
	return GetDatabaseName(_HeroUserDatabaseNbr)
}

// TableName Returns the table name
func (this HeroUser) TableName() string {
	return _HeroUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this HeroUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HERO_USER] ([shIndex], [strUserID], [strNation], [strClass], [strAchievement]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.UserId, false),
		GetOptionalVarCharVal(this.Nation, false),
		GetOptionalVarCharVal(this.ClassName, false),
		GetOptionalVarCharVal(this.Achievement, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this HeroUser) GetInsertHeader() string {
	return "INSERT INTO [HERO_USER] ([shIndex], [strUserID], [strNation], [strClass], [strAchievement]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this HeroUser) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.UserId, false),
		GetOptionalVarCharVal(this.Nation, false),
		GetOptionalVarCharVal(this.ClassName, false),
		GetOptionalVarCharVal(this.Achievement, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this HeroUser) GetCreateTableString() string {
	query := "CREATE TABLE [HERO_USER] (\n\t[shIndex] smallint NOT NULL,\n\t[strUserID] varchar(21),\n\t[strNation] varchar(20),\n\t[strClass] varchar(30),\n\t[strAchievement] varchar(50)\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this HeroUser) SelectClause() (selectClause clause.Select) {
	return _HeroUserSelectClause
}

// GetAllTableData Returns a list of all table data
func (this HeroUser) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []HeroUser{}
	rawSql := "SELECT [shIndex], [strUserID], [strNation], [strClass], [strAchievement] FROM [HERO_USER]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _HeroUserSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[shIndex]",
		},
		clause.Column{
			Name: "[strUserID]",
		},
		clause.Column{
			Name: "[strNation]",
		},
		clause.Column{
			Name: "[strClass]",
		},
		clause.Column{
			Name: "[strAchievement]",
		},
	},
}
