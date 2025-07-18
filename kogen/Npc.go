package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_NpcDatabaseNbr = "GAME"
	_NpcTableName   = "K_NPC"
)

func init() {
	ModelList = append(ModelList, &Npc{})
}

// Npc Non-Player Character
type Npc struct {
	NpcId           int16          `gorm:"column:sSid;type:smallint;primaryKey;not null" json:"sSid"`
	Name            *mssql.VarChar `gorm:"column:strName;type:varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"strName,omitempty"`
	PictureId       int16          `gorm:"column:sPid;type:smallint;not null" json:"sPid"`
	Size            int16          `gorm:"column:sSize;type:smallint;not null" json:"sSize"`
	Weapon1         int            `gorm:"column:iWeapon1;type:int;not null" json:"iWeapon1"`
	Weapon2         int            `gorm:"column:iWeapon2;type:int;not null" json:"iWeapon2"`
	Group           uint8          `gorm:"column:byGroup;type:tinyint;not null" json:"byGroup"`
	ActType         uint8          `gorm:"column:byActType;type:tinyint;not null" json:"byActType"`
	Type            uint8          `gorm:"column:byType;type:tinyint;not null" json:"byType"`
	Family          uint8          `gorm:"column:byFamily;type:tinyint;not null" json:"byFamily"`
	Rank            uint8          `gorm:"column:byRank;type:tinyint;not null" json:"byRank"`
	Title           uint8          `gorm:"column:byTitle;type:tinyint;not null" json:"byTitle"`
	SellingGroup    int            `gorm:"column:iSellingGroup;type:int;not null" json:"iSellingGroup"`
	Level           int16          `gorm:"column:sLevel;type:smallint;not null" json:"sLevel"`
	Exp             int            `gorm:"column:iExp;type:int;not null" json:"iExp"`
	Loyalty         int            `gorm:"column:iLoyalty;type:int;not null" json:"iLoyalty"`
	HitPoints       int            `gorm:"column:iHpPoint;type:int;not null" json:"iHpPoint"`
	ManaPoints      int16          `gorm:"column:sMpPoint;type:smallint;not null" json:"sMpPoint"`
	Attack          int16          `gorm:"column:sAtk;type:smallint;not null" json:"sAtk"`
	Armor           int16          `gorm:"column:sAc;type:smallint;not null" json:"sAc"`
	HitRate         int16          `gorm:"column:sHitRate;type:smallint;not null" json:"sHitRate"`
	EvadeRate       int16          `gorm:"column:sEvadeRate;type:smallint;not null" json:"sEvadeRate"`
	Damage          int16          `gorm:"column:sDamage;type:smallint;not null" json:"sDamage"`
	AttackDelay     int16          `gorm:"column:sAttackDelay;type:smallint;not null" json:"sAttackDelay"`
	WalkSpeed       uint8          `gorm:"column:bySpeed1;type:tinyint;not null" json:"bySpeed1"`
	RunSpeed        uint8          `gorm:"column:bySpeed2;type:tinyint;not null" json:"bySpeed2"`
	StandTime       int16          `gorm:"column:sStandtime;type:smallint;not null" json:"sStandtime"`
	Magic1          int            `gorm:"column:iMagic1;type:int;not null" json:"iMagic1"`
	Magic2          int            `gorm:"column:iMagic2;type:int;not null" json:"iMagic2"`
	Magic3          int            `gorm:"column:iMagic3;type:int;not null" json:"iMagic3"`
	FireResist      int16          `gorm:"column:sFireR;type:smallint;not null" json:"sFireR"`
	ColdResist      int16          `gorm:"column:sColdR;type:smallint;not null" json:"sColdR"`
	LightningResist int16          `gorm:"column:sLightningR;type:smallint;not null" json:"sLightningR"`
	MagicResist     int16          `gorm:"column:sMagicR;type:smallint;not null" json:"sMagicR"`
	DiseaseResist   int16          `gorm:"column:sDiseaseR;type:smallint;not null" json:"sDiseaseR"`
	PoisonResist    int16          `gorm:"column:sPoisonR;type:smallint;not null" json:"sPoisonR"`
	LightResist     int16          `gorm:"column:sLightR;type:smallint;not null" json:"sLightR"`
	Bulk            int16          `gorm:"column:sBulk;type:smallint;not null" json:"sBulk"`
	AttackRange     uint8          `gorm:"column:byAttackRange;type:tinyint;not null" json:"byAttackRange"`
	SearchRange     uint8          `gorm:"column:bySearchRange;type:tinyint;not null" json:"bySearchRange"`
	TracingRange    uint8          `gorm:"column:byTracingRange;type:tinyint;not null" json:"byTracingRange"`
	Money           int            `gorm:"column:iMoney;type:int;not null" json:"iMoney"`
	Item            int16          `gorm:"column:sItem;type:smallint;not null" json:"sItem"`
	DirectAttack    uint8          `gorm:"column:byDirectAttack;type:tinyint;not null" json:"byDirectAttack"`
	MagicAttack     uint8          `gorm:"column:byMagicAttack;type:tinyint;not null" json:"byMagicAttack"`
	MoneyType       uint8          `gorm:"column:byMoneyType;type:tinyint;not null" json:"byMoneyType"`
}

