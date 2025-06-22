package kogen

import (
	"fmt"
)

const (
	_CouponSerialListDatabaseNbr = 0
	_CouponSerialListTableName   = "COUPON_SERIAL_LIST"
)

func init() {
	ModelList = append(ModelList, &CouponSerialList{})
}

// CouponSerialList Coupon Serial List
type CouponSerialList struct {
	Index      int      `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	SerialNum  [16]byte `gorm:"column:strSerialNum;type:char(16);not null" json:"strSerialNum"`
	ItemNumber int      `gorm:"column:nItemNum;type:int;not null" json:"nItemNum"`
	ItemCount  int16    `gorm:"column:sItemCount;type:smallint;not null" json:"sItemCount"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *CouponSerialList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CouponSerialListDatabaseNbr))
}

// GetTableName Returns the table name
func (this *CouponSerialList) GetTableName() string {
	return _CouponSerialListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *CouponSerialList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COUPON_SERIAL_LIST] (nIndex, strSerialNum, nItemNum, sItemCount) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalBinaryVal(this.SerialNum),
		GetOptionalDecVal(&this.ItemNumber),
		GetOptionalDecVal(&this.ItemCount))
}

// GetCreateTableString Returns the create table statement for this object
func (this *CouponSerialList) GetCreateTableString() string {
	query := "CREATE TABLE [COUPON_SERIAL_LIST] (\n\t[nIndex] int NOT NULL,\n\t[strSerialNum] char(16) NOT NULL,\n\t[nItemNum] int NOT NULL,\n\t[sItemCount] smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
