package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_CouponSerialListDatabaseNbr = "GAME"
	_CouponSerialListTableName   = "COUPON_SERIAL_LIST"
)

func init() {
	ModelList = append(ModelList, &CouponSerialList{})
}

// CouponSerialList Coupon Serial List
type CouponSerialList struct {
	Index      int           `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	SerialNum  mssql.VarChar `gorm:"column:strSerialNum;type:varchar(16);not null" json:"strSerialNum"`
	ItemNumber int           `gorm:"column:nItemNum;type:int;not null" json:"nItemNum"`
	ItemCount  int16         `gorm:"column:sItemCount;type:smallint;not null" json:"sItemCount"`
}

// GetDatabaseName Returns the table's database name
func (this CouponSerialList) GetDatabaseName() string {
	return GetDatabaseName(_CouponSerialListDatabaseNbr)
}

// TableName Returns the table name
func (this CouponSerialList) TableName() string {
	return _CouponSerialListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this CouponSerialList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COUPON_SERIAL_LIST] ([nIndex], [strSerialNum], [nItemNum], [sItemCount]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(&this.SerialNum, false),
		GetOptionalDecVal(&this.ItemNumber),
		GetOptionalDecVal(&this.ItemCount))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this CouponSerialList) GetInsertHeader() string {
	return "INSERT INTO [COUPON_SERIAL_LIST] ([nIndex], [strSerialNum], [nItemNum], [sItemCount]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this CouponSerialList) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalVarCharVal(&this.SerialNum, false),
		GetOptionalDecVal(&this.ItemNumber),
		GetOptionalDecVal(&this.ItemCount))
}

// GetCreateTableString Returns the create table statement for this object
func (this CouponSerialList) GetCreateTableString() string {
	query := "CREATE TABLE [COUPON_SERIAL_LIST] (\n\t[nIndex] int NOT NULL,\n\t[strSerialNum] varchar(16) NOT NULL,\n\t[nItemNum] int NOT NULL,\n\t[sItemCount] smallint NOT NULL\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this CouponSerialList) SelectClause() (selectClause clause.Select) {
	return _CouponSerialListSelectClause
}

// GetAllTableData Returns a list of all table data
func (this CouponSerialList) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []CouponSerialList{}
	rawSql := "SELECT [nIndex], [strSerialNum], [nItemNum], [sItemCount] FROM [COUPON_SERIAL_LIST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _CouponSerialListSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nIndex]",
		},
		clause.Column{
			Name: "[strSerialNum]",
		},
		clause.Column{
			Name: "[nItemNum]",
		},
		clause.Column{
			Name: "[sItemCount]",
		},
	},
}
