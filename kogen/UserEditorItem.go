package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_UserEditorItemDatabaseNbr = "GAME"
	_UserEditorItemTableName   = "USER_EDITOR_ITEM"
)

func init() {
	ModelList = append(ModelList, &UserEditorItem{})
}

// UserEditorItem User editor item
type UserEditorItem struct {
	CharId     mssql.VarChar `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	AccountId  mssql.VarChar `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	OpId       mssql.VarChar `gorm:"column:strOpID;type:varchar(21);not null" json:"strOpID"`
	OpIP       mssql.VarChar `gorm:"column:strOpIP;type:varchar(21);not null" json:"strOpIP"`
	DbIndex    int16         `gorm:"column:sDBIndex;type:smallint;not null" json:"sDBIndex"`
	Pos        int16         `gorm:"column:sPos;type:smallint;not null" json:"sPos"`
	Type       uint8         `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	ItemId1    int           `gorm:"column:nItemID1;type:int;not null" json:"nItemID1"`
	ItemId2    int           `gorm:"column:nItemID2;type:int;not null" json:"nItemID2"`
	UpdateTime *time.Time    `gorm:"column:UpdateTime;type:smalldatetime" json:"UpdateTime,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this UserEditorItem) GetDatabaseName() string {
	return GetDatabaseName(_UserEditorItemDatabaseNbr)
}

// TableName Returns the table name
func (this UserEditorItem) TableName() string {
	return _UserEditorItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserEditorItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_EDITOR_ITEM] ([strCharID], [strAccountID], [strOpID], [strOpIP], [sDBIndex], [sPos], [byType], [nItemID1], [nItemID2], [UpdateTime]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.OpId, false),
		GetOptionalVarCharVal(&this.OpIP, false),
		GetOptionalDecVal(&this.DbIndex),
		GetOptionalDecVal(&this.Pos),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ItemId1),
		GetOptionalDecVal(&this.ItemId2),
		GetDateTimeExportFmt(this.UpdateTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserEditorItem) GetInsertHeader() string {
	return "INSERT INTO [USER_EDITOR_ITEM] ([strCharID], [strAccountID], [strOpID], [strOpIP], [sDBIndex], [sPos], [byType], [nItemID1], [nItemID2], [UpdateTime]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserEditorItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.OpId, false),
		GetOptionalVarCharVal(&this.OpIP, false),
		GetOptionalDecVal(&this.DbIndex),
		GetOptionalDecVal(&this.Pos),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ItemId1),
		GetOptionalDecVal(&this.ItemId2),
		GetDateTimeExportFmt(this.UpdateTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserEditorItem) GetCreateTableString() string {
	query := "CREATE TABLE [USER_EDITOR_ITEM] (\n\t[strCharID] varchar(21) NOT NULL,\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strOpID] varchar(21) NOT NULL,\n\t[strOpIP] varchar(21) NOT NULL,\n\t[sDBIndex] smallint NOT NULL,\n\t[sPos] smallint NOT NULL,\n\t[byType] tinyint NOT NULL,\n\t[nItemID1] int NOT NULL,\n\t[nItemID2] int NOT NULL,\n\t[UpdateTime] smalldatetime\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserEditorItem) SelectClause() (selectClause clause.Select) {
	return _UserEditorItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserEditorItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserEditorItem{}
	rawSql := "SELECT [strCharID], [strAccountID], [strOpID], [strOpIP], [sDBIndex], [sPos], [byType], [nItemID1], [nItemID2], [UpdateTime] FROM [USER_EDITOR_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserEditorItemSelectClause = clause.Select{
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
			Name: "[sDBIndex]",
		},
		clause.Column{
			Name: "[sPos]",
		},
		clause.Column{
			Name: "[byType]",
		},
		clause.Column{
			Name: "[nItemID1]",
		},
		clause.Column{
			Name: "[nItemID2]",
		},
		clause.Column{
			Name: "[UpdateTime]",
		},
	},
}
