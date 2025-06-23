package kogen

import (
	"fmt"
)

const (
	_ZoneInfoDatabaseNbr = 0
	_ZoneInfoTableName   = "ZONE_INFO"
)

func init() {
	ModelList = append(ModelList, &ZoneInfo{})
}

// ZoneInfo Zone (map) information
type ZoneInfo struct {
	ServerId  uint8    `gorm:"column:ServerNo;type:tinyint;not null" json:"ServerNo"`
	ZoneId    int16    `gorm:"column:ZoneNo;type:smallint;not null" json:"ZoneNo"`
	Name      [50]byte `gorm:"column:strZoneName;type:varchar(50);not null" json:"strZoneName"`
	InitX     int      `gorm:"column:InitX;type:int;not null" json:"InitX"`
	InitZ     int      `gorm:"column:InitZ;type:int;not null" json:"InitZ"`
	InitY     int      `gorm:"column:InitY;type:int;not null" json:"InitY"`
	Type      uint8    `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	RoomEvent uint8    `gorm:"column:RoomEvent;type:tinyint;not null" json:"RoomEvent"`
	Bz        [50]byte `gorm:"column:bz;type:varchar(50)" json:"bz,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ZoneInfo) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ZoneInfoDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ZoneInfo) GetTableName() string {
	return _ZoneInfoTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ZoneInfo) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ZONE_INFO] (ServerNo, ZoneNo, strZoneName, InitX, InitZ, InitY, Type, RoomEvent, bz) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ZoneId),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(&this.InitX),
		GetOptionalDecVal(&this.InitZ),
		GetOptionalDecVal(&this.InitY),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.RoomEvent),
		GetOptionalBinaryVal(this.Bz))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ZoneInfo) GetCreateTableString() string {
	query := "CREATE TABLE [ZONE_INFO] (\n\t[ServerNo] tinyint NOT NULL,\n\t[ZoneNo] smallint NOT NULL,\n\t[strZoneName] varchar(50) NOT NULL,\n\t[InitX] int NOT NULL,\n\t[InitZ] int NOT NULL,\n\t[InitY] int NOT NULL,\n\t[Type] tinyint NOT NULL,\n\t[RoomEvent] tinyint NOT NULL,\n\t[bz] varchar(50)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
