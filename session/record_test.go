package session

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/KinoHui/project-jorm/dialect"
	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func NewSession() *Session {
	return New(TestDB, TestDial)
}

type User struct {
	Name string `jorm:"PRIMARY KEY"`
	Age  int
}

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Sam", 25}
	user3 = &User{"Jack", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("failed init test records")
	}
	return s
}

func TestMain(m *testing.M) {
	db, err := sql.Open("sqlite3", "../jin.db")
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	TestDB = db
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInit(t)
	affected, err := s.Insert(user3)
	if err != nil || affected != 1 {
		t.Fatal("failed to create record")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); err != nil || len(users) != 2 {
		t.Fatal("failed to query all")
	}
}
