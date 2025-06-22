package kogen

import (
	"fmt"
)

const (
	_MagicType5TableName   = "MAGIC_TYPE5"
	_MagicType5DatabaseNbr = 0
)

func init() {
	ModelList = append(ModelList, &MagicType5{})
}

// MagicType5 Type 5 supports recovery skills
type MagicType5 struct {
	MagicNumber int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name        [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Type        uint8     `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	ExpRecover  uint8     `gorm:"column:ExpRecover;type:tinyint;not null" json:"ExpRecover"`
	NeedStone   int16     `gorm:"column:NeedStone;type:smallint;not null" json:"NeedStone"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType5) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType5DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType5) GetTableName() string {
	return _MagicType5TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType5) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE5] (iNum, Name, Description, Type, ExpRecover, NeedStone) \nVALUES (%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.ExpRecover),
		GetOptionalDecVal(&this.NeedStone))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType5) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE5] (\n\t\"iNum\" int NOT NULL,\n\t\"Name\" varchar(50),\n\t\"Description\" varchar(100),\n\t\"Type\" tinyint NOT NULL,\n\t\"ExpRecover\" tinyint NOT NULL,\n\t\"NeedStone\" smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
