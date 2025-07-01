package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KnightsSiegeWarfareDatabaseNbr = "GAME"
	_KnightsSiegeWarfareTableName   = "KNIGHTS_SIEGE_WARFARE"
)

func init() {
	ModelList = append(ModelList, &KnightsSiegeWarfare{})
}

// KnightsSiegeWarfare Knights Siege Warfare
type KnightsSiegeWarfare struct {
	CastleIndex        int16   `gorm:"column:sCastleIndex;type:smallint;not null" json:"sCastleIndex"`
	MasterKnights      int16   `gorm:"column:sMasterKnights;type:smallint;not null" json:"sMasterKnights"`
	SiegeType          uint8   `gorm:"column:bySiegeType;type:tinyint;not null" json:"bySiegeType"`
	WarDay             uint8   `gorm:"column:byWarDay;type:tinyint;not null" json:"byWarDay"`
	WarHour            uint8   `gorm:"column:byWarTime;type:tinyint;not null" json:"byWarTime"`
	WarMinute          uint8   `gorm:"column:byWarMinute;type:tinyint;not null" json:"byWarMinute"`
	ChallengeList1     int16   `gorm:"column:sChallengeList_1;type:smallint;not null" json:"sChallengeList_1"`
	ChallengeList2     int16   `gorm:"column:sChallengeList_2;type:smallint;not null" json:"sChallengeList_2"`
	ChallengeList3     int16   `gorm:"column:sChallengeList_3;type:smallint;not null" json:"sChallengeList_3"`
	ChallengeList4     int16   `gorm:"column:sChallengeList_4;type:smallint;not null" json:"sChallengeList_4"`
	ChallengeList5     int16   `gorm:"column:sChallengeList_5;type:smallint;not null" json:"sChallengeList_5"`
	ChallengeList6     int16   `gorm:"column:sChallengeList_6;type:smallint;not null" json:"sChallengeList_6"`
	ChallengeList7     int16   `gorm:"column:sChallengeList_7;type:smallint;not null;default:0" json:"sChallengeList_7"`
	ChallengeList8     int16   `gorm:"column:sChallengeList_8;type:smallint;not null;default:0" json:"sChallengeList_8"`
	ChallengeList9     int16   `gorm:"column:sChallengeList_9;type:smallint;not null;default:0" json:"sChallengeList_9"`
	ChallengeList10    int16   `gorm:"column:sChallengeList_10;type:smallint;not null;default:0" json:"sChallengeList_10"`
	WarRequestDay      uint8   `gorm:"column:byWarRequestDay;type:tinyint;not null;default:3" json:"byWarRequestDay"`
	WarRequestTime     uint8   `gorm:"column:byWarRequestTime;type:tinyint;not null;default:9" json:"byWarRequestTime"`
	WarRequestMinute   uint8   `gorm:"column:byWarRequestMinute;type:tinyint;not null;default:0" json:"byWarRequestMinute"`
	GuerrillaWarDay    uint8   `gorm:"column:byGuerrillaWarDay;type:tinyint;not null;default:1" json:"byGuerrillaWarDay"`
	GuerrillaWarTime   uint8   `gorm:"column:byGuerrillaWarTime;type:tinyint;not null;default:20" json:"byGuerrillaWarTime"`
	GuerrillaWarMinute uint8   `gorm:"column:byGuerrillaWarMinute;type:tinyint;not null;default:0" json:"byGuerrillaWarMinute"`
	ChallengeList      *[]byte `gorm:"column:strChallengeList;type:char(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strChallengeList,omitempty"`
	MoradonTariff      int16   `gorm:"column:sMoradonTariff;type:smallint;not null;default:0" json:"sMoradonTariff"`
	DelosTariff        int16   `gorm:"column:sDellosTariff;type:smallint;not null;default:0" json:"sDellosTariff"`
	DungeonCharge      int     `gorm:"column:nDungeonCharge;type:int;not null;default:0" json:"nDungeonCharge"`
	MoradonTax         int     `gorm:"column:nMoradonTax;type:int;not null;default:0" json:"nMoradonTax"`
	DelosTax           int     `gorm:"column:nDellosTax;type:int;not null;default:0" json:"nDellosTax"`
	RequestList1       int16   `gorm:"column:sRequestList_1;type:smallint;not null;default:0" json:"sRequestList_1"`
	RequestList2       int16   `gorm:"column:sRequestList_2;type:smallint;not null;default:0" json:"sRequestList_2"`
	RequestList3       int16   `gorm:"column:sRequestList_3;type:smallint;not null;default:0" json:"sRequestList_3"`
	RequestList4       int16   `gorm:"column:sRequestList_4;type:smallint;not null;default:0" json:"sRequestList_4"`
	RequestList5       int16   `gorm:"column:sRequestList_5;type:smallint;not null;default:0" json:"sRequestList_5"`
	RequestList6       int16   `gorm:"column:sRequestList_6;type:smallint;not null;default:0" json:"sRequestList_6"`
	RequestList7       int16   `gorm:"column:sRequestList_7;type:smallint;not null;default:0" json:"sRequestList_7"`
	RequestList8       int16   `gorm:"column:sRequestList_8;type:smallint;not null;default:0" json:"sRequestList_8"`
	RequestList9       int16   `gorm:"column:sRequestList_9;type:smallint;not null;default:0" json:"sRequestList_9"`
	RequestList10      int16   `gorm:"column:sRequestList_10;type:smallint;not null;default:0" json:"sRequestList_10"`
}

