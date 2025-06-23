package kogen

import (
	"fmt"
)

const (
	_UserKnightsRankDatabaseNbr = 0
	_UserKnightsRankTableName   = "USER_KNIGHTS_RANK"
)

func init() {
	ModelList = append(ModelList, &UserKnightsRank{})
}

// UserKnightsRank User Knights Ranking
type UserKnightsRank struct {
	Index            int16    `gorm:"column:shIndex;type:smallint;primaryKey;not null" json:"shIndex"`
	Name             [21]byte `gorm:"column:strName;type:varchar(21);not null" json:"strName"`
	ElmoUserId       [21]byte `gorm:"column:strElmoUserID;type:varchar(21)" json:"strElmoUserID,omitempty"`
	ElmoKnightsName  [21]byte `gorm:"column:strElmoKnightsName;type:varchar(21)" json:"strElmoKnightsName,omitempty"`
	ElmoLoyalty      *int     `gorm:"column:nElmoLoyalty;type:int" json:"nElmoLoyalty,omitempty"`
	KarusUserId      [21]byte `gorm:"column:strKarusUserID;type:varchar(21)" json:"strKarusUserID,omitempty"`
	KarusKnightsName [21]byte `gorm:"column:strKarusKnightsName;type:varchar(21)" json:"strKarusKnightsName,omitempty"`
	KarusLoyalty     *int     `gorm:"column:nKarusLoyalty;type:int" json:"nKarusLoyalty,omitempty"`
	Money            int      `gorm:"column:nMoney;type:int;not null" json:"nMoney"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserKnightsRank) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserKnightsRankDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserKnightsRank) GetTableName() string {
	return _UserKnightsRankTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserKnightsRank) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_KNIGHTS_RANK] (shIndex, strName, strElmoUserID, strElmoKnightsName, nElmoLoyalty, strKarusUserID, strKarusKnightsName, nKarusLoyalty, nMoney) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.ElmoUserId),
		GetOptionalBinaryVal(this.ElmoKnightsName),
		GetOptionalDecVal(this.ElmoLoyalty),
		GetOptionalBinaryVal(this.KarusUserId),
		GetOptionalBinaryVal(this.KarusKnightsName),
		GetOptionalDecVal(this.KarusLoyalty),
		GetOptionalDecVal(&this.Money))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserKnightsRank) GetCreateTableString() string {
	query := "CREATE TABLE [USER_KNIGHTS_RANK] (\n\t[shIndex] smallint NOT NULL,\n\t[strName] varchar(21) NOT NULL,\n\t[strElmoUserID] varchar(21),\n\t[strElmoKnightsName] varchar(21),\n\t[nElmoLoyalty] int,\n\t[strKarusUserID] varchar(21),\n\t[strKarusKnightsName] varchar(21),\n\t[nKarusLoyalty] int,\n\t[nMoney] int NOT NULL\n\tCONSTRAINT [PK_USER_KNIGHTS_RANK] PRIMARY KEY ([shIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
