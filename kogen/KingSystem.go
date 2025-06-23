package kogen

import (
	"fmt"
)

const (
	_KingSystemDatabaseNbr = 0
	_KingSystemTableName   = "KING_SYSTEM"
)

func init() {
	ModelList = append(ModelList, &KingSystem{})
}

// KingSystem King System
type KingSystem struct {
	Nation            uint8    `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	Type              uint8    `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	Year              int16    `gorm:"column:sYear;type:smallint;not null" json:"sYear"`
	Month             uint8    `gorm:"column:byMonth;type:tinyint;not null" json:"byMonth"`
	Day               uint8    `gorm:"column:byDay;type:tinyint;not null" json:"byDay"`
	Hour              uint8    `gorm:"column:byHour;type:tinyint;not null" json:"byHour"`
	Minute            uint8    `gorm:"column:byMinute;type:tinyint;not null" json:"byMinute"`
	ImType            uint8    `gorm:"column:byImType;type:tinyint;not null" json:"byImType"`
	ImYear            int16    `gorm:"column:sImYear;type:smallint;not null" json:"sImYear"`
	ImMonth           uint8    `gorm:"column:byImMonth;type:tinyint;not null" json:"byImMonth"`
	ImDay             uint8    `gorm:"column:byImDay;type:tinyint;not null" json:"byImDay"`
	ImHour            uint8    `gorm:"column:byImHour;type:tinyint;not null" json:"byImHour"`
	ImMinute          uint8    `gorm:"column:byImMinute;type:tinyint;not null" json:"byImMinute"`
	NoahEvent         uint8    `gorm:"column:byNoahEvent;type:tinyint;not null;default:0" json:"byNoahEvent"`
	NoahEventDay      uint8    `gorm:"column:byNoahEvent_Day;type:tinyint;not null" json:"byNoahEvent_Day"`
	NoahEventHour     uint8    `gorm:"column:byNoahEvent_Hour;type:tinyint;not null" json:"byNoahEvent_Hour"`
	NoahEventMinute   uint8    `gorm:"column:byNoahEvent_Minute;type:tinyint;not null" json:"byNoahEvent_Minute"`
	NoahEventDuration int16    `gorm:"column:sNoahEvent_Duration;type:smallint;not null" json:"sNoahEvent_Duration"`
	ExpEvent          uint8    `gorm:"column:byExpEvent;type:tinyint;not null;default:0" json:"byExpEvent"`
	ExpEventDay       uint8    `gorm:"column:byExpEvent_Day;type:tinyint;not null" json:"byExpEvent_Day"`
	ExpEventHour      uint8    `gorm:"column:byExpEvent_Hour;type:tinyint;not null" json:"byExpEvent_Hour"`
	ExpEventMinute    uint8    `gorm:"column:byExpEvent_Minute;type:tinyint;not null" json:"byExpEvent_Minute"`
	ExpEventDuration  int16    `gorm:"column:sExpEvent_Duration;type:smallint;not null" json:"sExpEvent_Duration"`
	Tribute           int      `gorm:"column:nTribute;type:int;not null;default:0" json:"nTribute"`
	TerritoryTariff   uint8    `gorm:"column:byTerritoryTariff;type:tinyint;not null;default:0" json:"byTerritoryTariff"`
	TerritoryTax      int      `gorm:"column:nTerritoryTax;type:int;not null;default:0" json:"nTerritoryTax"`
	NationalTreasury  int      `gorm:"column:nNationalTreasury;type:int;not null;default:0" json:"nNationalTreasury"`
	KingName          [21]byte `gorm:"column:strKingName;type:varchar(21)" json:"strKingName,omitempty"`
	ImRequestId       [21]byte `gorm:"column:strImRequestID;type:varchar(21)" json:"strImRequestID,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KingSystem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KingSystemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KingSystem) GetTableName() string {
	return _KingSystemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KingSystem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_SYSTEM] (byNation, byType, sYear, byMonth, byDay, byHour, byMinute, byImType, sImYear, byImMonth, byImDay, byImHour, byImMinute, byNoahEvent, byNoahEvent_Day, byNoahEvent_Hour, byNoahEvent_Minute, sNoahEvent_Duration, byExpEvent, byExpEvent_Day, byExpEvent_Hour, byExpEvent_Minute, sExpEvent_Duration, nTribute, byTerritoryTariff, nTerritoryTax, nNationalTreasury, strKingName, strImRequestID) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Year),
		GetOptionalDecVal(&this.Month),
		GetOptionalDecVal(&this.Day),
		GetOptionalDecVal(&this.Hour),
		GetOptionalDecVal(&this.Minute),
		GetOptionalDecVal(&this.ImType),
		GetOptionalDecVal(&this.ImYear),
		GetOptionalDecVal(&this.ImMonth),
		GetOptionalDecVal(&this.ImDay),
		GetOptionalDecVal(&this.ImHour),
		GetOptionalDecVal(&this.ImMinute),
		GetOptionalDecVal(&this.NoahEvent),
		GetOptionalDecVal(&this.NoahEventDay),
		GetOptionalDecVal(&this.NoahEventHour),
		GetOptionalDecVal(&this.NoahEventMinute),
		GetOptionalDecVal(&this.NoahEventDuration),
		GetOptionalDecVal(&this.ExpEvent),
		GetOptionalDecVal(&this.ExpEventDay),
		GetOptionalDecVal(&this.ExpEventHour),
		GetOptionalDecVal(&this.ExpEventMinute),
		GetOptionalDecVal(&this.ExpEventDuration),
		GetOptionalDecVal(&this.Tribute),
		GetOptionalDecVal(&this.TerritoryTariff),
		GetOptionalDecVal(&this.TerritoryTax),
		GetOptionalDecVal(&this.NationalTreasury),
		GetOptionalBinaryVal(this.KingName),
		GetOptionalBinaryVal(this.ImRequestId))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KingSystem) GetCreateTableString() string {
	query := "CREATE TABLE [KING_SYSTEM] (\n\t[byNation] tinyint NOT NULL,\n\t[byType] tinyint NOT NULL,\n\t[sYear] smallint NOT NULL,\n\t[byMonth] tinyint NOT NULL,\n\t[byDay] tinyint NOT NULL,\n\t[byHour] tinyint NOT NULL,\n\t[byMinute] tinyint NOT NULL,\n\t[byImType] tinyint NOT NULL,\n\t[sImYear] smallint NOT NULL,\n\t[byImMonth] tinyint NOT NULL,\n\t[byImDay] tinyint NOT NULL,\n\t[byImHour] tinyint NOT NULL,\n\t[byImMinute] tinyint NOT NULL,\n\t[byNoahEvent] tinyint NOT NULL,\n\t[byNoahEvent_Day] tinyint NOT NULL,\n\t[byNoahEvent_Hour] tinyint NOT NULL,\n\t[byNoahEvent_Minute] tinyint NOT NULL,\n\t[sNoahEvent_Duration] smallint NOT NULL,\n\t[byExpEvent] tinyint NOT NULL,\n\t[byExpEvent_Day] tinyint NOT NULL,\n\t[byExpEvent_Hour] tinyint NOT NULL,\n\t[byExpEvent_Minute] tinyint NOT NULL,\n\t[sExpEvent_Duration] smallint NOT NULL,\n\t[nTribute] int NOT NULL,\n\t[byTerritoryTariff] tinyint NOT NULL,\n\t[nTerritoryTax] int NOT NULL,\n\t[nNationalTreasury] int NOT NULL,\n\t[strKingName] varchar(21),\n\t[strImRequestID] varchar(21)\n\n)\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_byNoahEvent] DEFAULT 0 FOR [byNoahEvent]\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_byExpEvent] DEFAULT 0 FOR [byExpEvent]\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_nTribute] DEFAULT 0 FOR [nTribute]\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_byTerritoryTariff] DEFAULT 0 FOR [byTerritoryTariff]\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_nTerritoryTax] DEFAULT 0 FOR [nTerritoryTax]\nGO\nALTER TABLE [KING_SYSTEM] ADD CONSTRAINT [DF_KING_SYSTEM_nNationalTreasury] DEFAULT 0 FOR [nNationalTreasury]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
