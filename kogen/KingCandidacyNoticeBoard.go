package kogen

import (
	"fmt"
)

const (
	_KingCandidacyNoticeBoardDatabaseNbr = 0
	_KingCandidacyNoticeBoardTableName   = "KING_CANDIDACY_NOTICE_BOARD"
)

func init() {
	ModelList = append(ModelList, &KingCandidacyNoticeBoard{})
}

// KingCandidacyNoticeBoard King candidacy notice board
type KingCandidacyNoticeBoard struct {
	CandidateId  [21]byte   `gorm:"column:strUserID;type:varchar(21);primaryKey;not null" json:"strUserID"`
	Nation       uint8      `gorm:"column:byNation;type:tinyint;not null" json:"byNation"`
	NoticeLength int16      `gorm:"column:sNoticeLen;type:smallint;not null" json:"sNoticeLen"`
	Notice       [1024]byte `gorm:"column:strNotice;type:varbinary(1024);not null" json:"strNotice"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *KingCandidacyNoticeBoard) GetDatabaseName() string {
	return GetDatabaseName(DbType(_KingCandidacyNoticeBoardDatabaseNbr))
}

// GetTableName Returns the table name
func (this *KingCandidacyNoticeBoard) GetTableName() string {
	return _KingCandidacyNoticeBoardTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *KingCandidacyNoticeBoard) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [KING_CANDIDACY_NOTICE_BOARD] (strUserID, byNation, sNoticeLen, strNotice) \nVALUES (%s, %s, %s, %s)", GetOptionalBinaryVal(this.CandidateId),
		GetOptionalDecVal(&this.Nation),
		GetOptionalDecVal(&this.NoticeLength),
		GetOptionalBinaryVal(this.Notice))
}

// GetCreateTableString Returns the create table statement for this object
func (this *KingCandidacyNoticeBoard) GetCreateTableString() string {
	query := "CREATE TABLE [KING_CANDIDACY_NOTICE_BOARD] (\n\t[strUserID] varchar(21) NOT NULL,\n\t[byNation] tinyint NOT NULL,\n\t[sNoticeLen] smallint NOT NULL,\n\t[strNotice] varbinary(1024) NOT NULL\n\tCONSTRAINT [PK_KING_CANDIDACY_NOTICE_BOARD] PRIMARY KEY ([strUserID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
