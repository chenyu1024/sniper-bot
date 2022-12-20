package runner

import (
	"bytes"
	"fmt"
	"log"
	"math/big"
)

func getBoxListByQuality(boxQuality, tableName string, amount int) ([]*big.Int, []string, []string) {
	sqlStr50 := fmt.Sprintf(
		"(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'C' and is_mint = 0 ORDER BY id ASC LIMIT 30) ")
	//" UNION " +
	//"(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'B' and is_mint = 0 ORDER BY id ASC LIMIT 25) ")
	stmt, err := db.Prepare(sqlStr50)
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
		err := rows.Scan(&u.Id, &u.TokenId, &u.JsonUri, &u.BoxUri, &u.Quality)
		if err != nil {
			log.Println(err.Error())
		}
		tokenIds = append(tokenIds, big.NewInt(int64(u.TokenId)))
		jsonUris = append(jsonUris, string(bytes.Repeat(u.JsonUri, 1)))
		boxUris = append(boxUris, string(bytes.Repeat(u.BoxUri, 1)))
	}
	//log.Println(tokenIds)
	//log.Println(jsonUris)
	//log.Println(boxUris)
	return tokenIds, jsonUris, boxUris
}

func GetTestBoxAndShoe(boxQuality string) ([]*big.Int, []string, []string) {
	sqlStr50 := fmt.Sprintf(
		"(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = '%s' ORDER BY id ASC LIMIT 10) ", boxQuality)
	stmt, err := db.Prepare(sqlStr50)
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
		err := rows.Scan(&u.Id, &u.TokenId, &u.JsonUri, &u.BoxUri, &u.Quality)
		if err != nil {
			log.Println(err.Error())
		}
		tokenIds = append(tokenIds, big.NewInt(int64(u.TokenId)))
		jsonUris = append(jsonUris, string(bytes.Repeat(u.JsonUri, 1)))
		boxUris = append(boxUris, string(bytes.Repeat(u.BoxUri, 1)))
	}
	//log.Println(tokenIds)
	//log.Println(jsonUris)
	//log.Println(boxUris)
	return tokenIds, jsonUris, boxUris
}
func GetFreeMintData() ([]*big.Int, []string, []string) {
	sqlStr50 := fmt.Sprintf("SELECT\nid,token_id,json_uri,box_uri,quality\nFROM\n(\n(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'C' and is_mint = 0 ORDER BY id ASC LIMIT 97)\nUNION\n(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'B' and is_mint = 0 ORDER BY id ASC LIMIT 66)\nUNION \n(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'A' and is_mint = 0 ORDER BY id ASC LIMIT 25)\nUNION \n(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'S' and is_mint = 0 ORDER BY id ASC LIMIT 10)\nUNION \n(SELECT id,token_id,json_uri,box_uri,quality FROM sneaker_info_main_copy1 WHERE quality = 'SR' and is_mint = 0 ORDER BY id ASC LIMIT 2)\n) as res\nORDER BY token_id ASC")
	stmt, err := db.Prepare(sqlStr50)
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
		err := rows.Scan(&u.Id, &u.TokenId, &u.JsonUri, &u.BoxUri, &u.Quality)
		if err != nil {
			log.Println(err.Error())
		}
		tokenIds = append(tokenIds, big.NewInt(int64(u.TokenId)))
		jsonUris = append(jsonUris, string(bytes.Repeat(u.JsonUri, 1)))
		boxUris = append(boxUris, string(bytes.Repeat(u.BoxUri, 1)))
	}
	//log.Println(tokenIds)
	//log.Println(jsonUris)
	//log.Println(boxUris)
	return tokenIds, boxUris, jsonUris
}
