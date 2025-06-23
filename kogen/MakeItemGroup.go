package kogen

import (
	"fmt"
)

const (
	_MakeItemGroupDatabaseNbr = 0
	_MakeItemGroupTableName   = "MAKE_ITEM_GROUP"
)

func init() {
	ModelList = append(ModelList, &MakeItemGroup{})
}

// MakeItemGroup Make item group
type MakeItemGroup struct {
	ItemGroupNumber int `gorm:"column:iItemGroupNum;type:int;not null" json:"iItemGroupNum"`
	Item1           int `gorm:"column:iItem_1;type:int;not null" json:"iItem_1"`
	Item2           int `gorm:"column:iItem_2;type:int;not null" json:"iItem_2"`
	Item3           int `gorm:"column:iItem_3;type:int;not null" json:"iItem_3"`
	Item4           int `gorm:"column:iItem_4;type:int;not null" json:"iItem_4"`
	Item5           int `gorm:"column:iItem_5;type:int;not null" json:"iItem_5"`
	Item6           int `gorm:"column:iItem_6;type:int;not null" json:"iItem_6"`
	Item7           int `gorm:"column:iItem_7;type:int;not null" json:"iItem_7"`
	Item8           int `gorm:"column:iItem_8;type:int;not null" json:"iItem_8"`
	Item9           int `gorm:"column:iItem_9;type:int;not null" json:"iItem_9"`
	Item10          int `gorm:"column:iItem_10;type:int;not null" json:"iItem_10"`
	Item11          int `gorm:"column:iItem_11;type:int;not null" json:"iItem_11"`
	Item12          int `gorm:"column:iItem_12;type:int;not null" json:"iItem_12"`
	Item13          int `gorm:"column:iItem_13;type:int;not null" json:"iItem_13"`
	Item14          int `gorm:"column:iItem_14;type:int;not null" json:"iItem_14"`
	Item15          int `gorm:"column:iItem_15;type:int;not null" json:"iItem_15"`
	Item16          int `gorm:"column:iItem_16;type:int;not null" json:"iItem_16"`
	Item17          int `gorm:"column:iItem_17;type:int;not null" json:"iItem_17"`
	Item18          int `gorm:"column:iItem_18;type:int;not null" json:"iItem_18"`
	Item19          int `gorm:"column:iItem_19;type:int;not null" json:"iItem_19"`
	Item20          int `gorm:"column:iItem_20;type:int;not null" json:"iItem_20"`
	Item21          int `gorm:"column:iItem_21;type:int;not null" json:"iItem_21"`
	Item22          int `gorm:"column:iItem_22;type:int;not null" json:"iItem_22"`
	Item23          int `gorm:"column:iItem_23;type:int;not null" json:"iItem_23"`
	Item24          int `gorm:"column:iItem_24;type:int;not null" json:"iItem_24"`
	Item25          int `gorm:"column:iItem_25;type:int;not null" json:"iItem_25"`
	Item26          int `gorm:"column:iItem_26;type:int;not null" json:"iItem_26"`
	Item27          int `gorm:"column:iItem_27;type:int;not null" json:"iItem_27"`
	Item28          int `gorm:"column:iItem_28;type:int;not null" json:"iItem_28"`
	Item29          int `gorm:"column:iItem_29;type:int;not null" json:"iItem_29"`
	Item30          int `gorm:"column:iItem_30;type:int;not null" json:"iItem_30"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MakeItemGroup) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MakeItemGroupDatabaseNbr))
}

// GetTableName Returns the table name
func (this *MakeItemGroup) GetTableName() string {
	return _MakeItemGroupTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MakeItemGroup) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAKE_ITEM_GROUP] (iItemGroupNum, iItem_1, iItem_2, iItem_3, iItem_4, iItem_5, iItem_6, iItem_7, iItem_8, iItem_9, iItem_10, iItem_11, iItem_12, iItem_13, iItem_14, iItem_15, iItem_16, iItem_17, iItem_18, iItem_19, iItem_20, iItem_21, iItem_22, iItem_23, iItem_24, iItem_25, iItem_26, iItem_27, iItem_28, iItem_29, iItem_30) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ItemGroupNumber),
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
func (this *MakeItemGroup) GetCreateTableString() string {
	query := "CREATE TABLE [MAKE_ITEM_GROUP] (\n\t[iItemGroupNum] int NOT NULL,\n\t[iItem_1] int NOT NULL,\n\t[iItem_2] int NOT NULL,\n\t[iItem_3] int NOT NULL,\n\t[iItem_4] int NOT NULL,\n\t[iItem_5] int NOT NULL,\n\t[iItem_6] int NOT NULL,\n\t[iItem_7] int NOT NULL,\n\t[iItem_8] int NOT NULL,\n\t[iItem_9] int NOT NULL,\n\t[iItem_10] int NOT NULL,\n\t[iItem_11] int NOT NULL,\n\t[iItem_12] int NOT NULL,\n\t[iItem_13] int NOT NULL,\n\t[iItem_14] int NOT NULL,\n\t[iItem_15] int NOT NULL,\n\t[iItem_16] int NOT NULL,\n\t[iItem_17] int NOT NULL,\n\t[iItem_18] int NOT NULL,\n\t[iItem_19] int NOT NULL,\n\t[iItem_20] int NOT NULL,\n\t[iItem_21] int NOT NULL,\n\t[iItem_22] int NOT NULL,\n\t[iItem_23] int NOT NULL,\n\t[iItem_24] int NOT NULL,\n\t[iItem_25] int NOT NULL,\n\t[iItem_26] int NOT NULL,\n\t[iItem_27] int NOT NULL,\n\t[iItem_28] int NOT NULL,\n\t[iItem_29] int NOT NULL,\n\t[iItem_30] int NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
