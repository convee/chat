package mysql_test

import (
	"chat/mysql"
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
	var test []Test
	err := mysql.New("test").Select(&test, "select * from test")
	if err != nil {
		t.Error(err)
	} else {
		t.Log("test")
	}

}
