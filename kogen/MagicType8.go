package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType8DatabaseNbr = "GAME"
	_MagicType8TableName   = "MAGIC_TYPE8"
)

func init() {
	ModelList = append(ModelList, &MagicType8{})
}

// MagicType8 Supports warp magic
type MagicType8 struct {
	MagicNumber int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name        *mssql.VarChar `gorm:"column:Name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Name,omitempty"`
	Description *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	Target      uint8          `gorm:"column:Target;type:tinyint;not null" json:"Target"`
	Radius      int16          `gorm:"column:Radius;type:smallint;not null" json:"Radius"`
	WarpType    uint8          `gorm:"column:WarpType;type:tinyint;not null" json:"WarpType"`
	ExpRecover  int16          `gorm:"column:ExpRecover;type:smallint;not null" json:"ExpRecover"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType8) GetDatabaseName() string {
	return GetDatabaseName(_MagicType8DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType8) TableName() string {
	return _MagicType8TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType8) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE8] ([iNum], [Name], [Description], [Target], [Radius], [WarpType], [ExpRecover]) VALUES\n(%s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Target),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.WarpType),
		GetOptionalDecVal(&this.ExpRecover))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType8) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE8] ([iNum], [Name], [Description], [Target], [Radius], [WarpType], [ExpRecover]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType8) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Target),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.WarpType),
		GetOptionalDecVal(&this.ExpRecover))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType8) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE8] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Target] tinyint NOT NULL,\n\t[Radius] smallint NOT NULL,\n\t[WarpType] tinyint NOT NULL,\n\t[ExpRecover] smallint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE8] PRIMARY KEY CLUSTERED ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType8) SelectClause() (selectClause clause.Select) {
	return _MagicType8SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType8) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType8{}
	rawSql := "SELECT [iNum], [Name], [Description], [Target], [Radius], [WarpType], [ExpRecover] FROM [MAGIC_TYPE8]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType8SelectClause = clause.Select{
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
			Name: "[Target]",
		},
		clause.Column{
			Name: "[Radius]",
		},
		clause.Column{
			Name: "[WarpType]",
		},
		clause.Column{
			Name: "[ExpRecover]",
		},
	},
}
