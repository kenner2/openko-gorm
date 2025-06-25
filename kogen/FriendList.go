package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	UserId   mssql.VarChar  `gorm:"column:strUserID;type:varchar(21);primaryKey;not null" json:"strUserID"`
	Friend1  *mssql.VarChar `gorm:"column:strFriend1;type:varchar(21)" json:"strFriend1,omitempty"`
	Friend2  *mssql.VarChar `gorm:"column:strFriend2;type:varchar(21)" json:"strFriend2,omitempty"`
	Friend3  *mssql.VarChar `gorm:"column:strFriend3;type:varchar(21)" json:"strFriend3,omitempty"`
	Friend4  *mssql.VarChar `gorm:"column:strFriend4;type:varchar(21)" json:"strFriend4,omitempty"`
	Friend5  *mssql.VarChar `gorm:"column:strFriend5;type:varchar(21)" json:"strFriend5,omitempty"`
	Friend6  *mssql.VarChar `gorm:"column:strFriend6;type:varchar(21)" json:"strFriend6,omitempty"`
	Friend7  *mssql.VarChar `gorm:"column:strFriend7;type:varchar(21)" json:"strFriend7,omitempty"`
	Friend8  *mssql.VarChar `gorm:"column:strFriend8;type:varchar(21)" json:"strFriend8,omitempty"`
	Friend9  *mssql.VarChar `gorm:"column:strFriend9;type:varchar(21)" json:"strFriend9,omitempty"`
	Friend10 *mssql.VarChar `gorm:"column:strFriend10;type:varchar(21)" json:"strFriend10,omitempty"`
	Friend11 *mssql.VarChar `gorm:"column:strFriend11;type:varchar(21)" json:"strFriend11,omitempty"`
	Friend12 *mssql.VarChar `gorm:"column:strFriend12;type:varchar(21)" json:"strFriend12,omitempty"`
	Friend13 *mssql.VarChar `gorm:"column:strFriend13;type:varchar(21)" json:"strFriend13,omitempty"`
	Friend14 *mssql.VarChar `gorm:"column:strFriend14;type:varchar(21)" json:"strFriend14,omitempty"`
	Friend15 *mssql.VarChar `gorm:"column:strFriend15;type:varchar(21)" json:"strFriend15,omitempty"`
	Friend16 *mssql.VarChar `gorm:"column:strFriend16;type:varchar(21)" json:"strFriend16,omitempty"`
	Friend17 *mssql.VarChar `gorm:"column:strFriend17;type:varchar(21)" json:"strFriend17,omitempty"`
	Friend18 *mssql.VarChar `gorm:"column:strFriend18;type:varchar(21)" json:"strFriend18,omitempty"`
	Friend19 *mssql.VarChar `gorm:"column:strFriend19;type:varchar(21)" json:"strFriend19,omitempty"`
	Friend20 *mssql.VarChar `gorm:"column:strFriend20;type:varchar(21)" json:"strFriend20,omitempty"`
	Friend21 *mssql.VarChar `gorm:"column:strFriend21;type:varchar(21)" json:"strFriend21,omitempty"`
	Friend22 *mssql.VarChar `gorm:"column:strFriend22;type:varchar(21)" json:"strFriend22,omitempty"`
	Friend23 *mssql.VarChar `gorm:"column:strFriend23;type:varchar(21)" json:"strFriend23,omitempty"`
	Friend24 *mssql.VarChar `gorm:"column:strFriend24;type:varchar(21)" json:"strFriend24,omitempty"`
}

// GetDatabaseName Returns the table's database name
func (this FriendList) GetDatabaseName() string {
	return GetDatabaseName(DbType(_FriendListDatabaseNbr))
}

