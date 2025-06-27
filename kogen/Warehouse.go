package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_WarehouseDatabaseNbr = "GAME"
	_WarehouseTableName   = "WAREHOUSE"
)

func init() {
	ModelList = append(ModelList, &Warehouse{})
}

// Warehouse The warehouse system is referred to as the Inn in-game.  It is account-level storage for a user
type Warehouse struct {
	AccountId mssql.VarChar  `gorm:"column:strAccountID;type:varchar(21);primaryKey;not null" json:"strAccountID"`
	Money     int            `gorm:"column:nMoney;type:int;not null;default:0" json:"nMoney"`
	DwTime    int            `gorm:"column:dwTime;type:int;not null;default:0" json:"dwTime"`
	ItemData  *mssql.VarChar `gorm:"column:WarehouseData;type:varchar(1600)" json:"WarehouseData,omitempty"`
	Serial    *mssql.VarChar `gorm:"column:strSerial;type:varchar(1600)" json:"strSerial,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this Warehouse) GetDatabaseName() string {
	return GetDatabaseName(_WarehouseDatabaseNbr)
}

// TableName Returns the table name
func (this Warehouse) TableName() string {
	return _WarehouseTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Warehouse) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WAREHOUSE] ([strAccountID], [nMoney], [dwTime], [WarehouseData], [strSerial]) VALUES\n(%s, %s, %s, CONVERT(varchar(1600), %s), CONVERT(varchar(1600), %s))", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.Money),
		GetOptionalDecVal(&this.DwTime),
		GetOptionalVarCharVal(this.ItemData, true),
		GetOptionalVarCharVal(this.Serial, true))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Warehouse) GetInsertHeader() string {
	return "INSERT INTO [WAREHOUSE] ([strAccountID], [nMoney], [dwTime], [WarehouseData], [strSerial]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Warehouse) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, CONVERT(varchar(1600), %s), CONVERT(varchar(1600), %s))", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.Money),
		GetOptionalDecVal(&this.DwTime),
		GetOptionalVarCharVal(this.ItemData, true),
		GetOptionalVarCharVal(this.Serial, true))
}

// GetCreateTableString Returns the create table statement for this object
func (this Warehouse) GetCreateTableString() string {
	query := "CREATE TABLE [WAREHOUSE] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[nMoney] int NOT NULL,\n\t[dwTime] int NOT NULL,\n\t[WarehouseData] varchar(1600),\n\t[strSerial] varchar(1600)\n\tCONSTRAINT [PK_WAREHOUSE] PRIMARY KEY CLUSTERED ([strAccountID])\n)\nGO\nALTER TABLE [WAREHOUSE] ADD CONSTRAINT [DF_WAREHOUSE_nMoney] DEFAULT 0 FOR [nMoney]\nGO\nALTER TABLE [WAREHOUSE] ADD CONSTRAINT [DF_WAREHOUSE_dwTime] DEFAULT 0 FOR [dwTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Warehouse) SelectClause() (selectClause clause.Select) {
	return _WarehouseSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Warehouse) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Warehouse{}
	rawSql := "SELECT [strAccountID], [nMoney], [dwTime], CONVERT(VARBINARY(1600), [WarehouseData]) as [WarehouseData], CONVERT(VARBINARY(1600), [strSerial]) as [strSerial] FROM [WAREHOUSE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _WarehouseSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[nMoney]",
		},
		clause.Column{
			Name: "[dwTime]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1600), [WarehouseData]) as [WarehouseData]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1600), [strSerial]) as [strSerial]",
		},
	},
}
