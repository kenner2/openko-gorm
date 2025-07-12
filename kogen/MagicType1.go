package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType1DatabaseNbr = "GAME"
	_MagicType1TableName   = "MAGIC_TYPE1"
)

func init() {
	ModelList = append(ModelList, &MagicType1{})
}

// MagicType1 Supports melee abilities
type MagicType1 struct {
	MagicNumber int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name        *mssql.VarChar `gorm:"column:Name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Name,omitempty"`
	Description *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	Type        uint8          `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	HitRateMod  int16          `gorm:"column:HitRate;type:smallint;not null" json:"HitRate"`
	DamageMod   int16          `gorm:"column:Hit;type:smallint;not null" json:"Hit"`
	AddDamage   int16          `gorm:"column:AddDamage;type:smallint;not null" json:"AddDamage"`
	Delay       uint8          `gorm:"column:Delay;type:tinyint;not null" json:"Delay"`
	ComboType   uint8          `gorm:"column:ComboType;type:tinyint;not null" json:"ComboType"`
	ComboCount  uint8          `gorm:"column:ComboCount;type:tinyint;not null" json:"ComboCount"`
	ComboDamage int16          `gorm:"column:ComboDamage;type:smallint;not null" json:"ComboDamage"`
	Range       int16          `gorm:"column:Range;type:smallint;not null" json:"Range"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType1) GetDatabaseName() string {
	return GetDatabaseName(_MagicType1DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType1) TableName() string {
	return _MagicType1TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType1) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE1] ([iNum], [Name], [Description], [Type], [HitRate], [Hit], [AddDamage], [Delay], [ComboType], [ComboCount], [ComboDamage], [Range]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Delay),
		GetOptionalDecVal(&this.ComboType),
		GetOptionalDecVal(&this.ComboCount),
		GetOptionalDecVal(&this.ComboDamage),
		GetOptionalDecVal(&this.Range))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType1) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE1] ([iNum], [Name], [Description], [Type], [HitRate], [Hit], [AddDamage], [Delay], [ComboType], [ComboCount], [ComboDamage], [Range]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType1) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.HitRateMod),
		GetOptionalDecVal(&this.DamageMod),
		GetOptionalDecVal(&this.AddDamage),
		GetOptionalDecVal(&this.Delay),
		GetOptionalDecVal(&this.ComboType),
		GetOptionalDecVal(&this.ComboCount),
		GetOptionalDecVal(&this.ComboDamage),
		GetOptionalDecVal(&this.Range))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType1) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE1] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Type] tinyint NOT NULL,\n\t[HitRate] smallint NOT NULL,\n\t[Hit] smallint NOT NULL,\n\t[AddDamage] smallint NOT NULL,\n\t[Delay] tinyint NOT NULL,\n\t[ComboType] tinyint NOT NULL,\n\t[ComboCount] tinyint NOT NULL,\n\t[ComboDamage] smallint NOT NULL,\n\t[Range] smallint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE1] PRIMARY KEY CLUSTERED ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType1) SelectClause() (selectClause clause.Select) {
	return _MagicType1SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType1) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType1{}
	rawSql := "SELECT [iNum], [Name], [Description], [Type], [HitRate], [Hit], [AddDamage], [Delay], [ComboType], [ComboCount], [ComboDamage], [Range] FROM [MAGIC_TYPE1]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType1SelectClause = clause.Select{
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
			Name: "[Type]",
		},
		clause.Column{
			Name: "[HitRate]",
		},
		clause.Column{
			Name: "[Hit]",
		},
		clause.Column{
			Name: "[AddDamage]",
		},
		clause.Column{
			Name: "[Delay]",
		},
		clause.Column{
			Name: "[ComboType]",
		},
		clause.Column{
			Name: "[ComboCount]",
		},
		clause.Column{
			Name: "[ComboDamage]",
		},
		clause.Column{
			Name: "[Range]",
		},
	},
}
