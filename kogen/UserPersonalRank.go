package kogen

import (
	"fmt"
)

const (
	_UserPersonalRankDatabaseNbr = 0
	_UserPersonalRankTableName   = "USER_PERSONAL_RANK"
)

func init() {
	ModelList = append(ModelList, &UserPersonalRank{})
}

// UserPersonalRank User personal ranking
type UserPersonalRank struct {
	Rank                int16    `gorm:"column:nRank;type:smallint;primaryKey;not null" json:"nRank"`
	Position            [21]byte `gorm:"column:strPosition;type:varchar(21);not null" json:"strPosition"`
	ElmoUp              int16    `gorm:"column:nElmoUP;type:smallint;not null" json:"nElmoUP"`
	ElmoUserId          [21]byte `gorm:"column:strElmoUserID;type:varchar(21)" json:"strElmoUserID,omitempty"`
	ElmoLoyaltyMonthly  *int     `gorm:"column:nElmoLoyaltyMonthly;type:int" json:"nElmoLoyaltyMonthly,omitempty"`
	ElmoCheck           int      `gorm:"column:nElmoCheck;type:int;not null;default:0" json:"nElmoCheck"`
	KarusUp             int16    `gorm:"column:nKarusUP;type:smallint;not null" json:"nKarusUP"`
	KarusUserId         [21]byte `gorm:"column:strKarusUserID;type:varchar(21)" json:"strKarusUserID,omitempty"`
	KarusLoyaltyMonthly *int     `gorm:"column:nKarusLoyaltyMonthly;type:int" json:"nKarusLoyaltyMonthly,omitempty"`
	KarusCheck          int      `gorm:"column:nKarusCheck;type:int;not null;default:0" json:"nKarusCheck"`
	Salary              int      `gorm:"column:nSalary;type:int;not null" json:"nSalary"`
	UpdateDate          int      `gorm:"column:UpdateDate;type:smalldatetime;not null" json:"UpdateDate"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserPersonalRank) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserPersonalRankDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserPersonalRank) GetTableName() string {
	return _UserPersonalRankTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserPersonalRank) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_PERSONAL_RANK] (nRank, strPosition, nElmoUP, strElmoUserID, nElmoLoyaltyMonthly, nElmoCheck, nKarusUP, strKarusUserID, nKarusLoyaltyMonthly, nKarusCheck, nSalary, UpdateDate) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalBinaryVal(this.Position),
		GetOptionalDecVal(&this.ElmoUp),
		GetOptionalBinaryVal(this.ElmoUserId),
		GetOptionalDecVal(this.ElmoLoyaltyMonthly),
		GetOptionalDecVal(&this.ElmoCheck),
		GetOptionalDecVal(&this.KarusUp),
		GetOptionalBinaryVal(this.KarusUserId),
		GetOptionalDecVal(this.KarusLoyaltyMonthly),
		GetOptionalDecVal(&this.KarusCheck),
		GetOptionalDecVal(&this.Salary),
		GetOptionalDecVal(&this.UpdateDate))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserPersonalRank) GetCreateTableString() string {
	query := "CREATE TABLE [USER_PERSONAL_RANK] (\n\t\"nRank\" smallint NOT NULL,\n\t\"strPosition\" varchar(21) NOT NULL,\n\t\"nElmoUP\" smallint NOT NULL,\n\t\"strElmoUserID\" varchar(21),\n\t\"nElmoLoyaltyMonthly\" int,\n\t\"nElmoCheck\" int NOT NULL,\n\t\"nKarusUP\" smallint NOT NULL,\n\t\"strKarusUserID\" varchar(21),\n\t\"nKarusLoyaltyMonthly\" int,\n\t\"nKarusCheck\" int NOT NULL,\n\t\"nSalary\" int NOT NULL,\n\t\"UpdateDate\" smalldatetime NOT NULL\n\tPRIMARY KEY (\"nRank\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
