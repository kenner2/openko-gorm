package kogen

import (
	"fmt"
)

const (
	_ConcurrentDatabaseNbr = 0
	_ConcurrentTableName   = "CONCURRENT"
)

func init() {
	ModelList = append(ModelList, &Concurrent{})
}

// Concurrent: Keeps track of concurrent user counts
type Concurrent struct {
	ServerId   uint8   `gorm:"column:serverid;not null" json:"serverid"`
	Zone1Count *int16  `gorm:"column:zone1_count" json:"zone1_count,omitempty"`
	Zone2Count *int16  `gorm:"column:zone2_count" json:"zone2_count,omitempty"`
	Zone3Count *int16  `gorm:"column:zone3_count" json:"zone3_count,omitempty"`
	Bz         *string `gorm:"column:bz" json:"bz,omitempty"`
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
		GetOptionalStringVal(this.Bz))
}
