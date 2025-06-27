package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_RentalItemDatabaseNbr = "GAME"
	_RentalItemTableName   = "RENTAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &RentalItem{})
}

// RentalItem Rental item
type RentalItem struct {
	RentalIndex       int            `gorm:"column:nRentalIndex;type:int;primaryKey;not null" json:"nRentalIndex"`
	ItemIndex         int            `gorm:"column:nItemIndex;type:int;not null" json:"nItemIndex"`
	Durability        int16          `gorm:"column:sDurability;type:smallint;not null;default:0" json:"sDurability"`
	SerialNumber      int64          `gorm:"column:nSerialNumber;type:bigint;not null" json:"nSerialNumber"`
	RegType           uint8          `gorm:"column:byRegType;type:tinyint;not null;default:0" json:"byRegType"`
	ItemType          uint8          `gorm:"column:byItemType;type:tinyint;not null" json:"byItemType"`
	Class             uint8          `gorm:"column:byClass;type:tinyint;not null" json:"byClass"`
	RentalTime        int16          `gorm:"column:sRentalTime;type:smallint;not null" json:"sRentalTime"`
	RentalMoney       int            `gorm:"column:nRentalMoney;type:int;not null" json:"nRentalMoney"`
	LenderCharId      mssql.VarChar  `gorm:"column:strLenderCharID;type:varchar(21);not null" json:"strLenderCharID"`
	LenderAccountId   mssql.VarChar  `gorm:"column:strLenderAcID;type:varchar(21);not null" json:"strLenderAcID"`
	BorrowerCharId    *mssql.VarChar `gorm:"column:strBorrowerCharID;type:varchar(21)" json:"strBorrowerCharID,omitempty"`
	BorrowerAccountId *mssql.VarChar `gorm:"column:strBorrowerAcID;type:varchar(21)" json:"strBorrowerAcID,omitempty"`
	LendTime          *time.Time     `gorm:"column:timeLender;type:smalldatetime" json:"timeLender,omitempty"`
	RegisterTime      time.Time      `gorm:"column:timeRegister;type:smalldatetime;not null;default:getdate()" json:"timeRegister"`
}

// GetDatabaseName Returns the table's database name
func (this RentalItem) GetDatabaseName() string {
	return GetDatabaseName(_RentalItemDatabaseNbr)
}

// TableName Returns the table name
func (this RentalItem) TableName() string {
	return _RentalItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this RentalItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [RENTAL_ITEM] ([nRentalIndex], [nItemIndex], [sDurability], [nSerialNumber], [byRegType], [byItemType], [byClass], [sRentalTime], [nRentalMoney], [strLenderCharID], [strLenderAcID], [strBorrowerCharID], [strBorrowerAcID], [timeLender], [timeRegister]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RegType),
		GetOptionalDecVal(&this.ItemType),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalVarCharVal(&this.LenderCharId, false),
		GetOptionalVarCharVal(&this.LenderAccountId, false),
		GetOptionalVarCharVal(this.BorrowerCharId, false),
		GetOptionalVarCharVal(this.BorrowerAccountId, false),
		GetDateTimeExportFmt(this.LendTime),
		GetDateTimeExportFmt(&this.RegisterTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this RentalItem) GetInsertHeader() string {
	return "INSERT INTO [RENTAL_ITEM] ([nRentalIndex], [nItemIndex], [sDurability], [nSerialNumber], [byRegType], [byItemType], [byClass], [sRentalTime], [nRentalMoney], [strLenderCharID], [strLenderAcID], [strBorrowerCharID], [strBorrowerAcID], [timeLender], [timeRegister]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this RentalItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.RentalIndex),
		GetOptionalDecVal(&this.ItemIndex),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.SerialNumber),
		GetOptionalDecVal(&this.RegType),
		GetOptionalDecVal(&this.ItemType),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.RentalTime),
		GetOptionalDecVal(&this.RentalMoney),
		GetOptionalVarCharVal(&this.LenderCharId, false),
		GetOptionalVarCharVal(&this.LenderAccountId, false),
		GetOptionalVarCharVal(this.BorrowerCharId, false),
		GetOptionalVarCharVal(this.BorrowerAccountId, false),
		GetDateTimeExportFmt(this.LendTime),
		GetDateTimeExportFmt(&this.RegisterTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this RentalItem) GetCreateTableString() string {
	query := "CREATE TABLE [RENTAL_ITEM] (\n\t[nRentalIndex] int NOT NULL,\n\t[nItemIndex] int NOT NULL,\n\t[sDurability] smallint NOT NULL,\n\t[nSerialNumber] bigint NOT NULL,\n\t[byRegType] tinyint NOT NULL,\n\t[byItemType] tinyint NOT NULL,\n\t[byClass] tinyint NOT NULL,\n\t[sRentalTime] smallint NOT NULL,\n\t[nRentalMoney] int NOT NULL,\n\t[strLenderCharID] varchar(21) NOT NULL,\n\t[strLenderAcID] varchar(21) NOT NULL,\n\t[strBorrowerCharID] varchar(21),\n\t[strBorrowerAcID] varchar(21),\n\t[timeLender] smalldatetime,\n\t[timeRegister] smalldatetime NOT NULL\n\tCONSTRAINT [PK_RENTAL_ITEM] PRIMARY KEY CLUSTERED ([nRentalIndex])\n)\nGO\nALTER TABLE [RENTAL_ITEM] ADD CONSTRAINT [DF_RENTAL_ITEM_sDurability] DEFAULT 0 FOR [sDurability]\nGO\nALTER TABLE [RENTAL_ITEM] ADD CONSTRAINT [DF_RENTAL_ITEM_byRegType] DEFAULT 0 FOR [byRegType]\nGO\nALTER TABLE [RENTAL_ITEM] ADD CONSTRAINT [DF_RENTAL_ITEM_timeRegister] DEFAULT getdate() FOR [timeRegister]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this RentalItem) SelectClause() (selectClause clause.Select) {
	return _RentalItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this RentalItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []RentalItem{}
	rawSql := "SELECT [nRentalIndex], [nItemIndex], [sDurability], [nSerialNumber], [byRegType], [byItemType], [byClass], [sRentalTime], [nRentalMoney], [strLenderCharID], [strLenderAcID], [strBorrowerCharID], [strBorrowerAcID], [timeLender], [timeRegister] FROM [RENTAL_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _RentalItemSelectClause = clause.Select{
	Columns: []clause.Column{
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
			Name: "[byRegType]",
		},
		clause.Column{
			Name: "[byItemType]",
		},
		clause.Column{
			Name: "[byClass]",
		},
		clause.Column{
			Name: "[sRentalTime]",
		},
		clause.Column{
			Name: "[nRentalMoney]",
		},
		clause.Column{
			Name: "[strLenderCharID]",
		},
		clause.Column{
			Name: "[strLenderAcID]",
		},
		clause.Column{
			Name: "[strBorrowerCharID]",
		},
		clause.Column{
			Name: "[strBorrowerAcID]",
		},
		clause.Column{
			Name: "[timeLender]",
		},
		clause.Column{
			Name: "[timeRegister]",
		},
	},
}
