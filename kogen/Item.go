package kogen

import (
	"fmt"
)

const (
	_ItemDatabaseNbr = 0
	_ItemTableName   = "ITEM"
)

func init() {
	ModelList = append(ModelList, &Item{})
}

// Item: Item information
type Item struct {
	Number              int      `gorm:"column:Num;type:int;primaryKey;not null" json:"Num"`
	Name                [50]byte `gorm:"column:strName;type:varchar(50);not null" json:"strName"`
	Kind                uint8    `gorm:"column:Kind;type:tinyint;not null" json:"Kind"`
	Slot                uint8    `gorm:"column:Slot;type:tinyint;not null" json:"Slot"`
	Race                uint8    `gorm:"column:Race;type:tinyint;not null" json:"Race"`
	ClassId             uint8    `gorm:"column:Class;type:tinyint;not null" json:"Class"`
	Damage              int16    `gorm:"column:Damage;type:smallint;not null" json:"Damage"`
	Delay               int16    `gorm:"column:Delay;type:smallint;not null" json:"Delay"`
	Range               int16    `gorm:"column:Range;type:smallint;not null" json:"Range"`
	Weight              int16    `gorm:"column:Weight;type:smallint;not null" json:"Weight"`
	Durability          int16    `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	BuyPrice            int16    `gorm:"column:BuyPrice;type:smallint;not null" json:"BuyPrice"`
	SellPrice           int16    `gorm:"column:SellPrice;type:smallint;not null" json:"SellPrice"`
	Armor               int16    `gorm:"column:Ac;type:smallint;not null" json:"Ac"`
	Countable           uint8    `gorm:"column:Countable;type:tinyint;not null" json:"Countable"`
	MagicEffect         int      `gorm:"column:Effect1;type:int;not null" json:"Effect1"`
	SpecialEffect       int      `gorm:"column:Effect2;type:int;not null" json:"Effect2"`
	RequireLevel        uint8    `gorm:"column:ReqLevel;type:tinyint;not null" json:"ReqLevel"`
	MaxLevel            uint8    `gorm:"column:ReqLevelMax;type:tinyint;not null" json:"ReqLevelMax"`
	RequireRank         uint8    `gorm:"column:ReqRank;type:tinyint;not null" json:"ReqRank"`
	RequireTitle        uint8    `gorm:"column:ReqTitle;type:tinyint;not null" json:"ReqTitle"`
	RequireStrength     uint8    `gorm:"column:ReqStr;type:tinyint;not null" json:"ReqStr"`
	RequireStamina      uint8    `gorm:"column:ReqSta;type:tinyint;not null" json:"ReqSta"`
	RequireDexterity    uint8    `gorm:"column:ReqDex;type:tinyint;not null" json:"ReqDex"`
	RequireIntelligence uint8    `gorm:"column:ReqIntel;type:tinyint;not null" json:"ReqIntel"`
	RequireCharisma     uint8    `gorm:"column:ReqCha;type:tinyint;not null" json:"ReqCha"`
	SellingGroup        uint8    `gorm:"column:SellingGroup;type:tinyint;not null" json:"SellingGroup"`
	Type                uint8    `gorm:"column:ItemType;type:tinyint;not null" json:"ItemType"`
	HitRate             int16    `gorm:"column:Hitrate;type:smallint;not null" json:"Hitrate"`
	EvasionRate         int16    `gorm:"column:Evasionrate;type:smallint;not null" json:"Evasionrate"`
	DaggerArmor         int16    `gorm:"column:DaggerAc;type:smallint;not null" json:"DaggerAc"`
	SwordArmor          int16    `gorm:"column:SwordAc;type:smallint;not null" json:"SwordAc"`
	MaceArmor           int16    `gorm:"column:MaceAc;type:smallint;not null" json:"MaceAc"`
	AxeArmor            int16    `gorm:"column:AxeAc;type:smallint;not null" json:"AxeAc"`
	SpearArmor          int16    `gorm:"column:SpearAc;type:smallint;not null" json:"SpearAc"`
	BowArmor            int16    `gorm:"column:BowAc;type:smallint;not null" json:"BowAc"`
	FireDamage          uint8    `gorm:"column:FireDamage;type:tinyint;not null" json:"FireDamage"`
	IceDamage           uint8    `gorm:"column:IceDamage;type:tinyint;not null" json:"IceDamage"`
	LightningDamage     uint8    `gorm:"column:LightningDamage;type:tinyint;not null" json:"LightningDamage"`
	PoisonDamage        uint8    `gorm:"column:PoisonDamage;type:tinyint;not null" json:"PoisonDamage"`
	HpDrain             uint8    `gorm:"column:HPDrain;type:tinyint;not null" json:"HPDrain"`
	MpDamage            uint8    `gorm:"column:MPDamage;type:tinyint;not null" json:"MPDamage"`
	MpDrain             uint8    `gorm:"column:MPDrain;type:tinyint;not null" json:"MPDrain"`
	MirrorDamage        uint8    `gorm:"column:MirrorDamage;type:tinyint;not null" json:"MirrorDamage"`
	DropRate            uint8    `gorm:"column:Droprate;type:tinyint;not null" json:"Droprate"`
	StrengthBonus       int16    `gorm:"column:StrB;type:smallint;not null" json:"StrB"`
	StaminaBonus        int16    `gorm:"column:StaB;type:smallint;not null" json:"StaB"`
	DexterityBonus      int16    `gorm:"column:DexB;type:smallint;not null" json:"DexB"`
	IntelligenceBonus   int16    `gorm:"column:IntelB;type:smallint;not null" json:"IntelB"`
	CharismaBonus       int16    `gorm:"column:ChaB;type:smallint;not null" json:"ChaB"`
	MaxHpBonus          int16    `gorm:"column:MaxHpB;type:smallint;not null" json:"MaxHpB"`
	MaxMpBonus          int16    `gorm:"column:MaxMpB;type:smallint;not null" json:"MaxMpB"`
	FireResistance      int16    `gorm:"column:FireR;type:smallint;not null" json:"FireR"`
	ColdResistance      int16    `gorm:"column:ColdR;type:smallint;not null" json:"ColdR"`
	LightningResistance int16    `gorm:"column:LightningR;type:smallint;not null" json:"LightningR"`
	MagicResistance     int16    `gorm:"column:MagicR;type:smallint;not null" json:"MagicR"`
	PoisonResistance    int16    `gorm:"column:PoisonR;type:smallint;not null" json:"PoisonR"`
	CurseResistance     int16    `gorm:"column:CurseR;type:smallint;not null" json:"CurseR"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *Item) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemDatabaseNbr))
}

