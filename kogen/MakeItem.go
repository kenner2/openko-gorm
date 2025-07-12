package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MakeItemDatabaseNbr = "GAME"
	_MakeItemTableName   = "MAKE_ITEM"
)

func init() {
	ModelList = append(ModelList, &MakeItem{})
}

// MakeItem Make item
type MakeItem struct {
	Index     int16          `gorm:"column:sIndex;type:smallint;primaryKey;not null" json:"sIndex"`
	ItemInfo  *mssql.VarChar `gorm:"column:strItemInfo;type:varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strItemInfo,omitempty"`
	ItemCode  int            `gorm:"column:iItemCode;type:int;not null" json:"iItemCode"`
	ItemLevel uint8          `gorm:"column:byItemLevel;type:tinyint;not null;default:0" json:"byItemLevel"`
}

// GetDatabaseName Returns the table's database name
func (this MakeItem) GetDatabaseName() string {
	return GetDatabaseName(_MakeItemDatabaseNbr)
}

// TableName Returns the table name
func (this MakeItem) TableName() string {
	return _MakeItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MakeItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM] ([sIndex], [strItemInfo], [iItemCode], [byItemLevel]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.ItemInfo, false),
		GetOptionalDecVal(&this.ItemCode),
		GetOptionalDecVal(&this.ItemLevel))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MakeItem) GetInsertHeader() string {
	return "INSERT INTO [MAKE_ITEM] ([sIndex], [strItemInfo], [iItemCode], [byItemLevel]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MakeItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(this.ItemInfo, false),
		GetOptionalDecVal(&this.ItemCode),
		GetOptionalDecVal(&this.ItemLevel))
}

// GetCreateTableString Returns the create table statement for this object
func (this MakeItem) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM] (\n\t[sIndex] smallint NOT NULL,\n\t[strItemInfo] varchar(20) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[iItemCode] int NOT NULL,\n\t[byItemLevel] tinyint NOT NULL\n\tCONSTRAINT [PK_MAKE_ITEM] PRIMARY KEY CLUSTERED ([sIndex])\n)\nGO\nALTER TABLE [MAKE_ITEM] ADD CONSTRAINT [DF_MAKE_ITEM_byItemLevel] DEFAULT 0 FOR [byItemLevel]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MakeItem) SelectClause() (selectClause clause.Select) {
	return _MakeItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MakeItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MakeItem{}
	rawSql := "SELECT [sIndex], [strItemInfo], [iItemCode], [byItemLevel] FROM [MAKE_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MakeItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIndex]",
		},
		clause.Column{
			Name: "[strItemInfo]",
		},
		clause.Column{
			Name: "[iItemCode]",
		},
		clause.Column{
			Name: "[byItemLevel]",
		},
	},
}
