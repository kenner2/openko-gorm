package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_AccountCharDatabaseNbr = "GAME"
	_AccountCharTableName   = "ACCOUNT_CHAR"
)

func init() {
	ModelList = append(ModelList, &AccountChar{})
}

// AccountChar Represents the relationship between accounts and characters
type AccountChar struct {
	AccountId mssql.VarChar  `gorm:"column:strAccountID;type:varchar(21);primaryKey;not null" json:"strAccountID"`
	Nation    uint8          `gorm:"column:bNation;type:tinyint;not null" json:"bNation"`
	CharNum   uint8          `gorm:"column:bCharNum;type:tinyint;not null;default:0" json:"bCharNum"`
	CharId1   *mssql.VarChar `gorm:"column:strCharID1;type:varchar(21)" json:"strCharID1,omitempty"`
	CharId2   *mssql.VarChar `gorm:"column:strCharID2;type:varchar(21)" json:"strCharID2,omitempty"`
	CharId3   *mssql.VarChar `gorm:"column:strCharID3;type:varchar(21)" json:"strCharID3,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this AccountChar) GetDatabaseName() string {
	return GetDatabaseName(_AccountCharDatabaseNbr)
}

// TableName Returns the table name
func (this AccountChar) TableName() string {
	return _AccountCharTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this AccountChar) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ACCOUNT_CHAR] ([strAccountID], [bNation], [bCharNum], [strCharID1], [strCharID2], [strCharID3]) VALUES\n(%s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.CharNum),
		GetOptionalVarCharVal(this.CharId1, false),
		GetOptionalVarCharVal(this.CharId2, false),
		GetOptionalVarCharVal(this.CharId3, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this AccountChar) GetInsertHeader() string {
	return "INSERT INTO [ACCOUNT_CHAR] ([strAccountID], [bNation], [bCharNum], [strCharID1], [strCharID2], [strCharID3]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this AccountChar) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.CharNum),
		GetOptionalVarCharVal(this.CharId1, false),
		GetOptionalVarCharVal(this.CharId2, false),
		GetOptionalVarCharVal(this.CharId3, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this AccountChar) GetCreateTableString() string {
	query := "CREATE TABLE [ACCOUNT_CHAR] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[bNation] tinyint NOT NULL,\n\t[bCharNum] tinyint NOT NULL,\n\t[strCharID1] varchar(21),\n\t[strCharID2] varchar(21),\n\t[strCharID3] varchar(21)\n\tCONSTRAINT [PK_ACCOUNT_CHAR] PRIMARY KEY CLUSTERED ([strAccountID])\n)\nGO\nALTER TABLE [ACCOUNT_CHAR] ADD CONSTRAINT [DF_ACCOUNT_CHAR_bCharNum] DEFAULT 0 FOR [bCharNum]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this AccountChar) SelectClause() (selectClause clause.Select) {
	return _AccountCharSelectClause
}

// GetAllTableData Returns a list of all table data
func (this AccountChar) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []AccountChar{}
	rawSql := "SELECT [strAccountID], [bNation], [bCharNum], [strCharID1], [strCharID2], [strCharID3] FROM [ACCOUNT_CHAR]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _AccountCharSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[bNation]",
		},
		clause.Column{
			Name: "[bCharNum]",
		},
		clause.Column{
			Name: "[strCharID1]",
		},
		clause.Column{
			Name: "[strCharID2]",
		},
		clause.Column{
			Name: "[strCharID3]",
		},
	},
}
