package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MonsterChallengeDatabaseNbr = "GAME"
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

// GetDatabaseName Returns the table's database name
func (this MonsterChallenge) GetDatabaseName() string {
	return GetDatabaseName(_MonsterChallengeDatabaseNbr)
}

// TableName Returns the table name
func (this MonsterChallenge) TableName() string {
	return _MonsterChallengeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MonsterChallenge) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MONSTER_CHALLENGE] ([sIndex], [bStartTime1], [bStartTime2], [bStartTime3], [bLevelMin], [bLevelMax]) VALUES\n(%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.StartTime1),
		GetOptionalDecVal(&this.StartTime2),
		GetOptionalDecVal(&this.StartTime3),
		GetOptionalDecVal(&this.LevelMin),
		GetOptionalDecVal(&this.LevelMax))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MonsterChallenge) GetInsertHeader() string {
	return "INSERT INTO [MONSTER_CHALLENGE] ([sIndex], [bStartTime1], [bStartTime2], [bStartTime3], [bLevelMin], [bLevelMax]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MonsterChallenge) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.StartTime1),
		GetOptionalDecVal(&this.StartTime2),
		GetOptionalDecVal(&this.StartTime3),
		GetOptionalDecVal(&this.LevelMin),
		GetOptionalDecVal(&this.LevelMax))
}

// GetCreateTableString Returns the create table statement for this object
func (this MonsterChallenge) GetCreateTableString() string {
	query := "CREATE TABLE [MONSTER_CHALLENGE] (\n\t[sIndex] smallint NOT NULL,\n\t[bStartTime1] tinyint NOT NULL,\n\t[bStartTime2] tinyint NOT NULL,\n\t[bStartTime3] tinyint NOT NULL,\n\t[bLevelMin] tinyint NOT NULL,\n\t[bLevelMax] tinyint NOT NULL\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MonsterChallenge) SelectClause() (selectClause clause.Select) {
	return _MonsterChallengeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this MonsterChallenge) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MonsterChallenge{}
	rawSql := "SELECT [sIndex], [bStartTime1], [bStartTime2], [bStartTime3], [bLevelMin], [bLevelMax] FROM [MONSTER_CHALLENGE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MonsterChallengeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sIndex]",
		},
		clause.Column{
			Name: "[bStartTime1]",
		},
		clause.Column{
			Name: "[bStartTime2]",
		},
		clause.Column{
			Name: "[bStartTime3]",
		},
		clause.Column{
			Name: "[bLevelMin]",
		},
		clause.Column{
			Name: "[bLevelMax]",
		},
	},
}
