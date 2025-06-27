package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_NpcPosDatabaseNbr = "GAME"
	_NpcPosTableName   = "K_NPCPOS"
)

func init() {
	ModelList = append(ModelList, &NpcPos{})
}

// NpcPos NPC Spawn Positions
type NpcPos struct {
	ZoneId        *int16  `gorm:"column:ZoneID;type:smallint" json:"ZoneID,omitempty"`
	NpcId         *int    `gorm:"column:NpcID;type:int" json:"NpcID,omitempty"`
	ActType       *uint8  `gorm:"column:ActType;type:tinyint" json:"ActType,omitempty"`
	RegenType     *uint8  `gorm:"column:RegenType;type:tinyint" json:"RegenType,omitempty"`
	DungeonFamily *uint8  `gorm:"column:DungeonFamily;type:tinyint" json:"DungeonFamily,omitempty"`
	SpecialType   *uint8  `gorm:"column:SpecialType;type:tinyint" json:"SpecialType,omitempty"`
	TrapNumber    *uint8  `gorm:"column:TrapNumber;type:tinyint" json:"TrapNumber,omitempty"`
	LeftX         *int    `gorm:"column:LeftX;type:int" json:"LeftX,omitempty"`
	TopZ          *int    `gorm:"column:TopZ;type:int" json:"TopZ,omitempty"`
	RightX        *int    `gorm:"column:RightX;type:int" json:"RightX,omitempty"`
	BottomZ       *int    `gorm:"column:BottomZ;type:int" json:"BottomZ,omitempty"`
	LimitMinZ     *int    `gorm:"column:LimitMinZ;type:int" json:"LimitMinZ,omitempty"`
	LimitMinX     *int    `gorm:"column:LimitMinX;type:int" json:"LimitMinX,omitempty"`
	LimitMaxX     *int    `gorm:"column:LimitMaxX;type:int" json:"LimitMaxX,omitempty"`
	LimitMaxZ     *int    `gorm:"column:LimitMaxZ;type:int" json:"LimitMaxZ,omitempty"`
	NumNpc        *uint8  `gorm:"column:NumNPC;type:tinyint" json:"NumNPC,omitempty"`
	RespawnTime   *int16  `gorm:"column:RegTime;type:smallint" json:"RegTime,omitempty"`
	Direction     *int    `gorm:"column:byDirection;type:int" json:"byDirection,omitempty"`
	DotCount      *uint8  `gorm:"column:DotCnt;type:tinyint" json:"DotCnt,omitempty"`
	Path          *[]byte `gorm:"column:path;type:text" json:"path,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this NpcPos) GetDatabaseName() string {
	return GetDatabaseName(_NpcPosDatabaseNbr)
}

// TableName Returns the table name
func (this NpcPos) TableName() string {
	return _NpcPosTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this NpcPos) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [K_NPCPOS] ([ZoneID], [NpcID], [ActType], [RegenType], [DungeonFamily], [SpecialType], [TrapNumber], [LeftX], [TopZ], [RightX], [BottomZ], [LimitMinZ], [LimitMinX], [LimitMaxX], [LimitMaxZ], [NumNPC], [RegTime], [byDirection], [DotCnt], [path]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(this.ZoneId),
		GetOptionalDecVal(this.NpcId),
		GetOptionalDecVal(this.ActType),
		GetOptionalDecVal(this.RegenType),
		GetOptionalDecVal(this.DungeonFamily),
		GetOptionalDecVal(this.SpecialType),
		GetOptionalDecVal(this.TrapNumber),
		GetOptionalDecVal(this.LeftX),
		GetOptionalDecVal(this.TopZ),
		GetOptionalDecVal(this.RightX),
		GetOptionalDecVal(this.BottomZ),
		GetOptionalDecVal(this.LimitMinZ),
		GetOptionalDecVal(this.LimitMinX),
		GetOptionalDecVal(this.LimitMaxX),
		GetOptionalDecVal(this.LimitMaxZ),
		GetOptionalDecVal(this.NumNpc),
		GetOptionalDecVal(this.RespawnTime),
		GetOptionalDecVal(this.Direction),
		GetOptionalDecVal(this.DotCount),
		GetOptionalByteArrayVal(this.Path, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this NpcPos) GetInsertHeader() string {
	return "INSERT INTO [K_NPCPOS] ([ZoneID], [NpcID], [ActType], [RegenType], [DungeonFamily], [SpecialType], [TrapNumber], [LeftX], [TopZ], [RightX], [BottomZ], [LimitMinZ], [LimitMinX], [LimitMaxX], [LimitMaxZ], [NumNPC], [RegTime], [byDirection], [DotCnt], [path]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this NpcPos) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(this.ZoneId),
		GetOptionalDecVal(this.NpcId),
		GetOptionalDecVal(this.ActType),
		GetOptionalDecVal(this.RegenType),
		GetOptionalDecVal(this.DungeonFamily),
		GetOptionalDecVal(this.SpecialType),
		GetOptionalDecVal(this.TrapNumber),
		GetOptionalDecVal(this.LeftX),
		GetOptionalDecVal(this.TopZ),
		GetOptionalDecVal(this.RightX),
		GetOptionalDecVal(this.BottomZ),
		GetOptionalDecVal(this.LimitMinZ),
		GetOptionalDecVal(this.LimitMinX),
		GetOptionalDecVal(this.LimitMaxX),
		GetOptionalDecVal(this.LimitMaxZ),
		GetOptionalDecVal(this.NumNpc),
		GetOptionalDecVal(this.RespawnTime),
		GetOptionalDecVal(this.Direction),
		GetOptionalDecVal(this.DotCount),
		GetOptionalByteArrayVal(this.Path, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this NpcPos) GetCreateTableString() string {
	query := "CREATE TABLE [K_NPCPOS] (\n\t[ZoneID] smallint,\n\t[NpcID] int,\n\t[ActType] tinyint,\n\t[RegenType] tinyint,\n\t[DungeonFamily] tinyint,\n\t[SpecialType] tinyint,\n\t[TrapNumber] tinyint,\n\t[LeftX] int,\n\t[TopZ] int,\n\t[RightX] int,\n\t[BottomZ] int,\n\t[LimitMinZ] int,\n\t[LimitMinX] int,\n\t[LimitMaxX] int,\n\t[LimitMaxZ] int,\n\t[NumNPC] tinyint,\n\t[RegTime] smallint,\n\t[byDirection] int,\n\t[DotCnt] tinyint,\n\t[path] text\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this NpcPos) SelectClause() (selectClause clause.Select) {
	return _NpcPosSelectClause
}

// GetAllTableData Returns a list of all table data
func (this NpcPos) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []NpcPos{}
	rawSql := "SELECT [ZoneID], [NpcID], [ActType], [RegenType], [DungeonFamily], [SpecialType], [TrapNumber], [LeftX], [TopZ], [RightX], [BottomZ], [LimitMinZ], [LimitMinX], [LimitMaxX], [LimitMaxZ], [NumNPC], [RegTime], [byDirection], [DotCnt], [path] FROM [K_NPCPOS]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _NpcPosSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[ZoneID]",
		},
		clause.Column{
			Name: "[NpcID]",
		},
		clause.Column{
			Name: "[ActType]",
		},
		clause.Column{
			Name: "[RegenType]",
		},
		clause.Column{
			Name: "[DungeonFamily]",
		},
		clause.Column{
			Name: "[SpecialType]",
		},
		clause.Column{
			Name: "[TrapNumber]",
		},
		clause.Column{
			Name: "[LeftX]",
		},
		clause.Column{
			Name: "[TopZ]",
		},
		clause.Column{
			Name: "[RightX]",
		},
		clause.Column{
			Name: "[BottomZ]",
		},
		clause.Column{
			Name: "[LimitMinZ]",
		},
		clause.Column{
			Name: "[LimitMinX]",
		},
		clause.Column{
			Name: "[LimitMaxX]",
		},
		clause.Column{
			Name: "[LimitMaxZ]",
		},
		clause.Column{
			Name: "[NumNPC]",
		},
		clause.Column{
			Name: "[RegTime]",
		},
		clause.Column{
			Name: "[byDirection]",
		},
		clause.Column{
			Name: "[DotCnt]",
		},
		clause.Column{
			Name: "[path]",
		},
	},
}
