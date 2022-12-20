package runner

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"time"
)

type randFile struct {
	Id      int     `json:"id"`
	TokenId int     `json:"token_id"`
	JsonUri []uint8 `json:"json_uri"`
	BoxUri  []uint8 `json:"box_uri"`
	Quality string  `json:"quality"`
}

var db *sql.DB

func InitDB() {
	var err error
	config := mysql.Config{
		User:                 "go",
		Passwd:               "KJ6Gm4rFxCCa8xHX",
		Addr:                 "42.192.183.235:3306",
		Net:                  "tcp",
		DBName:               "go",
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalln(err.Error())
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	err = db.Ping()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

type dbVerify struct {
	vId string
}

func getFieldValue(tableName, jsonUri string) string {
	sqlStr := fmt.Sprintf("SELECT id FROM %s WHERE json_uri = '%s'", tableName, jsonUri)
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
	}
	defer rows.Close()
	var vvId string
	for rows.Next() {
		var dV dbVerify
		err = rows.Scan(&dV.vId)
		if err != nil {
			log.Fatalln(err.Error())
		}
		vvId = dV.vId
	}
	return vvId

}

// 闭区间
func getFields(leftScope, rightScope int64, tableName string) ([]*big.Int, []string, []string) {
	sqlStr := fmt.Sprintf("SELECT token_id ,json_uri, box_uri FROM %s WHERE id >= %d and id <= %d order by id asc", tableName, leftScope, rightScope)
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare sql failed, err:%v\n", err)
		return nil, nil, nil
	}
	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return nil, nil, nil
	}
	defer rows.Close()
	tokenIds := make([]*big.Int, 0)
	jsonUris := make([]string, 0)
	boxUris := make([]string, 0)
	for rows.Next() {
		var u randFile
		err := rows.Scan(&u.TokenId, &u.JsonUri, &u.BoxUri)
		if err != nil {
			log.Println(err.Error())
		}
		tokenIds = append(tokenIds, big.NewInt(int64(u.TokenId)))
		jsonUris = append(jsonUris, string(bytes.Repeat(u.JsonUri, 1)))
		boxUris = append(boxUris, string(bytes.Repeat(u.BoxUri, 1)))
	}
	log.Println(tokenIds)
	log.Println(jsonUris)
	log.Println(boxUris)
	return tokenIds, jsonUris, boxUris
}

type zcData struct {
	Id      string `json:"id"`
	JsonUri string `json:"json_uri"`
}

var zcDatas []zcData

func getJsonData(path string) []zcData {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = json.Unmarshal(jsonData, &zcDatas)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return zcDatas
	//for _, v := range zcDatas {
	//	log.Printf("%v", v)
	//}

}

func updateIsMint(tableName string, value, id int64) int64 {
	sqld := fmt.Sprintf("update %s set `is_mint` = ? where `token_id` = ?", tableName)
	stmt, err := db.Exec(sqld, value, id)
	if err != nil {
		log.Fatalln(err.Error())
	}
	rows, err := stmt.RowsAffected()
	if err != nil {
		log.Fatalln(err.Error())
	}
	return rows
	//defer stmt.Close()
	//rs, err := stmt.Exec(value, id)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//if id, err = rs.LastInsertId(); id > 0 {
	//	return id
	//}
}
