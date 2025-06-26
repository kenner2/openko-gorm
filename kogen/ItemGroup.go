package kogen

import (
	"fmt"
	mssql "github.com/microsoft/go-mssqldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	_ItemGroupDatabaseNbr = 1
	_ItemGroupTableName   = "ITEM_GROUP"
)

func init() {
	ModelList = append(ModelList, &ItemGroup{})
}

// ItemGroup TODO Doc; No Data in table
type ItemGroup struct {
	Group  int16          `gorm:"column:group;type:smallint;not null" json:"group"`
	Name   *mssql.VarChar `gorm:"column:name;type:varchar(30)" json:"name,omitempty"`
	Item1  int            `gorm:"column:item1;type:int;not null;default:0" json:"item1"`
	Item2  int            `gorm:"column:item2;type:int;not null;default:0" json:"item2"`
	Item3  int            `gorm:"column:item3;type:int;not null;default:0" json:"item3"`
	Item4  int            `gorm:"column:item4;type:int;not null;default:0" json:"item4"`
	Item5  int            `gorm:"column:item5;type:int;not null;default:0" json:"item5"`
	Item6  int            `gorm:"column:item6;type:int;not null;default:0" json:"item6"`
	Item7  int            `gorm:"column:item7;type:int;not null;default:0" json:"item7"`
	Item8  int            `gorm:"column:item8;type:int;not null;default:0" json:"item8"`
	Item9  int            `gorm:"column:item9;type:int;not null;default:0" json:"item9"`
	Item10 int            `gorm:"column:item10;type:int;not null;default:0" json:"item10"`
	Item11 int            `gorm:"column:item11;type:int;not null;default:0" json:"item11"`
	Item12 int            `gorm:"column:item12;type:int;not null;default:0" json:"item12"`
	Item13 int            `gorm:"column:item13;type:int;not null;default:0" json:"item13"`
	Item14 int            `gorm:"column:item14;type:int;not null;default:0" json:"item14"`
	Item15 int            `gorm:"column:item15;type:int;not null;default:0" json:"item15"`
	Item16 int            `gorm:"column:item16;type:int;not null;default:0" json:"item16"`
	Item17 int            `gorm:"column:item17;type:int;not null;default:0" json:"item17"`
	Item18 int            `gorm:"column:item18;type:int;not null;default:0" json:"item18"`
	Item19 int            `gorm:"column:item19;type:int;not null;default:0" json:"item19"`
	Item20 int            `gorm:"column:item20;type:int;not null;default:0" json:"item20"`
	Item21 int            `gorm:"column:item21;type:int;not null;default:0" json:"item21"`
	Item22 int            `gorm:"column:item22;type:int;not null;default:0" json:"item22"`
	Item23 int            `gorm:"column:item23;type:int;not null;default:0" json:"item23"`
	Item24 int            `gorm:"column:item24;type:int;not null;default:0" json:"item24"`
	Item25 int            `gorm:"column:item25;type:int;not null;default:0" json:"item25"`
	Item26 int            `gorm:"column:item26;type:int;not null;default:0" json:"item26"`
	Item27 int            `gorm:"column:item27;type:int;not null;default:0" json:"item27"`
	Item28 int            `gorm:"column:item28;type:int;not null;default:0" json:"item28"`
	Item29 int            `gorm:"column:item29;type:int;not null;default:0" json:"item29"`
	Item30 int            `gorm:"column:item30;type:int;not null;default:0" json:"item30"`
}

// GetDatabaseName Returns the table's database name
func (this ItemGroup) GetDatabaseName() string {
	return GetDatabaseName(DbType(_ItemGroupDatabaseNbr))
}

// TableName Returns the table name
func (this ItemGroup) TableName() string {
	return _ItemGroupTableName
}

