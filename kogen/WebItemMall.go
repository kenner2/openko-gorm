package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_WebItemMallDatabaseNbr = "GAME"
	_WebItemMallTableName   = "WEB_ITEMMALL"
)

func init() {
	ModelList = append(ModelList, &WebItemMall{})
}

// WebItemMall Power-up store purchases
type WebItemMall struct {
	AccountId   mssql.VarChar  `gorm:"column:strAccountID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strAccountID"`
	CharId      mssql.VarChar  `gorm:"column:strCharID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strCharID"`
	ServerId    int16          `gorm:"column:ServerNo;type:smallint;not null" json:"ServerNo"`
	ItemId      int            `gorm:"column:ItemID;type:int;not null" json:"ItemID"`
	ItemCount   int16          `gorm:"column:ItemCount;type:smallint;not null;default:1" json:"ItemCount"`
	BuyTime     time.Time      `gorm:"column:BuyTime;type:smalldatetime;not null;default:getdate()" json:"BuyTime"`
	ImgFileName *mssql.VarChar `gorm:"column:img_file_name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"img_file_name,omitempty"`
	ItemName    *mssql.VarChar `gorm:"column:strItemName;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strItemName,omitempty"`
	Price       *int           `gorm:"column:price;type:int" json:"price,omitempty"`
	PayType     *int           `gorm:"column:pay_type;type:int" json:"pay_type,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this WebItemMall) GetDatabaseName() string {
	return GetDatabaseName(_WebItemMallDatabaseNbr)
}

// TableName Returns the table name
func (this WebItemMall) TableName() string {
	return _WebItemMallTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this WebItemMall) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WEB_ITEMMALL] ([strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ItemId),
		GetOptionalDecVal(&this.ItemCount),
		GetDateTimeExportFmt(&this.BuyTime),
		GetOptionalVarCharVal(this.ImgFileName, false),
		GetOptionalVarCharVal(this.ItemName, false),
		GetOptionalDecVal(this.Price),
		GetOptionalDecVal(this.PayType))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this WebItemMall) GetInsertHeader() string {
	return "INSERT INTO [WEB_ITEMMALL] ([strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this WebItemMall) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
		GetOptionalVarCharVal(&this.CharId, false),
		GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ItemId),
		GetOptionalDecVal(&this.ItemCount),
		GetDateTimeExportFmt(&this.BuyTime),
		GetOptionalVarCharVal(this.ImgFileName, false),
		GetOptionalVarCharVal(this.ItemName, false),
		GetOptionalDecVal(this.Price),
		GetOptionalDecVal(this.PayType))
}

// GetCreateTableString Returns the create table statement for this object
func (this WebItemMall) GetCreateTableString() string {
	query := "CREATE TABLE [WEB_ITEMMALL] (\n\t[strAccountID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strCharID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[ServerNo] smallint NOT NULL,\n\t[ItemID] int NOT NULL,\n\t[ItemCount] smallint NOT NULL,\n\t[BuyTime] smalldatetime NOT NULL,\n\t[img_file_name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[strItemName] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[price] int,\n\t[pay_type] int\n)\nGO\nALTER TABLE [WEB_ITEMMALL] ADD CONSTRAINT [DF_WEB_ITEMMALL_ItemCount] DEFAULT 1 FOR [ItemCount]\nGO\nALTER TABLE [WEB_ITEMMALL] ADD CONSTRAINT [DF_WEB_ITEMMALL_BuyTime] DEFAULT getdate() FOR [BuyTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this WebItemMall) SelectClause() (selectClause clause.Select) {
	return _WebItemMallSelectClause
}

// GetAllTableData Returns a list of all table data
func (this WebItemMall) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []WebItemMall{}
	rawSql := "SELECT [strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type] FROM [WEB_ITEMMALL]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _WebItemMallSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strAccountID]",
		},
		clause.Column{
			Name: "[strCharID]",
		},
		clause.Column{
			Name: "[ServerNo]",
		},
		clause.Column{
			Name: "[ItemID]",
		},
		clause.Column{
			Name: "[ItemCount]",
		},
		clause.Column{
			Name: "[BuyTime]",
		},
		clause.Column{
			Name: "[img_file_name]",
		},
		clause.Column{
			Name: "[strItemName]",
		},
		clause.Column{
			Name: "[price]",
		},
		clause.Column{
			Name: "[pay_type]",
		},
	},
}
