package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type User struct {
	Id   int      `dynamo:"Id,hash"`
	Text []string `dynamo:"Text"`
}

func main() {
	// 初期化
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("ap-northeast-1"),
	})

	err := db.CreateTable("Users", User{}).Run()
	if err != nil {
		fmt.Println(err)
	}

	// ACTIVE待ち
	for {
		d, err := db.Table("Users").Describe().Run()
		if err != nil {
			fmt.Println(err)
		}

		if d.Status == dynamo.ActiveStatus {
			break
		}
	}
	table := db.Table("Users")

	// Put
	u1 := User{Id: 1, Text: []string{"aja"}}
	u2 := User{Id: 2, Text: []string{"aja", "ajaaja"}}
	u3 := User{Id: 3, Text: []string{"aja", "ajaaja", "ajaajaajaaja"}}
	u4 := User{Id: 4, Text: []string{"aja", "ajaaja", "ajaajaajaaja", "ajaajaajaajaaja"}}
	_ = table.Put(u1).Run()
	_ = table.Put(u2).Run()
	_ = table.Put(u3).Run()
	_ = table.Put(u4).Run()

	// GetBatchItem
	// 以下の例はパーティションキーのみで有効
	var r []User
	err2 := table.Batch("Id").Get([]dynamo.Keyed{dynamo.Keys{1}, dynamo.Keys{2}, dynamo.Keys{3}, dynamo.Keys{4}}...).All(&r)
	if err != nil {
		fmt.Println(err2)
	}

	fmt.Println(r)
}
