package kogen

import (
	"fmt"
)

const (
	_MagicType4DatabaseNbr = 0
	_MagicType4TableName   = "MAGIC_TYPE4"
)

func init() {
	ModelList = append(ModelList, &MagicType4{})
}

// MagicType4 Type 4 supports stat modification skills
type MagicType4 struct {
	MagicNumber     int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name            [50]byte  `gorm:"column:Name;type:varchar(50)" json:"Name,omitempty"`
	Description     [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	BuffType        uint8     `gorm:"column:BuffType;type:tinyint;not null" json:"BuffType"`
	Radius          uint8     `gorm:"column:Radius;type:tinyint;not null" json:"Radius"`
	Duration        int16     `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	AttackSpeed     uint8     `gorm:"column:AttackSpeed;type:tinyint;not null" json:"AttackSpeed"`
	Speed           uint8     `gorm:"column:Speed;type:tinyint;not null" json:"Speed"`
	Armor           int16     `gorm:"column:AC;type:smallint;not null" json:"AC"`
	ArmorPercent    int16     `gorm:"column:ACPct;type:smallint;not null" json:"ACPct"`
	AttackPower     uint8     `gorm:"column:Attack;type:tinyint;not null" json:"Attack"`
	MagicPower      uint8     `gorm:"column:MagicAttack;type:tinyint;not null" json:"MagicAttack"`
	MaxHp           int16     `gorm:"column:MaxHP;type:smallint;not null" json:"MaxHP"`
	MaxHpPercent    int16     `gorm:"column:MaxHpPct;type:smallint;not null" json:"MaxHpPct"`
	MaxMp           int16     `gorm:"column:MaxMP;type:smallint;not null" json:"MaxMP"`
	MaxMpPercent    int16     `gorm:"column:MaxMpPct;type:smallint;not null" json:"MaxMpPct"`
	HitRate         uint8     `gorm:"column:HitRate;type:tinyint;not null" json:"HitRate"`
	AvoidRate       int16     `gorm:"column:AvoidRate;type:smallint;not null" json:"AvoidRate"`
	Strength        int16     `gorm:"column:Str;type:smallint;not null" json:"Str"`
	Stamina         int16     `gorm:"column:Sta;type:smallint;not null" json:"Sta"`
	Dexterity       int16     `gorm:"column:Dex;type:smallint;not null" json:"Dex"`
	Intelligence    int16     `gorm:"column:Intel;type:smallint;not null" json:"Intel"`
	Charisma        int16     `gorm:"column:Cha;type:smallint;not null" json:"Cha"`
	FireResist      uint8     `gorm:"column:FireR;type:tinyint;not null" json:"FireR"`
	ColdResist      uint8     `gorm:"column:ColdR;type:tinyint;not null" json:"ColdR"`
	LightningResist uint8     `gorm:"column:LightningR;type:tinyint;not null" json:"LightningR"`
	MagicResist     uint8     `gorm:"column:MagicR;type:tinyint;not null" json:"MagicR"`
	DiseaseResist   uint8     `gorm:"column:DiseaseR;type:tinyint;not null" json:"DiseaseR"`
	PoisonResist    uint8     `gorm:"column:PoisonR;type:tinyint;not null" json:"PoisonR"`
	ExpPercent      uint8     `gorm:"column:ExpPct;type:tinyint;not null" json:"ExpPct"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType4) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType4DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType4) GetTableName() string {
	return _MagicType4TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType4) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE4] (iNum, Name, Description, BuffType, Radius, Duration, AttackSpeed, Speed, AC, ACPct, Attack, MagicAttack, MaxHP, MaxHpPct, MaxMP, MaxMpPct, HitRate, AvoidRate, Str, Sta, Dex, Intel, Cha, FireR, ColdR, LightningR, MagicR, DiseaseR, PoisonR, ExpPct) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.BuffType),
		GetOptionalDecVal(&this.Radius),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.AttackSpeed),
		GetOptionalDecVal(&this.Speed),
		GetOptionalDecVal(&this.Armor),
		GetOptionalDecVal(&this.ArmorPercent),
		GetOptionalDecVal(&this.AttackPower),
		GetOptionalDecVal(&this.MagicPower),
		GetOptionalDecVal(&this.MaxHp),
		GetOptionalDecVal(&this.MaxHpPercent),
		GetOptionalDecVal(&this.MaxMp),
		GetOptionalDecVal(&this.MaxMpPercent),
		GetOptionalDecVal(&this.HitRate),
		GetOptionalDecVal(&this.AvoidRate),
		GetOptionalDecVal(&this.Strength),
		GetOptionalDecVal(&this.Stamina),
		GetOptionalDecVal(&this.Dexterity),
		GetOptionalDecVal(&this.Intelligence),
		GetOptionalDecVal(&this.Charisma),
		GetOptionalDecVal(&this.FireResist),
		GetOptionalDecVal(&this.ColdResist),
		GetOptionalDecVal(&this.LightningResist),
		GetOptionalDecVal(&this.MagicResist),
		GetOptionalDecVal(&this.DiseaseResist),
		GetOptionalDecVal(&this.PoisonResist),
		GetOptionalDecVal(&this.ExpPercent))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType4) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE4] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50),\n\t[Description] varchar(100),\n\t[BuffType] tinyint NOT NULL,\n\t[Radius] tinyint NOT NULL,\n\t[Duration] smallint NOT NULL,\n\t[AttackSpeed] tinyint NOT NULL,\n\t[Speed] tinyint NOT NULL,\n\t[AC] smallint NOT NULL,\n\t[ACPct] smallint NOT NULL,\n\t[Attack] tinyint NOT NULL,\n\t[MagicAttack] tinyint NOT NULL,\n\t[MaxHP] smallint NOT NULL,\n\t[MaxHpPct] smallint NOT NULL,\n\t[MaxMP] smallint NOT NULL,\n\t[MaxMpPct] smallint NOT NULL,\n\t[HitRate] tinyint NOT NULL,\n\t[AvoidRate] smallint NOT NULL,\n\t[Str] smallint NOT NULL,\n\t[Sta] smallint NOT NULL,\n\t[Dex] smallint NOT NULL,\n\t[Intel] smallint NOT NULL,\n\t[Cha] smallint NOT NULL,\n\t[FireR] tinyint NOT NULL,\n\t[ColdR] tinyint NOT NULL,\n\t[LightningR] tinyint NOT NULL,\n\t[MagicR] tinyint NOT NULL,\n\t[DiseaseR] tinyint NOT NULL,\n\t[PoisonR] tinyint NOT NULL,\n\t[ExpPct] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
