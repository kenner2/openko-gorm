package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MakeWeaponDatabaseNbr = "GAME"
	_MakeWeaponTableName   = "MAKE_WEAPON"
)

func init() {
	ModelList = append(ModelList, &MakeWeapon{})
}

// MakeWeapon Make weapon
type MakeWeapon struct {
	Level   uint8  `gorm:"column:byLevel;type:tinyint;primaryKey;not null" json:"byLevel"`
	Class1  *int16 `gorm:"column:sClass_1;type:smallint" json:"sClass_1,omitempty"`
	Class2  *int16 `gorm:"column:sClass_2;type:smallint" json:"sClass_2,omitempty"`
	Class3  *int16 `gorm:"column:sClass_3;type:smallint" json:"sClass_3,omitempty"`
	Class4  *int16 `gorm:"column:sClass_4;type:smallint" json:"sClass_4,omitempty"`
	Class5  *int16 `gorm:"column:sClass_5;type:smallint" json:"sClass_5,omitempty"`
	Class6  *int16 `gorm:"column:sClass_6;type:smallint" json:"sClass_6,omitempty"`
	Class7  *int16 `gorm:"column:sClass_7;type:smallint" json:"sClass_7,omitempty"`
	Class8  *int16 `gorm:"column:sClass_8;type:smallint" json:"sClass_8,omitempty"`
	Class9  *int16 `gorm:"column:sClass_9;type:smallint" json:"sClass_9,omitempty"`
	Class10 *int16 `gorm:"column:sClass_10;type:smallint" json:"sClass_10,omitempty"`
	Class11 *int16 `gorm:"column:sClass_11;type:smallint" json:"sClass_11,omitempty"`
	Class12 *int16 `gorm:"column:sClass_12;type:smallint" json:"sClass_12,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this MakeWeapon) GetDatabaseName() string {
	return GetDatabaseName(_MakeWeaponDatabaseNbr)
}

// TableName Returns the table name
func (this MakeWeapon) TableName() string {
	return _MakeWeaponTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MakeWeapon) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_WEAPON] ([byLevel], [sClass_1], [sClass_2], [sClass_3], [sClass_4], [sClass_5], [sClass_6], [sClass_7], [sClass_8], [sClass_9], [sClass_10], [sClass_11], [sClass_12]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(this.Class1),
		GetOptionalDecVal(this.Class2),
		GetOptionalDecVal(this.Class3),
		GetOptionalDecVal(this.Class4),
		GetOptionalDecVal(this.Class5),
		GetOptionalDecVal(this.Class6),
		GetOptionalDecVal(this.Class7),
		GetOptionalDecVal(this.Class8),
		GetOptionalDecVal(this.Class9),
		GetOptionalDecVal(this.Class10),
		GetOptionalDecVal(this.Class11),
		GetOptionalDecVal(this.Class12))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MakeWeapon) GetInsertHeader() string {
	return "INSERT INTO [MAKE_WEAPON] ([byLevel], [sClass_1], [sClass_2], [sClass_3], [sClass_4], [sClass_5], [sClass_6], [sClass_7], [sClass_8], [sClass_9], [sClass_10], [sClass_11], [sClass_12]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MakeWeapon) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(this.Class1),
		GetOptionalDecVal(this.Class2),
		GetOptionalDecVal(this.Class3),
		GetOptionalDecVal(this.Class4),
		GetOptionalDecVal(this.Class5),
		GetOptionalDecVal(this.Class6),
		GetOptionalDecVal(this.Class7),
		GetOptionalDecVal(this.Class8),
		GetOptionalDecVal(this.Class9),
		GetOptionalDecVal(this.Class10),
		GetOptionalDecVal(this.Class11),
		GetOptionalDecVal(this.Class12))
}

// GetCreateTableString Returns the create table statement for this object
func (this MakeWeapon) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_WEAPON] (\n\t[byLevel] tinyint NOT NULL,\n\t[sClass_1] smallint,\n\t[sClass_2] smallint,\n\t[sClass_3] smallint,\n\t[sClass_4] smallint,\n\t[sClass_5] smallint,\n\t[sClass_6] smallint,\n\t[sClass_7] smallint,\n\t[sClass_8] smallint,\n\t[sClass_9] smallint,\n\t[sClass_10] smallint,\n\t[sClass_11] smallint,\n\t[sClass_12] smallint\n\tCONSTRAINT [PK_MAKE_WEAPON] PRIMARY KEY CLUSTERED ([byLevel])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MakeWeapon) SelectClause() (selectClause clause.Select) {
	return _MakeWeaponSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MakeWeapon) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MakeWeapon{}
	rawSql := "SELECT [byLevel], [sClass_1], [sClass_2], [sClass_3], [sClass_4], [sClass_5], [sClass_6], [sClass_7], [sClass_8], [sClass_9], [sClass_10], [sClass_11], [sClass_12] FROM [MAKE_WEAPON]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MakeWeaponSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[byLevel]",
		},
		clause.Column{
			Name: "[sClass_1]",
		},
		clause.Column{
			Name: "[sClass_2]",
		},
		clause.Column{
			Name: "[sClass_3]",
		},
		clause.Column{
			Name: "[sClass_4]",
		},
		clause.Column{
			Name: "[sClass_5]",
		},
		clause.Column{
			Name: "[sClass_6]",
		},
		clause.Column{
			Name: "[sClass_7]",
		},
		clause.Column{
			Name: "[sClass_8]",
		},
		clause.Column{
			Name: "[sClass_9]",
		},
		clause.Column{
			Name: "[sClass_10]",
		},
		clause.Column{
			Name: "[sClass_11]",
		},
		clause.Column{
			Name: "[sClass_12]",
		},
	},
}
