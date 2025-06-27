package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

const (
	_UserDataDatabaseNbr = "GAME"
	_UserDataTableName   = "USERDATA"
)

func init() {
	ModelList = append(ModelList, &UserData{})
}

// UserData User data contains saved character information
type UserData struct {
	UserId         mssql.VarChar  `gorm:"column:strUserId;type:varchar(21);primaryKey;not null" json:"strUserId"`
	Nation         uint8          `gorm:"column:Nation;type:tinyint;not null;default:0" json:"Nation"`
	Race           uint8          `gorm:"column:Race;type:tinyint;not null;default:1" json:"Race"`
	Class          int16          `gorm:"column:Class;type:smallint;not null;default:0" json:"Class"`
	HairColor      uint8          `gorm:"column:HairColor;type:tinyint;not null;default:0" json:"HairColor"`
	Rank           uint8          `gorm:"column:Rank;type:tinyint;not null;default:0" json:"Rank"`
	Title          uint8          `gorm:"column:Title;type:tinyint;not null;default:0" json:"Title"`
	Level          uint8          `gorm:"column:Level;type:tinyint;not null;default:1" json:"Level"`
	Exp            int            `gorm:"column:Exp;type:int;not null;default:5" json:"Exp"`
	Loyalty        int            `gorm:"column:Loyalty;type:int;not null;default:500" json:"Loyalty"`
	Face           uint8          `gorm:"column:Face;type:tinyint;not null;default:0" json:"Face"`
	City           uint8          `gorm:"column:City;type:tinyint;not null;default:0" json:"City"`
	KnightsId      int16          `gorm:"column:Knights;type:smallint;not null;default:0" json:"Knights"`
	Fame           uint8          `gorm:"column:Fame;type:tinyint;not null;default:0" json:"Fame"`
	Hp             int16          `gorm:"column:Hp;type:smallint;not null;default:100" json:"Hp"`
	Mp             int16          `gorm:"column:Mp;type:smallint;not null;default:100" json:"Mp"`
	Sp             int16          `gorm:"column:Sp;type:smallint;not null;default:100" json:"Sp"`
	Strength       uint8          `gorm:"column:Strong;type:tinyint;not null;default:0" json:"Strong"`
	Stamina        uint8          `gorm:"column:Sta;type:tinyint;not null;default:0" json:"Sta"`
	Dexterity      uint8          `gorm:"column:Dex;type:tinyint;not null;default:0" json:"Dex"`
	Intelligence   uint8          `gorm:"column:Intel;type:tinyint;not null;default:0" json:"Intel"`
	Charisma       uint8          `gorm:"column:Cha;type:tinyint;not null;default:0" json:"Cha"`
	Authority      uint8          `gorm:"column:Authority;type:tinyint;not null;default:1" json:"Authority"`
	StatPoints     uint8          `gorm:"column:Points;type:tinyint;not null;default:0" json:"Points"`
	Gold           int            `gorm:"column:Gold;type:int;not null;default:50000" json:"Gold"`
	Zone           uint8          `gorm:"column:Zone;type:tinyint;not null;default:1" json:"Zone"`
	Bind           *int16         `gorm:"column:Bind;type:smallint" json:"Bind,omitempty"`
	PosX           int            `gorm:"column:PX;type:int;not null;default:268100" json:"PX"`
	PosZ           int            `gorm:"column:PZ;type:int;not null;default:131000" json:"PZ"`
	PosY           int            `gorm:"column:PY;type:int;not null;default:0" json:"PY"`
	DwTime         int            `gorm:"column:dwTime;type:int;not null;default:0" json:"dwTime"`
	SkillData      *mssql.VarChar `gorm:"column:strSkill;type:varchar(10);default:0x00" json:"strSkill,omitempty"`
	ItemData       *mssql.VarChar `gorm:"column:strItem;type:varchar(400)" json:"strItem,omitempty"`
	Serial         *mssql.VarChar `gorm:"column:strSerial;type:varchar(400)" json:"strSerial,omitempty"`
	QuestCount     int16          `gorm:"column:sQuestCount;type:smallint;not null;default:0" json:"sQuestCount"`
	QuestData      *mssql.VarChar `gorm:"column:strQuest;type:varchar(400)" json:"strQuest,omitempty"`
	MannerPoint    int            `gorm:"column:MannerPoint;type:int;not null;default:0" json:"MannerPoint"`
	LoyaltyMonthly int            `gorm:"column:LoyaltyMonthly;type:int;not null;default:0" json:"LoyaltyMonthly"`
	CreateTime     time.Time      `gorm:"column:CreateTime;type:smalldatetime;not null;default:getdate()" json:"CreateTime"`
	UpdateTime     *time.Time     `gorm:"column:UpdateTime;type:smalldatetime" json:"UpdateTime,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this UserData) GetDatabaseName() string {
	return GetDatabaseName(_UserDataDatabaseNbr)
}

// TableName Returns the table name
func (this UserData) TableName() string {
	return _UserDataTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this UserData) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USERDATA] ([strUserId], [Nation], [Race], [Class], [HairColor], [Rank], [Title], [Level], [Exp], [Loyalty], [Face], [City], [Knights], [Fame], [Hp], [Mp], [Sp], [Strong], [Sta], [Dex], [Intel], [Cha], [Authority], [Points], [Gold], [Zone], [Bind], [PX], [PZ], [PY], [dwTime], [strSkill], [strItem], [strSerial], [sQuestCount], [strQuest], [MannerPoint], [LoyaltyMonthly], [CreateTime], [UpdateTime]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(varchar(10), %s), CONVERT(varchar(400), %s), CONVERT(varchar(400), %s), %s, CONVERT(varchar(400), %s), %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Race),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.HairColor),
		GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(&this.Title),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Exp),
		GetOptionalDecVal(&this.Loyalty),
		GetOptionalDecVal(&this.Face),
		GetOptionalDecVal(&this.City),
		GetOptionalDecVal(&this.KnightsId),
		GetOptionalDecVal(&this.Fame),
		GetOptionalDecVal(&this.Hp),
		GetOptionalDecVal(&this.Mp),
		GetOptionalDecVal(&this.Sp),
		GetOptionalDecVal(&this.Strength),
		GetOptionalDecVal(&this.Stamina),
		GetOptionalDecVal(&this.Dexterity),
		GetOptionalDecVal(&this.Intelligence),
		GetOptionalDecVal(&this.Charisma),
		GetOptionalDecVal(&this.Authority),
		GetOptionalDecVal(&this.StatPoints),
		GetOptionalDecVal(&this.Gold),
		GetOptionalDecVal(&this.Zone),
		GetOptionalDecVal(this.Bind),
		GetOptionalDecVal(&this.PosX),
		GetOptionalDecVal(&this.PosZ),
		GetOptionalDecVal(&this.PosY),
		GetOptionalDecVal(&this.DwTime),
		GetOptionalVarCharVal(this.SkillData, true),
		GetOptionalVarCharVal(this.ItemData, true),
		GetOptionalVarCharVal(this.Serial, true),
		GetOptionalDecVal(&this.QuestCount),
		GetOptionalVarCharVal(this.QuestData, true),
		GetOptionalDecVal(&this.MannerPoint),
		GetOptionalDecVal(&this.LoyaltyMonthly),
		GetDateTimeExportFmt(&this.CreateTime),
		GetDateTimeExportFmt(this.UpdateTime))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this UserData) GetInsertHeader() string {
	return "INSERT INTO [USERDATA] ([strUserId], [Nation], [Race], [Class], [HairColor], [Rank], [Title], [Level], [Exp], [Loyalty], [Face], [City], [Knights], [Fame], [Hp], [Mp], [Sp], [Strong], [Sta], [Dex], [Intel], [Cha], [Authority], [Points], [Gold], [Zone], [Bind], [PX], [PZ], [PY], [dwTime], [strSkill], [strItem], [strSerial], [sQuestCount], [strQuest], [MannerPoint], [LoyaltyMonthly], [CreateTime], [UpdateTime]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this UserData) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(varchar(10), %s), CONVERT(varchar(400), %s), CONVERT(varchar(400), %s), %s, CONVERT(varchar(400), %s), %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.Race),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.HairColor),
		GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(&this.Title),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Exp),
		GetOptionalDecVal(&this.Loyalty),
		GetOptionalDecVal(&this.Face),
		GetOptionalDecVal(&this.City),
		GetOptionalDecVal(&this.KnightsId),
		GetOptionalDecVal(&this.Fame),
		GetOptionalDecVal(&this.Hp),
		GetOptionalDecVal(&this.Mp),
		GetOptionalDecVal(&this.Sp),
		GetOptionalDecVal(&this.Strength),
		GetOptionalDecVal(&this.Stamina),
		GetOptionalDecVal(&this.Dexterity),
		GetOptionalDecVal(&this.Intelligence),
		GetOptionalDecVal(&this.Charisma),
		GetOptionalDecVal(&this.Authority),
		GetOptionalDecVal(&this.StatPoints),
		GetOptionalDecVal(&this.Gold),
		GetOptionalDecVal(&this.Zone),
		GetOptionalDecVal(this.Bind),
		GetOptionalDecVal(&this.PosX),
		GetOptionalDecVal(&this.PosZ),
		GetOptionalDecVal(&this.PosY),
		GetOptionalDecVal(&this.DwTime),
		GetOptionalVarCharVal(this.SkillData, true),
		GetOptionalVarCharVal(this.ItemData, true),
		GetOptionalVarCharVal(this.Serial, true),
		GetOptionalDecVal(&this.QuestCount),
		GetOptionalVarCharVal(this.QuestData, true),
		GetOptionalDecVal(&this.MannerPoint),
		GetOptionalDecVal(&this.LoyaltyMonthly),
		GetDateTimeExportFmt(&this.CreateTime),
		GetDateTimeExportFmt(this.UpdateTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this UserData) GetCreateTableString() string {
	query := "CREATE TABLE [USERDATA] (\n\t[strUserId] varchar(21) NOT NULL,\n\t[Nation] tinyint NOT NULL,\n\t[Race] tinyint NOT NULL,\n\t[Class] smallint NOT NULL,\n\t[HairColor] tinyint NOT NULL,\n\t[Rank] tinyint NOT NULL,\n\t[Title] tinyint NOT NULL,\n\t[Level] tinyint NOT NULL,\n\t[Exp] int NOT NULL,\n\t[Loyalty] int NOT NULL,\n\t[Face] tinyint NOT NULL,\n\t[City] tinyint NOT NULL,\n\t[Knights] smallint NOT NULL,\n\t[Fame] tinyint NOT NULL,\n\t[Hp] smallint NOT NULL,\n\t[Mp] smallint NOT NULL,\n\t[Sp] smallint NOT NULL,\n\t[Strong] tinyint NOT NULL,\n\t[Sta] tinyint NOT NULL,\n\t[Dex] tinyint NOT NULL,\n\t[Intel] tinyint NOT NULL,\n\t[Cha] tinyint NOT NULL,\n\t[Authority] tinyint NOT NULL,\n\t[Points] tinyint NOT NULL,\n\t[Gold] int NOT NULL,\n\t[Zone] tinyint NOT NULL,\n\t[Bind] smallint,\n\t[PX] int NOT NULL,\n\t[PZ] int NOT NULL,\n\t[PY] int NOT NULL,\n\t[dwTime] int NOT NULL,\n\t[strSkill] varchar(10),\n\t[strItem] varchar(400),\n\t[strSerial] varchar(400),\n\t[sQuestCount] smallint NOT NULL,\n\t[strQuest] varchar(400),\n\t[MannerPoint] int NOT NULL,\n\t[LoyaltyMonthly] int NOT NULL,\n\t[CreateTime] smalldatetime NOT NULL,\n\t[UpdateTime] smalldatetime\n\tCONSTRAINT [PK_USERDATA] PRIMARY KEY CLUSTERED ([strUserId])\n)\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Nation] DEFAULT 0 FOR [Nation]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Race] DEFAULT 1 FOR [Race]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Class] DEFAULT 0 FOR [Class]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_HairColor] DEFAULT 0 FOR [HairColor]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Rank] DEFAULT 0 FOR [Rank]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Title] DEFAULT 0 FOR [Title]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Level] DEFAULT 1 FOR [Level]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Exp] DEFAULT 5 FOR [Exp]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Loyalty] DEFAULT 500 FOR [Loyalty]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Face] DEFAULT 0 FOR [Face]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_City] DEFAULT 0 FOR [City]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Knights] DEFAULT 0 FOR [Knights]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Fame] DEFAULT 0 FOR [Fame]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Hp] DEFAULT 100 FOR [Hp]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Mp] DEFAULT 100 FOR [Mp]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Sp] DEFAULT 100 FOR [Sp]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Strong] DEFAULT 0 FOR [Strong]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Sta] DEFAULT 0 FOR [Sta]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Dex] DEFAULT 0 FOR [Dex]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Intel] DEFAULT 0 FOR [Intel]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Cha] DEFAULT 0 FOR [Cha]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Authority] DEFAULT 1 FOR [Authority]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Points] DEFAULT 0 FOR [Points]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Gold] DEFAULT 50000 FOR [Gold]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_Zone] DEFAULT 1 FOR [Zone]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_PX] DEFAULT 268100 FOR [PX]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_PZ] DEFAULT 131000 FOR [PZ]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_PY] DEFAULT 0 FOR [PY]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_dwTime] DEFAULT 0 FOR [dwTime]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_strSkill] DEFAULT 0x00 FOR [strSkill]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_sQuestCount] DEFAULT 0 FOR [sQuestCount]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_MannerPoint] DEFAULT 0 FOR [MannerPoint]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_LoyaltyMonthly] DEFAULT 0 FOR [LoyaltyMonthly]\nGO\nALTER TABLE [USERDATA] ADD CONSTRAINT [DF_USERDATA_CreateTime] DEFAULT getdate() FOR [CreateTime]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this UserData) SelectClause() (selectClause clause.Select) {
	return _UserDataSelectClause
}

// GetAllTableData Returns a list of all table data
func (this UserData) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []UserData{}
	rawSql := "SELECT [strUserId], [Nation], [Race], [Class], [HairColor], [Rank], [Title], [Level], [Exp], [Loyalty], [Face], [City], [Knights], [Fame], [Hp], [Mp], [Sp], [Strong], [Sta], [Dex], [Intel], [Cha], [Authority], [Points], [Gold], [Zone], [Bind], [PX], [PZ], [PY], [dwTime], CONVERT(VARBINARY(10), [strSkill]) as [strSkill], CONVERT(VARBINARY(400), [strItem]) as [strItem], CONVERT(VARBINARY(400), [strSerial]) as [strSerial], [sQuestCount], CONVERT(VARBINARY(400), [strQuest]) as [strQuest], [MannerPoint], [LoyaltyMonthly], [CreateTime], [UpdateTime] FROM [USERDATA]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _UserDataSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strUserId]",
		},
		clause.Column{
			Name: "[Nation]",
		},
		clause.Column{
			Name: "[Race]",
		},
		clause.Column{
			Name: "[Class]",
		},
		clause.Column{
			Name: "[HairColor]",
		},
		clause.Column{
			Name: "[Rank]",
		},
		clause.Column{
			Name: "[Title]",
		},
		clause.Column{
			Name: "[Level]",
		},
		clause.Column{
			Name: "[Exp]",
		},
		clause.Column{
			Name: "[Loyalty]",
		},
		clause.Column{
			Name: "[Face]",
		},
		clause.Column{
			Name: "[City]",
		},
		clause.Column{
			Name: "[Knights]",
		},
		clause.Column{
			Name: "[Fame]",
		},
		clause.Column{
			Name: "[Hp]",
		},
		clause.Column{
			Name: "[Mp]",
		},
		clause.Column{
			Name: "[Sp]",
		},
		clause.Column{
			Name: "[Strong]",
		},
		clause.Column{
			Name: "[Sta]",
		},
		clause.Column{
			Name: "[Dex]",
		},
		clause.Column{
			Name: "[Intel]",
		},
		clause.Column{
			Name: "[Cha]",
		},
		clause.Column{
			Name: "[Authority]",
		},
		clause.Column{
			Name: "[Points]",
		},
		clause.Column{
			Name: "[Gold]",
		},
		clause.Column{
			Name: "[Zone]",
		},
		clause.Column{
			Name: "[Bind]",
		},
		clause.Column{
			Name: "[PX]",
		},
		clause.Column{
			Name: "[PZ]",
		},
		clause.Column{
			Name: "[PY]",
		},
		clause.Column{
			Name: "[dwTime]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(10), [strSkill]) as [strSkill]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(400), [strItem]) as [strItem]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(400), [strSerial]) as [strSerial]",
		},
		clause.Column{
			Name: "[sQuestCount]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(400), [strQuest]) as [strQuest]",
		},
		clause.Column{
			Name: "[MannerPoint]",
		},
		clause.Column{
			Name: "[LoyaltyMonthly]",
		},
		clause.Column{
			Name: "[CreateTime]",
		},
		clause.Column{
			Name: "[UpdateTime]",
		},
	},
}
