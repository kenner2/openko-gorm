package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_UserKnightsRankDatabaseNbr = "GAME"
	_UserKnightsRankTableName   = "USER_KNIGHTS_RANK"
)

func init() {
	ModelList = append(ModelList, &UserKnightsRank{})
}

// UserKnightsRank User Knights Ranking
type UserKnightsRank struct {
	Index            int16          `gorm:"column:shIndex;type:smallint;primaryKey;not null" json:"shIndex"`
	Name             mssql.VarChar  `gorm:"column:strName;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strName"`
	ElmoUserId       *mssql.VarChar `gorm:"column:strElmoUserID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strElmoUserID,omitempty"`
	ElmoKnightsName  *mssql.VarChar `gorm:"column:strElmoKnightsName;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strElmoKnightsName,omitempty"`
	ElmoLoyalty      *int           `gorm:"column:nElmoLoyalty;type:int" json:"nElmoLoyalty,omitempty"`
	KarusUserId      *mssql.VarChar `gorm:"column:strKarusUserID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strKarusUserID,omitempty"`
	KarusKnightsName *mssql.VarChar `gorm:"column:strKarusKnightsName;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strKarusKnightsName,omitempty"`
	KarusLoyalty     *int           `gorm:"column:nKarusLoyalty;type:int" json:"nKarusLoyalty,omitempty"`
	Money            int            `gorm:"column:nMoney;type:int;not null" json:"nMoney"`
}

// GetDatabaseName Returns the table's database name
func (this UserKnightsRank) GetDatabaseName() string {
	return GetDatabaseName(_UserKnightsRankDatabaseNbr)
}

// TableName Returns the table name
func (this UserKnightsRank) TableName() string {
	return _UserKnightsRankTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserKnightsRank) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_KNIGHTS_RANK] ([shIndex], [strName], [strElmoUserID], [strElmoKnightsName], [nElmoLoyalty], [strKarusUserID], [strKarusKnightsName], [nKarusLoyalty], [nMoney]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(this.ElmoUserId, false),
		GetOptionalVarCharVal(this.ElmoKnightsName, false),
		GetOptionalDecVal(this.ElmoLoyalty),
		GetOptionalVarCharVal(this.KarusUserId, false),
		GetOptionalVarCharVal(this.KarusKnightsName, false),
		GetOptionalDecVal(this.KarusLoyalty),
		GetOptionalDecVal(&this.Money))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserKnightsRank) GetInsertHeader() string {
	return "INSERT INTO [USER_KNIGHTS_RANK] ([shIndex], [strName], [strElmoUserID], [strElmoKnightsName], [nElmoLoyalty], [strKarusUserID], [strKarusKnightsName], [nKarusLoyalty], [nMoney]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserKnightsRank) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(this.ElmoUserId, false),
		GetOptionalVarCharVal(this.ElmoKnightsName, false),
		GetOptionalDecVal(this.ElmoLoyalty),
		GetOptionalVarCharVal(this.KarusUserId, false),
		GetOptionalVarCharVal(this.KarusKnightsName, false),
		GetOptionalDecVal(this.KarusLoyalty),
		GetOptionalDecVal(&this.Money))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserKnightsRank) GetCreateTableString() string {
	query := "CREATE TABLE [USER_KNIGHTS_RANK] (\n\t[shIndex] smallint NOT NULL,\n\t[strName] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strElmoUserID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[strElmoKnightsName] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[nElmoLoyalty] int,\n\t[strKarusUserID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[strKarusKnightsName] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[nKarusLoyalty] int,\n\t[nMoney] int NOT NULL\n\tCONSTRAINT [PK_USER_KNIGHTS_RANK] PRIMARY KEY CLUSTERED ([shIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserKnightsRank) SelectClause() (selectClause clause.Select) {
	return _UserKnightsRankSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserKnightsRank) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserKnightsRank{}
	rawSql := "SELECT [shIndex], [strName], [strElmoUserID], [strElmoKnightsName], [nElmoLoyalty], [strKarusUserID], [strKarusKnightsName], [nKarusLoyalty], [nMoney] FROM [USER_KNIGHTS_RANK]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserKnightsRankSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[shIndex]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[strElmoUserID]",
		},
		clause.Column{
			Name: "[strElmoKnightsName]",
		},
		clause.Column{
			Name: "[nElmoLoyalty]",
		},
		clause.Column{
			Name: "[strKarusUserID]",
		},
		clause.Column{
			Name: "[strKarusKnightsName]",
		},
		clause.Column{
			Name: "[nKarusLoyalty]",
		},
		clause.Column{
			Name: "[nMoney]",
		},
	},
}
