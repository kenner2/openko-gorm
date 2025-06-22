package kogen

import (
	"fmt"
)

const (
	_MonsterChallengeDatabaseNbr = 0
	_MonsterChallengeTableName   = "MONSTER_CHALLENGE"
)

func init() {
	ModelList = append(ModelList, &MonsterChallenge{})
}

// MonsterChallenge Monster challenge (Forgotten Temple)
type MonsterChallenge struct {
	Index      int16 `gorm:"column:sIndex;type:smallint;not null" json:"sIndex"`
	StartTime1 uint8 `gorm:"column:bStartTime1;type:tinyint;not null" json:"bStartTime1"`
	StartTime2 uint8 `gorm:"column:bStartTime2;type:tinyint;not null" json:"bStartTime2"`
	StartTime3 uint8 `gorm:"column:bStartTime3;type:tinyint;not null" json:"bStartTime3"`
	LevelMin   uint8 `gorm:"column:bLevelMin;type:tinyint;not null" json:"bLevelMin"`
	LevelMax   uint8 `gorm:"column:bLevelMax;type:tinyint;not null" json:"bLevelMax"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MonsterChallenge) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MonsterChallengeDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MonsterChallenge) GetTableName() string {
	return _MonsterChallengeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MonsterChallenge) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_CHALLENGE] (sIndex, bStartTime1, bStartTime2, bStartTime3, bLevelMin, bLevelMax) \nVALUES (%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.StartTime1),
		GetOptionalDecVal(&this.StartTime2),
		GetOptionalDecVal(&this.StartTime3),
		GetOptionalDecVal(&this.LevelMin),
		GetOptionalDecVal(&this.LevelMax))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MonsterChallenge) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_CHALLENGE] (\n\t\"sIndex\" smallint NOT NULL,\n\t\"bStartTime1\" tinyint NOT NULL,\n\t\"bStartTime2\" tinyint NOT NULL,\n\t\"bStartTime3\" tinyint NOT NULL,\n\t\"bLevelMin\" tinyint NOT NULL,\n\t\"bLevelMax\" tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
