package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicDatabaseNbr = "GAME"
	_MagicTableName   = "MAGIC"
)

func init() {
	ModelList = append(ModelList, &Magic{})
}

// Magic Contains the configuration for magic and abilities
type Magic struct {
	ID           int            `gorm:"column:MagicNum;type:int;primaryKey;not null" json:"MagicNum"`
	EnglishName  *mssql.VarChar `gorm:"column:EnName;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"EnName,omitempty"`
	KoreanName   *mssql.VarChar `gorm:"column:KrName;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"KrName,omitempty"`
	Description  *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	BeforeAction uint8          `gorm:"column:BeforeAction;type:tinyint;not null" json:"BeforeAction"`
	TargetAction uint8          `gorm:"column:TargetAction;type:tinyint;not null" json:"TargetAction"`
	SelfEffect   uint8          `gorm:"column:SelfEffect;type:tinyint;not null" json:"SelfEffect"`
	FlyingEffect uint8          `gorm:"column:FlyingEffect;type:tinyint;not null" json:"FlyingEffect"`
	TargetEffect int16          `gorm:"column:TargetEffect;type:smallint;not null" json:"TargetEffect"`
	Moral        uint8          `gorm:"column:Moral;type:tinyint;not null" json:"Moral"`
	SkillLevel   int16          `gorm:"column:SkillLevel;type:smallint;not null" json:"SkillLevel"`
	Skill        int16          `gorm:"column:Skill;type:smallint;not null" json:"Skill"`
	ManaCost     int16          `gorm:"column:Msp;type:smallint;not null" json:"Msp"`
	HpCost       int16          `gorm:"column:HP;type:smallint;not null" json:"HP"`
	ItemGroup    uint8          `gorm:"column:ItemGroup;type:tinyint;not null" json:"ItemGroup"`
	UseItem      int            `gorm:"column:UseItem;type:int;not null" json:"UseItem"`
	CastTime     uint8          `gorm:"column:CastTime;type:tinyint;not null" json:"CastTime"`
	RecastTime   uint8          `gorm:"column:ReCastTime;type:tinyint;not null" json:"ReCastTime"`
	SuccessRate  uint8          `gorm:"column:SuccessRate;type:tinyint;not null" json:"SuccessRate"`
	Type1        uint8          `gorm:"column:Type1;type:tinyint;not null" json:"Type1"`
	Type2        uint8          `gorm:"column:Type2;type:tinyint;not null" json:"Type2"`
	Range        int16          `gorm:"column:Range;type:smallint;not null" json:"Range"`
	Etc          uint8          `gorm:"column:Etc;type:tinyint;not null" json:"Etc"`
	Event        int            `gorm:"column:Event;type:int;not null" json:"Event"`
}

// GetDatabaseName Returns the table's database name
func (this Magic) GetDatabaseName() string {
	return GetDatabaseName(_MagicDatabaseNbr)
}

// TableName Returns the table name
func (this Magic) TableName() string {
	return _MagicTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Magic) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC] ([MagicNum], [EnName], [KrName], [Description], [BeforeAction], [TargetAction], [SelfEffect], [FlyingEffect], [TargetEffect], [Moral], [SkillLevel], [Skill], [Msp], [HP], [ItemGroup], [UseItem], [CastTime], [ReCastTime], [SuccessRate], [Type1], [Type2], [Range], [Etc], [Event]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ID),
		GetOptionalVarCharVal(this.EnglishName, false),
		GetOptionalVarCharVal(this.KoreanName, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.BeforeAction),
		GetOptionalDecVal(&this.TargetAction),
		GetOptionalDecVal(&this.SelfEffect),
		GetOptionalDecVal(&this.FlyingEffect),
		GetOptionalDecVal(&this.TargetEffect),
		GetOptionalDecVal(&this.Moral),
		GetOptionalDecVal(&this.SkillLevel),
		GetOptionalDecVal(&this.Skill),
		GetOptionalDecVal(&this.ManaCost),
		GetOptionalDecVal(&this.HpCost),
		GetOptionalDecVal(&this.ItemGroup),
		GetOptionalDecVal(&this.UseItem),
		GetOptionalDecVal(&this.CastTime),
		GetOptionalDecVal(&this.RecastTime),
		GetOptionalDecVal(&this.SuccessRate),
		GetOptionalDecVal(&this.Type1),
		GetOptionalDecVal(&this.Type2),
		GetOptionalDecVal(&this.Range),
		GetOptionalDecVal(&this.Etc),
		GetOptionalDecVal(&this.Event))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Magic) GetInsertHeader() string {
	return "INSERT INTO [MAGIC] ([MagicNum], [EnName], [KrName], [Description], [BeforeAction], [TargetAction], [SelfEffect], [FlyingEffect], [TargetEffect], [Moral], [SkillLevel], [Skill], [Msp], [HP], [ItemGroup], [UseItem], [CastTime], [ReCastTime], [SuccessRate], [Type1], [Type2], [Range], [Etc], [Event]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Magic) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ID),
		GetOptionalVarCharVal(this.EnglishName, false),
		GetOptionalVarCharVal(this.KoreanName, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.BeforeAction),
		GetOptionalDecVal(&this.TargetAction),
		GetOptionalDecVal(&this.SelfEffect),
		GetOptionalDecVal(&this.FlyingEffect),
		GetOptionalDecVal(&this.TargetEffect),
		GetOptionalDecVal(&this.Moral),
		GetOptionalDecVal(&this.SkillLevel),
		GetOptionalDecVal(&this.Skill),
		GetOptionalDecVal(&this.ManaCost),
		GetOptionalDecVal(&this.HpCost),
		GetOptionalDecVal(&this.ItemGroup),
		GetOptionalDecVal(&this.UseItem),
		GetOptionalDecVal(&this.CastTime),
		GetOptionalDecVal(&this.RecastTime),
		GetOptionalDecVal(&this.SuccessRate),
		GetOptionalDecVal(&this.Type1),
		GetOptionalDecVal(&this.Type2),
		GetOptionalDecVal(&this.Range),
		GetOptionalDecVal(&this.Etc),
		GetOptionalDecVal(&this.Event))
}

