package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType3DatabaseNbr = 1
	_MagicType3TableName   = "MAGIC_TYPE3"
)

func init() {
	ModelList = append(ModelList, &MagicType3{})
}

// MagicType3 Type 3 supports Area of Effect and Damage over Time effects
type MagicType3 struct {
	MagicNumber int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name        *mssql.VarChar `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description *mssql.VarChar `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Radius      uint8          `gorm:"column:Radius;type:tinyint;not null" json:"Radius"`
	Angle       int16          `gorm:"column:Angle;type:smallint;not null" json:"Angle"`
	DirectType  uint8          `gorm:"column:DirectType;type:tinyint;not null" json:"DirectType"`
	FirstDamage int16          `gorm:"column:FirstDamage;type:smallint;not null" json:"FirstDamage"`
	EndDamage   int16          `gorm:"column:EndDamage;type:smallint;not null" json:"EndDamage"`
	TimeDamage  int16          `gorm:"column:TimeDamage;type:smallint;not null" json:"TimeDamage"`
	Duration    uint8          `gorm:"column:Duration;type:tinyint;not null" json:"Duration"`
	Attribute   uint8          `gorm:"column:Attribute;type:tinyint;not null" json:"Attribute"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType3) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType3DatabaseNbr))
}

// TableName Returns the table name
func (this MagicType3) TableName() string {
	return _MagicType3TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType3) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE3] ([iNum], [Name], [Description], [Radius], [Angle], [DirectType], [FirstDamage], [EndDamage], [TimeDamage], [Duration], [Attribute]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.Angle),
		GetOptionalDecVal(&this.DirectType),
		GetOptionalDecVal(&this.FirstDamage),
		GetOptionalDecVal(&this.EndDamage),
		GetOptionalDecVal(&this.TimeDamage),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Attribute))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType3) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE3] (iNum, Name, Description, Radius, Angle, DirectType, FirstDamage, EndDamage, TimeDamage, Duration, Attribute) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType3) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.Angle),
		GetOptionalDecVal(&this.DirectType),
		GetOptionalDecVal(&this.FirstDamage),
		GetOptionalDecVal(&this.EndDamage),
		GetOptionalDecVal(&this.TimeDamage),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.Attribute))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType3) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE3] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[Radius] tinyint NOT NULL,\n\t[Angle] smallint NOT NULL,\n\t[DirectType] tinyint NOT NULL,\n\t[FirstDamage] smallint NOT NULL,\n\t[EndDamage] smallint NOT NULL,\n\t[TimeDamage] smallint NOT NULL,\n\t[Duration] tinyint NOT NULL,\n\t[Attribute] tinyint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE3] PRIMARY KEY ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType3) SelectClause() (selectClause clause.Select) {
	return _MagicType3SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType3) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType3{}
	rawSql := "SELECT [iNum], [Name], [Description], [Radius], [Angle], [DirectType], [FirstDamage], [EndDamage], [TimeDamage], [Duration], [Attribute] FROM [MAGIC_TYPE3]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType3SelectClause = clause.Select{
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
			Name: "[Radius]",
		},
		clause.Column{
			Name: "[Angle]",
		},
		clause.Column{
			Name: "[DirectType]",
		},
		clause.Column{
			Name: "[FirstDamage]",
		},
		clause.Column{
			Name: "[EndDamage]",
		},
		clause.Column{
			Name: "[TimeDamage]",
		},
		clause.Column{
			Name: "[Duration]",
		},
		clause.Column{
			Name: "[Attribute]",
		},
	},
}
