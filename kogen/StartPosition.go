package kogen

import (
	"fmt"
)

const (
	_StartPositionDatabaseNbr = 0
	_StartPositionTableName   = "START_POSITION"
)

func init() {
	ModelList = append(ModelList, &StartPosition{})
}

// StartPosition Start position
type StartPosition struct {
	ZoneId     int16 `gorm:"column:ZoneID;type:smallint;not null" json:"ZoneID"`
	KarusX     int16 `gorm:"column:sKarusX;type:smallint;not null" json:"sKarusX"`
	KarusZ     int16 `gorm:"column:sKarusZ;type:smallint;not null" json:"sKarusZ"`
	ElmoX      int16 `gorm:"column:sElmoradX;type:smallint;not null" json:"sElmoradX"`
	ElmoZ      int16 `gorm:"column:sElmoradZ;type:smallint;not null" json:"sElmoradZ"`
	RangeX     uint8 `gorm:"column:bRangeX;type:tinyint;not null" json:"bRangeX"`
	RangeZ     uint8 `gorm:"column:bRangeZ;type:tinyint;not null" json:"bRangeZ"`
	KarusGateX int16 `gorm:"column:sKarusGateX;type:smallint;not null" json:"sKarusGateX"`
	KarusGateZ int16 `gorm:"column:sKarusGateZ;type:smallint;not null" json:"sKarusGateZ"`
	ElmoGateX  int16 `gorm:"column:sElmoGateX;type:smallint;not null" json:"sElmoGateX"`
	ElmoGateZ  int16 `gorm:"column:sElmoGateZ;type:smallint;not null" json:"sElmoGateZ"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *StartPosition) GetDatabaseName() string {
	return GetDatabaseName(DbType(_StartPositionDatabaseNbr))
}

// GetTableName Returns the table name
func (this *StartPosition) GetTableName() string {
	return _StartPositionTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *StartPosition) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [START_POSITION] (ZoneID, sKarusX, sKarusZ, sElmoradX, sElmoradZ, bRangeX, bRangeZ, sKarusGateX, sKarusGateZ, sElmoGateX, sElmoGateZ) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneId),
		GetOptionalDecVal(&this.KarusX),
		GetOptionalDecVal(&this.KarusZ),
		GetOptionalDecVal(&this.ElmoX),
		GetOptionalDecVal(&this.ElmoZ),
		GetOptionalDecVal(&this.RangeX),
		GetOptionalDecVal(&this.RangeZ),
		GetOptionalDecVal(&this.KarusGateX),
		GetOptionalDecVal(&this.KarusGateZ),
		GetOptionalDecVal(&this.ElmoGateX),
		GetOptionalDecVal(&this.ElmoGateZ))
}

// GetCreateTableString Returns the create table statement for this object
func (this *StartPosition) GetCreateTableString() string {
	query := "CREATE TABLE [START_POSITION] (\n\t[ZoneID] smallint NOT NULL,\n\t[sKarusX] smallint NOT NULL,\n\t[sKarusZ] smallint NOT NULL,\n\t[sElmoradX] smallint NOT NULL,\n\t[sElmoradZ] smallint NOT NULL,\n\t[bRangeX] tinyint NOT NULL,\n\t[bRangeZ] tinyint NOT NULL,\n\t[sKarusGateX] smallint NOT NULL,\n\t[sKarusGateZ] smallint NOT NULL,\n\t[sElmoGateX] smallint NOT NULL,\n\t[sElmoGateZ] smallint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