// GetInsertString Returns the insert statement for the table populated with record from the object
func (this ItemGroup) GetInsertString() string {
	return fmt.Sprintf("INSERT INTO [ITEM_GROUP] ([group], [name], [item1], [item2], [item3], [item4], [item5], [item6], [item7], [item8], [item9], [item10], [item11], [item12], [item13], [item14], [item15], [item16], [item17], [item18], [item19], [item20], [item21], [item22], [item23], [item24], [item25], [item26], [item27], [item28], [item29], [item30]) VALUES\n(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Group),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.Item1),
		GetOptionalDecVal(&this.Item2),
		GetOptionalDecVal(&this.Item3),
		GetOptionalDecVal(&this.Item4),
		GetOptionalDecVal(&this.Item5),
		GetOptionalDecVal(&this.Item6),
		GetOptionalDecVal(&this.Item7),
		GetOptionalDecVal(&this.Item8),
		GetOptionalDecVal(&this.Item9),
		GetOptionalDecVal(&this.Item10),
		GetOptionalDecVal(&this.Item11),
		GetOptionalDecVal(&this.Item12),
		GetOptionalDecVal(&this.Item13),
		GetOptionalDecVal(&this.Item14),
		GetOptionalDecVal(&this.Item15),
		GetOptionalDecVal(&this.Item16),
		GetOptionalDecVal(&this.Item17),
		GetOptionalDecVal(&this.Item18),
		GetOptionalDecVal(&this.Item19),
		GetOptionalDecVal(&this.Item20),
		GetOptionalDecVal(&this.Item21),
		GetOptionalDecVal(&this.Item22),
		GetOptionalDecVal(&this.Item23),
		GetOptionalDecVal(&this.Item24),
		GetOptionalDecVal(&this.Item25),
		GetOptionalDecVal(&this.Item26),
		GetOptionalDecVal(&this.Item27),
		GetOptionalDecVal(&this.Item28),
		GetOptionalDecVal(&this.Item29),
		GetOptionalDecVal(&this.Item30))
}

// GetInsertHeader Returns the header for the table insert dump (insert into table (cols) values
func (this ItemGroup) GetInsertHeader() string {
	return "INSERT INTO [ITEM_GROUP] (group, name, item1, item2, item3, item4, item5, item6, item7, item8, item9, item10, item11, item12, item13, item14, item15, item16, item17, item18, item19, item20, item21, item22, item23, item24, item25, item26, item27, item28, item29, item30) VALUES\n"
}

// GetInsertData Returns the record data for the table insert dump
func (this ItemGroup) GetInsertData() string {
	return fmt.Sprintf("(%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)", GetOptionalDecVal(&this.Group),
		GetOptionalVarCharVal(this.Name, false),
		GetOptionalDecVal(&this.Item1),
		GetOptionalDecVal(&this.Item2),
		GetOptionalDecVal(&this.Item3),
		GetOptionalDecVal(&this.Item4),
		GetOptionalDecVal(&this.Item5),
		GetOptionalDecVal(&this.Item6),
		GetOptionalDecVal(&this.Item7),
		GetOptionalDecVal(&this.Item8),
		GetOptionalDecVal(&this.Item9),
		GetOptionalDecVal(&this.Item10),
		GetOptionalDecVal(&this.Item11),
		GetOptionalDecVal(&this.Item12),
		GetOptionalDecVal(&this.Item13),
		GetOptionalDecVal(&this.Item14),
		GetOptionalDecVal(&this.Item15),
		GetOptionalDecVal(&this.Item16),
		GetOptionalDecVal(&this.Item17),
		GetOptionalDecVal(&this.Item18),
		GetOptionalDecVal(&this.Item19),
		GetOptionalDecVal(&this.Item20),
		GetOptionalDecVal(&this.Item21),
		GetOptionalDecVal(&this.Item22),
		GetOptionalDecVal(&this.Item23),
		GetOptionalDecVal(&this.Item24),
		GetOptionalDecVal(&this.Item25),
		GetOptionalDecVal(&this.Item26),
		GetOptionalDecVal(&this.Item27),
		GetOptionalDecVal(&this.Item28),
		GetOptionalDecVal(&this.Item29),
		GetOptionalDecVal(&this.Item30))
}

