package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_EventDatabaseNbr = "GAME"
	_EventTableName   = "EVENT"
)

func init() {
	ModelList = append(ModelList, &Event{})
}

// Event Event Information
type Event struct {
	ZoneNumber  uint8          `gorm:"column:ZoneNum;type:tinyint;not null" json:"ZoneNum"`
	EventNumber int16          `gorm:"column:EventNum;type:smallint;not null" json:"EventNum"`
	EventType   uint8          `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	Condition1  *mssql.VarChar `gorm:"column:Cond1;type:varchar(128)" json:"Cond1,omitempty"`
	Condition2  *mssql.VarChar `gorm:"column:Cond2;type:varchar(128)" json:"Cond2,omitempty"`
	Condition3  *mssql.VarChar `gorm:"column:Cond3;type:varchar(128)" json:"Cond3,omitempty"`
	Condition4  *mssql.VarChar `gorm:"column:Cond4;type:varchar(128)" json:"Cond4,omitempty"`
	Condition5  *mssql.VarChar `gorm:"column:Cond5;type:varchar(128)" json:"Cond5,omitempty"`
	Execute1    *mssql.VarChar `gorm:"column:Exec1;type:varchar(128)" json:"Exec1,omitempty"`
	Execute2    *mssql.VarChar `gorm:"column:Exec2;type:varchar(128)" json:"Exec2,omitempty"`
	Execute3    *mssql.VarChar `gorm:"column:Exec3;type:varchar(128)" json:"Exec3,omitempty"`
	Execute4    *mssql.VarChar `gorm:"column:Exec4;type:varchar(128)" json:"Exec4,omitempty"`
	Execute5    *mssql.VarChar `gorm:"column:Exec5;type:varchar(128)" json:"Exec5,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this Event) GetDatabaseName() string {
	return GetDatabaseName(_EventDatabaseNbr)
}

// TableName Returns the table name
func (this Event) TableName() string {
	return _EventTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Event) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [EVENT] ([ZoneNum], [EventNum], [Type], [Cond1], [Cond2], [Cond3], [Cond4], [Cond5], [Exec1], [Exec2], [Exec3], [Exec4], [Exec5]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneNumber),
		GetOptionalDecVal(&this.EventNumber),
		GetOptionalDecVal(&this.EventType),
		GetOptionalVarCharVal(this.Condition1, false),
		GetOptionalVarCharVal(this.Condition2, false),
		GetOptionalVarCharVal(this.Condition3, false),
		GetOptionalVarCharVal(this.Condition4, false),
		GetOptionalVarCharVal(this.Condition5, false),
		GetOptionalVarCharVal(this.Execute1, false),
		GetOptionalVarCharVal(this.Execute2, false),
		GetOptionalVarCharVal(this.Execute3, false),
		GetOptionalVarCharVal(this.Execute4, false),
		GetOptionalVarCharVal(this.Execute5, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Event) GetInsertHeader() string {
	return "INSERT INTO [EVENT] ([ZoneNum], [EventNum], [Type], [Cond1], [Cond2], [Cond3], [Cond4], [Cond5], [Exec1], [Exec2], [Exec3], [Exec4], [Exec5]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Event) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneNumber),
		GetOptionalDecVal(&this.EventNumber),
		GetOptionalDecVal(&this.EventType),
		GetOptionalVarCharVal(this.Condition1, false),
		GetOptionalVarCharVal(this.Condition2, false),
		GetOptionalVarCharVal(this.Condition3, false),
		GetOptionalVarCharVal(this.Condition4, false),
		GetOptionalVarCharVal(this.Condition5, false),
		GetOptionalVarCharVal(this.Execute1, false),
		GetOptionalVarCharVal(this.Execute2, false),
		GetOptionalVarCharVal(this.Execute3, false),
		GetOptionalVarCharVal(this.Execute4, false),
		GetOptionalVarCharVal(this.Execute5, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this Event) GetCreateTableString() string {
	query := "CREATE TABLE [EVENT] (\n\t[ZoneNum] tinyint NOT NULL,\n\t[EventNum] smallint NOT NULL,\n\t[Type] tinyint NOT NULL,\n\t[Cond1] varchar(128),\n\t[Cond2] varchar(128),\n\t[Cond3] varchar(128),\n\t[Cond4] varchar(128),\n\t[Cond5] varchar(128),\n\t[Exec1] varchar(128),\n\t[Exec2] varchar(128),\n\t[Exec3] varchar(128),\n\t[Exec4] varchar(128),\n\t[Exec5] varchar(128)\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Event) SelectClause() (selectClause clause.Select) {
	return _EventSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Event) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Event{}
	rawSql := "SELECT [ZoneNum], [EventNum], [Type], [Cond1], [Cond2], [Cond3], [Cond4], [Cond5], [Exec1], [Exec2], [Exec3], [Exec4], [Exec5] FROM [EVENT]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _EventSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[ZoneNum]",
		},
		clause.Column{
			Name: "[EventNum]",
		},
		clause.Column{
			Name: "[Type]",
		},
		clause.Column{
			Name: "[Cond1]",
		},
		clause.Column{
			Name: "[Cond2]",
		},
		clause.Column{
			Name: "[Cond3]",
		},
		clause.Column{
			Name: "[Cond4]",
		},
		clause.Column{
			Name: "[Cond5]",
		},
		clause.Column{
			Name: "[Exec1]",
		},
		clause.Column{
			Name: "[Exec2]",
		},
		clause.Column{
			Name: "[Exec3]",
		},
		clause.Column{
			Name: "[Exec4]",
		},
		clause.Column{
			Name: "[Exec5]",
		},
	},
}
