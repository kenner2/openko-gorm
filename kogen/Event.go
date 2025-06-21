package kogen

import (
	"fmt"
)

const (
	_EventDatabaseNbr = 0
	_EventTableName   = "EVENT"
)

func init() {
	ModelList = append(ModelList, &Event{})
}

// Event: Event Information
type Event struct {
	ZoneNumber  uint8     `gorm:"column:ZoneNum;type:tinyint;not null" json:"ZoneNum"`
	EventNumber int16     `gorm:"column:EventNum;type:smallint;not null" json:"EventNum"`
	EventType   uint8     `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	Condition1  [128]byte `gorm:"column:Cond1;type:varchar(128)" json:"Cond1,omitempty"`
	Condition2  [128]byte `gorm:"column:Cond2;type:varchar(128)" json:"Cond2,omitempty"`
	Condition3  [128]byte `gorm:"column:Cond3;type:varchar(128)" json:"Cond3,omitempty"`
	Condition4  [128]byte `gorm:"column:Cond4;type:varchar(128)" json:"Cond4,omitempty"`
	Condition5  [128]byte `gorm:"column:Cond5;type:varchar(128)" json:"Cond5,omitempty"`
	Execute1    [128]byte `gorm:"column:Exec1;type:varchar(128)" json:"Exec1,omitempty"`
	Execute2    [128]byte `gorm:"column:Exec2;type:varchar(128)" json:"Exec2,omitempty"`
	Execute3    [128]byte `gorm:"column:Exec3;type:varchar(128)" json:"Exec3,omitempty"`
	Execute4    [128]byte `gorm:"column:Exec4;type:varchar(128)" json:"Exec4,omitempty"`
	Execute5    [128]byte `gorm:"column:Exec5;type:varchar(128)" json:"Exec5,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Event) GetDatabaseName() string {
	return GetDatabaseName(DbType(_EventDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Event) GetTableName() string {
	return _EventTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Event) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [EVENT] (ZoneNum, EventNum, Type, Cond1, Cond2, Cond3, Cond4, Cond5, Exec1, Exec2, Exec3, Exec4, Exec5) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneNumber),
		GetOptionalDecVal(&this.EventNumber),
		GetOptionalDecVal(&this.EventType),
		GetOptionalBinaryVal(this.Condition1),
		GetOptionalBinaryVal(this.Condition2),
		GetOptionalBinaryVal(this.Condition3),
		GetOptionalBinaryVal(this.Condition4),
		GetOptionalBinaryVal(this.Condition5),
		GetOptionalBinaryVal(this.Execute1),
		GetOptionalBinaryVal(this.Execute2),
		GetOptionalBinaryVal(this.Execute3),
		GetOptionalBinaryVal(this.Execute4),
		GetOptionalBinaryVal(this.Execute5))
}

// GetCreateTableString Returns the create table statement for this object
func (this *Event) GetCreateTableString() string {
	query := "CREATE TABLE \"EVENT\" (\n\t\"ZoneNum\" tinyint NOT NULL,\n\t\"EventNum\" smallint NOT NULL,\n\t\"Type\" tinyint NOT NULL,\n\t\"Cond1\" varchar(128),\n\t\"Cond2\" varchar(128),\n\t\"Cond3\" varchar(128),\n\t\"Cond4\" varchar(128),\n\t\"Cond5\" varchar(128),\n\t\"Exec1\" varchar(128),\n\t\"Exec2\" varchar(128),\n\t\"Exec3\" varchar(128),\n\t\"Exec4\" varchar(128),\n\t\"Exec5\" varchar(128)\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
