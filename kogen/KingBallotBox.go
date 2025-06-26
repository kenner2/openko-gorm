package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KingBallotBoxDatabaseNbr = "GAME"
	_KingBallotBoxTableName   = "KING_BALLOT_BOX"
)

func init() {
	ModelList = append(ModelList, &KingBallotBox{})
}

// KingBallotBox King Ballot Box TODO
type KingBallotBox struct {
	AccountId   mssql.VarChar `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId      mssql.VarChar `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	Nation      uint8         `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	CandidateId mssql.VarChar `gorm:"column:strCandidacyID;type:varchar(21);not null" json:"strCandidacyID"`
}

// GetDatabaseName Returns the table's database name
func (this KingBallotBox) GetDatabaseName() string {
	return GetDatabaseName(_KingBallotBoxDatabaseNbr)
}

// TableName Returns the table name
func (this KingBallotBox) TableName() string {
	return _KingBallotBoxTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KingBallotBox) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_BALLOT_BOX] ([strAccountID], [strCharID], [byNation], [strCandidacyID]) VALUES\n(%s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalVarCharVal(&this.CandidateId, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KingBallotBox) GetInsertHeader() string {
	return "INSERT INTO [KING_BALLOT_BOX] ([strAccountID], [strCharID], [byNation], [strCandidacyID]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KingBallotBox) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalVarCharVal(&this.CandidateId, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this KingBallotBox) GetCreateTableString() string {
	query := "CREATE TABLE [KING_BALLOT_BOX] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[strCandidacyID] varchar(21) NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KingBallotBox) SelectClause() (selectClause clause.Select) {
	return _KingBallotBoxSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KingBallotBox) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KingBallotBox{}
	rawSql := "SELECT [strAccountID], [strCharID], [byNation], [strCandidacyID] FROM [KING_BALLOT_BOX]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KingBallotBoxSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[byNation]",
		},
		clause.Column{
			Name: "[strCandidacyID]",
		},
	},
}
