package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType5DatabaseNbr = 0
	_MagicType5TableName   = "MAGIC_TYPE5"
)

func init() {
	ModelList = append(ModelList, &MagicType5{})
}

// MagicType5 Type 5 supports recovery skills
type MagicType5 struct {
	MagicNumber int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name        *mssql.VarChar `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description *mssql.VarChar `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Type        uint8          `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	ExpRecover  uint8          `gorm:"column:ExpRecover;type:tinyint;not null" json:"ExpRecover"`
	NeedStone   int16          `gorm:"column:NeedStone;type:smallint;not null" json:"NeedStone"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType5) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType5DatabaseNbr))
}

// TableName Returns the table name
func (this MagicType5) TableName() string {
	return _MagicType5TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType5) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE5] ([iNum], [Name], [Description], [Type], [ExpRecover], [NeedStone]) VALUES\n(%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ExpRecover),
		GetOptionalDecVal(&this.NeedStone))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType5) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE5] (iNum, Name, Description, Type, ExpRecover, NeedStone) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType5) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ExpRecover),
		GetOptionalDecVal(&this.NeedStone))
}

// GetCreateTableString Returns the create table statement for this object
func (this MagicType5) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE5] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[Type] tinyint NOT NULL,\n\t[ExpRecover] tinyint NOT NULL,\n\t[NeedStone] smallint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE5] PRIMARY KEY ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType5) SelectClause() (selectClause clause.Select) {
	return _MagicType5SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType5) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType5{}
	rawSql := "SELECT [iNum], [Name], [Description], [Type], [ExpRecover], [NeedStone] FROM [MAGIC_TYPE5]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType5SelectClause = clause.Select{
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
			Name: "[ExpRecover]",
		},
		clause.Column{
			Name: "[NeedStone]",
		},
	},
}
