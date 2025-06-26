package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_KnightsDatabaseNbr = 1
	_KnightsTableName   = "KNIGHTS"
)

func init() {
	ModelList = append(ModelList, &Knights{})
}

// Knights Knights are the clan/guild system of the game
type Knights struct {
	IdNumber        int16          `gorm:"column:IDNum;type:smallint;primaryKey;not null" json:"IDNum"`
	Flag            uint8          `gorm:"column:Flag;type:tinyint;not null;default:1" json:"Flag"`
	Nation          uint8          `gorm:"column:Nation;type:tinyint;not null" json:"Nation"`
	Ranking         uint8          `gorm:"column:Ranking;type:tinyint;not null;default:0" json:"Ranking"`
	Name            mssql.VarChar  `gorm:"column:IDName;type:varchar(21);not null" json:"IDName"`
	Members         int16          `gorm:"column:Members;type:smallint;not null;default:1" json:"Members"`
	Chief           mssql.VarChar  `gorm:"column:Chief;type:varchar(21);not null" json:"Chief"`
	ViceChief1      *mssql.VarChar `gorm:"column:ViceChief_1;type:varchar(21)" json:"ViceChief_1,omitempty"`
	ViceChief2      *mssql.VarChar `gorm:"column:ViceChief_2;type:varchar(21)" json:"ViceChief_2,omitempty"`
	ViceChief3      *mssql.VarChar `gorm:"column:ViceChief_3;type:varchar(21)" json:"ViceChief_3,omitempty"`
	EnemyName       *mssql.VarChar `gorm:"column:strEnemyName;type:varchar(21)" json:"strEnemyName,omitempty"`
	OldWarResult    uint8          `gorm:"column:byOldWarResult;type:tinyint;not null;default:0" json:"byOldWarResult"`
	WarEnemyId      int            `gorm:"column:nWarEnemyID;type:int;not null;default:0" json:"nWarEnemyID"`
	Victory         int            `gorm:"column:nVictory;type:int;not null;default:0" json:"nVictory"`
	Lose            int            `gorm:"column:nLose;type:int;not null;default:0" json:"nLose"`
	Gold            int64          `gorm:"column:Gold;type:bigint;not null;default:0" json:"Gold"`
	Domination      int16          `gorm:"column:Domination;type:smallint;not null;default:0" json:"Domination"`
	Points          *int           `gorm:"column:Points;type:int;default:0" json:"Points,omitempty"`
	CreateTime      time.Time      `gorm:"column:CreateTime;type:smalldatetime;not null;default:getdate()" json:"CreateTime"`
	MarkVersion     int16          `gorm:"column:sMarkVersion;type:smallint;not null;default:0" json:"sMarkVersion"`
	MarkLength      int16          `gorm:"column:sMarkLen;type:smallint;not null;default:0" json:"sMarkLen"`
	Mark            *[]byte        `gorm:"column:Mark;type:image" json:"Mark,omitempty"`
	Stash           *mssql.VarChar `gorm:"column:Stash;type:varchar(1600)" json:"Stash,omitempty"`
	SiegeFlag       uint8          `gorm:"column:bySiegeFlag;type:tinyint;not null;default:0" json:"bySiegeFlag"`
	AllianceKnights int16          `gorm:"column:sAllianceKnights;type:smallint;not null;default:0" json:"sAllianceKnights"`
	Cape            int16          `gorm:"column:sCape;type:smallint;not null;default:-1" json:"sCape"`
}

// GetDatabaseName Returns the table's database name
func (this Knights) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsDatabaseNbr))
}

