package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KingElectionListDatabaseNbr = "GAME"
	_KingElectionListTableName   = "KING_ELECTION_LIST"
)

func init() {
	ModelList = append(ModelList, &KingElectionList{})
}

// KingElectionList King election list
type KingElectionList struct {
	Type    uint8         `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	Nation  uint8         `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	Knights int16         `gorm:"column:nKnights;type:smallint;not null" json:"nKnights"`
	Name    mssql.VarChar `gorm:"column:strName;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strName"`
	Money   int           `gorm:"column:nMoney;type:int;not null" json:"nMoney"`
}

// GetDatabaseName Returns the table's database name
func (this KingElectionList) GetDatabaseName() string {
	return GetDatabaseName(_KingElectionListDatabaseNbr)
}

// TableName Returns the table name
func (this KingElectionList) TableName() string {
	return _KingElectionListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KingElectionList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_ELECTION_LIST] ([byType], [byNation], [nKnights], [strName], [nMoney]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Knights),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Money))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KingElectionList) GetInsertHeader() string {
	return "INSERT INTO [KING_ELECTION_LIST] ([byType], [byNation], [nKnights], [strName], [nMoney]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KingElectionList) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Knights),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Money))
}

// GetCreateTableString Returns the create table statement for this object
func (this KingElectionList) GetCreateTableString() string {
	query := "CREATE TABLE [KING_ELECTION_LIST] (\n\t[byType] tinyint NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[nKnights] smallint NOT NULL,\n\t[strName] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[nMoney] int NOT NULL\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KingElectionList) SelectClause() (selectClause clause.Select) {
	return _KingElectionListSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KingElectionList) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KingElectionList{}
	rawSql := "SELECT [byType], [byNation], [nKnights], [strName], [nMoney] FROM [KING_ELECTION_LIST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KingElectionListSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[byType]",
		},
		clause.Column{
			Name: "[byNation]",
		},
		clause.Column{
			Name: "[nKnights]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[nMoney]",
		},
	},
}