// GetDatabaseName Returns the table's database name
func (this KnightsSiegeWarfare) GetDatabaseName() string {
	return GetDatabaseName(_KnightsSiegeWarfareDatabaseNbr)
}

// TableName Returns the table name
func (this KnightsSiegeWarfare) TableName() string {
	return _KnightsSiegeWarfareTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KnightsSiegeWarfare) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_SIEGE_WARFARE] ([sCastleIndex], [sMasterKnights], [bySiegeType], [byWarDay], [byWarTime], [byWarMinute], [sChallengeList_1], [sChallengeList_2], [sChallengeList_3], [sChallengeList_4], [sChallengeList_5], [sChallengeList_6], [sChallengeList_7], [sChallengeList_8], [sChallengeList_9], [sChallengeList_10], [byWarRequestDay], [byWarRequestTime], [byWarRequestMinute], [byGuerrillaWarDay], [byGuerrillaWarTime], [byGuerrillaWarMinute], [strChallengeList], [sMoradonTariff], [sDellosTariff], [nDungeonCharge], [nMoradonTax], [nDellosTax], [sRequestList_1], [sRequestList_2], [sRequestList_3], [sRequestList_4], [sRequestList_5], [sRequestList_6], [sRequestList_7], [sRequestList_8], [sRequestList_9], [sRequestList_10]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(char(50), %s), %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CastleIndex),
		GetOptionalDecVal(&this.MasterKnights),
		GetOptionalDecVal(&this.SiegeType),
		GetOptionalDecVal(&this.WarDay),
		GetOptionalDecVal(&this.WarHour),
		GetOptionalDecVal(&this.WarMinute),
		GetOptionalDecVal(&this.ChallengeList1),
		GetOptionalDecVal(&this.ChallengeList2),
		GetOptionalDecVal(&this.ChallengeList3),
		GetOptionalDecVal(&this.ChallengeList4),
		GetOptionalDecVal(&this.ChallengeList5),
		GetOptionalDecVal(&this.ChallengeList6),
		GetOptionalDecVal(&this.ChallengeList7),
		GetOptionalDecVal(&this.ChallengeList8),
		GetOptionalDecVal(&this.ChallengeList9),
		GetOptionalDecVal(&this.ChallengeList10),
		GetOptionalDecVal(&this.WarRequestDay),
		GetOptionalDecVal(&this.WarRequestTime),
		GetOptionalDecVal(&this.WarRequestMinute),
		GetOptionalDecVal(&this.GuerrillaWarDay),
		GetOptionalDecVal(&this.GuerrillaWarTime),
		GetOptionalDecVal(&this.GuerrillaWarMinute),
		GetOptionalByteArrayVal(this.ChallengeList, true),
		GetOptionalDecVal(&this.MoradonTariff),
		GetOptionalDecVal(&this.DelosTariff),
		GetOptionalDecVal(&this.DungeonCharge),
		GetOptionalDecVal(&this.MoradonTax),
		GetOptionalDecVal(&this.DelosTax),
		GetOptionalDecVal(&this.RequestList1),
		GetOptionalDecVal(&this.RequestList2),
		GetOptionalDecVal(&this.RequestList3),
		GetOptionalDecVal(&this.RequestList4),
		GetOptionalDecVal(&this.RequestList5),
		GetOptionalDecVal(&this.RequestList6),
		GetOptionalDecVal(&this.RequestList7),
		GetOptionalDecVal(&this.RequestList8),
		GetOptionalDecVal(&this.RequestList9),
		GetOptionalDecVal(&this.RequestList10))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KnightsSiegeWarfare) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS_SIEGE_WARFARE] ([sCastleIndex], [sMasterKnights], [bySiegeType], [byWarDay], [byWarTime], [byWarMinute], [sChallengeList_1], [sChallengeList_2], [sChallengeList_3], [sChallengeList_4], [sChallengeList_5], [sChallengeList_6], [sChallengeList_7], [sChallengeList_8], [sChallengeList_9], [sChallengeList_10], [byWarRequestDay], [byWarRequestTime], [byWarRequestMinute], [byGuerrillaWarDay], [byGuerrillaWarTime], [byGuerrillaWarMinute], [strChallengeList], [sMoradonTariff], [sDellosTariff], [nDungeonCharge], [nMoradonTax], [nDellosTax], [sRequestList_1], [sRequestList_2], [sRequestList_3], [sRequestList_4], [sRequestList_5], [sRequestList_6], [sRequestList_7], [sRequestList_8], [sRequestList_9], [sRequestList_10]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KnightsSiegeWarfare) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, CONVERT(char(50), %s), %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.CastleIndex),
		GetOptionalDecVal(&this.MasterKnights),
		GetOptionalDecVal(&this.SiegeType),
		GetOptionalDecVal(&this.WarDay),
		GetOptionalDecVal(&this.WarHour),
		GetOptionalDecVal(&this.WarMinute),
		GetOptionalDecVal(&this.ChallengeList1),
		GetOptionalDecVal(&this.ChallengeList2),
		GetOptionalDecVal(&this.ChallengeList3),
		GetOptionalDecVal(&this.ChallengeList4),
		GetOptionalDecVal(&this.ChallengeList5),
		GetOptionalDecVal(&this.ChallengeList6),
		GetOptionalDecVal(&this.ChallengeList7),
		GetOptionalDecVal(&this.ChallengeList8),
		GetOptionalDecVal(&this.ChallengeList9),
		GetOptionalDecVal(&this.ChallengeList10),
		GetOptionalDecVal(&this.WarRequestDay),
		GetOptionalDecVal(&this.WarRequestTime),
		GetOptionalDecVal(&this.WarRequestMinute),
		GetOptionalDecVal(&this.GuerrillaWarDay),
		GetOptionalDecVal(&this.GuerrillaWarTime),
		GetOptionalDecVal(&this.GuerrillaWarMinute),
		GetOptionalByteArrayVal(this.ChallengeList, true),
		GetOptionalDecVal(&this.MoradonTariff),
		GetOptionalDecVal(&this.DelosTariff),
		GetOptionalDecVal(&this.DungeonCharge),
		GetOptionalDecVal(&this.MoradonTax),
		GetOptionalDecVal(&this.DelosTax),
		GetOptionalDecVal(&this.RequestList1),
		GetOptionalDecVal(&this.RequestList2),
		GetOptionalDecVal(&this.RequestList3),
		GetOptionalDecVal(&this.RequestList4),
		GetOptionalDecVal(&this.RequestList5),
		GetOptionalDecVal(&this.RequestList6),
		GetOptionalDecVal(&this.RequestList7),
		GetOptionalDecVal(&this.RequestList8),
		GetOptionalDecVal(&this.RequestList9),
		GetOptionalDecVal(&this.RequestList10))
}

