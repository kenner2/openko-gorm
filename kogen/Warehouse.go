package kogen

import (
	"fmt"
)

const (
	_WarehouseDatabaseNbr = 0
	_WarehouseTableName   = "WAREHOUSE"
)

func init() {
	ModelList = append(ModelList, &Warehouse{})
}

// Warehouse The warehouse system is referred to as the Inn in-game.  It is account-level storage for a user
type Warehouse struct {
	AccountId [50]byte   `gorm:"column:strAccountID;type:varchar(50);primaryKey;not null" json:"strAccountID"`
	Money     int        `gorm:"column:nMoney;type:int;not null;default:0" json:"nMoney"`
	DwTime    int        `gorm:"column:dwTime;type:int;not null;default:0" json:"dwTime"`
	ItemData  [1600]byte `gorm:"column:WarehouseData;type:varchar(1600)" json:"WarehouseData,omitempty"`
	Serial    [1600]byte `gorm:"column:strSerial;type:varchar(1600)" json:"strSerial,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Warehouse) GetDatabaseName() string {
	return GetDatabaseName(DbType(_WarehouseDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Warehouse) GetTableName() string {
	return _WarehouseTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Warehouse) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WAREHOUSE] (strAccountID, nMoney, dwTime, WarehouseData, strSerial) \nVALUES (%s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalDecVal(&this.Money),
		GetOptionalDecVal(&this.DwTime),
		GetOptionalBinaryVal(this.ItemData),
		GetOptionalBinaryVal(this.Serial))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Warehouse) GetCreateTableString() string {
	query := "CREATE TABLE [WAREHOUSE] (\n\t[strAccountID] varchar(50) NOT NULL,\n\t[nMoney] int NOT NULL,\n\t[dwTime] int NOT NULL,\n\t[WarehouseData] varchar(1600),\n\t[strSerial] varchar(1600)\n\tPRIMARY KEY (\"strAccountID\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
