package db_test

import (
	"chat/db"
	"testing"
)

type Test struct {
	Id    int    `db:"id"`
	Test1 string `db:"test1"`
}

func Test_Test(t *testing.T) {
	t.Error("不通过测试")
}

func Test_Select(t *testing.T) {
	id := 1
	row := db.New("test").QueryRow("select * from test where id = ?", id)
	err := row.Scan(&id)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test")
	}

}

func Test_Insert(t *testing.T) {
	test1 := "test"
	ret, err := db.New("test").Exec("insert into test (test1) values (?)", test1)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(ret)
	}
}
