package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_EventDatabaseNbr = 0
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
	return GetDatabaseName(DbType(_EventDatabaseNbr))
}

// TableName Returns the table name
func (this Event) TableName() string {
	return _EventTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Event) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [EVENT] ([ZoneNum], [EventNum], [Type], [Cond1], [Cond2], [Cond3], [Cond4], [Cond5], [Exec1], [Exec2], [Exec3], [Exec4], [Exec5]) VALUES\n(%s, %s, %s, CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s))", GetOptionalDecVal(&this.ZoneNumber),
		GetOptionalDecVal(&this.EventNumber),
		GetOptionalDecVal(&this.EventType),
		GetOptionalVarCharVal(this.Condition1, true),
		GetOptionalVarCharVal(this.Condition2, true),
		GetOptionalVarCharVal(this.Condition3, true),
		GetOptionalVarCharVal(this.Condition4, true),
		GetOptionalVarCharVal(this.Condition5, true),
		GetOptionalVarCharVal(this.Execute1, true),
		GetOptionalVarCharVal(this.Execute2, true),
		GetOptionalVarCharVal(this.Execute3, true),
		GetOptionalVarCharVal(this.Execute4, true),
		GetOptionalVarCharVal(this.Execute5, true))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Event) GetInsertHeader() string {
	return "INSERT INTO [EVENT] (ZoneNum, EventNum, Type, Cond1, Cond2, Cond3, Cond4, Cond5, Exec1, Exec2, Exec3, Exec4, Exec5) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Event) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s), CONVERT(varchar(128), %s))", GetOptionalDecVal(&this.ZoneNumber),
		GetOptionalDecVal(&this.EventNumber),
		GetOptionalDecVal(&this.EventType),
		GetOptionalVarCharVal(this.Condition1, true),
		GetOptionalVarCharVal(this.Condition2, true),
		GetOptionalVarCharVal(this.Condition3, true),
		GetOptionalVarCharVal(this.Condition4, true),
		GetOptionalVarCharVal(this.Condition5, true),
		GetOptionalVarCharVal(this.Execute1, true),
		GetOptionalVarCharVal(this.Execute2, true),
		GetOptionalVarCharVal(this.Execute3, true),
		GetOptionalVarCharVal(this.Execute4, true),
		GetOptionalVarCharVal(this.Execute5, true))
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
	rawSql := "SELECT [ZoneNum], [EventNum], [Type], CONVERT(VARBINARY(128), [Cond1]) as [Cond1], CONVERT(VARBINARY(128), [Cond2]) as [Cond2], CONVERT(VARBINARY(128), [Cond3]) as [Cond3], CONVERT(VARBINARY(128), [Cond4]) as [Cond4], CONVERT(VARBINARY(128), [Cond5]) as [Cond5], CONVERT(VARBINARY(128), [Exec1]) as [Exec1], CONVERT(VARBINARY(128), [Exec2]) as [Exec2], CONVERT(VARBINARY(128), [Exec3]) as [Exec3], CONVERT(VARBINARY(128), [Exec4]) as [Exec4], CONVERT(VARBINARY(128), [Exec5]) as [Exec5] FROM [EVENT]"
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
			Name: "CONVERT(VARBINARY(128), [Cond1]) as [Cond1]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Cond2]) as [Cond2]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Cond3]) as [Cond3]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Cond4]) as [Cond4]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Cond5]) as [Cond5]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Exec1]) as [Exec1]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Exec2]) as [Exec2]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Exec3]) as [Exec3]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Exec4]) as [Exec4]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(128), [Exec5]) as [Exec5]",
		},
	},
}