// GetCreateTableString Returns the create table statement for this object
func (this Magic) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC] (\n\t[MagicNum] int NOT NULL,\n\t[EnName] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[KrName] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[BeforeAction] tinyint NOT NULL,\n\t[TargetAction] tinyint NOT NULL,\n\t[SelfEffect] tinyint NOT NULL,\n\t[FlyingEffect] tinyint NOT NULL,\n\t[TargetEffect] smallint NOT NULL,\n\t[Moral] tinyint NOT NULL,\n\t[SkillLevel] smallint NOT NULL,\n\t[Skill] smallint NOT NULL,\n\t[Msp] smallint NOT NULL,\n\t[HP] smallint NOT NULL,\n\t[ItemGroup] tinyint NOT NULL,\n\t[UseItem] int NOT NULL,\n\t[CastTime] tinyint NOT NULL,\n\t[ReCastTime] tinyint NOT NULL,\n\t[SuccessRate] tinyint NOT NULL,\n\t[Type1] tinyint NOT NULL,\n\t[Type2] tinyint NOT NULL,\n\t[Range] smallint NOT NULL,\n\t[Etc] tinyint NOT NULL,\n\t[Event] int NOT NULL\n\tCONSTRAINT [PK_MAGIC] PRIMARY KEY CLUSTERED ([MagicNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Magic) SelectClause() (selectClause clause.Select) {
	return _MagicSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Magic) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Magic{}
	rawSql := "SELECT [MagicNum], [EnName], [KrName], [Description], [BeforeAction], [TargetAction], [SelfEffect], [FlyingEffect], [TargetEffect], [Moral], [SkillLevel], [Skill], [Msp], [HP], [ItemGroup], [UseItem], [CastTime], [ReCastTime], [SuccessRate], [Type1], [Type2], [Range], [Etc], [Event] FROM [MAGIC]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[MagicNum]",
		},
		clause.Column{
			Name: "[EnName]",
		},
		clause.Column{
			Name: "[KrName]",
		},
		clause.Column{
			Name: "[Description]",
		},
		clause.Column{
			Name: "[BeforeAction]",
		},
		clause.Column{
			Name: "[TargetAction]",
		},
		clause.Column{
			Name: "[SelfEffect]",
		},
		clause.Column{
			Name: "[FlyingEffect]",
		},
		clause.Column{
			Name: "[TargetEffect]",
		},
		clause.Column{
			Name: "[Moral]",
		},
		clause.Column{
			Name: "[SkillLevel]",
		},
		clause.Column{
			Name: "[Skill]",
		},
		clause.Column{
			Name: "[Msp]",
		},
		clause.Column{
			Name: "[HP]",
		},
		clause.Column{
			Name: "[ItemGroup]",
		},
		clause.Column{
			Name: "[UseItem]",
		},
		clause.Column{
			Name: "[CastTime]",
		},
		clause.Column{
			Name: "[ReCastTime]",
		},
		clause.Column{
			Name: "[SuccessRate]",
		},
		clause.Column{
			Name: "[Type1]",
		},
		clause.Column{
			Name: "[Type2]",
		},
		clause.Column{
			Name: "[Range]",
		},
		clause.Column{
			Name: "[Etc]",
		},
		clause.Column{
			Name: "[Event]",
		},
	},
}
