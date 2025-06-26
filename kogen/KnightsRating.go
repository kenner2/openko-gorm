package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KnightsRatingDatabaseNbr = 1
	_KnightsRatingTableName   = "KNIGHTS_RATING"
)

func init() {
	ModelList = append(ModelList, &KnightsRating{})
}

// KnightsRating Knights Ratings
type KnightsRating struct {
	Rank   int            `gorm:"column:nRank;type:int;not null" json:"nRank"`
	Index  *int16         `gorm:"column:shIndex;type:smallint" json:"shIndex,omitempty"`
	Name   *mssql.VarChar `gorm:"column:strName;type:varchar(21)" json:"strName,omitempty"`
	Points *int           `gorm:"column:nPoints;type:int" json:"nPoints,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this KnightsRating) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KnightsRatingDatabaseNbr))
}

// TableName Returns the table name
func (this KnightsRating) TableName() string {
	return _KnightsRatingTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KnightsRating) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KNIGHTS_RATING] ([nRank], [shIndex], [strName], [nPoints]) VALUES\n(%s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(this.Index),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(this.Points))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KnightsRating) GetInsertHeader() string {
	return "INSERT INTO [KNIGHTS_RATING] (nRank, shIndex, strName, nPoints) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KnightsRating) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s)", GetOptionalDecVal(&this.Rank),
		GetOptionalDecVal(this.Index),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(this.Points))
}

// GetCreateTableString Returns the create table statement for this object
func (this KnightsRating) GetCreateTableString() string {
	query := "CREATE TABLE [KNIGHTS_RATING] (\n\t[nRank] int NOT NULL,\n\t[shIndex] smallint,\n\t[strName] varchar(21),\n\t[nPoints] int\n\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KnightsRating) SelectClause() (selectClause clause.Select) {
	return _KnightsRatingSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KnightsRating) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KnightsRating{}
	rawSql := "SELECT [nRank], [shIndex], [strName], [nPoints] FROM [KNIGHTS_RATING]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KnightsRatingSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[nRank]",
		},
		clause.Column{
			Name: "[shIndex]",
		},
		clause.Column{
			Name: "[strName]",
		},
		clause.Column{
			Name: "[nPoints]",
		},
	},
}
