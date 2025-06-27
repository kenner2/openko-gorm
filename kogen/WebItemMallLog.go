package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_WebItemMallLogDatabaseNbr = "GAME"
	_WebItemMallLogTableName   = "WEB_ITEMMALL_LOG"
)

func init() {
	ModelList = append(ModelList, &WebItemMallLog{})
}

// WebItemMallLog Power-up store purchase log
type WebItemMallLog struct {
	AccountId   mssql.VarChar  `gorm:"column:strAccountID;type:varchar(21);not null" json:"strAccountID"`
	CharId      mssql.VarChar  `gorm:"column:strCharID;type:varchar(21);not null" json:"strCharID"`
	ServerId    int16          `gorm:"column:ServerNo;type:smallint;not null" json:"ServerNo"`
	ItemId      int            `gorm:"column:ItemID;type:int;not null" json:"ItemID"`
	ItemCount   int16          `gorm:"column:ItemCount;type:smallint;not null" json:"ItemCount"`
	BuyTime     time.Time      `gorm:"column:BuyTime;type:smalldatetime;not null;default:getdate()" json:"BuyTime"`
	ImgFileName *mssql.VarChar `gorm:"column:img_file_name;type:varchar(50)" json:"img_file_name,omitempty"`
	ItemName    *mssql.VarChar `gorm:"column:strItemName;type:varchar(100)" json:"strItemName,omitempty"`
	Price       *int           `gorm:"column:price;type:int" json:"price,omitempty"`
	PayType     *int           `gorm:"column:pay_type;type:int" json:"pay_type,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this WebItemMallLog) GetDatabaseName() string {
	return GetDatabaseName(_WebItemMallLogDatabaseNbr)
}

// TableName Returns the table name
func (this WebItemMallLog) TableName() string {
	return _WebItemMallLogTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this WebItemMallLog) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [WEB_ITEMMALL_LOG] ([strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.AccountId, false),
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
func (this WebItemMallLog) GetInsertHeader() string {
	return "INSERT INTO [WEB_ITEMMALL_LOG] ([strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this WebItemMallLog) GetInsertData() string {
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
func (this WebItemMallLog) GetCreateTableString() string {
	query := "CREATE TABLE [WEB_ITEMMALL_LOG] (\n\t[strAccountID] varchar(21) NOT NULL,\n\t[strCharID] varchar(21) NOT NULL,\n\t[ServerNo] smallint NOT NULL,\n\t[ItemID] int NOT NULL,\n\t[ItemCount] smallint NOT NULL,\n\t[BuyTime] smalldatetime NOT NULL,\n\t[img_file_name] varchar(50),\n\t[strItemName] varchar(100),\n\t[price] int,\n\t[pay_type] int\n)\nGO\nALTER TABLE [WEB_ITEMMALL_LOG] ADD CONSTRAINT [DF_WEB_ITEMMALL_LOG_BuyTime] DEFAULT getdate() FOR [BuyTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this WebItemMallLog) SelectClause() (selectClause clause.Select) {
	return _WebItemMallLogSelectClause
}

// GetAllTableData Returns a list of all table data
func (this WebItemMallLog) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []WebItemMallLog{}
	rawSql := "SELECT [strAccountID], [strCharID], [ServerNo], [ItemID], [ItemCount], [BuyTime], [img_file_name], [strItemName], [price], [pay_type] FROM [WEB_ITEMMALL_LOG]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _WebItemMallLogSelectClause = clause.Select{
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
