package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_UserSavedMagicDatabaseNbr = "GAME"
	_UserSavedMagicTableName   = "USER_SAVED_MAGIC"
)

func init() {
	ModelList = append(ModelList, &UserSavedMagic{})
}

// UserSavedMagic User saved magic
type UserSavedMagic struct {
	CharId   mssql.VarChar `gorm:"column:strCharID;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS;primaryKey;not null" json:"strCharID"`
	Skill1   int           `gorm:"column:nSkill1;type:int;not null" json:"nSkill1"`
	During1  int16         `gorm:"column:nDuring1;type:smallint;not null" json:"nDuring1"`
	Skill2   int           `gorm:"column:nSkill2;type:int;not null" json:"nSkill2"`
	During2  int16         `gorm:"column:nDuring2;type:smallint;not null" json:"nDuring2"`
	Skill3   int           `gorm:"column:nSkill3;type:int;not null" json:"nSkill3"`
	During3  int16         `gorm:"column:nDuring3;type:smallint;not null" json:"nDuring3"`
	Skill4   int           `gorm:"column:nSkill4;type:int;not null" json:"nSkill4"`
	During4  int16         `gorm:"column:nDuring4;type:smallint;not null" json:"nDuring4"`
	Skill5   int           `gorm:"column:nSkill5;type:int;not null" json:"nSkill5"`
	During5  int16         `gorm:"column:nDuring5;type:smallint;not null" json:"nDuring5"`
	Skill6   int           `gorm:"column:nSkill6;type:int;not null" json:"nSkill6"`
	During6  int16         `gorm:"column:nDuring6;type:smallint;not null" json:"nDuring6"`
	Skill7   int           `gorm:"column:nSkill7;type:int;not null" json:"nSkill7"`
	During7  int16         `gorm:"column:nDuring7;type:smallint;not null" json:"nDuring7"`
	Skill8   int           `gorm:"column:nSkill8;type:int;not null" json:"nSkill8"`
	During8  int16         `gorm:"column:nDuring8;type:smallint;not null" json:"nDuring8"`
	Skill9   int           `gorm:"column:nSkill9;type:int;not null" json:"nSkill9"`
	During9  int16         `gorm:"column:nDuring9;type:smallint;not null" json:"nDuring9"`
	Skill10  int           `gorm:"column:nSkill10;type:int;not null" json:"nSkill10"`
	During10 int16         `gorm:"column:nDuring10;type:smallint;not null" json:"nDuring10"`
}

// GetDatabaseName Returns the table's database name
func (this UserSavedMagic) GetDatabaseName() string {
	return GetDatabaseName(_UserSavedMagicDatabaseNbr)
}

