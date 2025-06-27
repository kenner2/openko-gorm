package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_StartPositionDatabaseNbr = "GAME"
	_StartPositionTableName   = "START_POSITION"
)

func init() {
	ModelList = append(ModelList, &StartPosition{})
}

// StartPosition Start position
type StartPosition struct {
	ZoneId     int16 `gorm:"column:ZoneID;type:smallint;primaryKey;not null" json:"ZoneID"`
	KarusX     int16 `gorm:"column:sKarusX;type:smallint;not null" json:"sKarusX"`
	KarusZ     int16 `gorm:"column:sKarusZ;type:smallint;not null" json:"sKarusZ"`
	ElmoX      int16 `gorm:"column:sElmoradX;type:smallint;not null" json:"sElmoradX"`
	ElmoZ      int16 `gorm:"column:sElmoradZ;type:smallint;not null" json:"sElmoradZ"`
	RangeX     uint8 `gorm:"column:bRangeX;type:tinyint;not null" json:"bRangeX"`
	RangeZ     uint8 `gorm:"column:bRangeZ;type:tinyint;not null" json:"bRangeZ"`
	KarusGateX int16 `gorm:"column:sKarusGateX;type:smallint;not null" json:"sKarusGateX"`
	KarusGateZ int16 `gorm:"column:sKarusGateZ;type:smallint;not null" json:"sKarusGateZ"`
	ElmoGateX  int16 `gorm:"column:sElmoGateX;type:smallint;not null" json:"sElmoGateX"`
	ElmoGateZ  int16 `gorm:"column:sElmoGateZ;type:smallint;not null" json:"sElmoGateZ"`
}

// GetDatabaseName Returns the table's database name
func (this StartPosition) GetDatabaseName() string {
	return GetDatabaseName(_StartPositionDatabaseNbr)
}

// TableName Returns the table name
func (this StartPosition) TableName() string {
	return _StartPositionTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this StartPosition) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [START_POSITION] ([ZoneID], [sKarusX], [sKarusZ], [sElmoradX], [sElmoradZ], [bRangeX], [bRangeZ], [sKarusGateX], [sKarusGateZ], [sElmoGateX], [sElmoGateZ]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneId),
		GetOptionalDecVal(&this.KarusX),
		GetOptionalDecVal(&this.KarusZ),
		GetOptionalDecVal(&this.ElmoX),
		GetOptionalDecVal(&this.ElmoZ),
		GetOptionalDecVal(&this.RangeX),
		GetOptionalDecVal(&this.RangeZ),
		GetOptionalDecVal(&this.KarusGateX),
		GetOptionalDecVal(&this.KarusGateZ),
		GetOptionalDecVal(&this.ElmoGateX),
		GetOptionalDecVal(&this.ElmoGateZ))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this StartPosition) GetInsertHeader() string {
	return "INSERT INTO [START_POSITION] ([ZoneID], [sKarusX], [sKarusZ], [sElmoradX], [sElmoradZ], [bRangeX], [bRangeZ], [sKarusGateX], [sKarusGateZ], [sElmoGateX], [sElmoGateZ]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this StartPosition) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ZoneId),
		GetOptionalDecVal(&this.KarusX),
		GetOptionalDecVal(&this.KarusZ),
		GetOptionalDecVal(&this.ElmoX),
		GetOptionalDecVal(&this.ElmoZ),
		GetOptionalDecVal(&this.RangeX),
		GetOptionalDecVal(&this.RangeZ),
		GetOptionalDecVal(&this.KarusGateX),
		GetOptionalDecVal(&this.KarusGateZ),
		GetOptionalDecVal(&this.ElmoGateX),
		GetOptionalDecVal(&this.ElmoGateZ))
}

// GetCreateTableString Returns the create table statement for this object
func (this StartPosition) GetCreateTableString() string {
	query := "CREATE TABLE [START_POSITION] (\n\t[ZoneID] smallint NOT NULL,\n\t[sKarusX] smallint NOT NULL,\n\t[sKarusZ] smallint NOT NULL,\n\t[sElmoradX] smallint NOT NULL,\n\t[sElmoradZ] smallint NOT NULL,\n\t[bRangeX] tinyint NOT NULL,\n\t[bRangeZ] tinyint NOT NULL,\n\t[sKarusGateX] smallint NOT NULL,\n\t[sKarusGateZ] smallint NOT NULL,\n\t[sElmoGateX] smallint NOT NULL,\n\t[sElmoGateZ] smallint NOT NULL\n\tCONSTRAINT [PK_START_POSITION] PRIMARY KEY CLUSTERED ([ZoneID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this StartPosition) SelectClause() (selectClause clause.Select) {
	return _StartPositionSelectClause
}

// GetAllTableData Returns a list of all table data
func (this StartPosition) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []StartPosition{}
	rawSql := "SELECT [ZoneID], [sKarusX], [sKarusZ], [sElmoradX], [sElmoradZ], [bRangeX], [bRangeZ], [sKarusGateX], [sKarusGateZ], [sElmoGateX], [sElmoGateZ] FROM [START_POSITION]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _StartPositionSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[ZoneID]",
		},
		clause.Column{
			Name: "[sKarusX]",
		},
		clause.Column{
			Name: "[sKarusZ]",
		},
		clause.Column{
			Name: "[sElmoradX]",
		},
		clause.Column{
			Name: "[sElmoradZ]",
		},
		clause.Column{
			Name: "[bRangeX]",
		},
		clause.Column{
			Name: "[bRangeZ]",
		},
		clause.Column{
			Name: "[sKarusGateX]",
		},
		clause.Column{
			Name: "[sKarusGateZ]",
		},
		clause.Column{
			Name: "[sElmoGateX]",
		},
		clause.Column{
			Name: "[sElmoGateZ]",
		},
	},
}
