package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_NpcMoveItemDatabaseNbr = "GAME"
	_NpcMoveItemTableName   = "K_NPC_MOVE_ITEM"
)

func init() {
	ModelList = append(ModelList, &NpcMoveItem{})
}

// NpcMoveItem NPC Move Item: TODO
type NpcMoveItem struct {
	CastleIndex int16  `gorm:"column:sCastleIndex;type:smallint;primaryKey;not null" json:"sCastleIndex"`
	ChangeItem  *int   `gorm:"column:byChangeItem;type:int" json:"byChangeItem,omitempty"`
	ChangeId    *int   `gorm:"column:sChangeSid;type:int" json:"sChangeSid,omitempty"`
	MoveItem    *int   `gorm:"column:byMoveItem;type:int" json:"byMoveItem,omitempty"`
	MoveMinX    *int16 `gorm:"column:sMoveMinX;type:smallint" json:"sMoveMinX,omitempty"`
	MoveMinY    *int16 `gorm:"column:sMoveMinY;type:smallint" json:"sMoveMinY,omitempty"`
	MoveMaxX    *int16 `gorm:"column:sMoveMaxX;type:smallint" json:"sMoveMaxX,omitempty"`
	MoveMaxY    *int16 `gorm:"column:sMoveMaxY;type:smallint" json:"sMoveMaxY,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this NpcMoveItem) GetDatabaseName() string {
	return GetDatabaseName(_NpcMoveItemDatabaseNbr)
}

// TableName Returns the table name
func (this NpcMoveItem) TableName() string {
	return _NpcMoveItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this NpcMoveItem) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [K_NPC_MOVE_ITEM] ([sCastleIndex], [byChangeItem], [sChangeSid], [byMoveItem], [sMoveMinX], [sMoveMinY], [sMoveMaxX], [sMoveMaxY]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CastleIndex),
		GetOptionalDecVal(this.ChangeItem),
		GetOptionalDecVal(this.ChangeId),
		GetOptionalDecVal(this.MoveItem),
		GetOptionalDecVal(this.MoveMinX),
		GetOptionalDecVal(this.MoveMinY),
		GetOptionalDecVal(this.MoveMaxX),
		GetOptionalDecVal(this.MoveMaxY))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this NpcMoveItem) GetInsertHeader() string {
	return "INSERT INTO [K_NPC_MOVE_ITEM] ([sCastleIndex], [byChangeItem], [sChangeSid], [byMoveItem], [sMoveMinX], [sMoveMinY], [sMoveMaxX], [sMoveMaxY]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this NpcMoveItem) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CastleIndex),
		GetOptionalDecVal(this.ChangeItem),
		GetOptionalDecVal(this.ChangeId),
		GetOptionalDecVal(this.MoveItem),
		GetOptionalDecVal(this.MoveMinX),
		GetOptionalDecVal(this.MoveMinY),
		GetOptionalDecVal(this.MoveMaxX),
		GetOptionalDecVal(this.MoveMaxY))
}

// GetCreateTableString Returns the create table statement for this object
func (this NpcMoveItem) GetCreateTableString() string {
	query := "CREATE TABLE [K_NPC_MOVE_ITEM] (\n\t[sCastleIndex] smallint NOT NULL,\n\t[byChangeItem] int,\n\t[sChangeSid] int,\n\t[byMoveItem] int,\n\t[sMoveMinX] smallint,\n\t[sMoveMinY] smallint,\n\t[sMoveMaxX] smallint,\n\t[sMoveMaxY] smallint\n\tCONSTRAINT [PK_K_NPC_MOVE_ITEM] PRIMARY KEY ([sCastleIndex])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this NpcMoveItem) SelectClause() (selectClause clause.Select) {
	return _NpcMoveItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this NpcMoveItem) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []NpcMoveItem{}
	rawSql := "SELECT [sCastleIndex], [byChangeItem], [sChangeSid], [byMoveItem], [sMoveMinX], [sMoveMinY], [sMoveMaxX], [sMoveMaxY] FROM [K_NPC_MOVE_ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _NpcMoveItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sCastleIndex]",
		},
		clause.Column{
			Name: "[byChangeItem]",
		},
		clause.Column{
			Name: "[sChangeSid]",
		},
		clause.Column{
			Name: "[byMoveItem]",
		},
		clause.Column{
			Name: "[sMoveMinX]",
		},
		clause.Column{
			Name: "[sMoveMinY]",
		},
		clause.Column{
			Name: "[sMoveMaxX]",
		},
		clause.Column{
			Name: "[sMoveMaxY]",
		},
	},
}
