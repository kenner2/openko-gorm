package kogen

import (
	"fmt"
)

const (
	_MagicType8DatabaseNbr = 0
	_MagicType8TableName   = "MAGIC_TYPE8"
)

func init() {
	ModelList = append(ModelList, &MagicType8{})
}

// MagicType8 Type 8 supports warp magic
type MagicType8 struct {
	MagicNumber int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name        [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Target      uint8     `gorm:"column:Target;type:tinyint;not null" json:"Target"`
	Radius      int16     `gorm:"column:Radius;type:smallint;not null" json:"Radius"`
	WarpType    uint8     `gorm:"column:WarpType;type:tinyint;not null" json:"WarpType"`
	ExpRecover  int16     `gorm:"column:ExpRecover;type:smallint;not null" json:"ExpRecover"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType8) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType8DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType8) GetTableName() string {
	return _MagicType8TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType8) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE8] (iNum, Name, Description, Target, Radius, WarpType, ExpRecover) \nVALUES (%s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.Target),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.WarpType),
		GetOptionalDecVal(&this.ExpRecover))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType8) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE8] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[Target] tinyint NOT NULL,\n\t[Radius] smallint NOT NULL,\n\t[WarpType] tinyint NOT NULL,\n\t[ExpRecover] smallint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
