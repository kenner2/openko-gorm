package kogen

import (
	"fmt"
)

const (
	_KingElectionListDatabaseNbr = 0
	_KingElectionListTableName   = "KING_ELECTION_LIST"
)

func init() {
	ModelList = append(ModelList, &KingElectionList{})
}

// KingElectionList King election list
type KingElectionList struct {
	Type    uint8    `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	Nation  uint8    `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	Knights *int16   `gorm:"column:nKnights;type:smallint" json:"nKnights,omitempty"`
	Name    [21]byte `gorm:"column:strName;type:varchar(21)" json:"strName,omitempty"`
	Money   int      `gorm:"column:nMoney;type:int;not null" json:"nMoney"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KingElectionList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KingElectionListDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KingElectionList) GetTableName() string {
	return _KingElectionListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KingElectionList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_ELECTION_LIST] (byType, byNation, nKnights, strName, nMoney) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(this.Knights),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(&this.Money))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KingElectionList) GetCreateTableString() string {
	query := "CREATE TABLE [KING_ELECTION_LIST] (\n\t[byType] tinyint NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[nKnights] smallint,\n\t[strName] varchar(21),\n\t[nMoney] int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