// GetDatabaseName Returns the table's database name
func (this Npc) GetDatabaseName() string {
	return GetDatabaseName(_NpcDatabaseNbr)
}

// TableName Returns the table name
func (this Npc) TableName() string {
	return _NpcTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Npc) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [K_NPC] ([sSid], [strName], [sPid], [sSize], [iWeapon1], [iWeapon2], [byGroup], [byActType], [byType], [byFamily], [byRank], [byTitle], [iSellingGroup], [sLevel], [iExp], [iLoyalty], [iHpPoint], [sMpPoint], [sAtk], [sAc], [sHitRate], [sEvadeRate], [sDamage], [sAttackDelay], [bySpeed1], [bySpeed2], [sStandtime], [iMagic1], [iMagic2], [iMagic3], [sFireR], [sColdR], [sLightningR], [sMagicR], [sDiseaseR], [sPoisonR], [sLightR], [sBulk], [byAttackRange], [bySearchRange], [byTracingRange], [iMoney], [sItem], [byDirectAttack], [byMagicAttack], [byMoneyType]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.NpcId),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.PictureId),
		GetOptionalDecVal(&this.Size),
		GetOptionalDecVal(&this.Weapon1),
		GetOptionalDecVal(&this.Weapon2),
		GetOptionalDecVal(&this.Group),
		GetOptionalDecVal(&this.ActType),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Family),
		GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(&this.Title),
		GetOptionalDecVal(&this.SellingGroup),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Exp),
		GetOptionalDecVal(&this.Loyalty),
		GetOptionalDecVal(&this.HitPoints),
		GetOptionalDecVal(&this.ManaPoints),
		GetOptionalDecVal(&this.Attack),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.EvadeRate),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.AttackDelay),
		GetOptionalDecVal(&this.WalkSpeed),
		GetOptionalDecVal(&this.RunSpeed),
		GetOptionalDecVal(&this.StandTime),
		GetOptionalDecVal(&this.Magic1),
		GetOptionalDecVal(&this.Magic2),
		GetOptionalDecVal(&this.Magic3),
		GetOptionalDecVal(&this.FireResist),
		GetOptionalDecVal(&this.ColdResist),
		GetOptionalDecVal(&this.LightningResist),
		GetOptionalDecVal(&this.MagicResist),
		GetOptionalDecVal(&this.DiseaseResist),
		GetOptionalDecVal(&this.PoisonResist),
		GetOptionalDecVal(&this.LightResist),
		GetOptionalDecVal(&this.Bulk),
		GetOptionalDecVal(&this.AttackRange),
		GetOptionalDecVal(&this.SearchRange),
		GetOptionalDecVal(&this.TracingRange),
		GetOptionalDecVal(&this.Money),
		GetOptionalDecVal(&this.Item),
		GetOptionalDecVal(&this.DirectAttack),
		GetOptionalDecVal(&this.MagicAttack),
		GetOptionalDecVal(&this.MoneyType))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Npc) GetInsertHeader() string {
	return "INSERT INTO [K_NPC] ([sSid], [strName], [sPid], [sSize], [iWeapon1], [iWeapon2], [byGroup], [byActType], [byType], [byFamily], [byRank], [byTitle], [iSellingGroup], [sLevel], [iExp], [iLoyalty], [iHpPoint], [sMpPoint], [sAtk], [sAc], [sHitRate], [sEvadeRate], [sDamage], [sAttackDelay], [bySpeed1], [bySpeed2], [sStandtime], [iMagic1], [iMagic2], [iMagic3], [sFireR], [sColdR], [sLightningR], [sMagicR], [sDiseaseR], [sPoisonR], [sLightR], [sBulk], [byAttackRange], [bySearchRange], [byTracingRange], [iMoney], [sItem], [byDirectAttack], [byMagicAttack], [byMoneyType]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Npc) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.NpcId),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.PictureId),
		GetOptionalDecVal(&this.Size),
		GetOptionalDecVal(&this.Weapon1),
		GetOptionalDecVal(&this.Weapon2),
		GetOptionalDecVal(&this.Group),
		GetOptionalDecVal(&this.ActType),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.Family),
		GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(&this.Title),
		GetOptionalDecVal(&this.SellingGroup),
		GetOptionalDecVal(&this.Level),
		GetOptionalDecVal(&this.Exp),
		GetOptionalDecVal(&this.Loyalty),
		GetOptionalDecVal(&this.HitPoints),
		GetOptionalDecVal(&this.ManaPoints),
		GetOptionalDecVal(&this.Attack),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.EvadeRate),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.AttackDelay),
		GetOptionalDecVal(&this.WalkSpeed),
		GetOptionalDecVal(&this.RunSpeed),
		GetOptionalDecVal(&this.StandTime),
		GetOptionalDecVal(&this.Magic1),
		GetOptionalDecVal(&this.Magic2),
		GetOptionalDecVal(&this.Magic3),
		GetOptionalDecVal(&this.FireResist),
		GetOptionalDecVal(&this.ColdResist),
		GetOptionalDecVal(&this.LightningResist),
		GetOptionalDecVal(&this.MagicResist),
		GetOptionalDecVal(&this.DiseaseResist),
		GetOptionalDecVal(&this.PoisonResist),
		GetOptionalDecVal(&this.LightResist),
		GetOptionalDecVal(&this.Bulk),
		GetOptionalDecVal(&this.AttackRange),
		GetOptionalDecVal(&this.SearchRange),
		GetOptionalDecVal(&this.TracingRange),
		GetOptionalDecVal(&this.Money),
		GetOptionalDecVal(&this.Item),
		GetOptionalDecVal(&this.DirectAttack),
		GetOptionalDecVal(&this.MagicAttack),
		GetOptionalDecVal(&this.MoneyType))
}

