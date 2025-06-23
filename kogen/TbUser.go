package kogen

import (
	"fmt"
	"time"
)

const (
	_TbUserDatabaseNbr = 0
	_TbUserTableName   = "TB_USER"
)

func init() {
	ModelList = append(ModelList, &TbUser{})
}

// TbUser User Account Information
type TbUser struct {
	AccountId     [21]byte  `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	Password      [13]byte  `gorm:"column:strPasswd;type:varchar(13);not null" json:"strPasswd"`
	SocNo         [20]byte  `gorm:"column:strSocNo;type:varchar(20);not null;default:''" json:"strSocNo"`
	Email         [250]byte `gorm:"column:strEmail;type:varchar(250);not null;default:''" json:"strEmail"`
	Authority     uint8     `gorm:"column:strAuthority;type:tinyint;not null;default:1" json:"strAuthority"`
	PremiumExpire time.Time `gorm:"column:PremiumExpire;type:datetime;not null;default:getdate()+(3)" json:"PremiumExpire"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *TbUser) GetDatabaseName() string {
	return GetDatabaseName(DbType(_TbUserDatabaseNbr))
}

// GetTableName Returns the table name
func (this *TbUser) GetTableName() string {
	return _TbUserTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *TbUser) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [TB_USER] (strAccountID, strPasswd, strSocNo, strEmail, strAuthority, PremiumExpire) \nVALUES (%s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.Password),
		GetOptionalBinaryVal(this.SocNo),
		GetOptionalBinaryVal(this.Email),
		GetOptionalDecVal(&this.Authority),
		GetDateTimeExportFmt(&this.PremiumExpire))
}

// GetCreateTableString Returns the create table statement for this object
func (this *TbUser) GetCreateTableString() string {
	query := "CREATE TABLE [TB_USER] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strPasswd] varchar(13) NOT NULL,\n\t[strSocNo] varchar(20) NOT NULL,\n\t[strEmail] varchar(250) NOT NULL,\n\t[strAuthority] tinyint NOT NULL,\n\t[PremiumExpire] datetime NOT NULL\n\n)\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strSocNo] DEFAULT '' FOR [strSocNo]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strEmail] DEFAULT '' FOR [strEmail]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_strAuthority] DEFAULT 1 FOR [strAuthority]\nGO\nALTER TABLE [TB_USER] ADD CONSTRAINT [DF_TB_USER_PremiumExpire] DEFAULT getdate()+(3) FOR [PremiumExpire]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
