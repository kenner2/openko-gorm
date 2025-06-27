package kogen

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ItemUpProbabilityDatabaseNbr = "GAME"
	_ItemUpProbabilityTableName   = "ITEMUP_PROBABILITY"
)

func init() {
	ModelList = append(ModelList, &ItemUpProbability{})
}

// ItemUpProbability TODO: Doc usage
type ItemUpProbability struct {
	Type           uint8 `gorm:"column:bType;type:tinyint;primaryKey;not null" json:"bType"`
	MaxSuccess     int16 `gorm:"column:nMaxSuccess;type:smallint;not null;default:0" json:"nMaxSuccess"`
	MaxFail        int16 `gorm:"column:nMaxFail;type:smallint;not null;default:0" json:"nMaxFail"`
	CurrentSuccess int16 `gorm:"column:nCurSuccess;type:smallint;not null;default:0" json:"nCurSuccess"`
	CurrentFailure int16 `gorm:"column:nCurFail;type:smallint;not null;default:0" json:"nCurFail"`
}

// GetDatabaseName Returns the table's database name
func (this ItemUpProbability) GetDatabaseName() string {
	return GetDatabaseName(_ItemUpProbabilityDatabaseNbr)
}

// TableName Returns the table name
func (this ItemUpProbability) TableName() string {
	return _ItemUpProbabilityTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ItemUpProbability) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEMUP_PROBABILITY] ([bType], [nMaxSuccess], [nMaxFail], [nCurSuccess], [nCurFail]) VALUES\n(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.MaxSuccess),
		GetOptionalDecVal(&this.MaxFail),
		GetOptionalDecVal(&this.CurrentSuccess),
		GetOptionalDecVal(&this.CurrentFailure))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ItemUpProbability) GetInsertHeader() string {
	return "INSERT INTO [ITEMUP_PROBABILITY] ([bType], [nMaxSuccess], [nMaxFail], [nCurSuccess], [nCurFail]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ItemUpProbability) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Type),
		GetOptionalDecVal(&this.MaxSuccess),
		GetOptionalDecVal(&this.MaxFail),
		GetOptionalDecVal(&this.CurrentSuccess),
		GetOptionalDecVal(&this.CurrentFailure))
}

// GetCreateTableString Returns the create table statement for this object
func (this ItemUpProbability) GetCreateTableString() string {
	query := "CREATE TABLE [ITEMUP_PROBABILITY] (\n\t[bType] tinyint NOT NULL,\n\t[nMaxSuccess] smallint NOT NULL,\n\t[nMaxFail] smallint NOT NULL,\n\t[nCurSuccess] smallint NOT NULL,\n\t[nCurFail] smallint NOT NULL\n\tCONSTRAINT [PK_ITEMUP_PROBABILITY] PRIMARY KEY CLUSTERED ([bType])\n)\nGO\nALTER TABLE [ITEMUP_PROBABILITY] ADD CONSTRAINT [DF_ITEMUP_PROBABILITY_nMaxSuccess] DEFAULT 0 FOR [nMaxSuccess]\nGO\nALTER TABLE [ITEMUP_PROBABILITY] ADD CONSTRAINT [DF_ITEMUP_PROBABILITY_nMaxFail] DEFAULT 0 FOR [nMaxFail]\nGO\nALTER TABLE [ITEMUP_PROBABILITY] ADD CONSTRAINT [DF_ITEMUP_PROBABILITY_nCurSuccess] DEFAULT 0 FOR [nCurSuccess]\nGO\nALTER TABLE [ITEMUP_PROBABILITY] ADD CONSTRAINT [DF_ITEMUP_PROBABILITY_nCurFail] DEFAULT 0 FOR [nCurFail]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ItemUpProbability) SelectClause() (selectClause clause.Select) {
	return _ItemUpProbabilitySelectClause
}

// GetAllTableData Returns a list of all table data
func (this ItemUpProbability) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ItemUpProbability{}
	rawSql := "SELECT [bType], [nMaxSuccess], [nMaxFail], [nCurSuccess], [nCurFail] FROM [ITEMUP_PROBABILITY]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ItemUpProbabilitySelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[bType]",
		},
		clause.Column{
			Name: "[nMaxSuccess]",
		},
		clause.Column{
			Name: "[nMaxFail]",
		},
		clause.Column{
			Name: "[nCurSuccess]",
		},
		clause.Column{
			Name: "[nCurFail]",
		},
	},
}
