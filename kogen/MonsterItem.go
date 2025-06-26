package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MonsterItemDatabaseNbr = "GAME"
	_MonsterItemTableName   = "K_MONSTER_ITEM"
)

func init() {
	ModelList = append(ModelList, &MonsterItem{})
}

// MonsterItem Monster loot table
type MonsterItem struct {
	MonsterId   int16  `gorm:"column:sIndex;type:smallint;primaryKey;not null" json:"sIndex"`
	ItemId1     *int   `gorm:"column:iItem01;type:int" json:"iItem01,omitempty"`
	DropChance1 *int16 `gorm:"column:sPersent01;type:smallint" json:"sPersent01,omitempty"`
	ItemId2     *int   `gorm:"column:iItem02;type:int" json:"iItem02,omitempty"`
	DropChance2 *int16 `gorm:"column:sPersent02;type:smallint" json:"sPersent02,omitempty"`
	ItemId3     *int   `gorm:"column:iItem03;type:int" json:"iItem03,omitempty"`
	DropChance3 *int16 `gorm:"column:sPersent03;type:smallint" json:"sPersent03,omitempty"`
	ItemId4     *int   `gorm:"column:iItem04;type:int" json:"iItem04,omitempty"`
	DropChance4 *int16 `gorm:"column:sPersent04;type:smallint" json:"sPersent04,omitempty"`
	ItemId5     *int   `gorm:"column:iItem05;type:int" json:"iItem05,omitempty"`
	DropChance5 *int16 `gorm:"column:sPersent05;type:smallint" json:"sPersent05,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this MonsterItem) GetDatabaseName() string {
	return GetDatabaseName(_MonsterItemDatabaseNbr)
}

// TableName Returns the table name
func (this MonsterItem) TableName() string {
	return _MonsterItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MonsterItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [K_MONSTER_ITEM] ([sIndex], [iItem01], [sPersent01], [iItem02], [sPersent02], [iItem03], [sPersent03], [iItem04], [sPersent04], [iItem05], [sPersent05]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MonsterId),
		GetOptionalDecVal(this.ItemId1),
		GetOptionalDecVal(this.DropChance1),
		GetOptionalDecVal(this.ItemId2),
		GetOptionalDecVal(this.DropChance2),
		GetOptionalDecVal(this.ItemId3),
		GetOptionalDecVal(this.DropChance3),
		GetOptionalDecVal(this.ItemId4),
		GetOptionalDecVal(this.DropChance4),
		GetOptionalDecVal(this.ItemId5),
		GetOptionalDecVal(this.DropChance5))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MonsterItem) GetInsertHeader() string {
	return "INSERT INTO [K_MONSTER_ITEM] ([sIndex], [iItem01], [sPersent01], [iItem02], [sPersent02], [iItem03], [sPersent03], [iItem04], [sPersent04], [iItem05], [sPersent05]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MonsterItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MonsterId),
		GetOptionalDecVal(this.ItemId1),
		GetOptionalDecVal(this.DropChance1),
		GetOptionalDecVal(this.ItemId2),
		GetOptionalDecVal(this.DropChance2),
		GetOptionalDecVal(this.ItemId3),
		GetOptionalDecVal(this.DropChance3),
		GetOptionalDecVal(this.ItemId4),
		GetOptionalDecVal(this.DropChance4),
		GetOptionalDecVal(this.ItemId5),
		GetOptionalDecVal(this.DropChance5))
}

// GetCreateTableString Returns the create table statement for this object
func (this MonsterItem) GetCreateTableString() string {
	query := "CREATE TABLE [K_MONSTER_ITEM] (\n\t[sIndex] smallint NOT NULL,\n\t[iItem01] int,\n\t[sPersent01] smallint,\n\t[iItem02] int,\n\t[sPersent02] smallint,\n\t[iItem03] int,\n\t[sPersent03] smallint,\n\t[iItem04] int,\n\t[sPersent04] smallint,\n\t[iItem05] int,\n\t[sPersent05] smallint\n\tCONSTRAINT [PK_K_MONSTER_ITEM] PRIMARY KEY ([sIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MonsterItem) SelectClause() (selectClause clause.Select) {
	return _MonsterItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MonsterItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MonsterItem{}
	rawSql := "SELECT [sIndex], [iItem01], [sPersent01], [iItem02], [sPersent02], [iItem03], [sPersent03], [iItem04], [sPersent04], [iItem05], [sPersent05] FROM [K_MONSTER_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MonsterItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIndex]",
		},
		clause.Column{
			Name: "[iItem01]",
		},
		clause.Column{
			Name: "[sPersent01]",
		},
		clause.Column{
			Name: "[iItem02]",
		},
		clause.Column{
			Name: "[sPersent02]",
		},
		clause.Column{
			Name: "[iItem03]",
		},
		clause.Column{
			Name: "[sPersent03]",
		},
		clause.Column{
			Name: "[iItem04]",
		},
		clause.Column{
			Name: "[sPersent04]",
		},
		clause.Column{
			Name: "[iItem05]",
		},
		clause.Column{
			Name: "[sPersent05]",
		},
	},
}
