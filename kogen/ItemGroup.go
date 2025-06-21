package kogen

import (
	"fmt"
)

const (
	_ItemGroupDatabaseNbr = 0
	_ItemGroupTableName   = "ITEM_GROUP"
)

func init() {
	ModelList = append(ModelList, &ItemGroup{})
}

// ItemGroup: TODO Doc; No Data in table
type ItemGroup struct {
	Group  int16    `gorm:"column:group;type:smallint;not null" json:"group"`
	Name   [30]byte `gorm:"column:name;type:varchar(30)" json:"name,omitempty"`
	Item1  int      `gorm:"column:item1;type:int;not null" json:"item1"`
	Item2  int      `gorm:"column:item2;type:int;not null" json:"item2"`
	Item3  int      `gorm:"column:item3;type:int;not null" json:"item3"`
	Item4  int      `gorm:"column:item4;type:int;not null" json:"item4"`
	Item5  int      `gorm:"column:item5;type:int;not null" json:"item5"`
	Item6  int      `gorm:"column:item6;type:int;not null" json:"item6"`
	Item7  int      `gorm:"column:item7;type:int;not null" json:"item7"`
	Item8  int      `gorm:"column:item8;type:int;not null" json:"item8"`
	Item9  int      `gorm:"column:item9;type:int;not null" json:"item9"`
	Item10 int      `gorm:"column:item10;type:int;not null" json:"item10"`
	Item11 int      `gorm:"column:item11;type:int;not null" json:"item11"`
	Item12 int      `gorm:"column:item12;type:int;not null" json:"item12"`
	Item13 int      `gorm:"column:item13;type:int;not null" json:"item13"`
	Item14 int      `gorm:"column:item14;type:int;not null" json:"item14"`
	Item15 int      `gorm:"column:item15;type:int;not null" json:"item15"`
	Item16 int      `gorm:"column:item16;type:int;not null" json:"item16"`
	Item17 int      `gorm:"column:item17;type:int;not null" json:"item17"`
	Item18 int      `gorm:"column:item18;type:int;not null" json:"item18"`
	Item19 int      `gorm:"column:item19;type:int;not null" json:"item19"`
	Item20 int      `gorm:"column:item20;type:int;not null" json:"item20"`
	Item21 int      `gorm:"column:item21;type:int;not null" json:"item21"`
	Item22 int      `gorm:"column:item22;type:int;not null" json:"item22"`
	Item23 int      `gorm:"column:item23;type:int;not null" json:"item23"`
	Item24 int      `gorm:"column:item24;type:int;not null" json:"item24"`
	Item25 int      `gorm:"column:item25;type:int;not null" json:"item25"`
	Item26 int      `gorm:"column:item26;type:int;not null" json:"item26"`
	Item27 int      `gorm:"column:item27;type:int;not null" json:"item27"`
	Item28 int      `gorm:"column:item28;type:int;not null" json:"item28"`
	Item29 int      `gorm:"column:item29;type:int;not null" json:"item29"`
	Item30 int      `gorm:"column:item30;type:int;not null" json:"item30"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *ItemGroup) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemGroupDatabaseNbr))
}

// GetTableName Returns the table name
func (this *ItemGroup) GetTableName() string {
	return _ItemGroupTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *ItemGroup) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM_GROUP] (group, name, item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12, item13, item14, item15, item16, item17, item18, item19, item20, item21, item22, item23, item24, item25, item26, item27, item28, item29, item30) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Group),
		GetOptionalBinaryVal(this.Name),
		GetOptionalDecVal(&this.Item1),
		GetOptionalDecVal(&this.Item2),
		GetOptionalDecVal(&this.Item3),
		GetOptionalDecVal(&this.Item4),
		GetOptionalDecVal(&this.Item5),
		GetOptionalDecVal(&this.Item6),
		GetOptionalDecVal(&this.Item7),
		GetOptionalDecVal(&this.Item8),
		GetOptionalDecVal(&this.Item9),
		GetOptionalDecVal(&this.Item10),
		GetOptionalDecVal(&this.Item11),
		GetOptionalDecVal(&this.Item12),
		GetOptionalDecVal(&this.Item13),
		GetOptionalDecVal(&this.Item14),
		GetOptionalDecVal(&this.Item15),
		GetOptionalDecVal(&this.Item16),
		GetOptionalDecVal(&this.Item17),
		GetOptionalDecVal(&this.Item18),
		GetOptionalDecVal(&this.Item19),
		GetOptionalDecVal(&this.Item20),
		GetOptionalDecVal(&this.Item21),
		GetOptionalDecVal(&this.Item22),
		GetOptionalDecVal(&this.Item23),
		GetOptionalDecVal(&this.Item24),
		GetOptionalDecVal(&this.Item25),
		GetOptionalDecVal(&this.Item26),
		GetOptionalDecVal(&this.Item27),
		GetOptionalDecVal(&this.Item28),
		GetOptionalDecVal(&this.Item29),
		GetOptionalDecVal(&this.Item30))
}

// GetCreateTableString Returns the create table statement for this object
func (this *ItemGroup) GetCreateTableString() string {
	query := "CREATE TABLE \"ITEM_GROUP\" (\n\t\"group\" smallint NOT NULL,\n\t\"name\" varchar(30),\n\t\"item1\" int NOT NULL,\n\t\"item2\" int NOT NULL,\n\t\"item3\" int NOT NULL,\n\t\"item4\" int NOT NULL,\n\t\"item5\" int NOT NULL,\n\t\"item6\" int NOT NULL,\n\t\"item7\" int NOT NULL,\n\t\"item8\" int NOT NULL,\n\t\"item9\" int NOT NULL,\n\t\"item10\" int NOT NULL,\n\t\"item11\" int NOT NULL,\n\t\"item12\" int NOT NULL,\n\t\"item13\" int NOT NULL,\n\t\"item14\" int NOT NULL,\n\t\"item15\" int NOT NULL,\n\t\"item16\" int NOT NULL,\n\t\"item17\" int NOT NULL,\n\t\"item18\" int NOT NULL,\n\t\"item19\" int NOT NULL,\n\t\"item20\" int NOT NULL,\n\t\"item21\" int NOT NULL,\n\t\"item22\" int NOT NULL,\n\t\"item23\" int NOT NULL,\n\t\"item24\" int NOT NULL,\n\t\"item25\" int NOT NULL,\n\t\"item26\" int NOT NULL,\n\t\"item27\" int NOT NULL,\n\t\"item28\" int NOT NULL,\n\t\"item29\" int NOT NULL,\n\t\"item30\" int NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
