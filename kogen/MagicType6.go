package kogen

import (
	"fmt"
)

const (
	_MagicType6DatabaseNbr = 0
	_MagicType6TableName   = "MAGIC_TYPE6"
)

func init() {
	ModelList = append(ModelList, &MagicType6{})
}

// MagicType6 Type 6 supports transformation magic
type MagicType6 struct {
	MagicNumber          int       `gorm:"column:iNum;type:int;not null" json:"iNum"`
	Name                 [50]byte  `gorm:"column:Name;type:varchar(50);not null" json:"Name"`
	Description          [100]byte `gorm:"column:Description;type:varchar(100)" json:"Description,omitempty"`
	Size                 int16     `gorm:"column:Size;type:smallint;not null" json:"Size"`
	TransformId          int16     `gorm:"column:TransformID;type:smallint;not null" json:"TransformID"`
	Duration             int16     `gorm:"column:Duration;type:smallint;not null" json:"Duration"`
	MaxHp                int16     `gorm:"column:MaxHp;type:smallint;not null" json:"MaxHp"`
	MaxMp                int16     `gorm:"column:MaxMp;type:smallint;not null" json:"MaxMp"`
	Speed                uint8     `gorm:"column:Speed;type:tinyint;not null" json:"Speed"`
	AttackSpeed          int16     `gorm:"column:AttackSpeed;type:smallint;not null" json:"AttackSpeed"`
	TotalHit             int16     `gorm:"column:TotalHit;type:smallint;not null" json:"TotalHit"`
	TotalArmor           int16     `gorm:"column:TotalAc;type:smallint;not null" json:"TotalAc"`
	TotalHitRate         int16     `gorm:"column:TotalHitRate;type:smallint;not null" json:"TotalHitRate"`
	TotalEvasionRate     int16     `gorm:"column:TotalEvasionRate;type:smallint;not null" json:"TotalEvasionRate"`
	TotalFireResist      int16     `gorm:"column:TotalFireR;type:smallint;not null" json:"TotalFireR"`
	TotalColdResist      int16     `gorm:"column:TotalColdR;type:smallint;not null" json:"TotalColdR"`
	TotalLightningResist int16     `gorm:"column:TotalLightningR;type:smallint;not null" json:"TotalLightningR"`
	TotalMagicResist     int16     `gorm:"column:TotalMagicR;type:smallint;not null" json:"TotalMagicR"`
	TotalDiseaseResist   int16     `gorm:"column:TotalDiseaseR;type:smallint;not null" json:"TotalDiseaseR"`
	TotalPoisonResist    int16     `gorm:"column:TotalPoisonR;type:smallint;not null" json:"TotalPoisonR"`
	Class                int16     `gorm:"column:Class;type:smallint;not null" json:"Class"`
	UserSkillUse         uint8     `gorm:"column:UserSkillUse;type:tinyint;not null" json:"UserSkillUse"`
	NeedItem             uint8     `gorm:"column:NeedItem;type:tinyint;not null" json:"NeedItem"`
	SkillSuccessRate     uint8     `gorm:"column:SkillSuccessRate;type:tinyint;not null" json:"SkillSuccessRate"`
	MonsterFriendly      uint8     `gorm:"column:MonsterFriendly;type:tinyint;not null" json:"MonsterFriendly"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *MagicType6) GetDatabaseName() string {
	return GetDatabaseName(DbType(_MagicType6DatabaseNbr))
}

// GetTableName Returns the table name
func (this *MagicType6) GetTableName() string {
	return _MagicType6TableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *MagicType6) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [MAGIC_TYPE6] (iNum, Name, Description, Size, TransformID, Duration, MaxHp, MaxMp, Speed, AttackSpeed, TotalHit, TotalAc, TotalHitRate, TotalEvasionRate, TotalFireR, TotalColdR, TotalLightningR, TotalMagicR, TotalDiseaseR, TotalPoisonR, Class, UserSkillUse, NeedItem, SkillSuccessRate, MonsterFriendly) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.MagicNumber),
		GetOptionalBinaryVal(this.Name),
		GetOptionalBinaryVal(this.Description),
		GetOptionalDecVal(&this.Size),
		GetOptionalDecVal(&this.TransformId),
		GetOptionalDecVal(&this.Duration),
		GetOptionalDecVal(&this.MaxHp),
		GetOptionalDecVal(&this.MaxMp),
		GetOptionalDecVal(&this.Speed),
		GetOptionalDecVal(&this.AttackSpeed),
		GetOptionalDecVal(&this.TotalHit),
		GetOptionalDecVal(&this.TotalArmor),
		GetOptionalDecVal(&this.TotalHitRate),
		GetOptionalDecVal(&this.TotalEvasionRate),
		GetOptionalDecVal(&this.TotalFireResist),
		GetOptionalDecVal(&this.TotalColdResist),
		GetOptionalDecVal(&this.TotalLightningResist),
		GetOptionalDecVal(&this.TotalMagicResist),
		GetOptionalDecVal(&this.TotalDiseaseResist),
		GetOptionalDecVal(&this.TotalPoisonResist),
		GetOptionalDecVal(&this.Class),
		GetOptionalDecVal(&this.UserSkillUse),
		GetOptionalDecVal(&this.NeedItem),
		GetOptionalDecVal(&this.SkillSuccessRate),
		GetOptionalDecVal(&this.MonsterFriendly))
}

// GetCreateTableString Returns the create table statement for this object
func (this *MagicType6) GetCreateTableString() string {
	query := "CREATE TABLE [MAGIC_TYPE6] (\n\t[iNum] int NOT NULL,\n\t[Name] varchar(50) NOT NULL,\n\t[Description] varchar(100),\n\t[Size] smallint NOT NULL,\n\t[TransformID] smallint NOT NULL,\n\t[Duration] smallint NOT NULL,\n\t[MaxHp] smallint NOT NULL,\n\t[MaxMp] smallint NOT NULL,\n\t[Speed] tinyint NOT NULL,\n\t[AttackSpeed] smallint NOT NULL,\n\t[TotalHit] smallint NOT NULL,\n\t[TotalAc] smallint NOT NULL,\n\t[TotalHitRate] smallint NOT NULL,\n\t[TotalEvasionRate] smallint NOT NULL,\n\t[TotalFireR] smallint NOT NULL,\n\t[TotalColdR] smallint NOT NULL,\n\t[TotalLightningR] smallint NOT NULL,\n\t[TotalMagicR] smallint NOT NULL,\n\t[TotalDiseaseR] smallint NOT NULL,\n\t[TotalPoisonR] smallint NOT NULL,\n\t[Class] smallint NOT NULL,\n\t[UserSkillUse] tinyint NOT NULL,\n\t[NeedItem] tinyint NOT NULL,\n\t[SkillSuccessRate] tinyint NOT NULL,\n\t[MonsterFriendly] tinyint NOT NULL\n\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
