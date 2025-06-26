package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_HomeDatabaseNbr = "GAME"
	_HomeTableName   = "HOME"
)

func init() {
	ModelList = append(ModelList, &Home{})
}

// Home TODO Doc
type Home struct {
	Nation        uint8 `gorm:"column:Nation;type:tinyint;not null" json:"Nation"`
	ElmoZoneX     int   `gorm:"column:ElmoZoneX;type:int;not null" json:"ElmoZoneX"`
	ElmoZoneZ     int   `gorm:"column:ElmoZoneZ;type:int;not null" json:"ElmoZoneZ"`
	ElmoZoneLX    uint8 `gorm:"column:ElmoZoneLX;type:tinyint;not null" json:"ElmoZoneLX"`
	ElmoZoneLZ    uint8 `gorm:"column:ElmoZoneLZ;type:tinyint;not null" json:"ElmoZoneLZ"`
	KarusZoneX    int   `gorm:"column:KarusZoneX;type:int;not null" json:"KarusZoneX"`
	KarusZoneZ    int   `gorm:"column:KarusZoneZ;type:int;not null" json:"KarusZoneZ"`
	KarusZoneLX   uint8 `gorm:"column:KarusZoneLX;type:tinyint;not null" json:"KarusZoneLX"`
	KarusZoneLZ   uint8 `gorm:"column:KarusZoneLZ;type:tinyint;not null" json:"KarusZoneLZ"`
	FreeZoneX     int   `gorm:"column:FreeZoneX;type:int;not null" json:"FreeZoneX"`
	FreeZoneZ     int   `gorm:"column:FreeZoneZ;type:int;not null" json:"FreeZoneZ"`
	FreeZoneLX    uint8 `gorm:"column:FreeZoneLX;type:tinyint;not null" json:"FreeZoneLX"`
	FreeZoneLZ    uint8 `gorm:"column:FreeZoneLZ;type:tinyint;not null" json:"FreeZoneLZ"`
	BattleZoneX   int   `gorm:"column:BattleZoneX;type:int;not null" json:"BattleZoneX"`
	BattleZoneZ   int   `gorm:"column:BattleZoneZ;type:int;not null" json:"BattleZoneZ"`
	BattleZoneLX  uint8 `gorm:"column:BattleZoneLX;type:tinyint;not null" json:"BattleZoneLX"`
	BattleZoneLZ  uint8 `gorm:"column:BattleZoneLZ;type:tinyint;not null" json:"BattleZoneLZ"`
	BattleZone2X  int   `gorm:"column:BattleZone2X;type:int;not null" json:"BattleZone2X"`
	BattleZone2Z  int   `gorm:"column:BattleZone2Z;type:int;not null" json:"BattleZone2Z"`
	BattleZone2LX uint8 `gorm:"column:BattleZone2LX;type:tinyint;not null" json:"BattleZone2LX"`
	BattleZone2LZ uint8 `gorm:"column:BattleZone2LZ;type:tinyint;not null" json:"BattleZone2LZ"`
}

// GetDatabaseName Returns the table's database name
func (this Home) GetDatabaseName() string {
	return GetDatabaseName(_HomeDatabaseNbr)
}

// TableName Returns the table name
func (this Home) TableName() string {
	return _HomeTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Home) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [HOME] ([Nation], [ElmoZoneX], [ElmoZoneZ], [ElmoZoneLX], [ElmoZoneLZ], [KarusZoneX], [KarusZoneZ], [KarusZoneLX], [KarusZoneLZ], [FreeZoneX], [FreeZoneZ], [FreeZoneLX], [FreeZoneLZ], [BattleZoneX], [BattleZoneZ], [BattleZoneLX], [BattleZoneLZ], [BattleZone2X], [BattleZone2Z], [BattleZone2LX], [BattleZone2LZ]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.ElmoZoneX),
		GetOptionalDecVal(&this.ElmoZoneZ),
		GetOptionalDecVal(&this.ElmoZoneLX),
		GetOptionalDecVal(&this.ElmoZoneLZ),
		GetOptionalDecVal(&this.KarusZoneX),
		GetOptionalDecVal(&this.KarusZoneZ),
		GetOptionalDecVal(&this.KarusZoneLX),
		GetOptionalDecVal(&this.KarusZoneLZ),
		GetOptionalDecVal(&this.FreeZoneX),
		GetOptionalDecVal(&this.FreeZoneZ),
		GetOptionalDecVal(&this.FreeZoneLX),
		GetOptionalDecVal(&this.FreeZoneLZ),
		GetOptionalDecVal(&this.BattleZoneX),
		GetOptionalDecVal(&this.BattleZoneZ),
		GetOptionalDecVal(&this.BattleZoneLX),
		GetOptionalDecVal(&this.BattleZoneLZ),
		GetOptionalDecVal(&this.BattleZone2X),
		GetOptionalDecVal(&this.BattleZone2Z),
		GetOptionalDecVal(&this.BattleZone2LX),
		GetOptionalDecVal(&this.BattleZone2LZ))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Home) GetInsertHeader() string {
	return "INSERT INTO [HOME] ([Nation], [ElmoZoneX], [ElmoZoneZ], [ElmoZoneLX], [ElmoZoneLZ], [KarusZoneX], [KarusZoneZ], [KarusZoneLX], [KarusZoneLZ], [FreeZoneX], [FreeZoneZ], [FreeZoneLX], [FreeZoneLZ], [BattleZoneX], [BattleZoneZ], [BattleZoneLX], [BattleZoneLZ], [BattleZone2X], [BattleZone2Z], [BattleZone2LX], [BattleZone2LZ]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Home) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.ElmoZoneX),
		GetOptionalDecVal(&this.ElmoZoneZ),
		GetOptionalDecVal(&this.ElmoZoneLX),
		GetOptionalDecVal(&this.ElmoZoneLZ),
		GetOptionalDecVal(&this.KarusZoneX),
		GetOptionalDecVal(&this.KarusZoneZ),
		GetOptionalDecVal(&this.KarusZoneLX),
		GetOptionalDecVal(&this.KarusZoneLZ),
		GetOptionalDecVal(&this.FreeZoneX),
		GetOptionalDecVal(&this.FreeZoneZ),
		GetOptionalDecVal(&this.FreeZoneLX),
		GetOptionalDecVal(&this.FreeZoneLZ),
		GetOptionalDecVal(&this.BattleZoneX),
		GetOptionalDecVal(&this.BattleZoneZ),
		GetOptionalDecVal(&this.BattleZoneLX),
		GetOptionalDecVal(&this.BattleZoneLZ),
		GetOptionalDecVal(&this.BattleZone2X),
		GetOptionalDecVal(&this.BattleZone2Z),
		GetOptionalDecVal(&this.BattleZone2LX),
		GetOptionalDecVal(&this.BattleZone2LZ))
}

