package kogen

import (
	"fmt"
)

const (
	_ServerResourceDatabaseNbr = 0
	_ServerResourceTableName   = "SERVER_RESOURCE"
)

func init() {
	ModelList = append(ModelList, &ServerResource{})
}

// ServerResource Server resource
type ServerResource struct {
	ResourceId int       `gorm:"column:nResourceID;type:int;not null" json:"nResourceID"`
	Name       [50]byte  `gorm:"column:strName;type:varchar(50);not null" json:"strName"`
	Resource   [100]byte `gorm:"column:strResource;type:varchar(100)" json:"strResource,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ServerResource) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ServerResourceDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ServerResource) GetTableName() string {
	return _ServerResourceTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ServerResource) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [SERVER_RESOURCE] (nResourceID, strName, strResource) \nVALUES (%s, %s, %s)", GetOptionalDecVal(&this.ResourceId),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Resource))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ServerResource) GetCreateTableString() string {
	query := "CREATE TABLE [SERVER_RESOURCE] (\n\t\"nResourceID\" int NOT NULL,\n\t\"strName\" varchar(50) NOT NULL,\n\t\"strResource\" varchar(100)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
