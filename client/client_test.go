package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

const testHost = "http://localhost"
const baseURL = testHost + "/simple-tree"

var testDb *sql.DB
var fixture []ItemList = nil
var testServer, testDsn string

func setupTest(tb testing.TB) {
	testServer = os.Getenv("TEST_SERVER")
	testDsn = os.Getenv("TEST_DSN")
	genFixture()
	err := fixtureDatabase()
	assert.Nil(tb, err)
}

func genFixture() {
	u1 := uint32(1)
	u2 := uint32(2)
	u3 := uint32(3)
	u4 := uint32(4)
	u6 := uint32(6)
	u7 := uint32(7)
	u8 := uint32(8)
	recTime, _ := time.Parse(time.RFC3339, "2024-11-18T02:53:44Z")
	fixture = make([]ItemList, 50)
	for i := 1; i <= 50; i++ {
		fixture[i-1] = ItemList{
			Id:        uint32(i),
			Name:      fmt.Sprintf("name %d", i),
			CreatedAt: &recTime,
		}
	}
	setParent(2, 3, u1)
	setParent(4, 6, u2)
	setParent(7, 9, u3)
	setParent(10, 19, u4)
	setParent(20, 29, u7)
	setParent(30, 39, u6)
	setParent(40, 50, u8)
}

func fixtureDatabase() (err error) {
	if nil == testDb {
		testDb, err = sql.Open("mysql", testDsn)
		if err != nil {
			return
		}
	}
	rows := make([]string, 50)
	for i := 1; i <= 50; i++ {
		pid := "NULL"
		if fixture[i-1].ParentId != nil {
			pid = fmt.Sprintf("%d", *fixture[i-1].ParentId)
		}
		rows[i-1] = fmt.Sprintf(
			"(%d, %s, 'name %d', '2024-11-18 02:53:44')", i, pid, i,
		)
	}
	//goland:noinspection SqlNoDataSourceInspection,SqlWithoutWhere,SqlResolve
	_, err = testDb.Exec("TRUNCATE TABLE `items`")
	if err != nil {
		return
	}
	_, err = testDb.Exec(
		"INSERT INTO `items` (`id`, `parent_id`, `name`, `created_at`) VALUES" +
			strings.Join(rows, ","),
	)
	if err != nil {
		return
	}
	return
}

func setParent(from, to int, parent uint32) {
	for i := from - 1; i < to; i++ {
		fixture[i].ParentId = &parent
	}
}

func assertJsonEquals(tb testing.TB, expected interface{}, actual interface{}) {
	a, err := json.Marshal(actual)
	assert.Nil(tb, err)
	e, err := json.Marshal(expected)
	assert.Nil(tb, err)
	assert.JSONEq(tb, string(e), string(a))
}
