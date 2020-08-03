package sql2struct

import (
	"database/sql"
	"errors"
	"fmt"
)

type DBModel struct {
	DBEngine *sql.DB
	DBInfo   *DBInfo
}

type DBInfo struct {
	DBType   string
	Host     string
	Username string
	Password string
	Charset  string
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	columnComment string
}

func NewDBModel(info *DBInfo) *DBModel {
	return &DBModel{DBInfo: info}
}

//连接mysql数据库
func (m *DBModel) Connect() error {
	var err error
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=local",
		m.DBInfo.Username,
		m.DBInfo.Password,
		m.DBInfo.Host,
		m.DBInfo.Charset,
		)
	m.DBEngine, err = sql.Open(m.DBInfo.DBType, dsn)
	if err != nil {
		return err
	}
	return nil
}

//获取列表中的信息
func (m *DBModel) GetColumns(dbName, tableName string) ([] *TableColumn, error) {
	query := "SELECT " +
		"COLUMN_NAME, DATE_TYPE, COLUMN_KEY, IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
	rows, err := m.DBEngine.Query(query, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("无数据")
	}
	defer rows.Close()

	var columns []*TableColumn
	for rows.Next(){
		var column TableColumn
		err := rows.Scan(&column.ColumnName, &column.DataType,
			&column.ColumnKey, &column.IsNullable, &column.ColumnType,
			&column.columnComment)
		if err != nil {
			return nil, err
		}
		columns = append(columns, &column)
	}
	return columns, nil
}

//表字段类型映射
var DBTypeToStructType = map[string] string {
	"int": "int32",
	"tinyint": "int8",
	"smallint": "int",
	"mediumint": "int64",
	"bigint": "int64",
	"bit": "int",
	"bool": "bool",
	"enum": "string",
	"set": "string",
	"varchar": "string",
}
