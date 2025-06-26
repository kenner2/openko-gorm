package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_UserRentalItemDatabaseNbr = 1
	_UserRentalItemTableName   = "USER_RENTAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &UserRentalItem{})
}

// UserRentalItem User rental item
type UserRentalItem struct {
	UserId          mssql.VarChar `gorm:"column:strUserID;type:varchar(50);not null" json:"strUserID"`
	AccountId       mssql.VarChar `gorm:"column:strAccountID;type:varchar(50);not null" json:"strAccountID"`
	RentalType      uint8         `gorm:"column:byRentalType;type:tinyint;not null" json:"byRentalType"`
	RegTime         uint8         `gorm:"column:byRegType;type:tinyint;not null;default:0" json:"byRegType"`
	RentalIndex     int           `gorm:"column:nRentalIndex;type:int;not null" json:"nRentalIndex"`
	ItemIndex       int           `gorm:"column:nItemIndex;type:int;not null" json:"nItemIndex"`
	Durability      int16         `gorm:"column:sDurability;type:smallint;not null;default:0" json:"sDurability"`
	SerialNumber    int64         `gorm:"column:nSerialNumber;type:bigint;not null" json:"nSerialNumber"`
	RentalMoney     int           `gorm:"column:nRentalMoney;type:int;not null" json:"nRentalMoney"`
	RentalTime      int16         `gorm:"column:sRentalTime;type:smallint;not null" json:"sRentalTime"`
	DuringTime      int16         `gorm:"column:sDuringTime;type:smallint;not null" json:"sDuringTime"`
	RentalTimestamp *time.Time    `gorm:"column:timeRental;type:smalldatetime" json:"timeRental,omitempty"`
	RegisterTime    *time.Time    `gorm:"column:timeRegister;type:smalldatetime;default:getdate()" json:"timeRegister,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this UserRentalItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserRentalItemDatabaseNbr))
}

// TableName Returns the table name
func (this UserRentalItem) TableName() string {
	return _UserRentalItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserRentalItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USER_RENTAL_ITEM] ([strUserID], [strAccountID], [byRentalType], [byRegType], [nRentalIndex], [nItemIndex], [sDurability], [nSerialNumber], [nRentalMoney], [sRentalTime], [sDuringTime], [timeRental], [timeRegister]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.RentalType),
		GetOptionalDecVal(&this.RegTime),
		GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.DuringTime),
		GetDateTimeExportFmt(this.RentalTimestamp),
		GetDateTimeExportFmt(this.RegisterTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserRentalItem) GetInsertHeader() string {
	return "INSERT INTO [USER_RENTAL_ITEM] (strUserID, strAccountID, byRentalType, byRegType, nRentalIndex, nItemIndex, sDurability, nSerialNumber, nRentalMoney, sRentalTime, sDuringTime, timeRental, timeRegister) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserRentalItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalDecVal(&this.RentalType),
		GetOptionalDecVal(&this.RegTime),
		GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.DuringTime),
		GetDateTimeExportFmt(this.RentalTimestamp),
		GetDateTimeExportFmt(this.RegisterTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserRentalItem) GetCreateTableString() string {
	query := "CREATE TABLE [USER_RENTAL_ITEM] (\n\t[strUserID] varchar(50) NOT NULL,\n\t[strAccountID] varchar(50) NOT NULL,\n\t[byRentalType] tinyint NOT NULL,\n\t[byRegType] tinyint NOT NULL,\n\t[nRentalIndex] int NOT NULL,\n\t[nItemIndex] int NOT NULL,\n\t[sDurability] smallint NOT NULL,\n\t[nSerialNumber] bigint NOT NULL,\n\t[nRentalMoney] int NOT NULL,\n\t[sRentalTime] smallint NOT NULL,\n\t[sDuringTime] smallint NOT NULL,\n\t[timeRental] smalldatetime,\n\t[timeRegister] smalldatetime\n\n)\nGO\nALTER TABLE [USER_RENTAL_ITEM] ADD CONSTRAINT [DF_USER_RENTAL_ITEM_byRegType] DEFAULT 0 FOR [byRegType]\nGO\nALTER TABLE [USER_RENTAL_ITEM] ADD CONSTRAINT [DF_USER_RENTAL_ITEM_sDurability] DEFAULT 0 FOR [sDurability]\nGO\nALTER TABLE [USER_RENTAL_ITEM] ADD CONSTRAINT [DF_USER_RENTAL_ITEM_timeRegister] DEFAULT getdate() FOR [timeRegister]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserRentalItem) SelectClause() (selectClause clause.Select) {
	return _UserRentalItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserRentalItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserRentalItem{}
	rawSql := "SELECT [strUserID], [strAccountID], [byRentalType], [byRegType], [nRentalIndex], [nItemIndex], [sDurability], [nSerialNumber], [nRentalMoney], [sRentalTime], [sDuringTime], [timeRental], [timeRegister] FROM [USER_RENTAL_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserRentalItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strUserID]",
		},
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[byRentalType]",
		},
		clause.Column{
			Name: "[byRegType]",
		},
		clause.Column{
			Name: "[nRentalIndex]",
		},
		clause.Column{
			Name: "[nItemIndex]",
		},
		clause.Column{
			Name: "[sDurability]",
		},
		clause.Column{
			Name: "[nSerialNumber]",
		},
		clause.Column{
			Name: "[nRentalMoney]",
		},
		clause.Column{
			Name: "[sRentalTime]",
		},
		clause.Column{
			Name: "[sDuringTime]",
		},
		clause.Column{
			Name: "[timeRental]",
		},
		clause.Column{
			Name: "[timeRegister]",
		},
	},
}
