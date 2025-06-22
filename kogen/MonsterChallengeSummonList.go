package kogen

import (
	"fmt"
)

const (
	_MonsterChallengeSummonListDatabaseNbr = 0
	_MonsterChallengeSummonListTableName   = "MONSTER_CHALLENGE_SUMMON_LIST"
)

func init() {
	ModelList = append(ModelList, &MonsterChallengeSummonList{})
}

// MonsterChallengeSummonList Forgotten Temple summon list
type MonsterChallengeSummonList struct {
	Index      int16 `gorm:"column:sIndex;type:smallint;not null" json:"sIndex"`
	Level      uint8 `gorm:"column:bLevel;type:tinyint;not null" json:"bLevel"`
	Stage      uint8 `gorm:"column:bStage;type:tinyint;not null" json:"bStage"`
	StageLevel uint8 `gorm:"column:bStageLevel;type:tinyint;not null;default:0" json:"bStageLevel"`
	Time       int16 `gorm:"column:sTime;type:smallint;not null" json:"sTime"`
	MonsterId  int16 `gorm:"column:sSid;type:smallint;not null" json:"sSid"`
	Count      int16 `gorm:"column:sCount;type:smallint;not null" json:"sCount"`
	PosX       int16 `gorm:"column:sPosX;type:smallint;not null" json:"sPosX"`
	PosZ       int16 `gorm:"column:sPosZ;type:smallint;not null" json:"sPosZ"`
	Range      uint8 `gorm:"column:bRange;type:tinyint;not null" json:"bRange"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MonsterChallengeSummonList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MonsterChallengeSummonListDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MonsterChallengeSummonList) GetTableName() string {
	return _MonsterChallengeSummonListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MonsterChallengeSummonList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_CHALLENGE_SUMMON_LIST] (sIndex, bLevel, bStage, bStageLevel, sTime, sSid, sCount, sPosX, sPosZ, bRange) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Stage),
		GetOptionalDecVal(&this.StageLevel),
		GetOptionalDecVal(&this.Time),
		GetOptionalDecVal(&this.MonsterId),
		GetOptionalDecVal(&this.Count),
		GetOptionalDecVal(&this.PosX),
		GetOptionalDecVal(&this.PosZ),
		GetOptionalDecVal(&this.Range))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MonsterChallengeSummonList) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_CHALLENGE_SUMMON_LIST] (\n\t\"sIndex\" smallint NOT NULL,\n\t\"bLevel\" tinyint NOT NULL,\n\t\"bStage\" tinyint NOT NULL,\n\t\"bStageLevel\" tinyint NOT NULL,\n\t\"sTime\" smallint NOT NULL,\n\t\"sSid\" smallint NOT NULL,\n\t\"sCount\" smallint NOT NULL,\n\t\"sPosX\" smallint NOT NULL,\n\t\"sPosZ\" smallint NOT NULL,\n\t\"bRange\" tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
