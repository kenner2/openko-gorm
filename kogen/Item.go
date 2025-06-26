package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ItemDatabaseNbr = 1
	_ItemTableName   = "ITEM"
)

func init() {
	ModelList = append(ModelList, &Item{})
}

// Item Item information
type Item struct {
	Number              int           `gorm:"column:Num;type:int;primaryKey;not null" json:"Num"`
	Name                mssql.VarChar `gorm:"column:strName;type:varchar(50);not null" json:"strName"`
	Kind                uint8         `gorm:"column:Kind;type:tinyint;not null;default:0" json:"Kind"`
	Slot                uint8         `gorm:"column:Slot;type:tinyint;not null;default:0" json:"Slot"`
	Race                uint8         `gorm:"column:Race;type:tinyint;not null;default:0" json:"Race"`
	ClassId             uint8         `gorm:"column:Class;type:tinyint;not null;default:1" json:"Class"`
	Damage              int16         `gorm:"column:Damage;type:smallint;not null;default:0" json:"Damage"`
	Delay               int16         `gorm:"column:Delay;type:smallint;not null;default:0" json:"Delay"`
	Range               int16         `gorm:"column:Range;type:smallint;not null;default:0" json:"Range"`
	Weight              int16         `gorm:"column:Weight;type:smallint;not null;default:0" json:"Weight"`
	Durability          int16         `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	BuyPrice            int           `gorm:"column:BuyPrice;type:int;not null;default:0" json:"BuyPrice"`
	SellPrice           int           `gorm:"column:SellPrice;type:int;not null" json:"SellPrice"`
	Armor               int16         `gorm:"column:Ac;type:smallint;not null;default:0" json:"Ac"`
	Countable           uint8         `gorm:"column:Countable;type:tinyint;not null" json:"Countable"`
	MagicEffect         int           `gorm:"column:Effect1;type:int;not null" json:"Effect1"`
	SpecialEffect       int           `gorm:"column:Effect2;type:int;not null" json:"Effect2"`
	RequireLevel        uint8         `gorm:"column:ReqLevel;type:tinyint;not null" json:"ReqLevel"`
	MaxLevel            uint8         `gorm:"column:ReqLevelMax;type:tinyint;not null" json:"ReqLevelMax"`
	RequireRank         uint8         `gorm:"column:ReqRank;type:tinyint;not null" json:"ReqRank"`
	RequireTitle        uint8         `gorm:"column:ReqTitle;type:tinyint;not null" json:"ReqTitle"`
	RequireStrength     uint8         `gorm:"column:ReqStr;type:tinyint;not null" json:"ReqStr"`
	RequireStamina      uint8         `gorm:"column:ReqSta;type:tinyint;not null" json:"ReqSta"`
	RequireDexterity    uint8         `gorm:"column:ReqDex;type:tinyint;not null" json:"ReqDex"`
	RequireIntelligence uint8         `gorm:"column:ReqIntel;type:tinyint;not null" json:"ReqIntel"`
	RequireCharisma     uint8         `gorm:"column:ReqCha;type:tinyint;not null" json:"ReqCha"`
	SellingGroup        uint8         `gorm:"column:SellingGroup;type:tinyint;not null" json:"SellingGroup"`
	Type                uint8         `gorm:"column:ItemType;type:tinyint;not null" json:"ItemType"`
	HitRate             int16         `gorm:"column:Hitrate;type:smallint;not null" json:"Hitrate"`
	EvasionRate         int16         `gorm:"column:Evasionrate;type:smallint;not null" json:"Evasionrate"`
	DaggerArmor         int16         `gorm:"column:DaggerAc;type:smallint;not null;default:0" json:"DaggerAc"`
	SwordArmor          int16         `gorm:"column:SwordAc;type:smallint;not null;default:0" json:"SwordAc"`
	MaceArmor           int16         `gorm:"column:MaceAc;type:smallint;not null;default:0" json:"MaceAc"`
	AxeArmor            int16         `gorm:"column:AxeAc;type:smallint;not null;default:0" json:"AxeAc"`
	SpearArmor          int16         `gorm:"column:SpearAc;type:smallint;not null;default:0" json:"SpearAc"`
	BowArmor            int16         `gorm:"column:BowAc;type:smallint;not null;default:0" json:"BowAc"`
	FireDamage          uint8         `gorm:"column:FireDamage;type:tinyint;not null" json:"FireDamage"`
	IceDamage           uint8         `gorm:"column:IceDamage;type:tinyint;not null" json:"IceDamage"`
	LightningDamage     uint8         `gorm:"column:LightningDamage;type:tinyint;not null" json:"LightningDamage"`
	PoisonDamage        uint8         `gorm:"column:PoisonDamage;type:tinyint;not null" json:"PoisonDamage"`
	HpDrain             uint8         `gorm:"column:HPDrain;type:tinyint;not null" json:"HPDrain"`
	MpDamage            uint8         `gorm:"column:MPDamage;type:tinyint;not null" json:"MPDamage"`
	MpDrain             uint8         `gorm:"column:MPDrain;type:tinyint;not null" json:"MPDrain"`
	MirrorDamage        uint8         `gorm:"column:MirrorDamage;type:tinyint;not null;default:0" json:"MirrorDamage"`
	DropRate            uint8         `gorm:"column:Droprate;type:tinyint;not null;default:0" json:"Droprate"`
	StrengthBonus       int16         `gorm:"column:StrB;type:smallint;not null;default:0" json:"StrB"`
	StaminaBonus        int16         `gorm:"column:StaB;type:smallint;not null;default:0" json:"StaB"`
	DexterityBonus      int16         `gorm:"column:DexB;type:smallint;not null;default:0" json:"DexB"`
	IntelligenceBonus   int16         `gorm:"column:IntelB;type:smallint;not null;default:0" json:"IntelB"`
	CharismaBonus       int16         `gorm:"column:ChaB;type:smallint;not null;default:0" json:"ChaB"`
	MaxHpBonus          int16         `gorm:"column:MaxHpB;type:smallint;not null" json:"MaxHpB"`
	MaxMpBonus          int16         `gorm:"column:MaxMpB;type:smallint;not null" json:"MaxMpB"`
	FireResistance      int16         `gorm:"column:FireR;type:smallint;not null;default:0" json:"FireR"`
	ColdResistance      int16         `gorm:"column:ColdR;type:smallint;not null;default:0" json:"ColdR"`
	LightningResistance int16         `gorm:"column:LightningR;type:smallint;not null;default:0" json:"LightningR"`
	MagicResistance     int16         `gorm:"column:MagicR;type:smallint;not null;default:0" json:"MagicR"`
	PoisonResistance    int16         `gorm:"column:PoisonR;type:smallint;not null;default:0" json:"PoisonR"`
	CurseResistance     int16         `gorm:"column:CurseR;type:smallint;not null;default:0" json:"CurseR"`
}

// GetDatabaseName Returns the table's database name
func (this Item) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemDatabaseNbr))
}

// TableName Returns the table name
func (this Item) TableName() string {
	return _ItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this Item) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM] ([Num], [strName], [Kind], [Slot], [Race], [Class], [Damage], [Delay], [Range], [Weight], [Duration], [BuyPrice], [SellPrice], [Ac], [Countable], [Effect1], [Effect2], [ReqLevel], [ReqLevelMax], [ReqRank], [ReqTitle], [ReqStr], [ReqSta], [ReqDex], [ReqIntel], [ReqCha], [SellingGroup], [ItemType], [Hitrate], [Evasionrate], [DaggerAc], [SwordAc], [MaceAc], [AxeAc], [SpearAc], [BowAc], [FireDamage], [IceDamage], [LightningDamage], [PoisonDamage], [HPDrain], [MPDamage], [MPDrain], [MirrorDamage], [Droprate], [StrB], [StaB], [DexB], [IntelB], [ChaB], [MaxHpB], [MaxMpB], [FireR], [ColdR], [LightningR], [MagicR], [PoisonR], [CurseR]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Number),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Kind),
		GetOptionalDecVal(&this.Slot),
		GetOptionalDecVal(&this.Race),
		GetOptionalDecVal(&this.ClassId),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.Delay),
		GetOptionalDecVal(&this.Range),
		GetOptionalDecVal(&this.Weight),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.BuyPrice),
		GetOptionalDecVal(&this.SellPrice),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.Countable),
		GetOptionalDecVal(&this.MagicEffect),
		GetOptionalDecVal(&this.SpecialEffect),
		GetOptionalDecVal(&this.RequireLevel),
		GetOptionalDecVal(&this.MaxLevel),
		GetOptionalDecVal(&this.RequireRank),
		GetOptionalDecVal(&this.RequireTitle),
		GetOptionalDecVal(&this.RequireStrength),
		GetOptionalDecVal(&this.RequireStamina),
		GetOptionalDecVal(&this.RequireDexterity),
		GetOptionalDecVal(&this.RequireIntelligence),
		GetOptionalDecVal(&this.RequireCharisma),
		GetOptionalDecVal(&this.SellingGroup),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.EvasionRate),
		GetOptionalDecVal(&this.DaggerArmor),
		GetOptionalDecVal(&this.SwordArmor),
		GetOptionalDecVal(&this.MaceArmor),
		GetOptionalDecVal(&this.AxeArmor),
		GetOptionalDecVal(&this.SpearArmor),
		GetOptionalDecVal(&this.BowArmor),
		GetOptionalDecVal(&this.FireDamage),
		GetOptionalDecVal(&this.IceDamage),
		GetOptionalDecVal(&this.LightningDamage),
		GetOptionalDecVal(&this.PoisonDamage),
		GetOptionalDecVal(&this.HpDrain),
		GetOptionalDecVal(&this.MpDamage),
		GetOptionalDecVal(&this.MpDrain),
		GetOptionalDecVal(&this.MirrorDamage),
		GetOptionalDecVal(&this.DropRate),
		GetOptionalDecVal(&this.StrengthBonus),
		GetOptionalDecVal(&this.StaminaBonus),
		GetOptionalDecVal(&this.DexterityBonus),
		GetOptionalDecVal(&this.IntelligenceBonus),
		GetOptionalDecVal(&this.CharismaBonus),
		GetOptionalDecVal(&this.MaxHpBonus),
		GetOptionalDecVal(&this.MaxMpBonus),
		GetOptionalDecVal(&this.FireResistance),
		GetOptionalDecVal(&this.ColdResistance),
		GetOptionalDecVal(&this.LightningResistance),
		GetOptionalDecVal(&this.MagicResistance),
		GetOptionalDecVal(&this.PoisonResistance),
		GetOptionalDecVal(&this.CurseResistance))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this Item) GetInsertHeader() string {
	return "INSERT INTO [ITEM] (Num, strName, Kind, Slot, Race, Class, Damage, Delay, Range, Weight, Duration, BuyPrice, SellPrice, Ac, Countable, Effect1, Effect2, ReqLevel, ReqLevelMax, ReqRank, ReqTitle, ReqStr, ReqSta, ReqDex, ReqIntel, ReqCha, SellingGroup, ItemType, Hitrate, Evasionrate, DaggerAc, SwordAc, MaceAc, AxeAc, SpearAc, BowAc, FireDamage, IceDamage, LightningDamage, PoisonDamage, HPDrain, MPDamage, MPDrain, MirrorDamage, Droprate, StrB, StaB, DexB, IntelB, ChaB, MaxHpB, MaxMpB, FireR, ColdR, LightningR, MagicR, PoisonR, CurseR) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this Item) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Number),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.Kind),
		GetOptionalDecVal(&this.Slot),
		GetOptionalDecVal(&this.Race),
		GetOptionalDecVal(&this.ClassId),
		GetOptionalDecVal(&this.Damage),
		GetOptionalDecVal(&this.Delay),
		GetOptionalDecVal(&this.Range),
		GetOptionalDecVal(&this.Weight),
		GetOptionalDecVal(&this.Durability),
		GetOptionalDecVal(&this.BuyPrice),
		GetOptionalDecVal(&this.SellPrice),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.Countable),
		GetOptionalDecVal(&this.MagicEffect),
		GetOptionalDecVal(&this.SpecialEffect),
		GetOptionalDecVal(&this.RequireLevel),
		GetOptionalDecVal(&this.MaxLevel),
		GetOptionalDecVal(&this.RequireRank),
		GetOptionalDecVal(&this.RequireTitle),
		GetOptionalDecVal(&this.RequireStrength),
		GetOptionalDecVal(&this.RequireStamina),
		GetOptionalDecVal(&this.RequireDexterity),
		GetOptionalDecVal(&this.RequireIntelligence),
		GetOptionalDecVal(&this.RequireCharisma),
		GetOptionalDecVal(&this.SellingGroup),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.EvasionRate),
		GetOptionalDecVal(&this.DaggerArmor),
		GetOptionalDecVal(&this.SwordArmor),
		GetOptionalDecVal(&this.MaceArmor),
		GetOptionalDecVal(&this.AxeArmor),
		GetOptionalDecVal(&this.SpearArmor),
		GetOptionalDecVal(&this.BowArmor),
		GetOptionalDecVal(&this.FireDamage),
		GetOptionalDecVal(&this.IceDamage),
		GetOptionalDecVal(&this.LightningDamage),
		GetOptionalDecVal(&this.PoisonDamage),
		GetOptionalDecVal(&this.HpDrain),
		GetOptionalDecVal(&this.MpDamage),
		GetOptionalDecVal(&this.MpDrain),
		GetOptionalDecVal(&this.MirrorDamage),
		GetOptionalDecVal(&this.DropRate),
		GetOptionalDecVal(&this.StrengthBonus),
		GetOptionalDecVal(&this.StaminaBonus),
		GetOptionalDecVal(&this.DexterityBonus),
		GetOptionalDecVal(&this.IntelligenceBonus),
		GetOptionalDecVal(&this.CharismaBonus),
		GetOptionalDecVal(&this.MaxHpBonus),
		GetOptionalDecVal(&this.MaxMpBonus),
		GetOptionalDecVal(&this.FireResistance),
		GetOptionalDecVal(&this.ColdResistance),
		GetOptionalDecVal(&this.LightningResistance),
		GetOptionalDecVal(&this.MagicResistance),
		GetOptionalDecVal(&this.PoisonResistance),
		GetOptionalDecVal(&this.CurseResistance))
}

// GetCreateTableString Returns the create table statement for this object
func (this Item) GetCreateTableString() string {
	query := "CREATE TABLE [ITEM] (\n\t[Num] int NOT NULL,\n\t[strName] varchar(50) NOT NULL,\n\t[Kind] tinyint NOT NULL,\n\t[Slot] tinyint NOT NULL,\n\t[Race] tinyint NOT NULL,\n\t[Class] tinyint NOT NULL,\n\t[Damage] smallint NOT NULL,\n\t[Delay] smallint NOT NULL,\n\t[Range] smallint NOT NULL,\n\t[Weight] smallint NOT NULL,\n\t[Duration] smallint NOT NULL,\n\t[BuyPrice] int NOT NULL,\n\t[SellPrice] int NOT NULL,\n\t[Ac] smallint NOT NULL,\n\t[Countable] tinyint NOT NULL,\n\t[Effect1] int NOT NULL,\n\t[Effect2] int NOT NULL,\n\t[ReqLevel] tinyint NOT NULL,\n\t[ReqLevelMax] tinyint NOT NULL,\n\t[ReqRank] tinyint NOT NULL,\n\t[ReqTitle] tinyint NOT NULL,\n\t[ReqStr] tinyint NOT NULL,\n\t[ReqSta] tinyint NOT NULL,\n\t[ReqDex] tinyint NOT NULL,\n\t[ReqIntel] tinyint NOT NULL,\n\t[ReqCha] tinyint NOT NULL,\n\t[SellingGroup] tinyint NOT NULL,\n\t[ItemType] tinyint NOT NULL,\n\t[Hitrate] smallint NOT NULL,\n\t[Evasionrate] smallint NOT NULL,\n\t[DaggerAc] smallint NOT NULL,\n\t[SwordAc] smallint NOT NULL,\n\t[MaceAc] smallint NOT NULL,\n\t[AxeAc] smallint NOT NULL,\n\t[SpearAc] smallint NOT NULL,\n\t[BowAc] smallint NOT NULL,\n\t[FireDamage] tinyint NOT NULL,\n\t[IceDamage] tinyint NOT NULL,\n\t[LightningDamage] tinyint NOT NULL,\n\t[PoisonDamage] tinyint NOT NULL,\n\t[HPDrain] tinyint NOT NULL,\n\t[MPDamage] tinyint NOT NULL,\n\t[MPDrain] tinyint NOT NULL,\n\t[MirrorDamage] tinyint NOT NULL,\n\t[Droprate] tinyint NOT NULL,\n\t[StrB] smallint NOT NULL,\n\t[StaB] smallint NOT NULL,\n\t[DexB] smallint NOT NULL,\n\t[IntelB] smallint NOT NULL,\n\t[ChaB] smallint NOT NULL,\n\t[MaxHpB] smallint NOT NULL,\n\t[MaxMpB] smallint NOT NULL,\n\t[FireR] smallint NOT NULL,\n\t[ColdR] smallint NOT NULL,\n\t[LightningR] smallint NOT NULL,\n\t[MagicR] smallint NOT NULL,\n\t[PoisonR] smallint NOT NULL,\n\t[CurseR] smallint NOT NULL\n\tCONSTRAINT [PK_ITEM] PRIMARY KEY ([Num])\n)\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Kind] DEFAULT 0 FOR [Kind]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Slot] DEFAULT 0 FOR [Slot]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Race] DEFAULT 0 FOR [Race]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Class] DEFAULT 1 FOR [Class]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Damage] DEFAULT 0 FOR [Damage]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Delay] DEFAULT 0 FOR [Delay]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Range] DEFAULT 0 FOR [Range]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Weight] DEFAULT 0 FOR [Weight]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_BuyPrice] DEFAULT 0 FOR [BuyPrice]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Ac] DEFAULT 0 FOR [Ac]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_DaggerAc] DEFAULT 0 FOR [DaggerAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_SwordAc] DEFAULT 0 FOR [SwordAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_MaceAc] DEFAULT 0 FOR [MaceAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_AxeAc] DEFAULT 0 FOR [AxeAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_SpearAc] DEFAULT 0 FOR [SpearAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_BowAc] DEFAULT 0 FOR [BowAc]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_MirrorDamage] DEFAULT 0 FOR [MirrorDamage]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_Droprate] DEFAULT 0 FOR [Droprate]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_StrB] DEFAULT 0 FOR [StrB]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_StaB] DEFAULT 0 FOR [StaB]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_DexB] DEFAULT 0 FOR [DexB]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_IntelB] DEFAULT 0 FOR [IntelB]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_ChaB] DEFAULT 0 FOR [ChaB]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_FireR] DEFAULT 0 FOR [FireR]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_ColdR] DEFAULT 0 FOR [ColdR]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_LightningR] DEFAULT 0 FOR [LightningR]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_MagicR] DEFAULT 0 FOR [MagicR]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_PoisonR] DEFAULT 0 FOR [PoisonR]\nGO\nALTER TABLE [ITEM] ADD CONSTRAINT [DF_ITEM_CurseR] DEFAULT 0 FOR [CurseR]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this Item) SelectClause() (selectClause clause.Select) {
	return _ItemSelectClause
}

// GetAllTableData Returns a list of all table data
func (this Item) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []Item{}
	rawSql := "SELECT [Num], [strName], [Kind], [Slot], [Race], [Class], [Damage], [Delay], [Range], [Weight], [Duration], [BuyPrice], [SellPrice], [Ac], [Countable], [Effect1], [Effect2], [ReqLevel], [ReqLevelMax], [ReqRank], [ReqTitle], [ReqStr], [ReqSta], [ReqDex], [ReqIntel], [ReqCha], [SellingGroup], [ItemType], [Hitrate], [Evasionrate], [DaggerAc], [SwordAc], [MaceAc], [AxeAc], [SpearAc], [BowAc], [FireDamage], [IceDamage], [LightningDamage], [PoisonDamage], [HPDrain], [MPDamage], [MPDrain], [MirrorDamage], [Droprate], [StrB], [StaB], [DexB], [IntelB], [ChaB], [MaxHpB], [MaxMpB], [FireR], [ColdR], [LightningR], [MagicR], [PoisonR], [CurseR] FROM [ITEM]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ItemSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[Num]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[Kind]",
		},
		clause.Column{
			Name: "[Slot]",
		},
		clause.Column{
			Name: "[Race]",
		},
		clause.Column{
			Name: "[Class]",
		},
		clause.Column{
			Name: "[Damage]",
		},
		clause.Column{
			Name: "[Delay]",
		},
		clause.Column{
			Name: "[Range]",
		},
		clause.Column{
			Name: "[Weight]",
		},
		clause.Column{
			Name: "[Duration]",
		},
		clause.Column{
			Name: "[BuyPrice]",
		},
		clause.Column{
			Name: "[SellPrice]",
		},
		clause.Column{
			Name: "[Ac]",
		},
		clause.Column{
			Name: "[Countable]",
		},
		clause.Column{
			Name: "[Effect1]",
		},
		clause.Column{
			Name: "[Effect2]",
		},
		clause.Column{
			Name: "[ReqLevel]",
		},
		clause.Column{
			Name: "[ReqLevelMax]",
		},
		clause.Column{
			Name: "[ReqRank]",
		},
		clause.Column{
			Name: "[ReqTitle]",
		},
		clause.Column{
			Name: "[ReqStr]",
		},
		clause.Column{
			Name: "[ReqSta]",
		},
		clause.Column{
			Name: "[ReqDex]",
		},
		clause.Column{
			Name: "[ReqIntel]",
		},
		clause.Column{
			Name: "[ReqCha]",
		},
		clause.Column{
			Name: "[SellingGroup]",
		},
		clause.Column{
			Name: "[ItemType]",
		},
		clause.Column{
			Name: "[Hitrate]",
		},
		clause.Column{
			Name: "[Evasionrate]",
		},
		clause.Column{
			Name: "[DaggerAc]",
		},
		clause.Column{
			Name: "[SwordAc]",
		},
		clause.Column{
			Name: "[MaceAc]",
		},
		clause.Column{
			Name: "[AxeAc]",
		},
		clause.Column{
			Name: "[SpearAc]",
		},
		clause.Column{
			Name: "[BowAc]",
		},
		clause.Column{
			Name: "[FireDamage]",
		},
		clause.Column{
			Name: "[IceDamage]",
		},
		clause.Column{
			Name: "[LightningDamage]",
		},
		clause.Column{
			Name: "[PoisonDamage]",
		},
		clause.Column{
			Name: "[HPDrain]",
		},
		clause.Column{
			Name: "[MPDamage]",
		},
		clause.Column{
			Name: "[MPDrain]",
		},
		clause.Column{
			Name: "[MirrorDamage]",
		},
		clause.Column{
			Name: "[Droprate]",
		},
		clause.Column{
			Name: "[StrB]",
		},
		clause.Column{
			Name: "[StaB]",
		},
		clause.Column{
			Name: "[DexB]",
		},
		clause.Column{
			Name: "[IntelB]",
		},
		clause.Column{
			Name: "[ChaB]",
		},
		clause.Column{
			Name: "[MaxHpB]",
		},
		clause.Column{
			Name: "[MaxMpB]",
		},
		clause.Column{
			Name: "[FireR]",
		},
		clause.Column{
			Name: "[ColdR]",
		},
		clause.Column{
			Name: "[LightningR]",
		},
		clause.Column{
			Name: "[MagicR]",
		},
		clause.Column{
			Name: "[PoisonR]",
		},
		clause.Column{
			Name: "[CurseR]",
		},
	},
}
