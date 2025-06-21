package kogen

import (
	"fmt"
)

const (
	_ConcurrentTableName   = "CONCURRENT"
	_ConcurrentDatabaseNbr = 0
)

func init() {
	ModelList = append(ModelList, &Concurrent{})
}

// Concurrent: Keeps track of concurrent user counts
type Concurrent struct {
	ServerId   uint8    `gorm:"column:serverid;type:tinyint;not null" json:"serverid"`
	Zone1Count *int16   `gorm:"column:zone1_count;type:smallint" json:"zone1_count,omitempty"`
	Zone2Count *int16   `gorm:"column:zone2_count;type:smallint" json:"zone2_count,omitempty"`
	Zone3Count *int16   `gorm:"column:zone3_count;type:smallint" json:"zone3_count,omitempty"`
	Bz         [50]byte `gorm:"column:bz;type:varchar(50)" json:"bz,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Concurrent) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ConcurrentDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Concurrent) GetTableName() string {
	return _ConcurrentTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Concurrent) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [CONCURRENT] (serverid, zone1_count, zone2_count, zone3_count, bz) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(this.Zone1Count),
		GetOptionalDecVal(this.Zone2Count),
		GetOptionalDecVal(this.Zone3Count),
		GetOptionalBinaryVal(this.Bz))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Concurrent) GetCreateTableString() string {
	query := "CREATE TABLE \"CONCURRENT\" (\n\t\"serverid\" tinyint NOT NULL,\n\t\"zone1_count\" smallint,\n\t\"zone2_count\" smallint,\n\t\"zone3_count\" smallint,\n\t\"bz\" varchar(50)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
