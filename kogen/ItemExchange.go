package kogen

import (
	"fmt"
)

const (
	_ItemExchangeDatabaseNbr = 0
	_ItemExchangeTableName   = "ITEM_EXCHANGE"
)

func init() {
	ModelList = append(ModelList, &ItemExchange{})
}

// ItemExchange Enables players to be able to give items in exchange for an item from an NPC
type ItemExchange struct {
	Index               int       `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	NpcNumber           int16     `gorm:"column:nNpcNum;type:smallint;not null" json:"nNpcNum"`
	NpcName             [50]byte  `gorm:"column:strNpcName;type:varchar(50)" json:"strNpcName,omitempty"`
	Note                [100]byte `gorm:"column:strNote;type:varchar(100)" json:"strNote,omitempty"`
	RandomFlag          uint8     `gorm:"column:bRandomFlag;type:tinyint;not null" json:"bRandomFlag"`
	OriginItemNumber1   int       `gorm:"column:nOriginItemNum1;type:int;not null" json:"nOriginItemNum1"`
	OriginItemCount1    int16     `gorm:"column:nOriginItemCount1;type:smallint;not null" json:"nOriginItemCount1"`
	OriginItemNumber2   int       `gorm:"column:nOriginItemNum2;type:int;not null" json:"nOriginItemNum2"`
	OriginItemCount2    int16     `gorm:"column:nOriginItemCount2;type:smallint;not null" json:"nOriginItemCount2"`
	OriginItemNumber3   int       `gorm:"column:nOriginItemNum3;type:int;not null" json:"nOriginItemNum3"`
	OriginItemCount3    int16     `gorm:"column:nOriginItemCount3;type:smallint;not null" json:"nOriginItemCount3"`
	OriginItemNumber4   int       `gorm:"column:nOriginItemNum4;type:int;not null" json:"nOriginItemNum4"`
	OriginItemCount4    int16     `gorm:"column:nOriginItemCount4;type:smallint;not null" json:"nOriginItemCount4"`
	OriginItemNumber5   int       `gorm:"column:nOriginItemNum5;type:int;not null" json:"nOriginItemNum5"`
	OriginItemCount5    int16     `gorm:"column:nOriginItemCount5;type:smallint;not null" json:"nOriginItemCount5"`
	ExchangeItemNumber1 int       `gorm:"column:nExchangeItemNum1;type:int;not null" json:"nExchangeItemNum1"`
	ExchangeItemCount1  int16     `gorm:"column:nExchangeItemCount1;type:smallint;not null" json:"nExchangeItemCount1"`
	ExchangeItemNumber2 int       `gorm:"column:nExchangeItemNum2;type:int;not null" json:"nExchangeItemNum2"`
	ExchangeItemCount2  int16     `gorm:"column:nExchangeItemCount2;type:smallint;not null" json:"nExchangeItemCount2"`
	ExchangeItemNumber3 int       `gorm:"column:nExchangeItemNum3;type:int;not null" json:"nExchangeItemNum3"`
	ExchangeItemCount3  int16     `gorm:"column:nExchangeItemCount3;type:smallint;not null" json:"nExchangeItemCount3"`
	ExchangeItemNumber4 int       `gorm:"column:nExchangeItemNum4;type:int;not null" json:"nExchangeItemNum4"`
	ExchangeItemCount4  int16     `gorm:"column:nExchangeItemCount4;type:smallint;not null" json:"nExchangeItemCount4"`
	ExchangeItemNumber5 int       `gorm:"column:nExchangeItemNum5;type:int;not null" json:"nExchangeItemNum5"`
	ExchangeItemCount5  int16     `gorm:"column:nExchangeItemCount5;type:smallint;not null" json:"nExchangeItemCount5"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ItemExchange) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemExchangeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ItemExchange) GetTableName() string {
	return _ItemExchangeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ItemExchange) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM_EXCHANGE] (nIndex, nNpcNum, strNpcName, strNote, bRandomFlag, nOriginItemNum1, nOriginItemCount1, nOriginItemNum2, nOriginItemCount2, nOriginItemNum3, nOriginItemCount3, nOriginItemNum4, nOriginItemCount4, nOriginItemNum5, nOriginItemCount5, nExchangeItemNum1, nExchangeItemCount1, nExchangeItemNum2, nExchangeItemCount2, nExchangeItemNum3, nExchangeItemCount3, nExchangeItemNum4, nExchangeItemCount4, nExchangeItemNum5, nExchangeItemCount5) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcNumber),
		GetOptionalBinaryVal(this.NpcName),
		GetOptionalBinaryVal(this.Note),
		GetOptionalDecVal(&this.RandomFlag),
		GetOptionalDecVal(&this.OriginItemNumber1),
		GetOptionalDecVal(&this.OriginItemCount1),
		GetOptionalDecVal(&this.OriginItemNumber2),
		GetOptionalDecVal(&this.OriginItemCount2),
		GetOptionalDecVal(&this.OriginItemNumber3),
		GetOptionalDecVal(&this.OriginItemCount3),
		GetOptionalDecVal(&this.OriginItemNumber4),
		GetOptionalDecVal(&this.OriginItemCount4),
		GetOptionalDecVal(&this.OriginItemNumber5),
		GetOptionalDecVal(&this.OriginItemCount5),
		GetOptionalDecVal(&this.ExchangeItemNumber1),
		GetOptionalDecVal(&this.ExchangeItemCount1),
		GetOptionalDecVal(&this.ExchangeItemNumber2),
		GetOptionalDecVal(&this.ExchangeItemCount2),
		GetOptionalDecVal(&this.ExchangeItemNumber3),
		GetOptionalDecVal(&this.ExchangeItemCount3),
		GetOptionalDecVal(&this.ExchangeItemNumber4),
		GetOptionalDecVal(&this.ExchangeItemCount4),
		GetOptionalDecVal(&this.ExchangeItemNumber5),
		GetOptionalDecVal(&this.ExchangeItemCount5))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ItemExchange) GetCreateTableString() string {
	query := "CREATE TABLE [ITEM_EXCHANGE] (\n\t[nIndex] int NOT NULL,\n\t[nNpcNum] smallint NOT NULL,\n\t[strNpcName] varchar(50),\n\t[strNote] varchar(100),\n\t[bRandomFlag] tinyint NOT NULL,\n\t[nOriginItemNum1] int NOT NULL,\n\t[nOriginItemCount1] smallint NOT NULL,\n\t[nOriginItemNum2] int NOT NULL,\n\t[nOriginItemCount2] smallint NOT NULL,\n\t[nOriginItemNum3] int NOT NULL,\n\t[nOriginItemCount3] smallint NOT NULL,\n\t[nOriginItemNum4] int NOT NULL,\n\t[nOriginItemCount4] smallint NOT NULL,\n\t[nOriginItemNum5] int NOT NULL,\n\t[nOriginItemCount5] smallint NOT NULL,\n\t[nExchangeItemNum1] int NOT NULL,\n\t[nExchangeItemCount1] smallint NOT NULL,\n\t[nExchangeItemNum2] int NOT NULL,\n\t[nExchangeItemCount2] smallint NOT NULL,\n\t[nExchangeItemNum3] int NOT NULL,\n\t[nExchangeItemCount3] smallint NOT NULL,\n\t[nExchangeItemNum4] int NOT NULL,\n\t[nExchangeItemCount4] smallint NOT NULL,\n\t[nExchangeItemNum5] int NOT NULL,\n\t[nExchangeItemCount5] smallint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
