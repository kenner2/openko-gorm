package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ZoneInfoDatabaseNbr = "GAME"
	_ZoneInfoTableName   = "ZONE_INFO"
)

func init() {
	ModelList = append(ModelList, &ZoneInfo{})
}

// ZoneInfo Zone (map) information
type ZoneInfo struct {
	ServerId  uint8          `gorm:"column:ServerNo;type:tinyint;not null" json:"ServerNo"`
	ZoneId    int16          `gorm:"column:ZoneNo;type:smallint;primaryKey;not null" json:"ZoneNo"`
	Name      mssql.VarChar  `gorm:"column:strZoneName;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS;not null" json:"strZoneName"`
	InitX     int            `gorm:"column:InitX;type:int;not null" json:"InitX"`
	InitZ     int            `gorm:"column:InitZ;type:int;not null" json:"InitZ"`
	InitY     int            `gorm:"column:InitY;type:int;not null" json:"InitY"`
	Type      uint8          `gorm:"column:Type;type:tinyint;not null" json:"Type"`
	RoomEvent uint8          `gorm:"column:RoomEvent;type:tinyint;not null" json:"RoomEvent"`
	Bz        *mssql.VarChar `gorm:"column:bz;type:varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS" json:"bz,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this ZoneInfo) GetDatabaseName() string {
	return GetDatabaseName(_ZoneInfoDatabaseNbr)
}

// TableName Returns the table name
func (this ZoneInfo) TableName() string {
	return _ZoneInfoTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ZoneInfo) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ZONE_INFO] ([ServerNo], [ZoneNo], [strZoneName], [InitX], [InitZ], [InitY], [Type], [RoomEvent], [bz]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ZoneId),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.InitX),
		GetOptionalDecVal(&this.InitZ),
		GetOptionalDecVal(&this.InitY),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.RoomEvent),
		GetOptionalVarCharVal(this.Bz, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ZoneInfo) GetInsertHeader() string {
	return "INSERT INTO [ZONE_INFO] ([ServerNo], [ZoneNo], [strZoneName], [InitX], [InitZ], [InitY], [Type], [RoomEvent], [bz]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ZoneInfo) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.ServerId),
		GetOptionalDecVal(&this.ZoneId),
		GetOptionalVarCharVal(&this.Name, false),
		GetOptionalDecVal(&this.InitX),
		GetOptionalDecVal(&this.InitZ),
		GetOptionalDecVal(&this.InitY),
		GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.RoomEvent),
		GetOptionalVarCharVal(this.Bz, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this ZoneInfo) GetCreateTableString() string {
	query := "CREATE TABLE [ZONE_INFO] (\n\t[ServerNo] tinyint NOT NULL,\n\t[ZoneNo] smallint NOT NULL,\n\t[strZoneName] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[InitX] int NOT NULL,\n\t[InitZ] int NOT NULL,\n\t[InitY] int NOT NULL,\n\t[Type] tinyint NOT NULL,\n\t[RoomEvent] tinyint NOT NULL,\n\t[bz] varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS\n\tCONSTRAINT [PK_ZONE_INFO] PRIMARY KEY CLUSTERED ([ZoneNo])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ZoneInfo) SelectClause() (selectClause clause.Select) {
	return _ZoneInfoSelectClause
}

// GetAllTableData Returns a list of all table data
func (this ZoneInfo) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ZoneInfo{}
	rawSql := "SELECT [ServerNo], [ZoneNo], [strZoneName], [InitX], [InitZ], [InitY], [Type], [RoomEvent], [bz] FROM [ZONE_INFO]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ZoneInfoSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[ServerNo]",
		},
		clause.Column{
			Name: "[ZoneNo]",
		},
		clause.Column{
			Name: "[strZoneName]",
		},
		clause.Column{
			Name: "[InitX]",
		},
		clause.Column{
			Name: "[InitZ]",
		},
		clause.Column{
			Name: "[InitY]",
		},
		clause.Column{
			Name: "[Type]",
		},
		clause.Column{
			Name: "[RoomEvent]",
		},
		clause.Column{
			Name: "[bz]",
		},
	},
}
