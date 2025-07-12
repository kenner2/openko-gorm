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
	ZoneId        int16   `gorm:"column:ZoneID;type:smallint;not null" json:"ZoneID"`
	NpcId         int     `gorm:"column:NpcID;type:int;not null" json:"NpcID"`
	ActType       uint8   `gorm:"column:ActType;type:tinyint;not null" json:"ActType"`
	RegenType     uint8   `gorm:"column:RegenType;type:tinyint;not null" json:"RegenType"`
	DungeonFamily uint8   `gorm:"column:DungeonFamily;type:tinyint;not null" json:"DungeonFamily"`
	SpecialType   uint8   `gorm:"column:SpecialType;type:tinyint;not null" json:"SpecialType"`
	TrapNumber    uint8   `gorm:"column:TrapNumber;type:tinyint;not null" json:"TrapNumber"`
	LeftX         int     `gorm:"column:LeftX;type:int;not null" json:"LeftX"`
	TopZ          int     `gorm:"column:TopZ;type:int;not null" json:"TopZ"`
	RightX        int     `gorm:"column:RightX;type:int;not null" json:"RightX"`
	BottomZ       int     `gorm:"column:BottomZ;type:int;not null" json:"BottomZ"`
	LimitMinZ     int     `gorm:"column:LimitMinZ;type:int;not null" json:"LimitMinZ"`
	LimitMinX     int     `gorm:"column:LimitMinX;type:int;not null" json:"LimitMinX"`
	LimitMaxX     int     `gorm:"column:LimitMaxX;type:int;not null" json:"LimitMaxX"`
	LimitMaxZ     int     `gorm:"column:LimitMaxZ;type:int;not null" json:"LimitMaxZ"`
	NumNpc        uint8   `gorm:"column:NumNPC;type:tinyint;not null" json:"NumNPC"`
	RespawnTime   int16   `gorm:"column:RegTime;type:smallint;not null" json:"RegTime"`
	Direction     int     `gorm:"column:byDirection;type:int;not null" json:"byDirection"`
	DotCount      uint8   `gorm:"column:DotCnt;type:tinyint;not null" json:"DotCnt"`
	Path          *[]byte `gorm:"column:path;type:text COLLATE SQL_Latin1_General_CP1_CI_AS" json:"path,omitempty"`
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
	return fmt.Sprintf("INSERT INTO [K_NPCPOS] ([ZoneID], [NpcID], [ActType], [RegenType], [DungeonFamily], [SpecialType], [TrapNumber], [LeftX], [TopZ], [RightX], [BottomZ], [LimitMinZ], [LimitMinX], [LimitMaxX], [LimitMaxZ], [NumNPC], [RegTime], [byDirection], [DotCnt], [path]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneId),
		GetOptionalDecVal(&this.NpcId),
		GetOptionalDecVal(&this.ActType),
		GetOptionalDecVal(&this.RegenType),
		GetOptionalDecVal(&this.DungeonFamily),
		GetOptionalDecVal(&this.SpecialType),
		GetOptionalDecVal(&this.TrapNumber),
		GetOptionalDecVal(&this.LeftX),
		GetOptionalDecVal(&this.TopZ),
		GetOptionalDecVal(&this.RightX),
		GetOptionalDecVal(&this.BottomZ),
		GetOptionalDecVal(&this.LimitMinZ),
		GetOptionalDecVal(&this.LimitMinX),
		GetOptionalDecVal(&this.LimitMaxX),
		GetOptionalDecVal(&this.LimitMaxZ),
		GetOptionalDecVal(&this.NumNpc),
		GetOptionalDecVal(&this.RespawnTime),
		GetOptionalDecVal(&this.Direction),
		GetOptionalDecVal(&this.DotCount),
		GetOptionalByteArrayVal(this.Path, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this NpcPos) GetInsertHeader() string {
	return "INSERT INTO [K_NPCPOS] ([ZoneID], [NpcID], [ActType], [RegenType], [DungeonFamily], [SpecialType], [TrapNumber], [LeftX], [TopZ], [RightX], [BottomZ], [LimitMinZ], [LimitMinX], [LimitMaxX], [LimitMaxZ], [NumNPC], [RegTime], [byDirection], [DotCnt], [path]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this NpcPos) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneId),
		GetOptionalDecVal(&this.NpcId),
		GetOptionalDecVal(&this.ActType),
		GetOptionalDecVal(&this.RegenType),
		GetOptionalDecVal(&this.DungeonFamily),
		GetOptionalDecVal(&this.SpecialType),
		GetOptionalDecVal(&this.TrapNumber),
		GetOptionalDecVal(&this.LeftX),
		GetOptionalDecVal(&this.TopZ),
		GetOptionalDecVal(&this.RightX),
		GetOptionalDecVal(&this.BottomZ),
		GetOptionalDecVal(&this.LimitMinZ),
		GetOptionalDecVal(&this.LimitMinX),
		GetOptionalDecVal(&this.LimitMaxX),
		GetOptionalDecVal(&this.LimitMaxZ),
		GetOptionalDecVal(&this.NumNpc),
		GetOptionalDecVal(&this.RespawnTime),
		GetOptionalDecVal(&this.Direction),
		GetOptionalDecVal(&this.DotCount),
		GetOptionalByteArrayVal(this.Path, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this NpcPos) GetCreateTableString() string {
	query := "CREATE TABLE [K_NPCPOS] (\n\t[ZoneID] smallint NOT NULL,\n\t[NpcID] int NOT NULL,\n\t[ActType] tinyint NOT NULL,\n\t[RegenType] tinyint NOT NULL,\n\t[DungeonFamily] tinyint NOT NULL,\n\t[SpecialType] tinyint NOT NULL,\n\t[TrapNumber] tinyint NOT NULL,\n\t[LeftX] int NOT NULL,\n\t[TopZ] int NOT NULL,\n\t[RightX] int NOT NULL,\n\t[BottomZ] int NOT NULL,\n\t[LimitMinZ] int NOT NULL,\n\t[LimitMinX] int NOT NULL,\n\t[LimitMaxX] int NOT NULL,\n\t[LimitMaxZ] int NOT NULL,\n\t[NumNPC] tinyint NOT NULL,\n\t[RegTime] smallint NOT NULL,\n\t[byDirection] int NOT NULL,\n\t[DotCnt] tinyint NOT NULL,\n\t[path] text COLLATE SQL_Latin1_General_CP1_CI_AS\n)\nGO\n"
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