// TableName Returns the table name
func (this UserSavedMagic) TableName() string {
	return _UserSavedMagicTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserSavedMagic) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_SAVED_MAGIC] ([strCharID], [nSkill1], [nDuring1], [nSkill2], [nDuring2], [nSkill3], [nDuring3], [nSkill4], [nDuring4], [nSkill5], [nDuring5], [nSkill6], [nDuring6], [nSkill7], [nDuring7], [nSkill8], [nDuring8], [nSkill9], [nDuring9], [nSkill10], [nDuring10]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Skill1),
		GetOptionalDecVal(&this.During1),
		GetOptionalDecVal(&this.Skill2),
		GetOptionalDecVal(&this.During2),
		GetOptionalDecVal(&this.Skill3),
		GetOptionalDecVal(&this.During3),
		GetOptionalDecVal(&this.Skill4),
		GetOptionalDecVal(&this.During4),
		GetOptionalDecVal(&this.Skill5),
		GetOptionalDecVal(&this.During5),
		GetOptionalDecVal(&this.Skill6),
		GetOptionalDecVal(&this.During6),
		GetOptionalDecVal(&this.Skill7),
		GetOptionalDecVal(&this.During7),
		GetOptionalDecVal(&this.Skill8),
		GetOptionalDecVal(&this.During8),
		GetOptionalDecVal(&this.Skill9),
		GetOptionalDecVal(&this.During9),
		GetOptionalDecVal(&this.Skill10),
		GetOptionalDecVal(&this.During10))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserSavedMagic) GetInsertHeader() string {
	return "INSERT INTO [USER_SAVED_MAGIC] ([strCharID], [nSkill1], [nDuring1], [nSkill2], [nDuring2], [nSkill3], [nDuring3], [nSkill4], [nDuring4], [nSkill5], [nDuring5], [nSkill6], [nDuring6], [nSkill7], [nDuring7], [nSkill8], [nDuring8], [nSkill9], [nDuring9], [nSkill10], [nDuring10]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserSavedMagic) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.Skill1),
		GetOptionalDecVal(&this.During1),
		GetOptionalDecVal(&this.Skill2),
		GetOptionalDecVal(&this.During2),
		GetOptionalDecVal(&this.Skill3),
		GetOptionalDecVal(&this.During3),
		GetOptionalDecVal(&this.Skill4),
		GetOptionalDecVal(&this.During4),
		GetOptionalDecVal(&this.Skill5),
		GetOptionalDecVal(&this.During5),
		GetOptionalDecVal(&this.Skill6),
		GetOptionalDecVal(&this.During6),
		GetOptionalDecVal(&this.Skill7),
		GetOptionalDecVal(&this.During7),
		GetOptionalDecVal(&this.Skill8),
		GetOptionalDecVal(&this.During8),
		GetOptionalDecVal(&this.Skill9),
		GetOptionalDecVal(&this.During9),
		GetOptionalDecVal(&this.Skill10),
		GetOptionalDecVal(&this.During10))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserSavedMagic) GetCreateTableString() string {
	query := "CREATE TABLE [USER_SAVED_MAGIC] (\n\t[strCharID] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[nSkill1] int NOT NULL,\n\t[nDuring1] smallint NOT NULL,\n\t[nSkill2] int NOT NULL,\n\t[nDuring2] smallint NOT NULL,\n\t[nSkill3] int NOT NULL,\n\t[nDuring3] smallint NOT NULL,\n\t[nSkill4] int NOT NULL,\n\t[nDuring4] smallint NOT NULL,\n\t[nSkill5] int NOT NULL,\n\t[nDuring5] smallint NOT NULL,\n\t[nSkill6] int NOT NULL,\n\t[nDuring6] smallint NOT NULL,\n\t[nSkill7] int NOT NULL,\n\t[nDuring7] smallint NOT NULL,\n\t[nSkill8] int NOT NULL,\n\t[nDuring8] smallint NOT NULL,\n\t[nSkill9] int NOT NULL,\n\t[nDuring9] smallint NOT NULL,\n\t[nSkill10] int NOT NULL,\n\t[nDuring10] smallint NOT NULL\n\tCONSTRAINT [PK_USER_SAVED_MAGIC] PRIMARY KEY CLUSTERED ([strCharID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserSavedMagic) SelectClause() (selectClause clause.Select) {
	return _UserSavedMagicSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserSavedMagic) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserSavedMagic{}
	rawSql := "SELECT [strCharID], [nSkill1], [nDuring1], [nSkill2], [nDuring2], [nSkill3], [nDuring3], [nSkill4], [nDuring4], [nSkill5], [nDuring5], [nSkill6], [nDuring6], [nSkill7], [nDuring7], [nSkill8], [nDuring8], [nSkill9], [nDuring9], [nSkill10], [nDuring10] FROM [USER_SAVED_MAGIC]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserSavedMagicSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[nSkill1]",
		},
		clause.Column{
			Name: "[nDuring1]",
		},
		clause.Column{
			Name: "[nSkill2]",
		},
		clause.Column{
			Name: "[nDuring2]",
		},
		clause.Column{
			Name: "[nSkill3]",
		},
		clause.Column{
			Name: "[nDuring3]",
		},
		clause.Column{
			Name: "[nSkill4]",
		},
		clause.Column{
			Name: "[nDuring4]",
		},
		clause.Column{
			Name: "[nSkill5]",
		},
		clause.Column{
			Name: "[nDuring5]",
		},
		clause.Column{
			Name: "[nSkill6]",
		},
		clause.Column{
			Name: "[nDuring6]",
		},
		clause.Column{
			Name: "[nSkill7]",
		},
		clause.Column{
			Name: "[nDuring7]",
		},
		clause.Column{
			Name: "[nSkill8]",
		},
		clause.Column{
			Name: "[nDuring8]",
		},
		clause.Column{
			Name: "[nSkill9]",
		},
		clause.Column{
			Name: "[nDuring9]",
		},
		clause.Column{
			Name: "[nSkill10]",
		},
		clause.Column{
			Name: "[nDuring10]",
		},
	},
}
