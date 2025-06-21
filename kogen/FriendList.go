package kogen

import (
	"fmt"
)

const (
	_FriendListTableName   = "FRIEND_LIST"
	_FriendListDatabaseNbr = 0
)

func init() {
	ModelList = append(ModelList, &FriendList{})
}

// FriendList: User friend list
type FriendList struct {
	UserId   [21]byte `gorm:"column:strUserId;type:varchar(21);primaryKey;not null" json:"strUserId"`
	Friend1  [21]byte `gorm:"column:strFriend1;type:varchar(21);not null" json:"strFriend1"`
	Friend2  [21]byte `gorm:"column:strFriend2;type:varchar(21);not null" json:"strFriend2"`
	Friend3  [21]byte `gorm:"column:strFriend3;type:varchar(21);not null" json:"strFriend3"`
	Friend4  [21]byte `gorm:"column:strFriend4;type:varchar(21);not null" json:"strFriend4"`
	Friend5  [21]byte `gorm:"column:strFriend5;type:varchar(21);not null" json:"strFriend5"`
	Friend6  [21]byte `gorm:"column:strFriend6;type:varchar(21);not null" json:"strFriend6"`
	Friend7  [21]byte `gorm:"column:strFriend7;type:varchar(21);not null" json:"strFriend7"`
	Friend8  [21]byte `gorm:"column:strFriend8;type:varchar(21);not null" json:"strFriend8"`
	Friend9  [21]byte `gorm:"column:strFriend9;type:varchar(21);not null" json:"strFriend9"`
	Friend10 [21]byte `gorm:"column:strFriend10;type:varchar(21);not null" json:"strFriend10"`
	Friend11 [21]byte `gorm:"column:strFriend11;type:varchar(21);not null" json:"strFriend11"`
	Friend12 [21]byte `gorm:"column:strFriend12;type:varchar(21);not null" json:"strFriend12"`
	Friend13 [21]byte `gorm:"column:strFriend13;type:varchar(21);not null" json:"strFriend13"`
	Friend14 [21]byte `gorm:"column:strFriend14;type:varchar(21);not null" json:"strFriend14"`
	Friend15 [21]byte `gorm:"column:strFriend15;type:varchar(21);not null" json:"strFriend15"`
	Friend16 [21]byte `gorm:"column:strFriend16;type:varchar(21);not null" json:"strFriend16"`
	Friend17 [21]byte `gorm:"column:strFriend17;type:varchar(21);not null" json:"strFriend17"`
	Friend18 [21]byte `gorm:"column:strFriend18;type:varchar(21);not null" json:"strFriend18"`
	Friend19 [21]byte `gorm:"column:strFriend19;type:varchar(21);not null" json:"strFriend19"`
	Friend20 [21]byte `gorm:"column:strFriend20;type:varchar(21);not null" json:"strFriend20"`
	Friend21 [21]byte `gorm:"column:strFriend21;type:varchar(21);not null" json:"strFriend21"`
	Friend22 [21]byte `gorm:"column:strFriend22;type:varchar(21);not null" json:"strFriend22"`
	Friend23 [21]byte `gorm:"column:strFriend23;type:varchar(21);not null" json:"strFriend23"`
	Friend24 [21]byte `gorm:"column:strFriend24;type:varchar(21);not null" json:"strFriend24"`
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
	return fmt.Sprintf("INSERT INTO [FRIEND_LIST] (strUserId, strFriend1, strFriend2, strFriend3, strFriend4, strFriend5, strFriend6, strFriend7, strFriend8, strFriend9, strFriend10, strFriend11, strFriend12, strFriend13, strFriend14, strFriend15, strFriend16, strFriend17, strFriend18, strFriend19, strFriend20, strFriend21, strFriend22, strFriend23, strFriend24) \nVALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalBinaryVal(this.UserId),
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
	query := "CREATE TABLE \"FRIEND_LIST\" (\n\t\"strUserId\" varchar(21) NOT NULL,\n\t\"strFriend1\" varchar(21) NOT NULL,\n\t\"strFriend2\" varchar(21) NOT NULL,\n\t\"strFriend3\" varchar(21) NOT NULL,\n\t\"strFriend4\" varchar(21) NOT NULL,\n\t\"strFriend5\" varchar(21) NOT NULL,\n\t\"strFriend6\" varchar(21) NOT NULL,\n\t\"strFriend7\" varchar(21) NOT NULL,\n\t\"strFriend8\" varchar(21) NOT NULL,\n\t\"strFriend9\" varchar(21) NOT NULL,\n\t\"strFriend10\" varchar(21) NOT NULL,\n\t\"strFriend11\" varchar(21) NOT NULL,\n\t\"strFriend12\" varchar(21) NOT NULL,\n\t\"strFriend13\" varchar(21) NOT NULL,\n\t\"strFriend14\" varchar(21) NOT NULL,\n\t\"strFriend15\" varchar(21) NOT NULL,\n\t\"strFriend16\" varchar(21) NOT NULL,\n\t\"strFriend17\" varchar(21) NOT NULL,\n\t\"strFriend18\" varchar(21) NOT NULL,\n\t\"strFriend19\" varchar(21) NOT NULL,\n\t\"strFriend20\" varchar(21) NOT NULL,\n\t\"strFriend21\" varchar(21) NOT NULL,\n\t\"strFriend22\" varchar(21) NOT NULL,\n\t\"strFriend23\" varchar(21) NOT NULL,\n\t\"strFriend24\" varchar(21) NOT NULL\n\tPRIMARY KEY (\"strUserId\")\n)"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}
