package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ConcurrentDatabaseNbr = "GAME"
	_ConcurrentTableName   = "CONCURRENT"
)

func init() {
	ModelList = append(ModelList, &Concurrent{})
}

// Concurrent Keeps track of concurrent user counts
type Concurrent struct {
	ServerId   uint8          `gorm:"column:serverid;type:tinyint;primaryKey;not null" json:"serverid"`
	Zone1Count *int16         `gorm:"column:zone1_count;type:smallint" json:"zone1_count,omitempty"`
	Zone2Count *int16         `gorm:"column:zone2_count;type:smallint" json:"zone2_count,omitempty"`
	Zone3Count *int16         `gorm:"column:zone3_count;type:smallint" json:"zone3_count,omitempty"`
	Bz         *mssql.VarChar `gorm:"column:bz;type:varchar(50)" json:"bz,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this Concurrent) GetDatabaseName() string {
	return GetDatabaseName(_ConcurrentDatabaseNbr)
}

// TableName Returns the table name
func (this Concurrent) TableName() string {
	return _ConcurrentTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Concurrent) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [CONCURRENT] ([serverid], [zone1_count], [zone2_count], [zone3_count], [bz]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(this.Zone1Count),
		GetOptionalDecVal(this.Zone2Count),
		GetOptionalDecVal(this.Zone3Count),
		GetOptionalVarCharVal(this.Bz, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Concurrent) GetInsertHeader() string {
	return "INSERT INTO [CONCURRENT] ([serverid], [zone1_count], [zone2_count], [zone3_count], [bz]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Concurrent) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(this.Zone1Count),
		GetOptionalDecVal(this.Zone2Count),
		GetOptionalDecVal(this.Zone3Count),
		GetOptionalVarCharVal(this.Bz, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this Concurrent) GetCreateTableString() string {
	query := "CREATE TABLE [CONCURRENT] (\n\t[serverid] tinyint NOT NULL,\n\t[zone1_count] smallint,\n\t[zone2_count] smallint,\n\t[zone3_count] smallint,\n\t[bz] varchar(50)\n\tCONSTRAINT [PK_CONCURRENT] PRIMARY KEY CLUSTERED ([serverid])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Concurrent) SelectClause() (selectClause clause.Select) {
	return _ConcurrentSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Concurrent) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Concurrent{}
	rawSql := "SELECT [serverid], [zone1_count], [zone2_count], [zone3_count], [bz] FROM [CONCURRENT]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ConcurrentSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[serverid]",
		},
		clause.Column{
			Name: "[zone1_count]",
		},
		clause.Column{
			Name: "[zone2_count]",
		},
		clause.Column{
			Name: "[zone3_count]",
		},
		clause.Column{
			Name: "[bz]",
		},
	},
}
