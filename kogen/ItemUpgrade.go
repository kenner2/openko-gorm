package kogen

import (
	"fmt"
)

const (
	_ItemUpgradeDatabaseNbr = 0
	_ItemUpgradeTableName   = "ITEM_UPGRADE"
)

func init() {
	ModelList = append(ModelList, &ItemUpgrade{})
}

// ItemUpgrade Item upgrade configuration
type ItemUpgrade struct {
	Index        int       `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	NpcNumber    int16     `gorm:"column:nNPCNum;type:smallint;not null" json:"nNPCNum"`
	Name         [50]byte  `gorm:"column:strName;type:varchar(50)" json:"strName,omitempty"`
	Note         [100]byte `gorm:"column:strNote;type:varchar(100)" json:"strNote,omitempty"`
	OriginType   int16     `gorm:"column:nOriginType;type:smallint;not null" json:"nOriginType"`
	OriginItem   int16     `gorm:"column:nOriginItem;type:smallint;not null" json:"nOriginItem"`
	RequireItem1 int       `gorm:"column:nReqItem1;type:int;not null" json:"nReqItem1"`
	RequireItem2 int       `gorm:"column:nReqItem2;type:int;not null" json:"nReqItem2"`
	RequireItem3 int       `gorm:"column:nReqItem3;type:int;not null" json:"nReqItem3"`
	RequireItem4 int       `gorm:"column:nReqItem4;type:int;not null" json:"nReqItem4"`
	RequireItem5 int       `gorm:"column:nReqItem5;type:int;not null" json:"nReqItem5"`
	RequireItem6 int       `gorm:"column:nReqItem6;type:int;not null" json:"nReqItem6"`
	RequireItem7 int       `gorm:"column:nReqItem7;type:int;not null" json:"nReqItem7"`
	RequireItem8 int       `gorm:"column:nReqItem8;type:int;not null" json:"nReqItem8"`
	RequireCoin  int       `gorm:"column:nReqNoah;type:int;not null" json:"nReqNoah"`
	RateType     uint8     `gorm:"column:bRateType;type:tinyint;not null" json:"bRateType"`
	GenRate      int16     `gorm:"column:nGenRate;type:smallint;not null" json:"nGenRate"`
	GiveItem     int16     `gorm:"column:nGiveItem;type:smallint;not null" json:"nGiveItem"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ItemUpgrade) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemUpgradeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ItemUpgrade) GetTableName() string {
	return _ItemUpgradeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ItemUpgrade) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM_UPGRADE] (nIndex, nNPCNum, strName, strNote, nOriginType, nOriginItem, nReqItem1, nReqItem2, nReqItem3, nReqItem4, nReqItem5, nReqItem6, nReqItem7, nReqItem8, nReqNoah, bRateType, nGenRate, nGiveItem) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Note),
		GetOptionalDecVal(&this.OriginType),
		GetOptionalDecVal(&this.OriginItem),
		GetOptionalDecVal(&this.RequireItem1),
		GetOptionalDecVal(&this.RequireItem2),
		GetOptionalDecVal(&this.RequireItem3),
		GetOptionalDecVal(&this.RequireItem4),
		GetOptionalDecVal(&this.RequireItem5),
		GetOptionalDecVal(&this.RequireItem6),
		GetOptionalDecVal(&this.RequireItem7),
		GetOptionalDecVal(&this.RequireItem8),
		GetOptionalDecVal(&this.RequireCoin),
		GetOptionalDecVal(&this.RateType),
		GetOptionalDecVal(&this.GenRate),
		GetOptionalDecVal(&this.GiveItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ItemUpgrade) GetCreateTableString() string {
	query := "CREATE TABLE [ITEM_UPGRADE] (\n\t[nIndex] int NOT NULL,\n\t[nNPCNum] smallint NOT NULL,\n\t[strName] varchar(50),\n\t[strNote] varchar(100),\n\t[nOriginType] smallint NOT NULL,\n\t[nOriginItem] smallint NOT NULL,\n\t[nReqItem1] int NOT NULL,\n\t[nReqItem2] int NOT NULL,\n\t[nReqItem3] int NOT NULL,\n\t[nReqItem4] int NOT NULL,\n\t[nReqItem5] int NOT NULL,\n\t[nReqItem6] int NOT NULL,\n\t[nReqItem7] int NOT NULL,\n\t[nReqItem8] int NOT NULL,\n\t[nReqNoah] int NOT NULL,\n\t[bRateType] tinyint NOT NULL,\n\t[nGenRate] smallint NOT NULL,\n\t[nGiveItem] smallint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
