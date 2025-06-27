package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ServerResourceDatabaseNbr = "GAME"
	_ServerResourceTableName   = "SERVER_RESOURCE"
)

func init() {
	ModelList = append(ModelList, &ServerResource{})
}

// ServerResource Server resource
type ServerResource struct {
	ResourceId int            `gorm:"column:nResourceID;type:int;primaryKey;not null" json:"nResourceID"`
	Name       mssql.VarChar  `gorm:"column:strName;type:varchar(50);not null" json:"strName"`
	Resource   *mssql.VarChar `gorm:"column:strResource;type:varchar(100)" json:"strResource,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this ServerResource) GetDatabaseName() string {
	return GetDatabaseName(_ServerResourceDatabaseNbr)
}

// TableName Returns the table name
func (this ServerResource) TableName() string {
	return _ServerResourceTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ServerResource) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [SERVER_RESOURCE] ([nResourceID], [strName], [strResource]) VALUES\n(%s, %s, %s)", GetOptionalDecVal(&this.ResourceId),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(this.Resource, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ServerResource) GetInsertHeader() string {
	return "INSERT INTO [SERVER_RESOURCE] ([nResourceID], [strName], [strResource]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ServerResource) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s)", GetOptionalDecVal(&this.ResourceId),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(this.Resource, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this ServerResource) GetCreateTableString() string {
	query := "CREATE TABLE [SERVER_RESOURCE] (\n\t[nResourceID] int NOT NULL,\n\t[strName] varchar(50) NOT NULL,\n\t[strResource] varchar(100)\n\tCONSTRAINT [PK_SERVER_RESOURCE] PRIMARY KEY CLUSTERED ([nResourceID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ServerResource) SelectClause() (selectClause clause.Select) {
	return _ServerResourceSelectClause
}

// GetAllTableData Returns a list of all table data
func (this ServerResource) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ServerResource{}
	rawSql := "SELECT [nResourceID], [strName], [strResource] FROM [SERVER_RESOURCE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ServerResourceSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nResourceID]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[strResource]",
		},
	},
}
