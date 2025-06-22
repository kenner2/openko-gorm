package kogen

import (
	"fmt"
)

const (
	_BattleDatabaseNbr = 0
	_BattleTableName   = "BATTLE"
)

func init() {
	ModelList = append(ModelList, &Battle{})
}

// Battle Battle data for the game server
type Battle struct {
	Index          int16    `gorm:"column:sIndex;type:smallint;not null" json:"sIndex"`
	Nation         uint8    `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	UserName       [21]byte `gorm:"column:strUserName;type:varchar(21)" json:"strUserName,omitempty"`
	ElmoArea       uint8    `gorm:"column:byElmoArea;type:tinyint;not null;default:0" json:"byElmoArea"`
	KarusArea      uint8    `gorm:"column:byKarusArea;type:tinyint;not null;default:0" json:"byKarusArea"`
	ElmoAdvantage  uint8    `gorm:"column:byElmoAdvantage;type:tinyint;not null;default:0" json:"byElmoAdvantage"`
	KarusAdvantage uint8    `gorm:"column:byKarusAdvantage;type:tinyint;not null;default:0" json:"byKarusAdvantage"`
	Area1          uint8    `gorm:"column:byArea_1;type:tinyint;not null;default:0" json:"byArea_1"`
	Area2          uint8    `gorm:"column:byArea_2;type:tinyint;not null;default:0" json:"byArea_2"`
	Area3          uint8    `gorm:"column:byArea_3;type:tinyint;not null;default:0" json:"byArea_3"`
	Area4          uint8    `gorm:"column:byArea_4;type:tinyint;not null;default:0" json:"byArea_4"`
	Area5          uint8    `gorm:"column:byArea_5;type:tinyint;not null;default:0" json:"byArea_5"`
	Area6          uint8    `gorm:"column:byArea_6;type:tinyint;not null;default:0" json:"byArea_6"`
	Area7          uint8    `gorm:"column:byArea_7;type:tinyint;not null;default:0" json:"byArea_7"`
	Area8          uint8    `gorm:"column:byArea_8;type:tinyint;not null;default:0" json:"byArea_8"`
	Area9          uint8    `gorm:"column:byArea_9;type:tinyint;not null;default:0" json:"byArea_9"`
	Area10         uint8    `gorm:"column:byArea_10;type:tinyint;not null;default:0" json:"byArea_10"`
	Area11         uint8    `gorm:"column:byArea_11;type:tinyint;not null;default:0" json:"byArea_11"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Battle) GetDatabaseName() string {
	return GetDatabaseName(DbType(_BattleDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Battle) GetTableName() string {
	return _BattleTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Battle) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [BATTLE] (sIndex, byNation, strUserName, byElmoArea, byKarusArea, byElmoAdvantage, byKarusAdvantage, byArea_1, byArea_2, byArea_3, byArea_4, byArea_5, byArea_6, byArea_7, byArea_8, byArea_9, byArea_10, byArea_11) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.Nation),
		GetOptionalBinaryVal(this.UserName),
		GetOptionalDecVal(&this.ElmoArea),
		GetOptionalDecVal(&this.KarusArea),
		GetOptionalDecVal(&this.ElmoAdvantage),
		GetOptionalDecVal(&this.KarusAdvantage),
		GetOptionalDecVal(&this.Area1),
		GetOptionalDecVal(&this.Area2),
		GetOptionalDecVal(&this.Area3),
		GetOptionalDecVal(&this.Area4),
		GetOptionalDecVal(&this.Area5),
		GetOptionalDecVal(&this.Area6),
		GetOptionalDecVal(&this.Area7),
		GetOptionalDecVal(&this.Area8),
		GetOptionalDecVal(&this.Area9),
		GetOptionalDecVal(&this.Area10),
		GetOptionalDecVal(&this.Area11))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Battle) GetCreateTableString() string {
	query := "CREATE TABLE [BATTLE] (\n\t[sIndex] smallint NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[strUserName] varchar(21),\n\t[byElmoArea] tinyint NOT NULL,\n\t[byKarusArea] tinyint NOT NULL,\n\t[byElmoAdvantage] tinyint NOT NULL,\n\t[byKarusAdvantage] tinyint NOT NULL,\n\t[byArea_1] tinyint NOT NULL,\n\t[byArea_2] tinyint NOT NULL,\n\t[byArea_3] tinyint NOT NULL,\n\t[byArea_4] tinyint NOT NULL,\n\t[byArea_5] tinyint NOT NULL,\n\t[byArea_6] tinyint NOT NULL,\n\t[byArea_7] tinyint NOT NULL,\n\t[byArea_8] tinyint NOT NULL,\n\t[byArea_9] tinyint NOT NULL,\n\t[byArea_10] tinyint NOT NULL,\n\t[byArea_11] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