// GetCreateTableString Returns the create table statement for this object
func (this Home) GetCreateTableString() string {
	query := "CREATE TABLE [HOME] (\n\t[Nation] tinyint NOT NULL,\n\t[ElmoZoneX] int NOT NULL,\n\t[ElmoZoneZ] int NOT NULL,\n\t[ElmoZoneLX] tinyint NOT NULL,\n\t[ElmoZoneLZ] tinyint NOT NULL,\n\t[KarusZoneX] int NOT NULL,\n\t[KarusZoneZ] int NOT NULL,\n\t[KarusZoneLX] tinyint NOT NULL,\n\t[KarusZoneLZ] tinyint NOT NULL,\n\t[FreeZoneX] int NOT NULL,\n\t[FreeZoneZ] int NOT NULL,\n\t[FreeZoneLX] tinyint NOT NULL,\n\t[FreeZoneLZ] tinyint NOT NULL,\n\t[BattleZoneX] int NOT NULL,\n\t[BattleZoneZ] int NOT NULL,\n\t[BattleZoneLX] tinyint NOT NULL,\n\t[BattleZoneLZ] tinyint NOT NULL,\n\t[BattleZone2X] int NOT NULL,\n\t[BattleZone2Z] int NOT NULL,\n\t[BattleZone2LX] tinyint NOT NULL,\n\t[BattleZone2LZ] tinyint NOT NULL\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Home) SelectClause() (selectClause clause.Select) {
	return _HomeSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Home) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Home{}
	rawSql := "SELECT [Nation], [ElmoZoneX], [ElmoZoneZ], [ElmoZoneLX], [ElmoZoneLZ], [KarusZoneX], [KarusZoneZ], [KarusZoneLX], [KarusZoneLZ], [FreeZoneX], [FreeZoneZ], [FreeZoneLX], [FreeZoneLZ], [BattleZoneX], [BattleZoneZ], [BattleZoneLX], [BattleZoneLZ], [BattleZone2X], [BattleZone2Z], [BattleZone2LX], [BattleZone2LZ] FROM [HOME]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _HomeSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[Nation]",
		},
		clause.Column{
			Name: "[ElmoZoneX]",
		},
		clause.Column{
			Name: "[ElmoZoneZ]",
		},
		clause.Column{
			Name: "[ElmoZoneLX]",
		},
		clause.Column{
			Name: "[ElmoZoneLZ]",
		},
		clause.Column{
			Name: "[KarusZoneX]",
		},
		clause.Column{
			Name: "[KarusZoneZ]",
		},
		clause.Column{
			Name: "[KarusZoneLX]",
		},
		clause.Column{
			Name: "[KarusZoneLZ]",
		},
		clause.Column{
			Name: "[FreeZoneX]",
		},
		clause.Column{
			Name: "[FreeZoneZ]",
		},
		clause.Column{
			Name: "[FreeZoneLX]",
		},
		clause.Column{
			Name: "[FreeZoneLZ]",
		},
		clause.Column{
			Name: "[BattleZoneX]",
		},
		clause.Column{
			Name: "[BattleZoneZ]",
		},
		clause.Column{
			Name: "[BattleZoneLX]",
		},
		clause.Column{
			Name: "[BattleZoneLZ]",
		},
		clause.Column{
			Name: "[BattleZone2X]",
		},
		clause.Column{
			Name: "[BattleZone2Z]",
		},
		clause.Column{
			Name: "[BattleZone2LX]",
		},
		clause.Column{
			Name: "[BattleZone2LZ]",
		},
	},
}
