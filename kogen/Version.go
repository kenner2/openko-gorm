package kogen

import (
	"fmt"
)

const (
	_VersionDatabaseNbr = 0
	_VersionTableName   = "VERSION"
)

func init() {
	ModelList = append(ModelList, &Version{})
}

// Version Version data and patch management
type Version struct {
	Version        int16    `gorm:"column:sVersion;type:smallint;not null" json:"sVersion"`
	FileName       [50]byte `gorm:"column:strFileName;type:varchar(50);not null" json:"strFileName"`
	CompressName   [50]byte `gorm:"column:strCompressName;type:varchar(50);not null" json:"strCompressName"`
	HistoryVersion int16    `gorm:"column:sHistoryVersion;type:smallint;not null" json:"sHistoryVersion"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Version) GetDatabaseName() string {
	return GetDatabaseName(DbType(_VersionDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Version) GetTableName() string {
	return _VersionTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Version) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [VERSION] (sVersion, strFileName, strCompressName, sHistoryVersion) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Version),
		GetOptionalBinaryVal(this.FileName),
		GetOptionalBinaryVal(this.CompressName),
		GetOptionalDecVal(&this.HistoryVersion))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Version) GetCreateTableString() string {
	query := "CREATE TABLE [VERSION] (\n\t\"sVersion\" smallint NOT NULL,\n\t\"strFileName\" varchar(50) NOT NULL,\n\t\"strCompressName\" varchar(50) NOT NULL,\n\t\"sHistoryVersion\" smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
