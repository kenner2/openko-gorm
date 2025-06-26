package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_TermItemDatabaseNbr = 1
	_TermItemTableName   = "TERM_ITEM"
)

func init() {
	ModelList = append(ModelList, &TermItem{})
}

// TermItem Term item
type TermItem struct {
	UserId     *[]byte `gorm:"column:strUserId;type:char(21)" json:"strUserId,omitempty"`
	Type       *uint8  `gorm:"column:byType;type:tinyint" json:"byType,omitempty"`
	Pos        *int16  `gorm:"column:nPos;type:smallint" json:"nPos,omitempty"`
	ItemNumber *[]byte `gorm:"column:ItemNum;type:binary(4)" json:"ItemNum,omitempty"`
	ItemSerial *[]byte `gorm:"column:ItemSerial;type:binary(8)" json:"ItemSerial,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this TermItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_TermItemDatabaseNbr))
}

// TableName Returns the table name
func (this TermItem) TableName() string {
	return _TermItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this TermItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [TERM_ITEM] ([strUserId], [byType], [nPos], [ItemNum], [ItemSerial]) VALUES\n(%s, %s, %s, CONVERT(binary(4), %s), CONVERT(binary(8), %s))", GetOptionalByteArrayVal(this.UserId, false),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNumber),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this TermItem) GetInsertHeader() string {
	return "INSERT INTO [TERM_ITEM] (strUserId, byType, nPos, ItemNum, ItemSerial) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this TermItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, CONVERT(binary(4), %s), CONVERT(binary(8), %s))", GetOptionalByteArrayVal(this.UserId, false),
		GetOptionalDecVal(this.Type),
		GetOptionalDecVal(this.Pos),
		GetOptionalBinaryVal(this.ItemNumber),
		GetOptionalBinaryVal(this.ItemSerial))
}

// GetCreateTableString Returns the create table statement for this object
func (this TermItem) GetCreateTableString() string {
	query := "CREATE TABLE [TERM_ITEM] (\n\t[strUserId] char(21),\n\t[byType] tinyint,\n\t[nPos] smallint,\n\t[ItemNum] binary(4),\n\t[ItemSerial] binary(8)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this TermItem) SelectClause() (selectClause clause.Select) {
	return _TermItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this TermItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []TermItem{}
	rawSql := "SELECT [strUserId], [byType], [nPos], CONVERT(VARBINARY(4), [ItemNum]) as [ItemNum], CONVERT(VARBINARY(8), [ItemSerial]) as [ItemSerial] FROM [TERM_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _TermItemSelectClause = clause.Select{
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