// TableName Returns the table name
func (this Knights) TableName() string {
	return _KnightsTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Knights) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS] ([IDNum], [Flag], [Nation], [Ranking], [IDName], [Members], [Chief], [ViceChief_1], [ViceChief_2], [ViceChief_3], [strEnemyName], [byOldWarResult], [nWarEnemyID], [nVictory], [nLose], [Gold], [Domination], [Points], [CreateTime], [sMarkVersion], [sMarkLen], [Mark], [Stash], [bySiegeFlag], [sAllianceKnights], [sCape]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(varchar(1600), %s), %s, %s, %s)", GetOptionalDecVal(&this.IdNumber),
		GetOptionalDecVal(&this.Flag),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Ranking),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Members),
		GetOptionalVarCharVal(&this.Chief, false),
		GetOptionalVarCharVal(this.ViceChief1, false),
		GetOptionalVarCharVal(this.ViceChief2, false),
		GetOptionalVarCharVal(this.ViceChief3, false),
		GetOptionalVarCharVal(this.EnemyName, false),
		GetOptionalDecVal(&this.OldWarResult),
		GetOptionalDecVal(&this.WarEnemyId),
		GetOptionalDecVal(&this.Victory),
		GetOptionalDecVal(&this.Lose),
		GetOptionalDecVal(&this.Gold),
		GetOptionalDecVal(&this.Domination),
		GetOptionalDecVal(this.Points),
		GetDateTimeExportFmt(&this.CreateTime),
		GetOptionalDecVal(&this.MarkVersion),
		GetOptionalDecVal(&this.MarkLength),
		GetOptionalByteArrayVal(this.Mark, false),
		GetOptionalVarCharVal(this.Stash, true),
		GetOptionalDecVal(&this.SiegeFlag),
		GetOptionalDecVal(&this.AllianceKnights),
		GetOptionalDecVal(&this.Cape))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Knights) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS] (IDNum, Flag, Nation, Ranking, IDName, Members, Chief, ViceChief_1, ViceChief_2, ViceChief_3, strEnemyName, byOldWarResult, nWarEnemyID, nVictory, nLose, Gold, Domination, Points, CreateTime, sMarkVersion, sMarkLen, Mark, Stash, bySiegeFlag, sAllianceKnights, sCape) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Knights) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(varchar(1600), %s), %s, %s, %s)", GetOptionalDecVal(&this.IdNumber),
		GetOptionalDecVal(&this.Flag),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Ranking),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Members),
		GetOptionalVarCharVal(&this.Chief, false),
		GetOptionalVarCharVal(this.ViceChief1, false),
		GetOptionalVarCharVal(this.ViceChief2, false),
		GetOptionalVarCharVal(this.ViceChief3, false),
		GetOptionalVarCharVal(this.EnemyName, false),
		GetOptionalDecVal(&this.OldWarResult),
		GetOptionalDecVal(&this.WarEnemyId),
		GetOptionalDecVal(&this.Victory),
		GetOptionalDecVal(&this.Lose),
		GetOptionalDecVal(&this.Gold),
		GetOptionalDecVal(&this.Domination),
		GetOptionalDecVal(this.Points),
		GetDateTimeExportFmt(&this.CreateTime),
		GetOptionalDecVal(&this.MarkVersion),
		GetOptionalDecVal(&this.MarkLength),
		GetOptionalByteArrayVal(this.Mark, false),
		GetOptionalVarCharVal(this.Stash, true),
		GetOptionalDecVal(&this.SiegeFlag),
		GetOptionalDecVal(&this.AllianceKnights),
		GetOptionalDecVal(&this.Cape))
}

