package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType2DatabaseNbr = "GAME"
	_MagicType2TableName   = "MAGIC_TYPE2"
)

func init() {
	ModelList = append(ModelList, &MagicType2{})
}

// MagicType2 Supports bow abilities
type MagicType2 struct {
	ID            int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name          *mssql.VarChar `gorm:"column:Name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Name,omitempty"`
	Description   *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	HitType       uint8          `gorm:"column:HitType;type:tinyint;not null" json:"HitType"`
	HitRateMod    int16          `gorm:"column:HitRate;type:smallint;not null" json:"HitRate"`
	DamageMod     int16          `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	RangeMod      int16          `gorm:"column:AddRange;type:smallint;not null" json:"AddRange"`
	NeedArrow     uint8          `gorm:"column:NeedArrow;type:tinyint;not null" json:"NeedArrow"`
	AddDamagePlus int16          `gorm:"column:AddDamagePlus;type:smallint;not null" json:"AddDamagePlus"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType2) GetDatabaseName() string {
	return GetDatabaseName(_MagicType2DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType2) TableName() string {
	return _MagicType2TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType2) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE2] ([iNum], [Name], [Description], [HitType], [HitRate], [AddDamage], [AddRange], [NeedArrow], [AddDamagePlus]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ID),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.HitType),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.RangeMod),
		GetOptionalDecVal(&this.NeedArrow),
		GetOptionalDecVal(&this.AddDamagePlus))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType2) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE2] ([iNum], [Name], [Description], [HitType], [HitRate], [AddDamage], [AddRange], [NeedArrow], [AddDamagePlus]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType2) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ID),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.HitType),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.RangeMod),
		GetOptionalDecVal(&this.NeedArrow),
		GetOptionalDecVal(&this.AddDamagePlus))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType2) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE2] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[HitType] tinyint NOT NULL,\n\t[HitRate] smallint NOT NULL,\n\t[AddDamage] smallint NOT NULL,\n\t[AddRange] smallint NOT NULL,\n\t[NeedArrow] tinyint NOT NULL,\n\t[AddDamagePlus] smallint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE2] PRIMARY KEY CLUSTERED ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType2) SelectClause() (selectClause clause.Select) {
	return _MagicType2SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType2) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType2{}
	rawSql := "SELECT [iNum], [Name], [Description], [HitType], [HitRate], [AddDamage], [AddRange], [NeedArrow], [AddDamagePlus] FROM [MAGIC_TYPE2]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType2SelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[iNum]",
		},
		clause.Column{
			Name: "[Name]",
		},
		clause.Column{
			Name: "[Description]",
		},
		clause.Column{
			Name: "[HitType]",
		},
		clause.Column{
			Name: "[HitRate]",
		},
		clause.Column{
			Name: "[AddDamage]",
		},
		clause.Column{
			Name: "[AddRange]",
		},
		clause.Column{
			Name: "[NeedArrow]",
		},
		clause.Column{
			Name: "[AddDamagePlus]",
		},
	},
}
