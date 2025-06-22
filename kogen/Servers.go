package kogen

import (
	"fmt"
)

const (
	_ServersDatabaseNbr = 0
	_ServersTableName   = "SERVERS"
)

func init() {
	ModelList = append(ModelList, &Servers{})
}

// Servers Game server instances
type Servers struct {
	Id      int      `gorm:"column:id;type:int;not null" json:"id"`
	Name    [64]byte `gorm:"column:sName;type:varchar(64);not null" json:"sName"`
	Host    [64]byte `gorm:"column:sHost;type:varchar(64);not null" json:"sHost"`
	Players int      `gorm:"column:players;type:int;not null" json:"players"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Servers) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ServersDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Servers) GetTableName() string {
	return _ServersTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Servers) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [SERVERS] (id, sName, sHost, players) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Id),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Host),
		GetOptionalDecVal(&this.Players))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Servers) GetCreateTableString() string {
	query := "CREATE TABLE [SERVERS] (\n\t[id] int NOT NULL,\n\t[sName] varchar(64) NOT NULL,\n\t[sHost] varchar(64) NOT NULL,\n\t[players] int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
