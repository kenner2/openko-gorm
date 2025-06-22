package kogen

import (
	"fmt"
)

const (
	_WebItemMallLogDatabaseNbr = 0
	_WebItemMallLogTableName   = "WEB_ITEMMALL_LOG"
)

func init() {
	ModelList = append(ModelList, &WebItemMallLog{})
}

// WebItemMallLog Power-up store purchase log
type WebItemMallLog struct {
	AccountId   [21]byte  `gorm:"column:strAccountID;type:char(21);not null" json:"strAccountID"`
	CharId      [21]byte  `gorm:"column:strCharID;type:char(21);not null" json:"strCharID"`
	ServerId    int16     `gorm:"column:ServerNo;type:smallint;not null" json:"ServerNo"`
	ItemId      int       `gorm:"column:ItemID;type:int;not null" json:"ItemID"`
	ItemCount   int16     `gorm:"column:ItemCount;type:smallint;not null" json:"ItemCount"`
	BuyTime     int       `gorm:"column:BuyTime;type:smalldatetime;not null;default:getdate()" json:"BuyTime"`
	ImgFileName [50]byte  `gorm:"column:img_file_name;type:varchar(50)" json:"img_file_name,omitempty"`
	ItemName    [100]byte `gorm:"column:strItemName;type:varchar(100)" json:"strItemName,omitempty"`
	Price       *int      `gorm:"column:price;type:int" json:"price,omitempty"`
	PayType     *int      `gorm:"column:pay_type;type:int" json:"pay_type,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *WebItemMallLog) GetDatabaseName() string {
	return GetDatabaseName(DbType(_WebItemMallLogDatabaseNbr))
}

// GetTableName Returns the table name
func (this *WebItemMallLog) GetTableName() string {
	return _WebItemMallLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *WebItemMallLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WEB_ITEMMALL_LOG] (strAccountID, strCharID, ServerNo, ItemID, ItemCount, BuyTime, img_file_name, strItemName, price, pay_type) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.AccountId),
		GetOptionalBinaryVal(this.CharId),
		GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ItemId),
		GetOptionalDecVal(&this.ItemCount),
		GetOptionalDecVal(&this.BuyTime),
		GetOptionalBinaryVal(this.ImgFileName),
		GetOptionalBinaryVal(this.ItemName),
		GetOptionalDecVal(this.Price),
		GetOptionalDecVal(this.PayType))
}

// GetCreateTableString Returns the create table statement for this object
func (this *WebItemMallLog) GetCreateTableString() string {
	query := "CREATE TABLE [WEB_ITEMMALL_LOG] (\n\t\"strAccountID\" char(21) NOT NULL,\n\t\"strCharID\" char(21) NOT NULL,\n\t\"ServerNo\" smallint NOT NULL,\n\t\"ItemID\" int NOT NULL,\n\t\"ItemCount\" smallint NOT NULL,\n\t\"BuyTime\" smalldatetime NOT NULL,\n\t\"img_file_name\" varchar(50),\n\t\"strItemName\" varchar(100),\n\t\"price\" int,\n\t\"pay_type\" int\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