// GetTableName Returns the table name
func (this *Item) GetTableName() string {
	return _ItemTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *Item) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM] (Num, strName, Kind, Slot, Race, Class, Damage, Delay, Range, Weight, Duration, BuyPrice, SellPrice, Ac, Countable, Effect1, Effect2, ReqLevel, ReqLevelMax, ReqRank, ReqTitle, ReqStr, ReqSta, ReqDex, ReqIntel, ReqCha, SellingGroup, ItemType, Hitrate, Evasionrate, DaggerAc, SwordAc, MaceAc, AxeAc, SpearAc, BowAc, FireDamage, IceDamage, LightningDamage, PoisonDamage, HPDrain, MPDamage, MPDrain, MirrorDamage, Droprate, StrB, StaB, DexB, IntelB, ChaB, MaxHpB, MaxMpB, FireR, ColdR, LightningR, MagicR, PoisonR, CurseR) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Number),
		GetOptionalBinaryVal(this.Name),
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
func (this *Item) GetCreateTableString() string {
	query := "CREATE TABLE \"ITEM\" (\n\t\"Num\" int NOT NULL,\n\t\"strName\" varchar(50) NOT NULL,\n\t\"Kind\" tinyint NOT NULL,\n\t\"Slot\" tinyint NOT NULL,\n\t\"Race\" tinyint NOT NULL,\n\t\"Class\" tinyint NOT NULL,\n\t\"Damage\" smallint NOT NULL,\n\t\"Delay\" smallint NOT NULL,\n\t\"Range\" smallint NOT NULL,\n\t\"Weight\" smallint NOT NULL,\n\t\"Duration\" smallint NOT NULL,\n\t\"BuyPrice\" smallint NOT NULL,\n\t\"SellPrice\" smallint NOT NULL,\n\t\"Ac\" smallint NOT NULL,\n\t\"Countable\" tinyint NOT NULL,\n\t\"Effect1\" int NOT NULL,\n\t\"Effect2\" int NOT NULL,\n\t\"ReqLevel\" tinyint NOT NULL,\n\t\"ReqLevelMax\" tinyint NOT NULL,\n\t\"ReqRank\" tinyint NOT NULL,\n\t\"ReqTitle\" tinyint NOT NULL,\n\t\"ReqStr\" tinyint NOT NULL,\n\t\"ReqSta\" tinyint NOT NULL,\n\t\"ReqDex\" tinyint NOT NULL,\n\t\"ReqIntel\" tinyint NOT NULL,\n\t\"ReqCha\" tinyint NOT NULL,\n\t\"SellingGroup\" tinyint NOT NULL,\n\t\"ItemType\" tinyint NOT NULL,\n\t\"Hitrate\" smallint NOT NULL,\n\t\"Evasionrate\" smallint NOT NULL,\n\t\"DaggerAc\" smallint NOT NULL,\n\t\"SwordAc\" smallint NOT NULL,\n\t\"MaceAc\" smallint NOT NULL,\n\t\"AxeAc\" smallint NOT NULL,\n\t\"SpearAc\" smallint NOT NULL,\n\t\"BowAc\" smallint NOT NULL,\n\t\"FireDamage\" tinyint NOT NULL,\n\t\"IceDamage\" tinyint NOT NULL,\n\t\"LightningDamage\" tinyint NOT NULL,\n\t\"PoisonDamage\" tinyint NOT NULL,\n\t\"HPDrain\" tinyint NOT NULL,\n\t\"MPDamage\" tinyint NOT NULL,\n\t\"MPDrain\" tinyint NOT NULL,\n\t\"MirrorDamage\" tinyint NOT NULL,\n\t\"Droprate\" tinyint NOT NULL,\n\t\"StrB\" smallint NOT NULL,\n\t\"StaB\" smallint NOT NULL,\n\t\"DexB\" smallint NOT NULL,\n\t\"IntelB\" smallint NOT NULL,\n\t\"ChaB\" smallint NOT NULL,\n\t\"MaxHpB\" smallint NOT NULL,\n\t\"MaxMpB\" smallint NOT NULL,\n\t\"FireR\" smallint NOT NULL,\n\t\"ColdR\" smallint NOT NULL,\n\t\"LightningR\" smallint NOT NULL,\n\t\"MagicR\" smallint NOT NULL,\n\t\"PoisonR\" smallint NOT NULL,\n\t\"CurseR\" smallint NOT NULL\n\tPRIMARY KEY (\"Num\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
