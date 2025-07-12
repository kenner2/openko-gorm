package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ItemUpgradeDatabaseNbr = "GAME"
	_ItemUpgradeTableName   = "ITEM_UPGRADE"
)

func init() {
	ModelList = append(ModelList, &ItemUpgrade{})
}

// ItemUpgrade Item upgrade configuration
type ItemUpgrade struct {
	Index         int           `gorm:"column:nIndex;type:int;primaryKey;not null" json:"nIndex"`
	NpcNumber     int16         `gorm:"column:nNPCNum;type:smallint;not null" json:"nNPCNum"`
	Name          mssql.VarChar `gorm:"column:strName;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strName"`
	Note          mssql.VarChar `gorm:"column:strNote;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strNote"`
	OriginType    int16         `gorm:"column:nOriginType;type:smallint;not null" json:"nOriginType"`
	OriginItem    int16         `gorm:"column:nOriginItem;type:smallint;not null" json:"nOriginItem"`
	RequiredItem1 int           `gorm:"column:nReqItem1;type:int;not null" json:"nReqItem1"`
	RequiredItem2 int           `gorm:"column:nReqItem2;type:int;not null" json:"nReqItem2"`
	RequiredItem3 int           `gorm:"column:nReqItem3;type:int;not null" json:"nReqItem3"`
	RequiredItem4 int           `gorm:"column:nReqItem4;type:int;not null" json:"nReqItem4"`
	RequiredItem5 int           `gorm:"column:nReqItem5;type:int;not null" json:"nReqItem5"`
	RequiredItem6 int           `gorm:"column:nReqItem6;type:int;not null" json:"nReqItem6"`
	RequiredItem7 int           `gorm:"column:nReqItem7;type:int;not null" json:"nReqItem7"`
	RequiredItem8 int           `gorm:"column:nReqItem8;type:int;not null" json:"nReqItem8"`
	RequiredCoins int           `gorm:"column:nReqNoah;type:int;not null" json:"nReqNoah"`
	RateType      uint8         `gorm:"column:bRateType;type:tinyint;not null" json:"bRateType"`
	GenRate       int16         `gorm:"column:nGenRate;type:smallint;not null" json:"nGenRate"`
	GiveItem      int16         `gorm:"column:nGiveItem;type:smallint;not null" json:"nGiveItem"`
}

// GetDatabaseName Returns the table's database name
func (this ItemUpgrade) GetDatabaseName() string {
	return GetDatabaseName(_ItemUpgradeDatabaseNbr)
}

// TableName Returns the table name
func (this ItemUpgrade) TableName() string {
	return _ItemUpgradeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ItemUpgrade) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM_UPGRADE] ([nIndex], [nNPCNum], [strName], [strNote], [nOriginType], [nOriginItem], [nReqItem1], [nReqItem2], [nReqItem3], [nReqItem4], [nReqItem5], [nReqItem6], [nReqItem7], [nReqItem8], [nReqNoah], [bRateType], [nGenRate], [nGiveItem]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcNumber),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(&this.Note, false),
		GetOptionalDecVal(&this.OriginType),
		GetOptionalDecVal(&this.OriginItem),
		GetOptionalDecVal(&this.RequiredItem1),
		GetOptionalDecVal(&this.RequiredItem2),
		GetOptionalDecVal(&this.RequiredItem3),
		GetOptionalDecVal(&this.RequiredItem4),
		GetOptionalDecVal(&this.RequiredItem5),
		GetOptionalDecVal(&this.RequiredItem6),
		GetOptionalDecVal(&this.RequiredItem7),
		GetOptionalDecVal(&this.RequiredItem8),
		GetOptionalDecVal(&this.RequiredCoins),
		GetOptionalDecVal(&this.RateType),
		GetOptionalDecVal(&this.GenRate),
		GetOptionalDecVal(&this.GiveItem))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ItemUpgrade) GetInsertHeader() string {
	return "INSERT INTO [ITEM_UPGRADE] ([nIndex], [nNPCNum], [strName], [strNote], [nOriginType], [nOriginItem], [nReqItem1], [nReqItem2], [nReqItem3], [nReqItem4], [nReqItem5], [nReqItem6], [nReqItem7], [nReqItem8], [nReqNoah], [bRateType], [nGenRate], [nGiveItem]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ItemUpgrade) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcNumber),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalVarCharVal(&this.Note, false),
		GetOptionalDecVal(&this.OriginType),
		GetOptionalDecVal(&this.OriginItem),
		GetOptionalDecVal(&this.RequiredItem1),
		GetOptionalDecVal(&this.RequiredItem2),
		GetOptionalDecVal(&this.RequiredItem3),
		GetOptionalDecVal(&this.RequiredItem4),
		GetOptionalDecVal(&this.RequiredItem5),
		GetOptionalDecVal(&this.RequiredItem6),
		GetOptionalDecVal(&this.RequiredItem7),
		GetOptionalDecVal(&this.RequiredItem8),
		GetOptionalDecVal(&this.RequiredCoins),
		GetOptionalDecVal(&this.RateType),
		GetOptionalDecVal(&this.GenRate),
		GetOptionalDecVal(&this.GiveItem))
}

