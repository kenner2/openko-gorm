package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_UserDataSkillShortcutDatabaseNbr = 1
	_UserDataSkillShortcutTableName   = "USERDATA_SKILLSHORTCUT"
)

func init() {
	ModelList = append(ModelList, &UserDataSkillShortcut{})
}

// UserDataSkillShortcut User data skill shortcut
type UserDataSkillShortcut struct {
	CharId    mssql.VarChar `gorm:"column:strCharID;type:varchar(21);primaryKey;not null;default:''" json:"strCharID"`
	Count     int16         `gorm:"column:nCount;type:smallint;not null;default:0" json:"nCount"`
	SkillData mssql.VarChar `gorm:"column:strSkillData;type:varchar(260);not null;default:0x00" json:"strSkillData"`
}

// GetDatabaseName Returns the table's database name
func (this UserDataSkillShortcut) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserDataSkillShortcutDatabaseNbr))
}

// TableName Returns the table name
func (this UserDataSkillShortcut) TableName() string {
	return _UserDataSkillShortcutTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserDataSkillShortcut) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USERDATA_SKILLSHORTCUT] ([strCharID], [nCount], [strSkillData]) VALUES\n(%s, %s, CONVERT(varchar(260), %s))", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Count),
		GetOptionalVarCharVal(&this.SkillData, true))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserDataSkillShortcut) GetInsertHeader() string {
	return "INSERT INTO [USERDATA_SKILLSHORTCUT] (strCharID, nCount, strSkillData) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserDataSkillShortcut) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, CONVERT(varchar(260), %s))", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Count),
		GetOptionalVarCharVal(&this.SkillData, true))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserDataSkillShortcut) GetCreateTableString() string {
	query := "CREATE TABLE [USERDATA_SKILLSHORTCUT] (\n\t[strCharID] varchar(21) NOT NULL,\n\t[nCount] smallint NOT NULL,\n\t[strSkillData] varchar(260) NOT NULL\n\tCONSTRAINT [PK_USERDATA_SKILLSHORTCUT] PRIMARY KEY ([strCharID])\n)\nGO\nALTER TABLE [USERDATA_SKILLSHORTCUT] ADD CONSTRAINT [DF_USERDATA_SKILLSHORTCUT_strCharID] DEFAULT '' FOR [strCharID]\nGO\nALTER TABLE [USERDATA_SKILLSHORTCUT] ADD CONSTRAINT [DF_USERDATA_SKILLSHORTCUT_nCount] DEFAULT 0 FOR [nCount]\nGO\nALTER TABLE [USERDATA_SKILLSHORTCUT] ADD CONSTRAINT [DF_USERDATA_SKILLSHORTCUT_strSkillData] DEFAULT 0x00 FOR [strSkillData]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserDataSkillShortcut) SelectClause() (selectClause clause.Select) {
	return _UserDataSkillShortcutSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserDataSkillShortcut) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserDataSkillShortcut{}
	rawSql := "SELECT [strCharID], [nCount], CONVERT(VARBINARY(260), [strSkillData]) as [strSkillData] FROM [USERDATA_SKILLSHORTCUT]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserDataSkillShortcutSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[nCount]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(260), [strSkillData]) as [strSkillData]",
		},
	},
}
