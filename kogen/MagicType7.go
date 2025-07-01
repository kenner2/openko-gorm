package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType7DatabaseNbr = "GAME"
	_MagicType7TableName   = "MAGIC_TYPE7"
)

func init() {
	ModelList = append(ModelList, &MagicType7{})
}

// MagicType7 Type 7 supports targeting modifications
type MagicType7 struct {
	MagicNumber   int            `gorm:"column:nIndex;type:int;primaryKey;not null" json:"nIndex"`
	Name          *mssql.VarChar `gorm:"column:strName;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strName,omitempty"`
	Note          *mssql.VarChar `gorm:"column:strNote;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strNote,omitempty"`
	ValidGroup    uint8          `gorm:"column:byValidGroup;type:tinyint;not null" json:"byValidGroup"`
	NationChange  uint8          `gorm:"column:byNatoinChange;type:tinyint;not null" json:"byNatoinChange"`
	MonsterNumber int16          `gorm:"column:shMonsterNum;type:smallint;not null" json:"shMonsterNum"`
	TargetChange  uint8          `gorm:"column:byTargetChange;type:tinyint;not null" json:"byTargetChange"`
	StateChange   uint8          `gorm:"column:byStateChange;type:tinyint;not null" json:"byStateChange"`
	Radius        uint8          `gorm:"column:byRadius;type:tinyint;not null" json:"byRadius"`
	HitRate       int16          `gorm:"column:shHitrate;type:smallint;not null" json:"shHitrate"`
	Duration      int16          `gorm:"column:shDuration;type:smallint;not null" json:"shDuration"`
	Damage        int16          `gorm:"column:shDamage;type:smallint;not null" json:"shDamage"`
	Vision        uint8          `gorm:"column:byVisoin;type:tinyint;not null" json:"byVisoin"`
	NeedItem      int            `gorm:"column:nNeedItem;type:int;not null" json:"nNeedItem"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType7) GetDatabaseName() string {
	return GetDatabaseName(_MagicType7DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType7) TableName() string {
	return _MagicType7TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType7) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE7] ([nIndex], [strName], [strNote], [byValidGroup], [byNatoinChange], [shMonsterNum], [byTargetChange], [byStateChange], [byRadius], [shHitrate], [shDuration], [shDamage], [byVisoin], [nNeedItem]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Note, false),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType7) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE7] ([nIndex], [strName], [strNote], [byValidGroup], [byNatoinChange], [shMonsterNum], [byTargetChange], [byStateChange], [byRadius], [shHitrate], [shDuration], [shDamage], [byVisoin], [nNeedItem]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType7) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Note, false),
		GetOptionalDecVal(&this.ValidGroup),
		GetOptionalDecVal(&this.NationChange),
		GetOptionalDecVal(&this.MonsterNumber),
		GetOptionalDecVal(&this.TargetChange),
		GetOptionalDecVal(&this.StateChange),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.Vision),
		GetOptionalDecVal(&this.NeedItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType7) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE7] (\n\t[nIndex] int NOT NULL,\n\t[strName] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[strNote] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[byValidGroup] tinyint NOT NULL,\n\t[byNatoinChange] tinyint NOT NULL,\n\t[shMonsterNum] smallint NOT NULL,\n\t[byTargetChange] tinyint NOT NULL,\n\t[byStateChange] tinyint NOT NULL,\n\t[byRadius] tinyint NOT NULL,\n\t[shHitrate] smallint NOT NULL,\n\t[shDuration] smallint NOT NULL,\n\t[shDamage] smallint NOT NULL,\n\t[byVisoin] tinyint NOT NULL,\n\t[nNeedItem] int NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE7] PRIMARY KEY CLUSTERED ([nIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType7) SelectClause() (selectClause clause.Select) {
	return _MagicType7SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType7) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType7{}
	rawSql := "SELECT [nIndex], [strName], [strNote], [byValidGroup], [byNatoinChange], [shMonsterNum], [byTargetChange], [byStateChange], [byRadius], [shHitrate], [shDuration], [shDamage], [byVisoin], [nNeedItem] FROM [MAGIC_TYPE7]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType7SelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nIndex]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[strNote]",
		},
		clause.Column{
			Name: "[byValidGroup]",
		},
		clause.Column{
			Name: "[byNatoinChange]",
		},
		clause.Column{
			Name: "[shMonsterNum]",
		},
		clause.Column{
			Name: "[byTargetChange]",
		},
		clause.Column{
			Name: "[byStateChange]",
		},
		clause.Column{
			Name: "[byRadius]",
		},
		clause.Column{
			Name: "[shHitrate]",
		},
		clause.Column{
			Name: "[shDuration]",
		},
		clause.Column{
			Name: "[shDamage]",
		},
		clause.Column{
			Name: "[byVisoin]",
		},
		clause.Column{
			Name: "[nNeedItem]",
		},
	},
}
