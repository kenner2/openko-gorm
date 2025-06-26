package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ServersDatabaseNbr = 1
	_ServersTableName   = "SERVERS"
)

func init() {
	ModelList = append(ModelList, &Servers{})
}

// Servers Game server instances
type Servers struct {
	Id      int           `gorm:"column:id;type:int;primaryKey;not null" json:"id"`
	Name    mssql.VarChar `gorm:"column:sName;type:varchar(64);not null" json:"sName"`
	Host    mssql.VarChar `gorm:"column:sHost;type:varchar(64);not null" json:"sHost"`
	Players int           `gorm:"column:players;type:int;not null" json:"players"`
}

// GetDatabaseName Returns the table's database name
func (this Servers) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ServersDatabaseNbr))
}

// TableName Returns the table name
func (this Servers) TableName() string {
	return _ServersTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Servers) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [SERVERS] ([id], [sName], [sHost], [players]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(&this.Host, false),
		GetOptionalDecVal(&this.Players))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Servers) GetInsertHeader() string {
	return "INSERT INTO [SERVERS] (id, sName, sHost, players) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Servers) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(&this.Host, false),
		GetOptionalDecVal(&this.Players))
}

// GetCreateTableString Returns the create table statement for this object
func (this Servers) GetCreateTableString() string {
	query := "CREATE TABLE [SERVERS] (\n\t[id] int NOT NULL,\n\t[sName] varchar(64) NOT NULL,\n\t[sHost] varchar(64) NOT NULL,\n\t[players] int NOT NULL\n\tCONSTRAINT [PK_SERVERS] PRIMARY KEY ([id])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Servers) SelectClause() (selectClause clause.Select) {
	return _ServersSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Servers) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Servers{}
	rawSql := "SELECT [id], [sName], [sHost], [players] FROM [SERVERS]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ServersSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[id]",
		},
		clause.Column{
			Name: "[sName]",
		},
		clause.Column{
			Name: "[sHost]",
		},
		clause.Column{
			Name: "[players]",
		},
	},
}
