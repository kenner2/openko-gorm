package kogen

import (
	"fmt"
)

const (
	_NpcMoveItemDatabaseNbr = 0
	_NpcMoveItemTableName   = "K_NPC_MOVE_ITEM"
)

func init() {
	ModelList = append(ModelList, &NpcMoveItem{})
}

// NpcMoveItem NPC Move Item: TODO
type NpcMoveItem struct {
	CastleIndex int16  `gorm:"column:sCastleIndex;type:smallint;not null" json:"sCastleIndex"`
	ChangeItem  *int   `gorm:"column:byChangeItem;type:int" json:"byChangeItem,omitempty"`
	ChangeId    *int   `gorm:"column:sChangeSid;type:int" json:"sChangeSid,omitempty"`
	MoveItem    *int   `gorm:"column:byMoveItem;type:int" json:"byMoveItem,omitempty"`
	MoveMinX    *int16 `gorm:"column:sMoveMinX;type:smallint" json:"sMoveMinX,omitempty"`
	MoveMinY    *int16 `gorm:"column:sMoveMinY;type:smallint" json:"sMoveMinY,omitempty"`
	MoveMaxX    *int16 `gorm:"column:sMoveMaxX;type:smallint" json:"sMoveMaxX,omitempty"`
	MoveMaxY    *int16 `gorm:"column:sMoveMaxY;type:smallint" json:"sMoveMaxY,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *NpcMoveItem) GetDatabaseName() string {
	return GetDatabaseName(DbType(_NpcMoveItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *NpcMoveItem) GetTableName() string {
	return _NpcMoveItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *NpcMoveItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [K_NPC_MOVE_ITEM] (sCastleIndex, byChangeItem, sChangeSid, byMoveItem, sMoveMinX, sMoveMinY, sMoveMaxX, sMoveMaxY) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CastleIndex),
		GetOptionalDecVal(this.ChangeItem),
		GetOptionalDecVal(this.ChangeId),
		GetOptionalDecVal(this.MoveItem),
		GetOptionalDecVal(this.MoveMinX),
		GetOptionalDecVal(this.MoveMinY),
		GetOptionalDecVal(this.MoveMaxX),
		GetOptionalDecVal(this.MoveMaxY))
}

// GetCreateTableString Returns the create table statement for this object
func (this *NpcMoveItem) GetCreateTableString() string {
	query := "CREATE TABLE [K_NPC_MOVE_ITEM] (\n\t\"sCastleIndex\" smallint NOT NULL,\n\t\"byChangeItem\" int,\n\t\"sChangeSid\" int,\n\t\"byMoveItem\" int,\n\t\"sMoveMinX\" smallint,\n\t\"sMoveMinY\" smallint,\n\t\"sMoveMaxX\" smallint,\n\t\"sMoveMaxY\" smallint\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
