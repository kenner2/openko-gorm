package kogen

import (
	"fmt"
)

const (
	_HomeDatabaseNbr = 0
	_HomeTableName   = "HOME"
)

func init() {
	ModelList = append(ModelList, &Home{})
}

// Home TODO Doc
type Home struct {
	Nation        uint8 `gorm:"column:Nation;type:tinyint;not null" json:"Nation"`
	ElmoZoneX     int   `gorm:"column:ElmoZoneX;type:int;not null" json:"ElmoZoneX"`
	ElmoZoneZ     int   `gorm:"column:ElmoZoneZ;type:int;not null" json:"ElmoZoneZ"`
	ElmoZoneLX    uint8 `gorm:"column:ElmoZoneLX;type:tinyint;not null" json:"ElmoZoneLX"`
	ElmoZoneLZ    uint8 `gorm:"column:ElmoZoneLZ;type:tinyint;not null" json:"ElmoZoneLZ"`
	KarusZoneX    int   `gorm:"column:KarusZoneX;type:int;not null" json:"KarusZoneX"`
	KarusZoneZ    int   `gorm:"column:KarusZoneZ;type:int;not null" json:"KarusZoneZ"`
	KarusZoneLX   uint8 `gorm:"column:KarusZoneLX;type:tinyint;not null" json:"KarusZoneLX"`
	KarusZoneLZ   uint8 `gorm:"column:KarusZoneLZ;type:tinyint;not null" json:"KarusZoneLZ"`
	FreeZoneX     int   `gorm:"column:FreeZoneX;type:int;not null" json:"FreeZoneX"`
	FreeZoneZ     int   `gorm:"column:FreeZoneZ;type:int;not null" json:"FreeZoneZ"`
	FreeZoneLX    uint8 `gorm:"column:FreeZoneLX;type:tinyint;not null" json:"FreeZoneLX"`
	FreeZoneLZ    uint8 `gorm:"column:FreeZoneLZ;type:tinyint;not null" json:"FreeZoneLZ"`
	BattleZoneX   int   `gorm:"column:BattleZoneX;type:int;not null" json:"BattleZoneX"`
	BattleZoneZ   int   `gorm:"column:BattleZoneZ;type:int;not null" json:"BattleZoneZ"`
	BattleZoneLX  uint8 `gorm:"column:BattleZoneLX;type:tinyint;not null" json:"BattleZoneLX"`
	BattleZoneLZ  uint8 `gorm:"column:BattleZoneLZ;type:tinyint;not null" json:"BattleZoneLZ"`
	BattleZone2X  int   `gorm:"column:BattleZone2X;type:int;not null" json:"BattleZone2X"`
	BattleZone2Z  int   `gorm:"column:BattleZone2Z;type:int;not null" json:"BattleZone2Z"`
	BattleZone2LX uint8 `gorm:"column:BattleZone2LX;type:tinyint;not null" json:"BattleZone2LX"`
	BattleZone2LZ uint8 `gorm:"column:BattleZone2LZ;type:tinyint;not null" json:"BattleZone2LZ"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Home) GetDatabaseName() string {
	return GetDatabaseName(DbType(_HomeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Home) GetTableName() string {
	return _HomeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Home) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HOME] (Nation, ElmoZoneX, ElmoZoneZ, ElmoZoneLX, ElmoZoneLZ, KarusZoneX, KarusZoneZ, KarusZoneLX, KarusZoneLZ, FreeZoneX, FreeZoneZ, FreeZoneLX, FreeZoneLZ, BattleZoneX, BattleZoneZ, BattleZoneLX, BattleZoneLZ, BattleZone2X, BattleZone2Z, BattleZone2LX, BattleZone2LZ) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.ElmoZoneX),
		GetOptionalDecVal(&this.ElmoZoneZ),
		GetOptionalDecVal(&this.ElmoZoneLX),
		GetOptionalDecVal(&this.ElmoZoneLZ),
		GetOptionalDecVal(&this.KarusZoneX),
		GetOptionalDecVal(&this.KarusZoneZ),
		GetOptionalDecVal(&this.KarusZoneLX),
		GetOptionalDecVal(&this.KarusZoneLZ),
		GetOptionalDecVal(&this.FreeZoneX),
		GetOptionalDecVal(&this.FreeZoneZ),
		GetOptionalDecVal(&this.FreeZoneLX),
		GetOptionalDecVal(&this.FreeZoneLZ),
		GetOptionalDecVal(&this.BattleZoneX),
		GetOptionalDecVal(&this.BattleZoneZ),
		GetOptionalDecVal(&this.BattleZoneLX),
		GetOptionalDecVal(&this.BattleZoneLZ),
		GetOptionalDecVal(&this.BattleZone2X),
		GetOptionalDecVal(&this.BattleZone2Z),
		GetOptionalDecVal(&this.BattleZone2LX),
		GetOptionalDecVal(&this.BattleZone2LZ))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Home) GetCreateTableString() string {
	query := "CREATE TABLE [HOME] (\n\t[Nation] tinyint NOT NULL,\n\t[ElmoZoneX] int NOT NULL,\n\t[ElmoZoneZ] int NOT NULL,\n\t[ElmoZoneLX] tinyint NOT NULL,\n\t[ElmoZoneLZ] tinyint NOT NULL,\n\t[KarusZoneX] int NOT NULL,\n\t[KarusZoneZ] int NOT NULL,\n\t[KarusZoneLX] tinyint NOT NULL,\n\t[KarusZoneLZ] tinyint NOT NULL,\n\t[FreeZoneX] int NOT NULL,\n\t[FreeZoneZ] int NOT NULL,\n\t[FreeZoneLX] tinyint NOT NULL,\n\t[FreeZoneLZ] tinyint NOT NULL,\n\t[BattleZoneX] int NOT NULL,\n\t[BattleZoneZ] int NOT NULL,\n\t[BattleZoneLX] tinyint NOT NULL,\n\t[BattleZoneLZ] tinyint NOT NULL,\n\t[BattleZone2X] int NOT NULL,\n\t[BattleZone2Z] int NOT NULL,\n\t[BattleZone2LX] tinyint NOT NULL,\n\t[BattleZone2LZ] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
