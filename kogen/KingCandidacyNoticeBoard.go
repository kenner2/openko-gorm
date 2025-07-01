package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_KingCandidacyNoticeBoardDatabaseNbr = "GAME"
	_KingCandidacyNoticeBoardTableName   = "KING_CANDIDACY_NOTICE_BOARD"
)

func init() {
	ModelList = append(ModelList, &KingCandidacyNoticeBoard{})
}

// KingCandidacyNoticeBoard King candidacy notice board
type KingCandidacyNoticeBoard struct {
	CandidateId  mssql.VarChar `gorm:"column:strUserID;type:varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS;primaryKey;not null" json:"strUserID"`
	Nation       uint8         `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	NoticeLength int16         `gorm:"column:sNoticeLen;type:smallint;not null" json:"sNoticeLen"`
	Notice       []byte        `gorm:"column:strNotice;type:varbinary(1024);not null" json:"strNotice"`
}

// GetDatabaseName Returns the table's database name
func (this KingCandidacyNoticeBoard) GetDatabaseName() string {
	return GetDatabaseName(_KingCandidacyNoticeBoardDatabaseNbr)
}

// TableName Returns the table name
func (this KingCandidacyNoticeBoard) TableName() string {
	return _KingCandidacyNoticeBoardTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this KingCandidacyNoticeBoard) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_CANDIDACY_NOTICE_BOARD] ([strUserID], [byNation], [sNoticeLen], [strNotice]) VALUES\n(%s, %s, %s, CONVERT(varbinary(1024), %s))", GetOptionalVarCharVal(&this.CandidateId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.NoticeLength),
		GetOptionalBinaryVal(&this.Notice))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this KingCandidacyNoticeBoard) GetInsertHeader() string {
	return "INSERT INTO [KING_CANDIDACY_NOTICE_BOARD] ([strUserID], [byNation], [sNoticeLen], [strNotice]) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this KingCandidacyNoticeBoard) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, CONVERT(varbinary(1024), %s))", GetOptionalVarCharVal(&this.CandidateId, false),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.NoticeLength),
		GetOptionalBinaryVal(&this.Notice))
}

// GetCreateTableString Returns the create table statement for this object
func (this KingCandidacyNoticeBoard) GetCreateTableString() string {
	query := "CREATE TABLE [KING_CANDIDACY_NOTICE_BOARD] (\n\t[strUserID] varchar(21) COLLATE SQL_Latin1_General_CP1_CI_AS NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[sNoticeLen] smallint NOT NULL,\n\t[strNotice] varbinary(1024) NOT NULL\n\tCONSTRAINT [PK_KING_CANDIDACY_NOTICE_BOARD] PRIMARY KEY CLUSTERED ([strUserID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this KingCandidacyNoticeBoard) SelectClause() (selectClause clause.Select) {
	return _KingCandidacyNoticeBoardSelectClause
}

// GetAllTableData Returns a list of all table data
func (this KingCandidacyNoticeBoard) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []KingCandidacyNoticeBoard{}
	rawSql := "SELECT [strUserID], [byNation], [sNoticeLen], CONVERT(VARBINARY(1024), [strNotice]) as [strNotice] FROM [KING_CANDIDACY_NOTICE_BOARD]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _KingCandidacyNoticeBoardSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strUserID]",
		},
		clause.Column{
			Name: "[byNation]",
		},
		clause.Column{
			Name: "[sNoticeLen]",
		},
		clause.Column{
			Name: "CONVERT(VARBINARY(1024), [strNotice]) as [strNotice]",
		},
	},
}
