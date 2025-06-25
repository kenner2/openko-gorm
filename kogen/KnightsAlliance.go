package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KnightsAllianceDatabaseNbr = 0
	_KnightsAllianceTableName   = "KNIGHTS_ALLIANCE"
)

func init() {
	ModelList = append(ModelList, &KnightsAlliance{})
}

// KnightsAlliance Knights alliance formations
type KnightsAlliance struct {
	MainAllianceKnights int16 `gorm:"column:sMainAllianceKnights;type:smallint;primaryKey;not null" json:"sMainAllianceKnights"`
	SubAllianceKnights  int16 `gorm:"column:sSubAllianceKnights;type:smallint;not null" json:"sSubAllianceKnights"`
	MercenaryClan1      int16 `gorm:"column:sMercenaryClan_1;type:smallint;not null" json:"sMercenaryClan_1"`
	MercenaryClan2      int16 `gorm:"column:sMercenaryClan_2;type:smallint;not null" json:"sMercenaryClan_2"`
}

// GetDatabaseName Returns the table's database name
func (this KnightsAlliance) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsAllianceDatabaseNbr))
}

// TableName Returns the table name
func (this KnightsAlliance) TableName() string {
	return _KnightsAllianceTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KnightsAlliance) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_ALLIANCE] ([sMainAllianceKnights], [sSubAllianceKnights], [sMercenaryClan_1], [sMercenaryClan_2]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.MainAllianceKnights),
		GetOptionalDecVal(&this.SubAllianceKnights),
		GetOptionalDecVal(&this.MercenaryClan1),
		GetOptionalDecVal(&this.MercenaryClan2))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KnightsAlliance) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS_ALLIANCE] (sMainAllianceKnights, sSubAllianceKnights, sMercenaryClan_1, sMercenaryClan_2) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KnightsAlliance) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.MainAllianceKnights),
		GetOptionalDecVal(&this.SubAllianceKnights),
		GetOptionalDecVal(&this.MercenaryClan1),
		GetOptionalDecVal(&this.MercenaryClan2))
}

// GetCreateTableString Returns the create table statement for this object
func (this KnightsAlliance) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_ALLIANCE] (\n\t[sMainAllianceKnights] smallint NOT NULL,\n\t[sSubAllianceKnights] smallint NOT NULL,\n\t[sMercenaryClan_1] smallint NOT NULL,\n\t[sMercenaryClan_2] smallint NOT NULL\n\tCONSTRAINT [PK_KNIGHTS_ALLIANCE] PRIMARY KEY ([sMainAllianceKnights])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KnightsAlliance) SelectClause() (selectClause clause.Select) {
	return _KnightsAllianceSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KnightsAlliance) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KnightsAlliance{}
	rawSql := "SELECT [sMainAllianceKnights], [sSubAllianceKnights], [sMercenaryClan_1], [sMercenaryClan_2] FROM [KNIGHTS_ALLIANCE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsAllianceSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sMainAllianceKnights]",
		},
		clause.Column{
			Name: "[sSubAllianceKnights]",
		},
		clause.Column{
			Name: "[sMercenaryClan_1]",
		},
		clause.Column{
			Name: "[sMercenaryClan_2]",
		},
	},
}