// GetCreateTableString Returns the create table statement for this object
func (this ItemGroup) GetCreateTableString() string {
	query := "CREATE TABLE [ITEM_GROUP] (\n\t[group] smallint NOT NULL,\n\t[name] varchar(30),\n\t[item1] int NOT NULL,\n\t[item2] int NOT NULL,\n\t[item3] int NOT NULL,\n\t[item4] int NOT NULL,\n\t[item5] int NOT NULL,\n\t[item6] int NOT NULL,\n\t[item7] int NOT NULL,\n\t[item8] int NOT NULL,\n\t[item9] int NOT NULL,\n\t[item10] int NOT NULL,\n\t[item11] int NOT NULL,\n\t[item12] int NOT NULL,\n\t[item13] int NOT NULL,\n\t[item14] int NOT NULL,\n\t[item15] int NOT NULL,\n\t[item16] int NOT NULL,\n\t[item17] int NOT NULL,\n\t[item18] int NOT NULL,\n\t[item19] int NOT NULL,\n\t[item20] int NOT NULL,\n\t[item21] int NOT NULL,\n\t[item22] int NOT NULL,\n\t[item23] int NOT NULL,\n\t[item24] int NOT NULL,\n\t[item25] int NOT NULL,\n\t[item26] int NOT NULL,\n\t[item27] int NOT NULL,\n\t[item28] int NOT NULL,\n\t[item29] int NOT NULL,\n\t[item30] int NOT NULL\n\n)\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item1] DEFAULT 0 FOR [item1]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item2] DEFAULT 0 FOR [item2]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item3] DEFAULT 0 FOR [item3]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item4] DEFAULT 0 FOR [item4]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item5] DEFAULT 0 FOR [item5]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item6] DEFAULT 0 FOR [item6]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item7] DEFAULT 0 FOR [item7]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item8] DEFAULT 0 FOR [item8]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item9] DEFAULT 0 FOR [item9]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item10] DEFAULT 0 FOR [item10]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item11] DEFAULT 0 FOR [item11]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item12] DEFAULT 0 FOR [item12]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item13] DEFAULT 0 FOR [item13]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item14] DEFAULT 0 FOR [item14]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item15] DEFAULT 0 FOR [item15]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item16] DEFAULT 0 FOR [item16]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item17] DEFAULT 0 FOR [item17]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item18] DEFAULT 0 FOR [item18]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item19] DEFAULT 0 FOR [item19]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item20] DEFAULT 0 FOR [item20]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item21] DEFAULT 0 FOR [item21]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item22] DEFAULT 0 FOR [item22]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item23] DEFAULT 0 FOR [item23]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item24] DEFAULT 0 FOR [item24]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item25] DEFAULT 0 FOR [item25]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item26] DEFAULT 0 FOR [item26]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item27] DEFAULT 0 FOR [item27]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item28] DEFAULT 0 FOR [item28]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item29] DEFAULT 0 FOR [item29]\nGO\nALTER TABLE [ITEM_GROUP] ADD CONSTRAINT [DF_ITEM_GROUP_item30] DEFAULT 0 FOR [item30]\nGO\n"
	return fmt.Sprintf("USE [%[1]s]\nGO\n\n%[2]s", this.GetDatabaseName(), query)
}

// SelectClause Returns a safe select clause for the model
func (this ItemGroup) SelectClause() (selectClause clause.Select) {
	return _ItemGroupSelectClause
}

// GetAllTableData Returns a list of all table data
func (this ItemGroup) GetAllTableData(db *gorm.DB) (results []Model, err error) {
	res := []ItemGroup{}
	rawSql := "SELECT [group], [name], [item1], [item2], [item3], [item4], [item5], [item6], [item7], [item8], [item9], [item10], [item11], [item12], [item13], [item14], [item15], [item16], [item17], [item18], [item19], [item20], [item21], [item22], [item23], [item24], [item25], [item26], [item27], [item28], [item29], [item30] FROM [ITEM_GROUP]"
	err = db.Raw(rawSql).Find(&res).Error
	if err != nil {
		return nil, err
	}
	for i := range res {
		results = append(results, &res[i])
	}
	return results, nil
}

var _ItemGroupSelectClause = clause.Select{
	Columns: []clause.Column{
		clause.Column{
			Name: "[group]",
		},
		clause.Column{
			Name: "[name]",
		},
		clause.Column{
			Name: "[item1]",
		},
		clause.Column{
			Name: "[item2]",
		},
		clause.Column{
			Name: "[item3]",
		},
		clause.Column{
			Name: "[item4]",
		},
		clause.Column{
			Name: "[item5]",
		},
		clause.Column{
			Name: "[item6]",
		},
		clause.Column{
			Name: "[item7]",
		},
		clause.Column{
			Name: "[item8]",
		},
		clause.Column{
			Name: "[item9]",
		},
		clause.Column{
			Name: "[item10]",
		},
		clause.Column{
			Name: "[item11]",
		},
		clause.Column{
			Name: "[item12]",
		},
		clause.Column{
			Name: "[item13]",
		},
		clause.Column{
			Name: "[item14]",
		},
		clause.Column{
			Name: "[item15]",
		},
		clause.Column{
			Name: "[item16]",
		},
		clause.Column{
			Name: "[item17]",
		},
		clause.Column{
			Name: "[item18]",
		},
		clause.Column{
			Name: "[item19]",
		},
		clause.Column{
			Name: "[item20]",
		},
		clause.Column{
			Name: "[item21]",
		},
		clause.Column{
			Name: "[item22]",
		},
		clause.Column{
			Name: "[item23]",
		},
		clause.Column{
			Name: "[item24]",
		},
		clause.Column{
			Name: "[item25]",
		},
		clause.Column{
			Name: "[item26]",
		},
		clause.Column{
			Name: "[item27]",
		},
		clause.Column{
			Name: "[item28]",
		},
		clause.Column{
			Name: "[item29]",
		},
		clause.Column{
			Name: "[item30]",
		},
	},
}