// TableName Returns the table name
func (this FriendList) TableName() string {
	return _FriendListTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this FriendList) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [FRIEND_LIST] ([strUserID], [strFriend1], [strFriend2], [strFriend3], [strFriend4], [strFriend5], [strFriend6], [strFriend7], [strFriend8], [strFriend9], [strFriend10], [strFriend11], [strFriend12], [strFriend13], [strFriend14], [strFriend15], [strFriend16], [strFriend17], [strFriend18], [strFriend19], [strFriend20], [strFriend21], [strFriend22], [strFriend23], [strFriend24]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalVarCharVal(this.Friend1, false),
		GetOptionalVarCharVal(this.Friend2, false),
		GetOptionalVarCharVal(this.Friend3, false),
		GetOptionalVarCharVal(this.Friend4, false),
		GetOptionalVarCharVal(this.Friend5, false),
		GetOptionalVarCharVal(this.Friend6, false),
		GetOptionalVarCharVal(this.Friend7, false),
		GetOptionalVarCharVal(this.Friend8, false),
		GetOptionalVarCharVal(this.Friend9, false),
		GetOptionalVarCharVal(this.Friend10, false),
		GetOptionalVarCharVal(this.Friend11, false),
		GetOptionalVarCharVal(this.Friend12, false),
		GetOptionalVarCharVal(this.Friend13, false),
		GetOptionalVarCharVal(this.Friend14, false),
		GetOptionalVarCharVal(this.Friend15, false),
		GetOptionalVarCharVal(this.Friend16, false),
		GetOptionalVarCharVal(this.Friend17, false),
		GetOptionalVarCharVal(this.Friend18, false),
		GetOptionalVarCharVal(this.Friend19, false),
		GetOptionalVarCharVal(this.Friend20, false),
		GetOptionalVarCharVal(this.Friend21, false),
		GetOptionalVarCharVal(this.Friend22, false),
		GetOptionalVarCharVal(this.Friend23, false),
		GetOptionalVarCharVal(this.Friend24, false))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this FriendList) GetInsertHeader() string {
	return "INSERT INTO [FRIEND_LIST] (strUserID, strFriend1, strFriend2, strFriend3, strFriend4, strFriend5, strFriend6, strFriend7, strFriend8, strFriend9, strFriend10, strFriend11, strFriend12, strFriend13, strFriend14, strFriend15, strFriend16, strFriend17, strFriend18, strFriend19, strFriend20, strFriend21, strFriend22, strFriend23, strFriend24) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this FriendList) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalVarCharVal(&this.UserId, false),
		GetOptionalVarCharVal(this.Friend1, false),
		GetOptionalVarCharVal(this.Friend2, false),
		GetOptionalVarCharVal(this.Friend3, false),
		GetOptionalVarCharVal(this.Friend4, false),
		GetOptionalVarCharVal(this.Friend5, false),
		GetOptionalVarCharVal(this.Friend6, false),
		GetOptionalVarCharVal(this.Friend7, false),
		GetOptionalVarCharVal(this.Friend8, false),
		GetOptionalVarCharVal(this.Friend9, false),
		GetOptionalVarCharVal(this.Friend10, false),
		GetOptionalVarCharVal(this.Friend11, false),
		GetOptionalVarCharVal(this.Friend12, false),
		GetOptionalVarCharVal(this.Friend13, false),
		GetOptionalVarCharVal(this.Friend14, false),
		GetOptionalVarCharVal(this.Friend15, false),
		GetOptionalVarCharVal(this.Friend16, false),
		GetOptionalVarCharVal(this.Friend17, false),
		GetOptionalVarCharVal(this.Friend18, false),
		GetOptionalVarCharVal(this.Friend19, false),
		GetOptionalVarCharVal(this.Friend20, false),
		GetOptionalVarCharVal(this.Friend21, false),
		GetOptionalVarCharVal(this.Friend22, false),
		GetOptionalVarCharVal(this.Friend23, false),
		GetOptionalVarCharVal(this.Friend24, false))
}

// GetCreateTableString Returns the create table statement for this object
func (this FriendList) GetCreateTableString() string {
	query := "CREATE TABLE [FRIEND_LIST] (\n\t[strUserID] varchar(21) NOT NULL,\n\t[strFriend1] varchar(21),\n\t[strFriend2] varchar(21),\n\t[strFriend3] varchar(21),\n\t[strFriend4] varchar(21),\n\t[strFriend5] varchar(21),\n\t[strFriend6] varchar(21),\n\t[strFriend7] varchar(21),\n\t[strFriend8] varchar(21),\n\t[strFriend9] varchar(21),\n\t[strFriend10] varchar(21),\n\t[strFriend11] varchar(21),\n\t[strFriend12] varchar(21),\n\t[strFriend13] varchar(21),\n\t[strFriend14] varchar(21),\n\t[strFriend15] varchar(21),\n\t[strFriend16] varchar(21),\n\t[strFriend17] varchar(21),\n\t[strFriend18] varchar(21),\n\t[strFriend19] varchar(21),\n\t[strFriend20] varchar(21),\n\t[strFriend21] varchar(21),\n\t[strFriend22] varchar(21),\n\t[strFriend23] varchar(21),\n\t[strFriend24] varchar(21)\n\tCONSTRAINT [PK_FRIEND_LIST] PRIMARY KEY ([strUserID])\n)\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this FriendList) SelectClause() (selectClause clause.Select) {
	return _FriendListSelectClause
}

// GetAllTableData Returns a list of all table data
func (this FriendList) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []FriendList{}
	rawSql := "SELECT [strUserID], [strFriend1], [strFriend2], [strFriend3], [strFriend4], [strFriend5], [strFriend6], [strFriend7], [strFriend8], [strFriend9], [strFriend10], [strFriend11], [strFriend12], [strFriend13], [strFriend14], [strFriend15], [strFriend16], [strFriend17], [strFriend18], [strFriend19], [strFriend20], [strFriend21], [strFriend22], [strFriend23], [strFriend24] FROM [FRIEND_LIST]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _FriendListSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[strUserID]",
		},
		clause.Column{
			Name: "[strFriend1]",
		},
		clause.Column{
			Name: "[strFriend2]",
		},
		clause.Column{
			Name: "[strFriend3]",
		},
		clause.Column{
			Name: "[strFriend4]",
		},
		clause.Column{
			Name: "[strFriend5]",
		},
		clause.Column{
			Name: "[strFriend6]",
		},
		clause.Column{
			Name: "[strFriend7]",
		},
		clause.Column{
			Name: "[strFriend8]",
		},
		clause.Column{
			Name: "[strFriend9]",
		},
		clause.Column{
			Name: "[strFriend10]",
		},
		clause.Column{
			Name: "[strFriend11]",
		},
		clause.Column{
			Name: "[strFriend12]",
		},
		clause.Column{
			Name: "[strFriend13]",
		},
		clause.Column{
			Name: "[strFriend14]",
		},
		clause.Column{
			Name: "[strFriend15]",
		},
		clause.Column{
			Name: "[strFriend16]",
		},
		clause.Column{
			Name: "[strFriend17]",
		},
		clause.Column{
			Name: "[strFriend18]",
		},
		clause.Column{
			Name: "[strFriend19]",
		},
		clause.Column{
			Name: "[strFriend20]",
		},
		clause.Column{
			Name: "[strFriend21]",
		},
		clause.Column{
			Name: "[strFriend22]",
		},
		clause.Column{
			Name: "[strFriend23]",
		},
		clause.Column{
			Name: "[strFriend24]",
		},
	},
}
