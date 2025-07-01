package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_UserEditorDatabaseNbr = "GAME"
	_UserEditorTableName   = "USER_EDITOR"
)

func init() {
	ModelList = append(ModelList, &UserEditor{})
}

// UserEditor User editor
type UserEditor struct {
	CharId            mssql.VarChar `gorm:"column:strCharID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strCharID"`
	AccountId         mssql.VarChar `gorm:"column:strAccountID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strAccountID"`
	OpId              mssql.VarChar `gorm:"column:strOpID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOpID"`
	OpIP              mssql.VarChar `gorm:"column:strOpIP;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOpIP"`
	OldUserValue      []byte        `gorm:"column:strOldUserValue;type:char(600) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOldUserValue"`
	NewUserValue      []byte        `gorm:"column:strNewUserValue;type:char(600) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNewUserValue"`
	OldUserSkill      []byte        `gorm:"column:strOldUserSkill;type:char(10) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOldUserSkill"`
	NewUserSkill      []byte        `gorm:"column:strNewUserSkill;type:char(10) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNewUserSkill"`
	OldUserItem       []byte        `gorm:"column:strOldUserItem;type:char(400) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOldUserItem"`
	NewUserItem       []byte        `gorm:"column:strNewUserItem;type:char(400) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNewUserItem"`
	OldWarehouseValue []byte        `gorm:"column:strOldWHValue;type:char(100) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOldWHValue"`
	NewWarehouseValue []byte        `gorm:"column:strNewWHValue;type:char(100) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNewWHValue"`
	OldWarehouseItem  []byte        `gorm:"column:strOldWHItem;type:char(1600) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strOldWHItem"`
	NewWarehouseItem  []byte        `gorm:"column:strNewWHItem;type:char(1600) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNewWHItem"`
	EditorTime        time.Time     `gorm:"column:EditorTime;type:smalldatetime;not null;default:getdate()" json:"EditorTime"`
}

// GetDatabaseName Returns the table's database name
func (this UserEditor) GetDatabaseName() string {
	return GetDatabaseName(_UserEditorDatabaseNbr)
}

