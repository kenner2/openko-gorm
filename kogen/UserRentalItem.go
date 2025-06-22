package kogen

import (
	"fmt"
)

const (
	_UserRentalItemDatabaseNbr = 0
	_UserRentalItemTableName   = "USER_RENTAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &UserRentalItem{})
}

// UserRentalItem User rental item
type UserRentalItem struct {
	UserId          [50]byte `gorm:"column:strUserID;type:varchar(50);not null" json:"strUserID"`
	AccountId       [50]byte `gorm:"column:strAccountID;type:varchar(50);not null" json:"strAccountID"`
	RentalType      uint8    `gorm:"column:byRentalType;type:tinyint;not null" json:"byRentalType"`
	RegTime         uint8    `gorm:"column:byRegType;type:tinyint;not null;default:0" json:"byRegType"`
	RentalIndex     int      `gorm:"column:nRentalIndex;type:int;not null" json:"nRentalIndex"`
	ItemIndex       int      `gorm:"column:nItemIndex;type:int;not null" json:"nItemIndex"`
	Durability      int16    `gorm:"column:sDurability;type:smallint;not null;default:0" json:"sDurability"`
	SerialNumber    int64    `gorm:"column:nSerialNumber;type:bigint;not null" json:"nSerialNumber"`
	RentalMoney     int      `gorm:"column:nRentalMoney;type:int;not null" json:"nRentalMoney"`
	RentalTime      int16    `gorm:"column:sRentalTime;type:smallint;not null" json:"sRentalTime"`
	DuringTime      int16    `gorm:"column:sDuringTime;type:smallint;not null" json:"sDuringTime"`
	RentalTimestamp *int     `gorm:"column:timeRental;type:smalldatetime" json:"timeRental,omitempty"`
	RegisterTime    *int     `gorm:"column:timeRegister;type:smalldatetime;default:getdate()" json:"timeRegister,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserRentalItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserRentalItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserRentalItem) GetTableName() string {
	return _UserRentalItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserRentalItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_RENTAL_ITEM] (strUserID, strAccountID, byRentalType, byRegType, nRentalIndex, nItemIndex, sDurability, nSerialNumber, nRentalMoney, sRentalTime, sDuringTime, timeRental, timeRegister) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
		GetOptionalBinaryVal(this.AccountId),
		GetOptionalDecVal(&this.RentalType),
		GetOptionalDecVal(&this.RegTime),
		GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.DuringTime),
		GetOptionalDecVal(this.RentalTimestamp),
		GetOptionalDecVal(this.RegisterTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserRentalItem) GetCreateTableString() string {
	query := "CREATE TABLE [USER_RENTAL_ITEM] (\n\t\"strUserID\" varchar(50) NOT NULL,\n\t\"strAccountID\" varchar(50) NOT NULL,\n\t\"byRentalType\" tinyint NOT NULL,\n\t\"byRegType\" tinyint NOT NULL,\n\t\"nRentalIndex\" int NOT NULL,\n\t\"nItemIndex\" int NOT NULL,\n\t\"sDurability\" smallint NOT NULL,\n\t\"nSerialNumber\" bigint NOT NULL,\n\t\"nRentalMoney\" int NOT NULL,\n\t\"sRentalTime\" smallint NOT NULL,\n\t\"sDuringTime\" smallint NOT NULL,\n\t\"timeRental\" smalldatetime,\n\t\"timeRegister\" smalldatetime\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
