package main

import (
	"context"
	"fmt"
	"util/pkg/dao"
)

func main() {
	DbDsn := "root:Zkfc@2024!@tcp(172.10.60.75:3306)/malan-qa-dfcp_drone?timeout=10s&charset=utf8&collation=utf8_bin" +
		"&parseTime=True&loc=Local"

	conf := &dao.Config{
		DSN: DbDsn,
	}

	ctx := context.Background()
	db, err := dao.NewMysql(conf)

	if err != nil {
		return
	}

	taskModel := dao.NewTaskModel(db.DB)
	taskList, err := taskModel.CommonQuery(ctx, 1, -1, taskModel.WithID(42), taskModel.WithName("10kV后召线951_自适应巡检_202404221708"))

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(taskList)

}
