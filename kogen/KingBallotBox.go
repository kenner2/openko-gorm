package kogen

import (
	"fmt"
)

const (
	_KingBallotBoxDatabaseNbr = 0
	_KingBallotBoxTableName   = "KING_BALLOT_BOX"
)

func init() {
	ModelList = append(ModelList, &KingBallotBox{})
}

// KingBallotBox King Ballot Box TODO
type KingBallotBox struct {
	AccountId   [21]byte `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId      [21]byte `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	Nation      uint8    `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	CandidateId [21]byte `gorm:"column:strCandidacyID;type:varchar(21);not null" json:"strCandidacyID"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KingBallotBox) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KingBallotBoxDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KingBallotBox) GetTableName() string {
	return _KingBallotBoxTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KingBallotBox) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_BALLOT_BOX] (strAccountID, strCharID, byNation, strCandidacyID) \nVALUES (%s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.CharId),
		GetOptionalDecVal(&this.Nation),
		GetOptionalBinaryVal(this.CandidateId))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KingBallotBox) GetCreateTableString() string {
	query := "CREATE TABLE [KING_BALLOT_BOX] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[strCandidacyID] varchar(21) NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
