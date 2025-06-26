package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_EventTriggerDatabaseNbr = "GAME"
	_EventTriggerTableName   = "EVENT_TRIGGER"
)

func init() {
	ModelList = append(ModelList, &EventTrigger{})
}

// EventTrigger NPC Event Triggers
type EventTrigger struct {
	Index         int   `gorm:"column:nIndex;type:int;not null" json:"nIndex"`
	NpcType       uint8 `gorm:"column:bNpcType;type:tinyint;not null" json:"bNpcType"`
	NpcId         int16 `gorm:"column:sNpcID;type:smallint;not null" json:"sNpcID"`
	TriggerNumber int   `gorm:"column:nTriggerNum;type:int;not null" json:"nTriggerNum"`
}

// GetDatabaseName Returns the table's database name
func (this EventTrigger) GetDatabaseName() string {
	return GetDatabaseName(_EventTriggerDatabaseNbr)
}

// TableName Returns the table name
func (this EventTrigger) TableName() string {
	return _EventTriggerTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this EventTrigger) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [EVENT_TRIGGER] ([nIndex], [bNpcType], [sNpcID], [nTriggerNum]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcType),
		GetOptionalDecVal(&this.NpcId),
		GetOptionalDecVal(&this.TriggerNumber))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this EventTrigger) GetInsertHeader() string {
	return "INSERT INTO [EVENT_TRIGGER] ([nIndex], [bNpcType], [sNpcID], [nTriggerNum]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this EventTrigger) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcType),
		GetOptionalDecVal(&this.NpcId),
		GetOptionalDecVal(&this.TriggerNumber))
}

// GetCreateTableString Returns the create table statement for this object
func (this EventTrigger) GetCreateTableString() string {
	query := "CREATE TABLE [EVENT_TRIGGER] (\n\t[nIndex] int NOT NULL,\n\t[bNpcType] tinyint NOT NULL,\n\t[sNpcID] smallint NOT NULL,\n\t[nTriggerNum] int NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this EventTrigger) SelectClause() (selectClause clause.Select) {
	return _EventTriggerSelectClause
}

// GetAllTableData Returns a list of all table data
func (this EventTrigger) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []EventTrigger{}
	rawSql := "SELECT [nIndex], [bNpcType], [sNpcID], [nTriggerNum] FROM [EVENT_TRIGGER]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _EventTriggerSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nIndex]",
		},
		clause.Column{
			Name: "[bNpcType]",
		},
		clause.Column{
			Name: "[sNpcID]",
		},
		clause.Column{
			Name: "[nTriggerNum]",
		},
	},
}
