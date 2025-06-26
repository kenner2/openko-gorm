package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_CopySerialItemDatabaseNbr = 1
	_CopySerialItemTableName   = "COPY_SERIAL_ITEM"
)

func init() {
	ModelList = append(ModelList, &CopySerialItem{})
}

// CopySerialItem TODO: Doc
type CopySerialItem struct {
	UserId     *[]byte `gorm:"column:strUserId;type:char(21)" json:"strUserId,omitempty"`
	Type       *uint8  `gorm:"column:byType;type:tinyint" json:"byType,omitempty"`
	Pos        *int16  `gorm:"column:nPos;type:smallint" json:"nPos,omitempty"`
	ItemNum    *[]byte `gorm:"column:ItemNum;type:binary(4)" json:"ItemNum,omitempty"`
	ItemSerial *[]byte `gorm:"column:ItemSerial;type:binary(8)" json:"ItemSerial,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this CopySerialItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_CopySerialItemDatabaseNbr))
}

// TableName Returns the table name
func (this CopySerialItem) TableName() string {
	return _CopySerialItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this CopySerialItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [COPY_SERIAL_ITEM] ([strUserId], [byType], [nPos], [ItemNum], [ItemSerial]) VALUES\n(%s, %s, %s, CONVERT(binary(4), %s), CONVERT(binary(8), %s))", GetOptionalByteArrayVal(this.UserId, false),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNum),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this CopySerialItem) GetInsertHeader() string {
	return "INSERT INTO [COPY_SERIAL_ITEM] (strUserId, byType, nPos, ItemNum, ItemSerial) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this CopySerialItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, CONVERT(binary(4), %s), CONVERT(binary(8), %s))", GetOptionalByteArrayVal(this.UserId, false),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNum),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this CopySerialItem) GetCreateTableString() string {
	query := "CREATE TABLE [COPY_SERIAL_ITEM] (\n\t[strUserId] char(21),\n\t[byType] tinyint,\n\t[nPos] smallint,\n\t[ItemNum] binary(4),\n\t[ItemSerial] binary(8)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this CopySerialItem) SelectClause() (selectClause clause.Select) {
	return _CopySerialItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this CopySerialItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []CopySerialItem{}
	rawSql := "SELECT [strUserId], [byType], [nPos], CONVERT(VARBINARY(4), [ItemNum]) as [ItemNum], CONVERT(VARBINARY(8), [ItemSerial]) as [ItemSerial] FROM [COPY_SERIAL_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _CopySerialItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strUserId]",
		},
		clause.Column{
			Name: "[byType]",
		},
		clause.Column{
			Name: "[nPos]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(4), [ItemNum]) as [ItemNum]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(8), [ItemSerial]) as [ItemSerial]",
		},
	},
}