// GetCreateTableString Returns the create table statement for this object
func (this ItemUpgrade) GetCreateTableString() string {
	query := "CREATE TABLE [ITEM_UPGRADE] (\n\t[nIndex] int NOT NULL,\n\t[nNPCNum] smallint NOT NULL,\n\t[strName] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[strNote] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[nOriginType] smallint NOT NULL,\n\t[nOriginItem] smallint NOT NULL,\n\t[nReqItem1] int NOT NULL,\n\t[nReqItem2] int NOT NULL,\n\t[nReqItem3] int NOT NULL,\n\t[nReqItem4] int NOT NULL,\n\t[nReqItem5] int NOT NULL,\n\t[nReqItem6] int NOT NULL,\n\t[nReqItem7] int NOT NULL,\n\t[nReqItem8] int NOT NULL,\n\t[nReqNoah] int NOT NULL,\n\t[bRateType] tinyint NOT NULL,\n\t[nGenRate] smallint NOT NULL,\n\t[nGiveItem] smallint NOT NULL\n\tCONSTRAINT [PK_ITEM_UPGRADE] PRIMARY KEY CLUSTERED ([nIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ItemUpgrade) SelectClause() (selectClause clause.Select) {
	return _ItemUpgradeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this ItemUpgrade) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ItemUpgrade{}
	rawSql := "SELECT [nIndex], [nNPCNum], [strName], [strNote], [nOriginType], [nOriginItem], [nReqItem1], [nReqItem2], [nReqItem3], [nReqItem4], [nReqItem5], [nReqItem6], [nReqItem7], [nReqItem8], [nReqNoah], [bRateType], [nGenRate], [nGiveItem] FROM [ITEM_UPGRADE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ItemUpgradeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nIndex]",
		},
		clause.Column{
			Name: "[nNPCNum]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[strNote]",
		},
		clause.Column{
			Name: "[nOriginType]",
		},
		clause.Column{
			Name: "[nOriginItem]",
		},
		clause.Column{
			Name: "[nReqItem1]",
		},
		clause.Column{
			Name: "[nReqItem2]",
		},
		clause.Column{
			Name: "[nReqItem3]",
		},
		clause.Column{
			Name: "[nReqItem4]",
		},
		clause.Column{
			Name: "[nReqItem5]",
		},
		clause.Column{
			Name: "[nReqItem6]",
		},
		clause.Column{
			Name: "[nReqItem7]",
		},
		clause.Column{
			Name: "[nReqItem8]",
		},
		clause.Column{
			Name: "[nReqNoah]",
		},
		clause.Column{
			Name: "[bRateType]",
		},
		clause.Column{
			Name: "[nGenRate]",
		},
		clause.Column{
			Name: "[nGiveItem]",
		},
	},
}
