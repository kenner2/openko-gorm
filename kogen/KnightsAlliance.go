package kogen

import (
	"fmt"
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
	MainAllianceKnights int16 `gorm:"column:sMainAllianceKnights;type:smallint;not null" json:"sMainAllianceKnights"`
	SubAllianceKnights  int16 `gorm:"column:sSubAllianceKnights;type:smallint;not null" json:"sSubAllianceKnights"`
	MercenaryClan1      int16 `gorm:"column:sMercenaryClan_1;type:smallint;not null" json:"sMercenaryClan_1"`
	MercenaryClan2      int16 `gorm:"column:sMercenaryClan_2;type:smallint;not null" json:"sMercenaryClan_2"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KnightsAlliance) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsAllianceDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KnightsAlliance) GetTableName() string {
	return _KnightsAllianceTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KnightsAlliance) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_ALLIANCE] (sMainAllianceKnights, sSubAllianceKnights, sMercenaryClan_1, sMercenaryClan_2) \nVALUES (%s, %s, %s, %s)", GetOptionalDecVal(&this.MainAllianceKnights),
		GetOptionalDecVal(&this.SubAllianceKnights),
		GetOptionalDecVal(&this.MercenaryClan1),
		GetOptionalDecVal(&this.MercenaryClan2))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KnightsAlliance) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_ALLIANCE] (\n\t\"sMainAllianceKnights\" smallint NOT NULL,\n\t\"sSubAllianceKnights\" smallint NOT NULL,\n\t\"sMercenaryClan_1\" smallint NOT NULL,\n\t\"sMercenaryClan_2\" smallint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
