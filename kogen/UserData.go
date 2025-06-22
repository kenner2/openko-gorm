package kogen

import (
	"fmt"
)

const (
	_UserDataTableName   = "USERDATA"
	_UserDataDatabaseNbr = 0
)

func init() {
	ModelList = append(ModelList, &UserData{})
}

// UserData User data contains saved character information
type UserData struct {
	UserId         [21]byte  `gorm:"column:strUserId;type:varchar(21);primaryKey;not null" json:"strUserId"`
	Nation         uint8     `gorm:"column:Nation;type:tinyint;not null;default:0" json:"Nation"`
	Race           uint8     `gorm:"column:Race;type:tinyint;not null;default:1" json:"Race"`
	Class          int16     `gorm:"column:Class;type:smallint;not null;default:0" json:"Class"`
	HairColor      uint8     `gorm:"column:HairColor;type:tinyint;not null;default:0" json:"HairColor"`
	Rank           uint8     `gorm:"column:Rank;type:tinyint;not null;default:0" json:"Rank"`
	Title          uint8     `gorm:"column:Title;type:tinyint;not null;default:0" json:"Title"`
	Level          uint8     `gorm:"column:Level;type:tinyint;not null;default:1" json:"Level"`
	Exp            int       `gorm:"column:Exp;type:int;not null;default:5" json:"Exp"`
	Loyalty        int       `gorm:"column:Loyalty;type:int;not null;default:500" json:"Loyalty"`
	Face           uint8     `gorm:"column:Face;type:tinyint;not null;default:0" json:"Face"`
	City           uint8     `gorm:"column:City;type:tinyint;not null;default:0" json:"City"`
	KnightsId      int16     `gorm:"column:Knights;type:smallint;not null;default:0" json:"Knights"`
	Fame           uint8     `gorm:"column:Fame;type:tinyint;not null;default:0" json:"Fame"`
	Hp             int16     `gorm:"column:Hp;type:smallint;not null;default:100" json:"Hp"`
	Mp             int16     `gorm:"column:Mp;type:smallint;not null;default:100" json:"Mp"`
	Sp             int16     `gorm:"column:Sp;type:smallint;not null;default:100" json:"Sp"`
	Strength       uint8     `gorm:"column:Strong;type:tinyint;not null;default:0" json:"Strong"`
	Stamina        uint8     `gorm:"column:Sta;type:tinyint;not null;default:0" json:"Sta"`
	Dexterity      uint8     `gorm:"column:Dex;type:tinyint;not null;default:0" json:"Dex"`
	Intelligence   uint8     `gorm:"column:Intel;type:tinyint;not null;default:0" json:"Intel"`
	Charisma       uint8     `gorm:"column:Cha;type:tinyint;not null;default:0" json:"Cha"`
	Authority      uint8     `gorm:"column:Authority;type:tinyint;not null;default:1" json:"Authority"`
	StatPoints     uint8     `gorm:"column:Points;type:tinyint;not null;default:0" json:"Points"`
	Gold           int       `gorm:"column:Gold;type:int;not null;default:50000" json:"Gold"`
	Zone           uint8     `gorm:"column:Zone;type:tinyint;not null;default:1" json:"Zone"`
	Bind           *int16    `gorm:"column:Bind;type:smallint" json:"Bind,omitempty"`
	PosX           int       `gorm:"column:PX;type:int;not null;default:268100" json:"PX"`
	PosZ           int       `gorm:"column:PZ;type:int;not null;default:131000" json:"PZ"`
	PosY           int       `gorm:"column:PY;type:int;not null;default:0" json:"PY"`
	DwTime         int       `gorm:"column:dwTime;type:int;not null;default:0" json:"dwTime"`
	SkillData      [10]byte  `gorm:"column:strSkill;type:varchar(10);default:0x00" json:"strSkill,omitempty"`
	ItemData       [400]byte `gorm:"column:strItem;type:varchar(400)" json:"strItem,omitempty"`
	Serial         [400]byte `gorm:"column:strSerial;type:varchar(400)" json:"strSerial,omitempty"`
	QuestCount     int16     `gorm:"column:sQuestCount;type:smallint;not null;default:0" json:"sQuestCount"`
	QuestData      [400]byte `gorm:"column:strQuest;type:varchar(400)" json:"strQuest,omitempty"`
	MannerPoint    int       `gorm:"column:MannerPoint;type:int;not null;default:0" json:"MannerPoint"`
	LoyaltyMonthly int       `gorm:"column:LoyaltyMonthly;type:int;not null;default:0" json:"LoyaltyMonthly"`
	CreateTime     int       `gorm:"column:CreateTime;type:smalldatetime;not null;default:getdate()" json:"CreateTime"`
	UpdateTime     *int      `gorm:"column:UpdateTime;type:smalldatetime" json:"UpdateTime,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *UserData) GetDatabaseName() string {
	return GetDatabaseName(DbType(_UserDataDatabaseNbr))
}

// GetTableName Returns the table name
func (this *UserData) GetTableName() string {
	return _UserDataTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *UserData) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [USERDATA] (strUserId, Nation, Race, Class, HairColor, Rank, Title, Level, Exp, Loyalty, Face, City, Knights, Fame, Hp, Mp, Sp, Strong, Sta, Dex, Intel, Cha, Authority, Points, Gold, Zone, Bind, PX, PZ, PY, dwTime, strSkill, strItem, strSerial, sQuestCount, strQuest, MannerPoint, LoyaltyMonthly, CreateTime, UpdateTime) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
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
		GetOptionalBinaryVal(this.SkillData),
		GetOptionalBinaryVal(this.ItemData),
		GetOptionalBinaryVal(this.Serial),
		GetOptionalDecVal(&this.QuestCount),
		GetOptionalBinaryVal(this.QuestData),
		GetOptionalDecVal(&this.MannerPoint),
		GetOptionalDecVal(&this.LoyaltyMonthly),
		GetOptionalDecVal(&this.CreateTime),
		GetOptionalDecVal(this.UpdateTime))
}

// GetCreateTableString Returns the create table statement for this object
func (this *UserData) GetCreateTableString() string {
	query := "CREATE TABLE [USERDATA] (\n\t\"strUserId\" varchar(21) NOT NULL,\n\t\"Nation\" tinyint NOT NULL,\n\t\"Race\" tinyint NOT NULL,\n\t\"Class\" smallint NOT NULL,\n\t\"HairColor\" tinyint NOT NULL,\n\t\"Rank\" tinyint NOT NULL,\n\t\"Title\" tinyint NOT NULL,\n\t\"Level\" tinyint NOT NULL,\n\t\"Exp\" int NOT NULL,\n\t\"Loyalty\" int NOT NULL,\n\t\"Face\" tinyint NOT NULL,\n\t\"City\" tinyint NOT NULL,\n\t\"Knights\" smallint NOT NULL,\n\t\"Fame\" tinyint NOT NULL,\n\t\"Hp\" smallint NOT NULL,\n\t\"Mp\" smallint NOT NULL,\n\t\"Sp\" smallint NOT NULL,\n\t\"Strong\" tinyint NOT NULL,\n\t\"Sta\" tinyint NOT NULL,\n\t\"Dex\" tinyint NOT NULL,\n\t\"Intel\" tinyint NOT NULL,\n\t\"Cha\" tinyint NOT NULL,\n\t\"Authority\" tinyint NOT NULL,\n\t\"Points\" tinyint NOT NULL,\n\t\"Gold\" int NOT NULL,\n\t\"Zone\" tinyint NOT NULL,\n\t\"Bind\" smallint,\n\t\"PX\" int NOT NULL,\n\t\"PZ\" int NOT NULL,\n\t\"PY\" int NOT NULL,\n\t\"dwTime\" int NOT NULL,\n\t\"strSkill\" varchar(10),\n\t\"strItem\" varchar(400),\n\t\"strSerial\" varchar(400),\n\t\"sQuestCount\" smallint NOT NULL,\n\t\"strQuest\" varchar(400),\n\t\"MannerPoint\" int NOT NULL,\n\t\"LoyaltyMonthly\" int NOT NULL,\n\t\"CreateTime\" smalldatetime NOT NULL,\n\t\"UpdateTime\" smalldatetime\n\tPRIMARY KEY (\"strUserId\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
