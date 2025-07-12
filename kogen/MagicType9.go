package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType9DatabaseNbr = "GAME"
	_MagicType9TableName   = "MAGIC_TYPE9"
)

func init() {
	ModelList = append(ModelList, &MagicType9{})
}

// MagicType9 Supports stealth and detection abilities
type MagicType9 struct {
	MagicNumber   int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name          *mssql.VarChar `gorm:"column:Name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Name,omitempty"`
	Description   *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	ValidGroup    uint8          `gorm:"column:ValidGroup;type:tinyint;not null" json:"ValidGroup"`
	NationChange  uint8          `gorm:"column:NationChange;type:tinyint;not null" json:"NationChange"`
	MonsterNumber int16          `gorm:"column:MonsterNum;type:smallint;not null" json:"MonsterNum"`
	TargetChange  uint8          `gorm:"column:TargetChange;type:tinyint;not null" json:"TargetChange"`
	StateChange   uint8          `gorm:"column:StateChange;type:tinyint;not null" json:"StateChange"`
	Radius        int16          `gorm:"column:Radius;type:smallint;not null" json:"Radius"`
	HitRate       int16          `gorm:"column:Hitrate;type:smallint;not null" json:"Hitrate"`
	Duration      int16          `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	AddDamage     int16          `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	Vision        int16          `gorm:"column:Vision;type:smallint;not null" json:"Vision"`
	NeedItem      int            `gorm:"column:NeedItem;type:int;not null" json:"NeedItem"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType9) GetDatabaseName() string {
	return GetDatabaseName(_MagicType9DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType9) TableName() string {
	return _MagicType9TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType9) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE9] ([iNum], [Name], [Description], [ValidGroup], [NationChange], [MonsterNum], [TargetChange], [StateChange], [Radius], [Hitrate], [Duration], [AddDamage], [Vision], [NeedItem]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType9) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE9] ([iNum], [Name], [Description], [ValidGroup], [NationChange], [MonsterNum], [TargetChange], [StateChange], [Radius], [Hitrate], [Duration], [AddDamage], [Vision], [NeedItem]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType9) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType9) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE9] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[ValidGroup] tinyint NOT NULL,\n\t[NationChange] tinyint NOT NULL,\n\t[MonsterNum] smallint NOT NULL,\n\t[TargetChange] tinyint NOT NULL,\n\t[StateChange] tinyint NOT NULL,\n\t[Radius] smallint NOT NULL,\n\t[Hitrate] smallint NOT NULL,\n\t[Duration] smallint NOT NULL,\n\t[AddDamage] smallint NOT NULL,\n\t[Vision] smallint NOT NULL,\n\t[NeedItem] int NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE9] PRIMARY KEY CLUSTERED ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType9) SelectClause() (selectClause clause.Select) {
	return _MagicType9SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType9) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType9{}
	rawSql := "SELECT [iNum], [Name], [Description], [ValidGroup], [NationChange], [MonsterNum], [TargetChange], [StateChange], [Radius], [Hitrate], [Duration], [AddDamage], [Vision], [NeedItem] FROM [MAGIC_TYPE9]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType9SelectClause = clause.Select{
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
			Name: "[ValidGroup]",
		},
		clause.Column{
			Name: "[NationChange]",
		},
		clause.Column{
			Name: "[MonsterNum]",
		},
		clause.Column{
			Name: "[TargetChange]",
		},
		clause.Column{
			Name: "[StateChange]",
		},
		clause.Column{
			Name: "[Radius]",
		},
		clause.Column{
			Name: "[Hitrate]",
		},
		clause.Column{
			Name: "[Duration]",
		},
		clause.Column{
			Name: "[AddDamage]",
		},
		clause.Column{
			Name: "[Vision]",
		},
		clause.Column{
			Name: "[NeedItem]",
		},
	},
}
