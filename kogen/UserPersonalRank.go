package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_UserPersonalRankDatabaseNbr = "GAME"
	_UserPersonalRankTableName   = "USER_PERSONAL_RANK"
)

func init() {
	ModelList = append(ModelList, &UserPersonalRank{})
}

// UserPersonalRank User personal ranking
type UserPersonalRank struct {
	Rank                int16          `gorm:"column:nRank;type:smallint;primaryKey;not null" json:"nRank"`
	Position            mssql.VarChar  `gorm:"column:strPosition;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strPosition"`
	ElmoUp              int16          `gorm:"column:nElmoUP;type:smallint;not null" json:"nElmoUP"`
	ElmoUserId          *mssql.VarChar `gorm:"column:strElmoUserID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strElmoUserID,omitempty"`
	ElmoLoyaltyMonthly  *int           `gorm:"column:nElmoLoyaltyMonthly;type:int" json:"nElmoLoyaltyMonthly,omitempty"`
	ElmoCheck           int            `gorm:"column:nElmoCheck;type:int;not null;default:0" json:"nElmoCheck"`
	KarusUp             int16          `gorm:"column:nKarusUP;type:smallint;not null" json:"nKarusUP"`
	KarusUserId         *mssql.VarChar `gorm:"column:strKarusUserID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strKarusUserID,omitempty"`
	KarusLoyaltyMonthly *int           `gorm:"column:nKarusLoyaltyMonthly;type:int" json:"nKarusLoyaltyMonthly,omitempty"`
	KarusCheck          int            `gorm:"column:nKarusCheck;type:int;not null;default:0" json:"nKarusCheck"`
	Salary              int            `gorm:"column:nSalary;type:int;not null" json:"nSalary"`
	UpdateDate          time.Time      `gorm:"column:UpdateDate;type:smalldatetime;not null" json:"UpdateDate"`
}

// GetDatabaseName Returns the table's database name
func (this UserPersonalRank) GetDatabaseName() string {
	return GetDatabaseName(_UserPersonalRankDatabaseNbr)
}

// TableName Returns the table name
func (this UserPersonalRank) TableName() string {
	return _UserPersonalRankTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserPersonalRank) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_PERSONAL_RANK] ([nRank], [strPosition], [nElmoUP], [strElmoUserID], [nElmoLoyaltyMonthly], [nElmoCheck], [nKarusUP], [strKarusUserID], [nKarusLoyaltyMonthly], [nKarusCheck], [nSalary], [UpdateDate]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalVarCharVal(&this.Position, false),
		GetOptionalDecVal(&this.ElmoUp),
		GetOptionalVarCharVal(this.ElmoUserId, false),
		GetOptionalDecVal(this.ElmoLoyaltyMonthly),
		GetOptionalDecVal(&this.ElmoCheck),
		GetOptionalDecVal(&this.KarusUp),
		GetOptionalVarCharVal(this.KarusUserId, false),
		GetOptionalDecVal(this.KarusLoyaltyMonthly),
		GetOptionalDecVal(&this.KarusCheck),
		GetOptionalDecVal(&this.Salary),
		GetDateTimeExportFmt(&this.UpdateDate))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserPersonalRank) GetInsertHeader() string {
	return "INSERT INTO [USER_PERSONAL_RANK] ([nRank], [strPosition], [nElmoUP], [strElmoUserID], [nElmoLoyaltyMonthly], [nElmoCheck], [nKarusUP], [strKarusUserID], [nKarusLoyaltyMonthly], [nKarusCheck], [nSalary], [UpdateDate]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserPersonalRank) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalVarCharVal(&this.Position, false),
		GetOptionalDecVal(&this.ElmoUp),
		GetOptionalVarCharVal(this.ElmoUserId, false),
		GetOptionalDecVal(this.ElmoLoyaltyMonthly),
		GetOptionalDecVal(&this.ElmoCheck),
		GetOptionalDecVal(&this.KarusUp),
		GetOptionalVarCharVal(this.KarusUserId, false),
		GetOptionalDecVal(this.KarusLoyaltyMonthly),
		GetOptionalDecVal(&this.KarusCheck),
		GetOptionalDecVal(&this.Salary),
		GetDateTimeExportFmt(&this.UpdateDate))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserPersonalRank) GetCreateTableString() string {
	query := "CREATE TABLE [USER_PERSONAL_RANK] (\n\t[nRank] smallint NOT NULL,\n\t[strPosition] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[nElmoUP] smallint NOT NULL,\n\t[strElmoUserID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[nElmoLoyaltyMonthly] int,\n\t[nElmoCheck] int NOT NULL,\n\t[nKarusUP] smallint NOT NULL,\n\t[strKarusUserID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[nKarusLoyaltyMonthly] int,\n\t[nKarusCheck] int NOT NULL,\n\t[nSalary] int NOT NULL,\n\t[UpdateDate] smalldatetime NOT NULL\n\tCONSTRAINT [PK_USER_PERSONAL_RANK] PRIMARY KEY CLUSTERED ([nRank])\n)\nGO\nALTER TABLE [USER_PERSONAL_RANK] ADD CONSTRAINT [DF_USER_PERSONAL_RANK_nElmoCheck] DEFAULT 0 FOR [nElmoCheck]\nGO\nALTER TABLE [USER_PERSONAL_RANK] ADD CONSTRAINT [DF_USER_PERSONAL_RANK_nKarusCheck] DEFAULT 0 FOR [nKarusCheck]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserPersonalRank) SelectClause() (selectClause clause.Select) {
	return _UserPersonalRankSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserPersonalRank) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserPersonalRank{}
	rawSql := "SELECT [nRank], [strPosition], [nElmoUP], [strElmoUserID], [nElmoLoyaltyMonthly], [nElmoCheck], [nKarusUP], [strKarusUserID], [nKarusLoyaltyMonthly], [nKarusCheck], [nSalary], [UpdateDate] FROM [USER_PERSONAL_RANK]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserPersonalRankSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nRank]",
		},
		clause.Column{
			Name: "[strPosition]",
		},
		clause.Column{
			Name: "[nElmoUP]",
		},
		clause.Column{
			Name: "[strElmoUserID]",
		},
		clause.Column{
			Name: "[nElmoLoyaltyMonthly]",
		},
		clause.Column{
			Name: "[nElmoCheck]",
		},
		clause.Column{
			Name: "[nKarusUP]",
		},
		clause.Column{
			Name: "[strKarusUserID]",
		},
		clause.Column{
			Name: "[nKarusLoyaltyMonthly]",
		},
		clause.Column{
			Name: "[nKarusCheck]",
		},
		clause.Column{
			Name: "[nSalary]",
		},
		clause.Column{
			Name: "[UpdateDate]",
		},
	},
}
