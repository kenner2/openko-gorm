package kogen

import (
	"fmt"
)

const (
	_UserDataSkillShortcutDatabaseNbr = 0
	_UserDataSkillShortcutTableName   = "USERDATA_SKILLSHORTCUT"
)

func init() {
	ModelList = append(ModelList, &UserDataSkillShortcut{})
}

// UserDataSkillShortcut User data skill shortcut
type UserDataSkillShortcut struct {
	CharId    [21]byte  `gorm:"column:strCharID;type:varchar(21);primaryKey;not null;default:''" json:"strCharID"`
	Count     int16     `gorm:"column:nCount;type:smallint;not null;default:0" json:"nCount"`
	SkillData [260]byte `gorm:"column:strSkillData;type:varchar(260);not null;default:0x00" json:"strSkillData"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserDataSkillShortcut) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserDataSkillShortcutDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserDataSkillShortcut) GetTableName() string {
	return _UserDataSkillShortcutTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserDataSkillShortcut) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USERDATA_SKILLSHORTCUT] (strCharID, nCount, strSkillData) \nVALUES (%s, %s, %s)", GetOptionalBinaryVal(this.CharId),
		GetOptionalDecVal(&this.Count),
		GetOptionalBinaryVal(this.SkillData))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserDataSkillShortcut) GetCreateTableString() string {
	query := "CREATE TABLE [USERDATA_SKILLSHORTCUT] (\n\t[strCharID] varchar(21) NOT NULL,\n\t[nCount] smallint NOT NULL,\n\t[strSkillData] varchar(260) NOT NULL\n\tPRIMARY KEY (\"strCharID\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
