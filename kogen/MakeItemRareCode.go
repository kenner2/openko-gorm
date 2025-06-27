package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MakeItemRareCodeDatabaseNbr = "GAME"
	_MakeItemRareCodeTableName   = "MAKE_ITEM_LARECODE"
)

func init() {
	ModelList = append(ModelList, &MakeItemRareCode{})
}

// MakeItemRareCode Make item rarity codes
type MakeItemRareCode struct {
	LevelGrade  uint8 `gorm:"column:byLevelGrade;type:tinyint;not null" json:"byLevelGrade"`
	UpgradeItem int16 `gorm:"column:sUpgradeItem;type:smallint;not null;default:0" json:"sUpgradeItem"`
	RareItem    int16 `gorm:"column:sLareItem;type:smallint;not null;default:0" json:"sLareItem"`
	MagicItem   int16 `gorm:"column:sMagicItem;type:smallint;not null;default:0" json:"sMagicItem"`
	GeneralItem int16 `gorm:"column:sGereralItem;type:smallint;not null;default:0" json:"sGereralItem"`
}

// GetDatabaseName Returns the table's database name
func (this MakeItemRareCode) GetDatabaseName() string {
	return GetDatabaseName(_MakeItemRareCodeDatabaseNbr)
}

// TableName Returns the table name
func (this MakeItemRareCode) TableName() string {
	return _MakeItemRareCodeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MakeItemRareCode) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM_LARECODE] ([byLevelGrade], [sUpgradeItem], [sLareItem], [sMagicItem], [sGereralItem]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.LevelGrade),
		GetOptionalDecVal(&this.UpgradeItem),
		GetOptionalDecVal(&this.RareItem),
		GetOptionalDecVal(&this.MagicItem),
		GetOptionalDecVal(&this.GeneralItem))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MakeItemRareCode) GetInsertHeader() string {
	return "INSERT INTO [MAKE_ITEM_LARECODE] ([byLevelGrade], [sUpgradeItem], [sLareItem], [sMagicItem], [sGereralItem]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MakeItemRareCode) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.LevelGrade),
		GetOptionalDecVal(&this.UpgradeItem),
		GetOptionalDecVal(&this.RareItem),
		GetOptionalDecVal(&this.MagicItem),
		GetOptionalDecVal(&this.GeneralItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this MakeItemRareCode) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM_LARECODE] (\n\t[byLevelGrade] tinyint NOT NULL,\n\t[sUpgradeItem] smallint NOT NULL,\n\t[sLareItem] smallint NOT NULL,\n\t[sMagicItem] smallint NOT NULL,\n\t[sGereralItem] smallint NOT NULL\n)\nGO\nALTER TABLE [MAKE_ITEM_LARECODE] ADD CONSTRAINT [DF_MAKE_ITEM_LARECODE_sUpgradeItem] DEFAULT 0 FOR [sUpgradeItem]\nGO\nALTER TABLE [MAKE_ITEM_LARECODE] ADD CONSTRAINT [DF_MAKE_ITEM_LARECODE_sLareItem] DEFAULT 0 FOR [sLareItem]\nGO\nALTER TABLE [MAKE_ITEM_LARECODE] ADD CONSTRAINT [DF_MAKE_ITEM_LARECODE_sMagicItem] DEFAULT 0 FOR [sMagicItem]\nGO\nALTER TABLE [MAKE_ITEM_LARECODE] ADD CONSTRAINT [DF_MAKE_ITEM_LARECODE_sGereralItem] DEFAULT 0 FOR [sGereralItem]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MakeItemRareCode) SelectClause() (selectClause clause.Select) {
	return _MakeItemRareCodeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MakeItemRareCode) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MakeItemRareCode{}
	rawSql := "SELECT [byLevelGrade], [sUpgradeItem], [sLareItem], [sMagicItem], [sGereralItem] FROM [MAKE_ITEM_LARECODE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MakeItemRareCodeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[byLevelGrade]",
		},
		clause.Column{
			Name: "[sUpgradeItem]",
		},
		clause.Column{
			Name: "[sLareItem]",
		},
		clause.Column{
			Name: "[sMagicItem]",
		},
		clause.Column{
			Name: "[sGereralItem]",
		},
	},
}
