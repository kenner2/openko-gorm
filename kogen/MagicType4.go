package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_MagicType4DatabaseNbr = "GAME"
	_MagicType4TableName   = "MAGIC_TYPE4"
)

func init() {
	ModelList = append(ModelList, &MagicType4{})
}

// MagicType4 Supports stat modification skills
type MagicType4 struct {
	MagicNumber     int            `gorm:"column:iNum;type:int;primaryKey;not null" json:"iNum"`
	Name            *mssql.VarChar `gorm:"column:Name;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Name,omitempty"`
	Description     *mssql.VarChar `gorm:"column:Description;type:varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"Description,omitempty"`
	BuffType        uint8          `gorm:"column:BuffType;type:tinyint;not null" json:"BuffType"`
	Radius          uint8          `gorm:"column:Radius;type:tinyint;not null" json:"Radius"`
	Duration        int16          `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	AttackSpeed     uint8          `gorm:"column:AttackSpeed;type:tinyint;not null" json:"AttackSpeed"`
	Speed           uint8          `gorm:"column:Speed;type:tinyint;not null" json:"Speed"`
	Armor           int16          `gorm:"column:AC;type:smallint;not null" json:"AC"`
	ArmorPercent    int16          `gorm:"column:ACPct;type:smallint;not null" json:"ACPct"`
	AttackPower     uint8          `gorm:"column:Attack;type:tinyint;not null" json:"Attack"`
	MagicPower      uint8          `gorm:"column:MagicAttack;type:tinyint;not null" json:"MagicAttack"`
	MaxHp           int16          `gorm:"column:MaxHP;type:smallint;not null" json:"MaxHP"`
	MaxHpPercent    int16          `gorm:"column:MaxHpPct;type:smallint;not null" json:"MaxHpPct"`
	MaxMp           int16          `gorm:"column:MaxMP;type:smallint;not null" json:"MaxMP"`
	MaxMpPercent    int16          `gorm:"column:MaxMpPct;type:smallint;not null" json:"MaxMpPct"`
	HitRate         uint8          `gorm:"column:HitRate;type:tinyint;not null" json:"HitRate"`
	AvoidRate       int16          `gorm:"column:AvoidRate;type:smallint;not null" json:"AvoidRate"`
	Strength        int16          `gorm:"column:Str;type:smallint;not null" json:"Str"`
	Stamina         int16          `gorm:"column:Sta;type:smallint;not null" json:"Sta"`
	Dexterity       int16          `gorm:"column:Dex;type:smallint;not null" json:"Dex"`
	Intelligence    int16          `gorm:"column:Intel;type:smallint;not null" json:"Intel"`
	Charisma        int16          `gorm:"column:Cha;type:smallint;not null" json:"Cha"`
	FireResist      uint8          `gorm:"column:FireR;type:tinyint;not null" json:"FireR"`
	ColdResist      uint8          `gorm:"column:ColdR;type:tinyint;not null" json:"ColdR"`
	LightningResist uint8          `gorm:"column:LightningR;type:tinyint;not null" json:"LightningR"`
	MagicResist     uint8          `gorm:"column:MagicR;type:tinyint;not null" json:"MagicR"`
	DiseaseResist   uint8          `gorm:"column:DiseaseR;type:tinyint;not null" json:"DiseaseR"`
	PoisonResist    uint8          `gorm:"column:PoisonR;type:tinyint;not null" json:"PoisonR"`
	ExpPercent      uint8          `gorm:"column:ExpPct;type:tinyint;not null" json:"ExpPct"`
}

// GetDatabaseName Returns the table's database name
func (this MagicType4) GetDatabaseName() string {
	return GetDatabaseName(_MagicType4DatabaseNbr)
}

// TableName Returns the table name
func (this MagicType4) TableName() string {
	return _MagicType4TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this MagicType4) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE4] ([iNum], [Name], [Description], [BuffType], [Radius], [Duration], [AttackSpeed], [Speed], [AC], [ACPct], [Attack], [MagicAttack], [MaxHP], [MaxHpPct], [MaxMP], [MaxMpPct], [HitRate], [AvoidRate], [Str], [Sta], [Dex], [Intel], [Cha], [FireR], [ColdR], [LightningR], [MagicR], [DiseaseR], [PoisonR], [ExpPct]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
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

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this MagicType4) GetInsertHeader() string {
	return "INSERT INTO [MAGIC_TYPE4] ([iNum], [Name], [Description], [BuffType], [Radius], [Duration], [AttackSpeed], [Speed], [AC], [ACPct], [Attack], [MagicAttack], [MaxHP], [MaxHpPct], [MaxMP], [MaxMpPct], [HitRate], [AvoidRate], [Str], [Sta], [Dex], [Intel], [Cha], [FireR], [ColdR], [LightningR], [MagicR], [DiseaseR], [PoisonR], [ExpPct]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this MagicType4) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalVarCharVal(this.Description, false),
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
func (this MagicType4) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE4] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[Description] varchar(100) COLLATE SQL_Latin1_General_CP1_CI_AS,\n\t[BuffType] tinyint NOT NULL,\n\t[Radius] tinyint NOT NULL,\n\t[Duration] smallint NOT NULL,\n\t[AttackSpeed] tinyint NOT NULL,\n\t[Speed] tinyint NOT NULL,\n\t[AC] smallint NOT NULL,\n\t[ACPct] smallint NOT NULL,\n\t[Attack] tinyint NOT NULL,\n\t[MagicAttack] tinyint NOT NULL,\n\t[MaxHP] smallint NOT NULL,\n\t[MaxHpPct] smallint NOT NULL,\n\t[MaxMP] smallint NOT NULL,\n\t[MaxMpPct] smallint NOT NULL,\n\t[HitRate] tinyint NOT NULL,\n\t[AvoidRate] smallint NOT NULL,\n\t[Str] smallint NOT NULL,\n\t[Sta] smallint NOT NULL,\n\t[Dex] smallint NOT NULL,\n\t[Intel] smallint NOT NULL,\n\t[Cha] smallint NOT NULL,\n\t[FireR] tinyint NOT NULL,\n\t[ColdR] tinyint NOT NULL,\n\t[LightningR] tinyint NOT NULL,\n\t[MagicR] tinyint NOT NULL,\n\t[DiseaseR] tinyint NOT NULL,\n\t[PoisonR] tinyint NOT NULL,\n\t[ExpPct] tinyint NOT NULL\n\tCONSTRAINT [PK_MAGIC_TYPE4] PRIMARY KEY CLUSTERED ([iNum])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this MagicType4) SelectClause() (selectClause clause.Select) {
	return _MagicType4SelectClause
}

// GetAllTableData Returns a list of all table data
func (this MagicType4) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []MagicType4{}
	rawSql := "SELECT [iNum], [Name], [Description], [BuffType], [Radius], [Duration], [AttackSpeed], [Speed], [AC], [ACPct], [Attack], [MagicAttack], [MaxHP], [MaxHpPct], [MaxMP], [MaxMpPct], [HitRate], [AvoidRate], [Str], [Sta], [Dex], [Intel], [Cha], [FireR], [ColdR], [LightningR], [MagicR], [DiseaseR], [PoisonR], [ExpPct] FROM [MAGIC_TYPE4]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _MagicType4SelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[iNum]",
		},
		clause.Column{
			Name: "[Name]",
		},
		clause.Column{
			Name: "[Description]",
		},
		clause.Column{
			Name: "[BuffType]",
		},
		clause.Column{
			Name: "[Radius]",
		},
		clause.Column{
			Name: "[Duration]",
		},
		clause.Column{
			Name: "[AttackSpeed]",
		},
		clause.Column{
			Name: "[Speed]",
		},
		clause.Column{
			Name: "[AC]",
		},
		clause.Column{
			Name: "[ACPct]",
		},
		clause.Column{
			Name: "[Attack]",
		},
		clause.Column{
			Name: "[MagicAttack]",
		},
		clause.Column{
			Name: "[MaxHP]",
		},
		clause.Column{
			Name: "[MaxHpPct]",
		},
		clause.Column{
			Name: "[MaxMP]",
		},
		clause.Column{
			Name: "[MaxMpPct]",
		},
		clause.Column{
			Name: "[HitRate]",
		},
		clause.Column{
			Name: "[AvoidRate]",
		},
		clause.Column{
			Name: "[Str]",
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
			Name: "[DiseaseR]",
		},
		clause.Column{
			Name: "[PoisonR]",
		},
		clause.Column{
			Name: "[ExpPct]",
		},
	},
}