// TableName Returns the table name
func (this UserEditor) TableName() string {
	return _UserEditorTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserEditor) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_EDITOR] ([strCharID], [strAccountID], [strOpID], [strOpIP], [strOldUserValue], [strNewUserValue], [strOldUserSkill], [strNewUserSkill], [strOldUserItem], [strNewUserItem], [strOldWHValue], [strNewWHValue], [strOldWHItem], [strNewWHItem], [EditorTime]) VALUES\n(%s, %s, %s, %s, %s, %s, CONVERT(char(10), %s), CONVERT(char(10), %s), CONVERT(char(400), %s), CONVERT(char(400), %s), CONVERT(char(100), %s), CONVERT(char(100), %s), CONVERT(char(1600), %s), CONVERT(char(1600), %s), %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.OpId, false),
		GetOptionalVarCharVal(&this.OpIP, false),
		GetOptionalByteArrayVal(&this.OldUserValue, false),
		GetOptionalByteArrayVal(&this.NewUserValue, false),
		GetOptionalByteArrayVal(&this.OldUserSkill, true),
		GetOptionalByteArrayVal(&this.NewUserSkill, true),
		GetOptionalByteArrayVal(&this.OldUserItem, true),
		GetOptionalByteArrayVal(&this.NewUserItem, true),
		GetOptionalByteArrayVal(&this.OldWarehouseValue, true),
		GetOptionalByteArrayVal(&this.NewWarehouseValue, true),
		GetOptionalByteArrayVal(&this.OldWarehouseItem, true),
		GetOptionalByteArrayVal(&this.NewWarehouseItem, true),
		GetDateTimeExportFmt(&this.EditorTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserEditor) GetInsertHeader() string {
	return "INSERT INTO [USER_EDITOR] ([strCharID], [strAccountID], [strOpID], [strOpIP], [strOldUserValue], [strNewUserValue], [strOldUserSkill], [strNewUserSkill], [strOldUserItem], [strNewUserItem], [strOldWHValue], [strNewWHValue], [strOldWHItem], [strNewWHItem], [EditorTime]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserEditor) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, CONVERT(char(10), %s), CONVERT(char(10), %s), CONVERT(char(400), %s), CONVERT(char(400), %s), CONVERT(char(100), %s), CONVERT(char(100), %s), CONVERT(char(1600), %s), CONVERT(char(1600), %s), %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.OpId, false),
		GetOptionalVarCharVal(&this.OpIP, false),
		GetOptionalByteArrayVal(&this.OldUserValue, false),
		GetOptionalByteArrayVal(&this.NewUserValue, false),
		GetOptionalByteArrayVal(&this.OldUserSkill, true),
		GetOptionalByteArrayVal(&this.NewUserSkill, true),
		GetOptionalByteArrayVal(&this.OldUserItem, true),
		GetOptionalByteArrayVal(&this.NewUserItem, true),
		GetOptionalByteArrayVal(&this.OldWarehouseValue, true),
		GetOptionalByteArrayVal(&this.NewWarehouseValue, true),
		GetOptionalByteArrayVal(&this.OldWarehouseItem, true),
		GetOptionalByteArrayVal(&this.NewWarehouseItem, true),
		GetDateTimeExportFmt(&this.EditorTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserEditor) GetCreateTableString() string {
	query := "CREATE TABLE [USER_EDITOR] (\n\t[strCharID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strAccountID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOpID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOpIP] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOldUserValue] char(600) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNewUserValue] char(600) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOldUserSkill] char(10) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNewUserSkill] char(10) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOldUserItem] char(400) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNewUserItem] char(400) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOldWHValue] char(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNewWHValue] char(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strOldWHItem] char(1600) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNewWHItem] char(1600) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[EditorTime] smalldatetime NOT NULL\n)\nGO\nALTER TABLE [USER_EDITOR] ADD CONSTRAINT [DF_USER_EDITOR_EditorTime] DEFAULT getdate() FOR [EditorTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserEditor) SelectClause() (selectClause clause.Select) {
	return _UserEditorSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserEditor) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserEditor{}
	rawSql := "SELECT [strCharID], [strAccountID], [strOpID], [strOpIP], [strOldUserValue], [strNewUserValue], CONVERT(VARBINARY(10), [strOldUserSkill]) as [strOldUserSkill], CONVERT(VARBINARY(10), [strNewUserSkill]) as [strNewUserSkill], CONVERT(VARBINARY(400), [strOldUserItem]) as [strOldUserItem], CONVERT(VARBINARY(400), [strNewUserItem]) as [strNewUserItem], CONVERT(VARBINARY(100), [strOldWHValue]) as [strOldWHValue], CONVERT(VARBINARY(100), [strNewWHValue]) as [strNewWHValue], CONVERT(VARBINARY(1600), [strOldWHItem]) as [strOldWHItem], CONVERT(VARBINARY(1600), [strNewWHItem]) as [strNewWHItem], [EditorTime] FROM [USER_EDITOR]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserEditorSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strOpID]",
		},
		clause.Column{
			Name: "[strOpIP]",
		},
		clause.Column{
			Name: "[strOldUserValue]",
		},
		clause.Column{
			Name: "[strNewUserValue]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(10), [strOldUserSkill]) as [strOldUserSkill]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(10), [strNewUserSkill]) as [strNewUserSkill]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(400), [strOldUserItem]) as [strOldUserItem]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(400), [strNewUserItem]) as [strNewUserItem]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(100), [strOldWHValue]) as [strOldWHValue]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(100), [strNewWHValue]) as [strNewWHValue]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1600), [strOldWHItem]) as [strOldWHItem]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1600), [strNewWHItem]) as [strNewWHItem]",
		},
		clause.Column{
			Name: "[EditorTime]",
		},
	},
}