// GetCreateTableString Returns the create table statement for this object
func (this KnightsSiegeWarfare) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_SIEGE_WARFARE] (\n\t[sCastleIndex] smallint NOT NULL,\n\t[sMasterKnights] smallint NOT NULL,\n\t[bySiegeType] tinyint NOT NULL,\n\t[byWarDay] tinyint NOT NULL,\n\t[byWarTime] tinyint NOT NULL,\n\t[byWarMinute] tinyint NOT NULL,\n\t[sChallengeList_1] smallint NOT NULL,\n\t[sChallengeList_2] smallint NOT NULL,\n\t[sChallengeList_3] smallint NOT NULL,\n\t[sChallengeList_4] smallint NOT NULL,\n\t[sChallengeList_5] smallint NOT NULL,\n\t[sChallengeList_6] smallint NOT NULL,\n\t[sChallengeList_7] smallint NOT NULL,\n\t[sChallengeList_8] smallint NOT NULL,\n\t[sChallengeList_9] smallint NOT NULL,\n\t[sChallengeList_10] smallint NOT NULL,\n\t[byWarRequestDay] tinyint NOT NULL,\n\t[byWarRequestTime] tinyint NOT NULL,\n\t[byWarRequestMinute] tinyint NOT NULL,\n\t[byGuerrillaWarDay] tinyint NOT NULL,\n\t[byGuerrillaWarTime] tinyint NOT NULL,\n\t[byGuerrillaWarMinute] tinyint NOT NULL,\n\t[strChallengeList] char(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[sMoradonTariff] smallint NOT NULL,\n\t[sDellosTariff] smallint NOT NULL,\n\t[nDungeonCharge] int NOT NULL,\n\t[nMoradonTax] int NOT NULL,\n\t[nDellosTax] int NOT NULL,\n\t[sRequestList_1] smallint NOT NULL,\n\t[sRequestList_2] smallint NOT NULL,\n\t[sRequestList_3] smallint NOT NULL,\n\t[sRequestList_4] smallint NOT NULL,\n\t[sRequestList_5] smallint NOT NULL,\n\t[sRequestList_6] smallint NOT NULL,\n\t[sRequestList_7] smallint NOT NULL,\n\t[sRequestList_8] smallint NOT NULL,\n\t[sRequestList_9] smallint NOT NULL,\n\t[sRequestList_10] smallint NOT NULL\n)\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sChallengeList_7] DEFAULT 0 FOR [sChallengeList_7]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sChallengeList_8] DEFAULT 0 FOR [sChallengeList_8]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sChallengeList_9] DEFAULT 0 FOR [sChallengeList_9]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sChallengeList_10] DEFAULT 0 FOR [sChallengeList_10]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byWarRequestDay] DEFAULT 3 FOR [byWarRequestDay]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byWarRequestTime] DEFAULT 9 FOR [byWarRequestTime]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byWarRequestMinute] DEFAULT 0 FOR [byWarRequestMinute]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byGuerrillaWarDay] DEFAULT 1 FOR [byGuerrillaWarDay]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byGuerrillaWarTime] DEFAULT 20 FOR [byGuerrillaWarTime]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_byGuerrillaWarMinute] DEFAULT 0 FOR [byGuerrillaWarMinute]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sMoradonTariff] DEFAULT 0 FOR [sMoradonTariff]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sDellosTariff] DEFAULT 0 FOR [sDellosTariff]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_nDungeonCharge] DEFAULT 0 FOR [nDungeonCharge]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_nMoradonTax] DEFAULT 0 FOR [nMoradonTax]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_nDellosTax] DEFAULT 0 FOR [nDellosTax]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_1] DEFAULT 0 FOR [sRequestList_1]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_2] DEFAULT 0 FOR [sRequestList_2]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_3] DEFAULT 0 FOR [sRequestList_3]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_4] DEFAULT 0 FOR [sRequestList_4]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_5] DEFAULT 0 FOR [sRequestList_5]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_6] DEFAULT 0 FOR [sRequestList_6]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_7] DEFAULT 0 FOR [sRequestList_7]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_8] DEFAULT 0 FOR [sRequestList_8]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_9] DEFAULT 0 FOR [sRequestList_9]\nGO\nALTER TABLE [KNIGHTS_SIEGE_WARFARE] ADD CONSTRAINT [DF_KNIGHTS_SIEGE_WARFARE_sRequestList_10] DEFAULT 0 FOR [sRequestList_10]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KnightsSiegeWarfare) SelectClause() (selectClause clause.Select) {
	return _KnightsSiegeWarfareSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KnightsSiegeWarfare) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KnightsSiegeWarfare{}
	rawSql := "SELECT [sCastleIndex], [sMasterKnights], [bySiegeType], [byWarDay], [byWarTime], [byWarMinute], [sChallengeList_1], [sChallengeList_2], [sChallengeList_3], [sChallengeList_4], [sChallengeList_5], [sChallengeList_6], [sChallengeList_7], [sChallengeList_8], [sChallengeList_9], [sChallengeList_10], [byWarRequestDay], [byWarRequestTime], [byWarRequestMinute], [byGuerrillaWarDay], [byGuerrillaWarTime], [byGuerrillaWarMinute], CONVERT(VARBINARY(50), [strChallengeList]) as [strChallengeList], [sMoradonTariff], [sDellosTariff], [nDungeonCharge], [nMoradonTax], [nDellosTax], [sRequestList_1], [sRequestList_2], [sRequestList_3], [sRequestList_4], [sRequestList_5], [sRequestList_6], [sRequestList_7], [sRequestList_8], [sRequestList_9], [sRequestList_10] FROM [KNIGHTS_SIEGE_WARFARE]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsSiegeWarfareSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sCastleIndex]",
		},
		clause.Column{
			Name: "[sMasterKnights]",
		},
		clause.Column{
			Name: "[bySiegeType]",
		},
		clause.Column{
			Name: "[byWarDay]",
		},
		clause.Column{
			Name: "[byWarTime]",
		},
		clause.Column{
			Name: "[byWarMinute]",
		},
		clause.Column{
			Name: "[sChallengeList_1]",
		},
		clause.Column{
			Name: "[sChallengeList_2]",
		},
		clause.Column{
			Name: "[sChallengeList_3]",
		},
		clause.Column{
			Name: "[sChallengeList_4]",
		},
		clause.Column{
			Name: "[sChallengeList_5]",
		},
		clause.Column{
			Name: "[sChallengeList_6]",
		},
		clause.Column{
			Name: "[sChallengeList_7]",
		},
		clause.Column{
			Name: "[sChallengeList_8]",
		},
		clause.Column{
			Name: "[sChallengeList_9]",
		},
		clause.Column{
			Name: "[sChallengeList_10]",
		},
		clause.Column{
			Name: "[byWarRequestDay]",
		},
		clause.Column{
			Name: "[byWarRequestTime]",
		},
		clause.Column{
			Name: "[byWarRequestMinute]",
		},
		clause.Column{
			Name: "[byGuerrillaWarDay]",
		},
		clause.Column{
			Name: "[byGuerrillaWarTime]",
		},
		clause.Column{
			Name: "[byGuerrillaWarMinute]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(50), [strChallengeList]) as [strChallengeList]",
		},
		clause.Column{
			Name: "[sMoradonTariff]",
		},
		clause.Column{
			Name: "[sDellosTariff]",
		},
		clause.Column{
			Name: "[nDungeonCharge]",
		},
		clause.Column{
			Name: "[nMoradonTax]",
		},
		clause.Column{
			Name: "[nDellosTax]",
		},
		clause.Column{
			Name: "[sRequestList_1]",
		},
		clause.Column{
			Name: "[sRequestList_2]",
		},
		clause.Column{
			Name: "[sRequestList_3]",
		},
		clause.Column{
			Name: "[sRequestList_4]",
		},
		clause.Column{
			Name: "[sRequestList_5]",
		},
		clause.Column{
			Name: "[sRequestList_6]",
		},
		clause.Column{
			Name: "[sRequestList_7]",
		},
		clause.Column{
			Name: "[sRequestList_8]",
		},
		clause.Column{
			Name: "[sRequestList_9]",
		},
		clause.Column{
			Name: "[sRequestList_10]",
		},
	},
}