// GetCreateTableString Returns the create table statement for this object
func (this Knights) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS] (\n\t[IDNum] smallint NOT NULL,\n\t[Flag] tinyint NOT NULL,\n\t[Nation] tinyint NOT NULL,\n\t[Ranking] tinyint NOT NULL,\n\t[IDName] varchar(21) NOT NULL,\n\t[Members] smallint NOT NULL,\n\t[Chief] varchar(21) NOT NULL,\n\t[ViceChief_1] varchar(21),\n\t[ViceChief_2] varchar(21),\n\t[ViceChief_3] varchar(21),\n\t[strEnemyName] varchar(21),\n\t[byOldWarResult] tinyint NOT NULL,\n\t[nWarEnemyID] int NOT NULL,\n\t[nVictory] int NOT NULL,\n\t[nLose] int NOT NULL,\n\t[Gold] bigint NOT NULL,\n\t[Domination] smallint NOT NULL,\n\t[Points] int,\n\t[CreateTime] smalldatetime NOT NULL,\n\t[sMarkVersion] smallint NOT NULL,\n\t[sMarkLen] smallint NOT NULL,\n\t[Mark] image,\n\t[Stash] varchar(1600),\n\t[bySiegeFlag] tinyint NOT NULL,\n\t[sAllianceKnights] smallint NOT NULL,\n\t[sCape] smallint NOT NULL\n\tCONSTRAINT [PK_KNIGHTS] PRIMARY KEY ([IDNum]),\n\tCONSTRAINT [IX_KNIGHTS] UNIQUE ([IDName])\n)\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Flag] DEFAULT 1 FOR [Flag]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Ranking] DEFAULT 0 FOR [Ranking]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Members] DEFAULT 1 FOR [Members]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_byOldWarResult] DEFAULT 0 FOR [byOldWarResult]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_nWarEnemyID] DEFAULT 0 FOR [nWarEnemyID]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_nVictory] DEFAULT 0 FOR [nVictory]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_nLose] DEFAULT 0 FOR [nLose]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Gold] DEFAULT 0 FOR [Gold]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Domination] DEFAULT 0 FOR [Domination]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_Points] DEFAULT 0 FOR [Points]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_CreateTime] DEFAULT getdate() FOR [CreateTime]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_sMarkVersion] DEFAULT 0 FOR [sMarkVersion]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_sMarkLen] DEFAULT 0 FOR [sMarkLen]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_bySiegeFlag] DEFAULT 0 FOR [bySiegeFlag]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_sAllianceKnights] DEFAULT 0 FOR [sAllianceKnights]\nGO\nALTER TABLE [KNIGHTS] ADD CONSTRAINT [DF_KNIGHTS_sCape] DEFAULT -1 FOR [sCape]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Knights) SelectClause() (selectClause clause.Select) {
	return _KnightsSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Knights) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Knights{}
	rawSql := "SELECT [IDNum], [Flag], [Nation], [Ranking], [IDName], [Members], [Chief], [ViceChief_1], [ViceChief_2], [ViceChief_3], [strEnemyName], [byOldWarResult], [nWarEnemyID], [nVictory], [nLose], [Gold], [Domination], [Points], [CreateTime], [sMarkVersion], [sMarkLen], [Mark], CONVERT(VARBINARY(1600), [Stash]) as [Stash], [bySiegeFlag], [sAllianceKnights], [sCape] FROM [KNIGHTS]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[IDNum]",
		},
		clause.Column{
			Name: "[Flag]",
		},
		clause.Column{
			Name: "[Nation]",
		},
		clause.Column{
			Name: "[Ranking]",
		},
		clause.Column{
			Name: "[IDName]",
		},
		clause.Column{
			Name: "[Members]",
		},
		clause.Column{
			Name: "[Chief]",
		},
		clause.Column{
			Name: "[ViceChief_1]",
		},
		clause.Column{
			Name: "[ViceChief_2]",
		},
		clause.Column{
			Name: "[ViceChief_3]",
		},
		clause.Column{
			Name: "[strEnemyName]",
		},
		clause.Column{
			Name: "[byOldWarResult]",
		},
		clause.Column{
			Name: "[nWarEnemyID]",
		},
		clause.Column{
			Name: "[nVictory]",
		},
		clause.Column{
			Name: "[nLose]",
		},
		clause.Column{
			Name: "[Gold]",
		},
		clause.Column{
			Name: "[Domination]",
		},
		clause.Column{
			Name: "[Points]",
		},
		clause.Column{
			Name: "[CreateTime]",
		},
		clause.Column{
			Name: "[sMarkVersion]",
		},
		clause.Column{
			Name: "[sMarkLen]",
		},
		clause.Column{
			Name: "[Mark]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1600), [Stash]) as [Stash]",
		},
		clause.Column{
			Name: "[bySiegeFlag]",
		},
		clause.Column{
			Name: "[sAllianceKnights]",
		},
		clause.Column{
			Name: "[sCape]",
		},
	},
}
