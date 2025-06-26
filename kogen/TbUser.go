package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_TbUserDatabaseNbr = "GAME"
	_TbUserTableName   = "TB_USER"
)

func init() {
	ModelList = append(ModelList, &TbUser{})
}

// TbUser User Account Information
type TbUser struct {
	AccountId     mssql.VarChar `gorm:"column:strAccountID;type:varchar(21);primaryKey;not null" json:"strAccountID"`
	Password      mssql.VarChar `gorm:"column:strPasswd;type:varchar(13);not null" json:"strPasswd"`
	SocNo         mssql.VarChar `gorm:"column:strSocNo;type:varchar(20);not null;default:''" json:"strSocNo"`
	Email         mssql.VarChar `gorm:"column:strEmail;type:varchar(250);not null;default:''" json:"strEmail"`
	Authority     uint8         `gorm:"column:strAuthority;type:tinyint;not null;default:1" json:"strAuthority"`
	PremiumExpire time.Time     `gorm:"column:PremiumExpire;type:datetime;not null;default:getdate()+(3)" json:"PremiumExpire"`
}

// GetDatabaseName Returns the table's database name
func (this TbUser) GetDatabaseName() string {
	return GetDatabaseName(_TbUserDatabaseNbr)
}

// TableName Returns the table name
func (this TbUser) TableName() string {
	return _TbUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this TbUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [TB_USER] ([strAccountID], [strPasswd], [strSocNo], [strEmail], [strAuthority], [PremiumExpire]) VALUES\n(%s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.Password, false),
		GetOptionalVarCharVal(&this.SocNo, false),
		GetOptionalVarCharVal(&this.Email, false),
		GetOptionalDecVal(&this.Authority),
		GetDateTimeExportFmt(&this.PremiumExpire))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this TbUser) GetInsertHeader() string {
	return "INSERT INTO [TB_USER] ([strAccountID], [strPasswd], [strSocNo], [strEmail], [strAuthority], [PremiumExpire]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this TbUser) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.Password, false),
		GetOptionalVarCharVal(&this.SocNo, false),
		GetOptionalVarCharVal(&this.Email, false),
		GetOptionalDecVal(&this.Authority),
		GetDateTimeExportFmt(&this.PremiumExpire))
}

// GetCreateTableString Returns the create table statement for this object
func (this TbUser) GetCreateTableString() string {
	query := "CREATE TABLE [TB_USER] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strPasswd] varchar(13) NOT NULL,\n\t[strSocNo] varchar(20) NOT NULL,\n\t[strEmail] varchar(250) NOT NULL,\n\t[strAuthority] tinyint NOT NULL,\n\t[PremiumExpire] datetime NOT NULL\n\tCONSTRAINT [PK_TB_USER] PRIMARY KEY ([strAccountID])\n)\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strSocNo] DEFAULT '' FOR [strSocNo]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strEmail] DEFAULT '' FOR [strEmail]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strAuthority] DEFAULT 1 FOR [strAuthority]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_PremiumExpire] DEFAULT getdate()+(3) FOR [PremiumExpire]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this TbUser) SelectClause() (selectClause clause.Select) {
	return _TbUserSelectClause
}

// GetAllTableData Returns a list of all table data
func (this TbUser) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []TbUser{}
	rawSql := "SELECT [strAccountID], [strPasswd], [strSocNo], [strEmail], [strAuthority], [PremiumExpire] FROM [TB_USER]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _TbUserSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strPasswd]",
		},
		clause.Column{
			Name: "[strSocNo]",
		},
		clause.Column{
			Name: "[strEmail]",
		},
		clause.Column{
			Name: "[strAuthority]",
		},
		clause.Column{
			Name: "[PremiumExpire]",
		},
	},
}
