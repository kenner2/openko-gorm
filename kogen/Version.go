package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_VersionDatabaseNbr = 1
	_VersionTableName   = "VERSION"
)

func init() {
	ModelList = append(ModelList, &Version{})
}

// Version Version data and patch management
type Version struct {
	Version        int16         `gorm:"column:sVersion;type:smallint;primaryKey;not null" json:"sVersion"`
	FileName       mssql.VarChar `gorm:"column:strFileName;type:varchar(50);not null" json:"strFileName"`
	CompressName   mssql.VarChar `gorm:"column:strCompressName;type:varchar(50);not null" json:"strCompressName"`
	HistoryVersion int16         `gorm:"column:sHistoryVersion;type:smallint;not null" json:"sHistoryVersion"`
}

// GetDatabaseName Returns the table's database name
func (this Version) GetDatabaseName() string {
	return GetDatabaseName(DbType(_VersionDatabaseNbr))
}

// TableName Returns the table name
func (this Version) TableName() string {
	return _VersionTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Version) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [VERSION] ([sVersion], [strFileName], [strCompressName], [sHistoryVersion]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Version),
		GetOptionalVarCharVal(&this.FileName, false),
		GetOptionalVarCharVal(&this.CompressName, false),
		GetOptionalDecVal(&this.HistoryVersion))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Version) GetInsertHeader() string {
	return "INSERT INTO [VERSION] (sVersion, strFileName, strCompressName, sHistoryVersion) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Version) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Version),
		GetOptionalVarCharVal(&this.FileName, false),
		GetOptionalVarCharVal(&this.CompressName, false),
		GetOptionalDecVal(&this.HistoryVersion))
}

// GetCreateTableString Returns the create table statement for this object
func (this Version) GetCreateTableString() string {
	query := "CREATE TABLE [VERSION] (\n\t[sVersion] smallint NOT NULL,\n\t[strFileName] varchar(50) NOT NULL,\n\t[strCompressName] varchar(50) NOT NULL,\n\t[sHistoryVersion] smallint NOT NULL\n\tCONSTRAINT [PK_VERSION] PRIMARY KEY ([sVersion])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Version) SelectClause() (selectClause clause.Select) {
	return _VersionSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Version) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Version{}
	rawSql := "SELECT [sVersion], [strFileName], [strCompressName], [sHistoryVersion] FROM [VERSION]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _VersionSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sVersion]",
		},
		clause.Column{
			Name: "[strFileName]",
		},
		clause.Column{
			Name: "[strCompressName]",
		},
		clause.Column{
			Name: "[sHistoryVersion]",
		},
	},
}
