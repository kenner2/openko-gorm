package kogen

import (
	"fmt"
)

const (
	_EventTriggerDatabaseNbr = 0
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

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *EventTrigger) GetDatabaseName() string {
	return GetDatabaseName(DbType(_EventTriggerDatabaseNbr))
}

// GetTableName Returns the table name
func (this *EventTrigger) GetTableName() string {
	return _EventTriggerTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *EventTrigger) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [EVENT_TRIGGER] (nIndex, bNpcType, sNpcID, nTriggerNum) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.Index),
		GetOptionalDecVal(&this.NpcType),
		GetOptionalDecVal(&this.NpcId),
		GetOptionalDecVal(&this.TriggerNumber))
}

// GetCreateTableString Returns the create table statement for this object
func (this *EventTrigger) GetCreateTableString() string {
	query := "CREATE TABLE [EVENT_TRIGGER] (\n\t\"nIndex\" int NOT NULL,\n\t\"bNpcType\" tinyint NOT NULL,\n\t\"sNpcID\" smallint NOT NULL,\n\t\"nTriggerNum\" int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