// GetCreateTableString Returns the create table statement for this object
func (this Npc) GetCreateTableString() string {
	query := "CREATE TABLE [K_NPC] (\n\t[sSid] smallint NOT NULL,\n\t[strName] varchar(30) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[sPid] smallint NOT NULL,\n\t[sSize] smallint NOT NULL,\n\t[iWeapon1] int NOT NULL,\n\t[iWeapon2] int NOT NULL,\n\t[byGroup] tinyint NOT NULL,\n\t[byActType] tinyint NOT NULL,\n\t[byType] tinyint NOT NULL,\n\t[byFamily] tinyint NOT NULL,\n\t[byRank] tinyint NOT NULL,\n\t[byTitle] tinyint NOT NULL,\n\t[iSellingGroup] int NOT NULL,\n\t[sLevel] smallint NOT NULL,\n\t[iExp] int NOT NULL,\n\t[iLoyalty] int NOT NULL,\n\t[iHpPoint] int NOT NULL,\n\t[sMpPoint] smallint NOT NULL,\n\t[sAtk] smallint NOT NULL,\n\t[sAc] smallint NOT NULL,\n\t[sHitRate] smallint NOT NULL,\n\t[sEvadeRate] smallint NOT NULL,\n\t[sDamage] smallint NOT NULL,\n\t[sAttackDelay] smallint NOT NULL,\n\t[bySpeed1] tinyint NOT NULL,\n\t[bySpeed2] tinyint NOT NULL,\n\t[sStandtime] smallint NOT NULL,\n\t[iMagic1] int NOT NULL,\n\t[iMagic2] int NOT NULL,\n\t[iMagic3] int NOT NULL,\n\t[sFireR] smallint NOT NULL,\n\t[sColdR] smallint NOT NULL,\n\t[sLightningR] smallint NOT NULL,\n\t[sMagicR] smallint NOT NULL,\n\t[sDiseaseR] smallint NOT NULL,\n\t[sPoisonR] smallint NOT NULL,\n\t[sLightR] smallint NOT NULL,\n\t[sBulk] smallint NOT NULL,\n\t[byAttackRange] tinyint NOT NULL,\n\t[bySearchRange] tinyint NOT NULL,\n\t[byTracingRange] tinyint NOT NULL,\n\t[iMoney] int NOT NULL,\n\t[sItem] smallint NOT NULL,\n\t[byDirectAttack] tinyint NOT NULL,\n\t[byMagicAttack] tinyint NOT NULL,\n\t[byMoneyType] tinyint NOT NULL\n\tCONSTRAINT [PK_K_NPC] PRIMARY KEY CLUSTERED ([sSid])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Npc) SelectClause() (selectClause clause.Select) {
	return _NpcSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Npc) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Npc{}
	rawSql := "SELECT [sSid], [strName], [sPid], [sSize], [iWeapon1], [iWeapon2], [byGroup], [byActType], [byType], [byFamily], [byRank], [byTitle], [iSellingGroup], [sLevel], [iExp], [iLoyalty], [iHpPoint], [sMpPoint], [sAtk], [sAc], [sHitRate], [sEvadeRate], [sDamage], [sAttackDelay], [bySpeed1], [bySpeed2], [sStandtime], [iMagic1], [iMagic2], [iMagic3], [sFireR], [sColdR], [sLightningR], [sMagicR], [sDiseaseR], [sPoisonR], [sLightR], [sBulk], [byAttackRange], [bySearchRange], [byTracingRange], [iMoney], [sItem], [byDirectAttack], [byMagicAttack], [byMoneyType] FROM [K_NPC]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _NpcSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[sSid]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[sPid]",
		},
		clause.Column{
			Name: "[sSize]",
		},
		clause.Column{
			Name: "[iWeapon1]",
		},
		clause.Column{
			Name: "[iWeapon2]",
		},
		clause.Column{
			Name: "[byGroup]",
		},
		clause.Column{
			Name: "[byActType]",
		},
		clause.Column{
			Name: "[byType]",
		},
		clause.Column{
			Name: "[byFamily]",
		},
		clause.Column{
			Name: "[byRank]",
		},
		clause.Column{
			Name: "[byTitle]",
		},
		clause.Column{
			Name: "[iSellingGroup]",
		},
		clause.Column{
			Name: "[sLevel]",
		},
		clause.Column{
			Name: "[iExp]",
		},
		clause.Column{
			Name: "[iLoyalty]",
		},
		clause.Column{
			Name: "[iHpPoint]",
		},
		clause.Column{
			Name: "[sMpPoint]",
		},
		clause.Column{
			Name: "[sAtk]",
		},
		clause.Column{
			Name: "[sAc]",
		},
		clause.Column{
			Name: "[sHitRate]",
		},
		clause.Column{
			Name: "[sEvadeRate]",
		},
		clause.Column{
			Name: "[sDamage]",
		},
		clause.Column{
			Name: "[sAttackDelay]",
		},
		clause.Column{
			Name: "[bySpeed1]",
		},
		clause.Column{
			Name: "[bySpeed2]",
		},
		clause.Column{
			Name: "[sStandtime]",
		},
		clause.Column{
			Name: "[iMagic1]",
		},
		clause.Column{
			Name: "[iMagic2]",
		},
		clause.Column{
			Name: "[iMagic3]",
		},
		clause.Column{
			Name: "[sFireR]",
		},
		clause.Column{
			Name: "[sColdR]",
		},
		clause.Column{
			Name: "[sLightningR]",
		},
		clause.Column{
			Name: "[sMagicR]",
		},
		clause.Column{
			Name: "[sDiseaseR]",
		},
		clause.Column{
			Name: "[sPoisonR]",
		},
		clause.Column{
			Name: "[sLightR]",
		},
		clause.Column{
			Name: "[sBulk]",
		},
		clause.Column{
			Name: "[byAttackRange]",
		},
		clause.Column{
			Name: "[bySearchRange]",
		},
		clause.Column{
			Name: "[byTracingRange]",
		},
		clause.Column{
			Name: "[iMoney]",
		},
		clause.Column{
			Name: "[sItem]",
		},
		clause.Column{
			Name: "[byDirectAttack]",
		},
		clause.Column{
			Name: "[byMagicAttack]",
		},
		clause.Column{
			Name: "[byMoneyType]",
		},
	},
}
