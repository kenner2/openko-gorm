package kogen

import (
	"fmt"
)

const (
	_RentalItemDatabaseNbr = 0
	_RentalItemTableName   = "RENTAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &RentalItem{})
}

// RentalItem Rental item
type RentalItem struct {
	RentalIndex       int      `gorm:"column:nRentalIndex;type:int;primaryKey;not null" json:"nRentalIndex"`
	ItemIndex         int      `gorm:"column:nItemIndex;type:int;not null" json:"nItemIndex"`
	Durability        int16    `gorm:"column:sDurability;type:smallint;not null;default:0" json:"sDurability"`
	SerialNumber      int64    `gorm:"column:nSerialNumber;type:bigint;not null" json:"nSerialNumber"`
	RegType           uint8    `gorm:"column:byRegType;type:tinyint;not null;default:0" json:"byRegType"`
	ItemType          uint8    `gorm:"column:byItemType;type:tinyint;not null" json:"byItemType"`
	Class             uint8    `gorm:"column:byClass;type:tinyint;not null" json:"byClass"`
	RentalTime        int16    `gorm:"column:sRentalTime;type:smallint;not null" json:"sRentalTime"`
	RentalMoney       int      `gorm:"column:nRentalMoney;type:int;not null" json:"nRentalMoney"`
	LenderCharId      [21]byte `gorm:"column:strLenderCharID;type:char(21);not null" json:"strLenderCharID"`
	LenderAccountId   [21]byte `gorm:"column:strLenderAcID;type:char(21);not null" json:"strLenderAcID"`
	BorrowerCharId    [21]byte `gorm:"column:strBorrowerCharID;type:char(21)" json:"strBorrowerCharID,omitempty"`
	BorrowerAccountId [21]byte `gorm:"column:strBorrowerAcID;type:char(21)" json:"strBorrowerAcID,omitempty"`
	LendTime          *int     `gorm:"column:timeLender;type:smalldatetime" json:"timeLender,omitempty"`
	RegisterTime      int      `gorm:"column:timeRegister;type:smalldatetime;not null;default:getdate()" json:"timeRegister"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *RentalItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_RentalItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *RentalItem) GetTableName() string {
	return _RentalItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *RentalItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [RENTAL_ITEM] (nRentalIndex, nItemIndex, sDurability, nSerialNumber, byRegType, byItemType, byClass, sRentalTime, nRentalMoney, strLenderCharID, strLenderAcID, strBorrowerCharID, strBorrowerAcID, timeLender, timeRegister) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RegType),
		GetOptionalDecVal(&this.ItemType),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalBinaryVal(this.LenderCharId),
		GetOptionalBinaryVal(this.LenderAccountId),
		GetOptionalBinaryVal(this.BorrowerCharId),
		GetOptionalBinaryVal(this.BorrowerAccountId),
		GetOptionalDecVal(this.LendTime),
		GetOptionalDecVal(&this.RegisterTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *RentalItem) GetCreateTableString() string {
	query := "CREATE TABLE [RENTAL_ITEM] (\n\t\"nRentalIndex\" int NOT NULL,\n\t\"nItemIndex\" int NOT NULL,\n\t\"sDurability\" smallint NOT NULL,\n\t\"nSerialNumber\" bigint NOT NULL,\n\t\"byRegType\" tinyint NOT NULL,\n\t\"byItemType\" tinyint NOT NULL,\n\t\"byClass\" tinyint NOT NULL,\n\t\"sRentalTime\" smallint NOT NULL,\n\t\"nRentalMoney\" int NOT NULL,\n\t\"strLenderCharID\" char(21) NOT NULL,\n\t\"strLenderAcID\" char(21) NOT NULL,\n\t\"strBorrowerCharID\" char(21),\n\t\"strBorrowerAcID\" char(21),\n\t\"timeLender\" smalldatetime,\n\t\"timeRegister\" smalldatetime NOT NULL\n\tPRIMARY KEY (\"nRentalIndex\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
