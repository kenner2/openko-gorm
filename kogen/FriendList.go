package kogen

import (
	"fmt"
)

const (
	_FriendListDatabaseNbr = 0
	_FriendListTableName   = "FRIEND_LIST"
)

func init() {
	ModelList = append(ModelList, &FriendList{})
}

// FriendList User friend list
type FriendList struct {
	UserId   [21]byte `gorm:"column:strUserID;type:varchar(21);primaryKey;not null" json:"strUserID"`
	Friend1  [21]byte `gorm:"column:strFriend1;type:varchar(21)" json:"strFriend1,omitempty"`
	Friend2  [21]byte `gorm:"column:strFriend2;type:varchar(21)" json:"strFriend2,omitempty"`
	Friend3  [21]byte `gorm:"column:strFriend3;type:varchar(21)" json:"strFriend3,omitempty"`
	Friend4  [21]byte `gorm:"column:strFriend4;type:varchar(21)" json:"strFriend4,omitempty"`
	Friend5  [21]byte `gorm:"column:strFriend5;type:varchar(21)" json:"strFriend5,omitempty"`
	Friend6  [21]byte `gorm:"column:strFriend6;type:varchar(21)" json:"strFriend6,omitempty"`
	Friend7  [21]byte `gorm:"column:strFriend7;type:varchar(21)" json:"strFriend7,omitempty"`
	Friend8  [21]byte `gorm:"column:strFriend8;type:varchar(21)" json:"strFriend8,omitempty"`
	Friend9  [21]byte `gorm:"column:strFriend9;type:varchar(21)" json:"strFriend9,omitempty"`
	Friend10 [21]byte `gorm:"column:strFriend10;type:varchar(21)" json:"strFriend10,omitempty"`
	Friend11 [21]byte `gorm:"column:strFriend11;type:varchar(21)" json:"strFriend11,omitempty"`
	Friend12 [21]byte `gorm:"column:strFriend12;type:varchar(21)" json:"strFriend12,omitempty"`
	Friend13 [21]byte `gorm:"column:strFriend13;type:varchar(21)" json:"strFriend13,omitempty"`
	Friend14 [21]byte `gorm:"column:strFriend14;type:varchar(21)" json:"strFriend14,omitempty"`
	Friend15 [21]byte `gorm:"column:strFriend15;type:varchar(21)" json:"strFriend15,omitempty"`
	Friend16 [21]byte `gorm:"column:strFriend16;type:varchar(21)" json:"strFriend16,omitempty"`
	Friend17 [21]byte `gorm:"column:strFriend17;type:varchar(21)" json:"strFriend17,omitempty"`
	Friend18 [21]byte `gorm:"column:strFriend18;type:varchar(21)" json:"strFriend18,omitempty"`
	Friend19 [21]byte `gorm:"column:strFriend19;type:varchar(21)" json:"strFriend19,omitempty"`
	Friend20 [21]byte `gorm:"column:strFriend20;type:varchar(21)" json:"strFriend20,omitempty"`
	Friend21 [21]byte `gorm:"column:strFriend21;type:varchar(21)" json:"strFriend21,omitempty"`
	Friend22 [21]byte `gorm:"column:strFriend22;type:varchar(21)" json:"strFriend22,omitempty"`
	Friend23 [21]byte `gorm:"column:strFriend23;type:varchar(21)" json:"strFriend23,omitempty"`
	Friend24 [21]byte `gorm:"column:strFriend24;type:varchar(21)" json:"strFriend24,omitempty"`
}

/* Helper Functions */

// GetDatabaseName Returns the table's database name
func (this *FriendList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_FriendListDatabaseNbr))
}

// GetTableName Returns the table name
func (this *FriendList) GetTableName() string {
	return _FriendListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this *FriendList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [FRIEND_LIST] (strUserID, strFriend1, strFriend2, strFriend3, strFriend4, strFriend5, strFriend6, strFriend7, strFriend8, strFriend9, strFriend10, strFriend11, strFriend12, strFriend13, strFriend14, strFriend15, strFriend16, strFriend17, strFriend18, strFriend19, strFriend20, strFriend21, strFriend22, strFriend23, strFriend24) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
		GetOptionalBinaryVal(this.Friend1),
		GetOptionalBinaryVal(this.Friend2),
		GetOptionalBinaryVal(this.Friend3),
		GetOptionalBinaryVal(this.Friend4),
		GetOptionalBinaryVal(this.Friend5),
		GetOptionalBinaryVal(this.Friend6),
		GetOptionalBinaryVal(this.Friend7),
		GetOptionalBinaryVal(this.Friend8),
		GetOptionalBinaryVal(this.Friend9),
		GetOptionalBinaryVal(this.Friend10),
		GetOptionalBinaryVal(this.Friend11),
		GetOptionalBinaryVal(this.Friend12),
		GetOptionalBinaryVal(this.Friend13),
		GetOptionalBinaryVal(this.Friend14),
		GetOptionalBinaryVal(this.Friend15),
		GetOptionalBinaryVal(this.Friend16),
		GetOptionalBinaryVal(this.Friend17),
		GetOptionalBinaryVal(this.Friend18),
		GetOptionalBinaryVal(this.Friend19),
		GetOptionalBinaryVal(this.Friend20),
		GetOptionalBinaryVal(this.Friend21),
		GetOptionalBinaryVal(this.Friend22),
		GetOptionalBinaryVal(this.Friend23),
		GetOptionalBinaryVal(this.Friend24))
}

// GetCreateTableString Returns the create table statement for this object
func (this *FriendList) GetCreateTableString() string {
	query := "CREATE TABLE [FRIEND_LIST] (\n\t\"strUserID\" varchar(21) NOT NULL,\n\t\"strFriend1\" varchar(21),\n\t\"strFriend2\" varchar(21),\n\t\"strFriend3\" varchar(21),\n\t\"strFriend4\" varchar(21),\n\t\"strFriend5\" varchar(21),\n\t\"strFriend6\" varchar(21),\n\t\"strFriend7\" varchar(21),\n\t\"strFriend8\" varchar(21),\n\t\"strFriend9\" varchar(21),\n\t\"strFriend10\" varchar(21),\n\t\"strFriend11\" varchar(21),\n\t\"strFriend12\" varchar(21),\n\t\"strFriend13\" varchar(21),\n\t\"strFriend14\" varchar(21),\n\t\"strFriend15\" varchar(21),\n\t\"strFriend16\" varchar(21),\n\t\"strFriend17\" varchar(21),\n\t\"strFriend18\" varchar(21),\n\t\"strFriend19\" varchar(21),\n\t\"strFriend20\" varchar(21),\n\t\"strFriend21\" varchar(21),\n\t\"strFriend22\" varchar(21),\n\t\"strFriend23\" varchar(21),\n\t\"strFriend24\" varchar(21)\n\tPRIMARY KEY (\"strUserID\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
