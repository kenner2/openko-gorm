package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MonsterSummonListDatabaseNbr = "GAME"
	_MonsterSummonListTableName   = "MONSTER_SUMMON_LIST"
)

func init() {
	ModelList = append(ModelList, &MonsterSummonList{})
}

// MonsterSummonList Monster summon list
type MonsterSummonList struct {
	MonsterId   int16          `gorm:"column:sSid;type:smallint;not null" json:"sSid"`
	Name        *mssql.VarChar `gorm:"column:strName;type:varchar(31)" json:"strName,omitempty"`
	Level       int16          `gorm:"column:sLevel;type:smallint;not null" json:"sLevel"`
	Probability int16          `gorm:"column:sProbability;type:smallint;not null" json:"sProbability"`
	Type        uint8          `gorm:"column:bType;type:tinyint;not null;default:0" json:"bType"`
}

// GetDatabaseName Returns the table's database name
func (this MonsterSummonList) GetDatabaseName() string {
	return GetDatabaseName(_MonsterSummonListDatabaseNbr)
}

// TableName Returns the table name
func (this MonsterSummonList) TableName() string {
	return _MonsterSummonListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MonsterSummonList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_SUMMON_LIST] ([sSid], [strName], [sLevel], [sProbability], [bType]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MonsterId),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Probability),
		GetOptionalDecVal(&this.Type))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MonsterSummonList) GetInsertHeader() string {
	return "INSERT INTO [MONSTER_SUMMON_LIST] ([sSid], [strName], [sLevel], [sProbability], [bType]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MonsterSummonList) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MonsterId),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Probability),
		GetOptionalDecVal(&this.Type))
}

// GetCreateTableString Returns the create table statement for this object
func (this MonsterSummonList) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_SUMMON_LIST] (\n\t[sSid] smallint NOT NULL,\n\t[strName] varchar(31),\n\t[sLevel] smallint NOT NULL,\n\t[sProbability] smallint NOT NULL,\n\t[bType] tinyint NOT NULL\n\n)\nGO\nALTER TABLE [MONSTER_SUMMON_LIST] ADD CONSTRAINT [DF_MONSTER_SUMMON_LIST_bType] DEFAULT 0 FOR [bType]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MonsterSummonList) SelectClause() (selectClause clause.Select) {
	return _MonsterSummonListSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MonsterSummonList) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MonsterSummonList{}
	rawSql := "SELECT [sSid], [strName], [sLevel], [sProbability], [bType] FROM [MONSTER_SUMMON_LIST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MonsterSummonListSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sSid]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[sLevel]",
		},
		clause.Column{
			Name: "[sProbability]",
		},
		clause.Column{
			Name: "[bType]",
		},
	},
}
