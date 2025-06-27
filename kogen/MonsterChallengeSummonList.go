package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MonsterChallengeSummonListDatabaseNbr = "GAME"
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

// GetDatabaseName Returns the table's database name
func (this MonsterChallengeSummonList) GetDatabaseName() string {
	return GetDatabaseName(_MonsterChallengeSummonListDatabaseNbr)
}

// TableName Returns the table name
func (this MonsterChallengeSummonList) TableName() string {
	return _MonsterChallengeSummonListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MonsterChallengeSummonList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_CHALLENGE_SUMMON_LIST] ([sIndex], [bLevel], [bStage], [bStageLevel], [sTime], [sSid], [sCount], [sPosX], [sPosZ], [bRange]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
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

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MonsterChallengeSummonList) GetInsertHeader() string {
	return "INSERT INTO [MONSTER_CHALLENGE_SUMMON_LIST] ([sIndex], [bLevel], [bStage], [bStageLevel], [sTime], [sSid], [sCount], [sPosX], [sPosZ], [bRange]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MonsterChallengeSummonList) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
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
func (this MonsterChallengeSummonList) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_CHALLENGE_SUMMON_LIST] (\n\t[sIndex] smallint NOT NULL,\n\t[bLevel] tinyint NOT NULL,\n\t[bStage] tinyint NOT NULL,\n\t[bStageLevel] tinyint NOT NULL,\n\t[sTime] smallint NOT NULL,\n\t[sSid] smallint NOT NULL,\n\t[sCount] smallint NOT NULL,\n\t[sPosX] smallint NOT NULL,\n\t[sPosZ] smallint NOT NULL,\n\t[bRange] tinyint NOT NULL\n)\nGO\nALTER TABLE [MONSTER_CHALLENGE_SUMMON_LIST] ADD CONSTRAINT [DF_MONSTER_CHALLENGE_SUMMON_LIST_bStageLevel] DEFAULT 0 FOR [bStageLevel]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MonsterChallengeSummonList) SelectClause() (selectClause clause.Select) {
	return _MonsterChallengeSummonListSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MonsterChallengeSummonList) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MonsterChallengeSummonList{}
	rawSql := "SELECT [sIndex], [bLevel], [bStage], [bStageLevel], [sTime], [sSid], [sCount], [sPosX], [sPosZ], [bRange] FROM [MONSTER_CHALLENGE_SUMMON_LIST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MonsterChallengeSummonListSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIndex]",
		},
		clause.Column{
			Name: "[bLevel]",
		},
		clause.Column{
			Name: "[bStage]",
		},
		clause.Column{
			Name: "[bStageLevel]",
		},
		clause.Column{
			Name: "[sTime]",
		},
		clause.Column{
			Name: "[sSid]",
		},
		clause.Column{
			Name: "[sCount]",
		},
		clause.Column{
			Name: "[sPosX]",
		},
		clause.Column{
			Name: "[sPosZ]",
		},
		clause.Column{
			Name: "[bRange]",
		},
	},
}
