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

// Battle: Battle data for the game server
type Battle struct {
	Index          int16  `gorm:"column:sIndex;not null" json:"sIndex"`
	Nation         uint8  `gorm:"column:byNation;not null" json:"byNation"`
	UserName       *uint8 `gorm:"column:strUserName" json:"strUserName,omitempty"`
	ElmoArea       uint8  `gorm:"column:byElmoArea;not null" json:"byElmoArea"`
	KarusArea      uint8  `gorm:"column:byKarusArea;not null" json:"byKarusArea"`
	ElmoAdvantage  uint8  `gorm:"column:byElmoAdvantage;not null" json:"byElmoAdvantage"`
	KarusAdvantage uint8  `gorm:"column:byKarusAdvantage;not null" json:"byKarusAdvantage"`
	Area1          uint8  `gorm:"column:byArea_1;not null" json:"byArea_1"`
	Area2          uint8  `gorm:"column:byArea_2;not null" json:"byArea_2"`
	Area3          uint8  `gorm:"column:byArea_3;not null" json:"byArea_3"`
	Area4          uint8  `gorm:"column:byArea_4;not null" json:"byArea_4"`
	Area5          uint8  `gorm:"column:byArea_5;not null" json:"byArea_5"`
	Area6          uint8  `gorm:"column:byArea_6;not null" json:"byArea_6"`
	Area7          uint8  `gorm:"column:byArea_7;not null" json:"byArea_7"`
	Area8          uint8  `gorm:"column:byArea_8;not null" json:"byArea_8"`
	Area9          uint8  `gorm:"column:byArea_9;not null" json:"byArea_9"`
	Area10         uint8  `gorm:"column:byArea_10;not null" json:"byArea_10"`
	Area11         uint8  `gorm:"column:byArea_11;not null" json:"byArea_11"`
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
		GetOptionalDecVal(this.UserName),
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
