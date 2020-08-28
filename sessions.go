package main

import (
	"encoding/base64"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"

	_ "github.com/mattn/go-sqlite3" //DBのパッケージだが、操作はGORMで行うため、importだけして使わない
)

// セッション情報
type TempSession struct {
	gorm.Model
	SessionID string `gorm:"unique;not null"`
	UserID    string `gorm:"not null"`
}

// DBマイグレート
func sessionStoreInit() {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3") //第一引数：使用するDBのデバイス。第二引数：ファイル名
	if err != nil {
		panic("データベース開ません(sessionStoreinit)")
	}
	db.AutoMigrate(&TempSession{}) //ファイルがなければ、生成を行う。すでにあればマイグレート。すでにあってマイグレートされていれば何も行わない
	defer db.Close()
}

// 指定したsessionIDのセッションがあるか確認する
func SessionExist(sessionID string) bool {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(dbGetOne)")
	}
	var session TempSession
	var count int

	db.Where(&TempSession{SessionID: sessionID}).Find(&session).Count(&count)
	if count == 0 {
		return false
	}
	db.Close()
	return true
}

// 指定したuserIDのセッションがあるか確認する
func SessionExistByUserID(userID string) bool {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(dbGetOne)")
	}
	var session TempSession
	var count int
	db.Where(&TempSession{UserID: userID}).Find(&session).Count(&count)
	if count == 0 {
		return false
	}
	db.Close()
	return true
}

//指定したsessionIDのオブジェクトが削除される
func sessionDelete(sessionID string) {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(dbDelete)")
	}
	var session TempSession
	db.Where(&TempSession{SessionID: sessionID}).Limit(1).Find(&session)
	db.Delete(&session)
	db.Close()
}

// sessionを作成。sessionIDとuserIDの組みを格納し、sessionIDを返す
func createSession(userID string) string {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(createUser)")
	}
	defer db.Close()

	sessionID := LongSecureRandomBase64()

	// Insert処理
	if err := db.Create(&TempSession{SessionID: sessionID, UserID: userID}).Error; err != nil {

		return ""
	}
	return sessionID

}

// 指定したsessionIDのuserIDを返す
func getuserIDBySessionID(sessionID string) string {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(getUser)")
	}
	defer db.Close()
	var session TempSession
	db.Where(&TempSession{SessionID: sessionID}).Find(&session)

	return session.UserID
}

// getUserById は、指定されたIDを持つユーザーを一つ返します。
// ユーザーが存在しない場合、空のレコードが返る?(GORMの仕様を要確認)
func getSessionIDByuserID(userID string) string {
	db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
	if err != nil {
		panic("データベース開ません(getUserById)")
	}
	defer db.Close()
	var session TempSession
	db.Where(&TempSession{UserID: userID}).Limit(1).Find(&session)
	return session.SessionID
}

//期限切れのセッション情報を削除
func sessionStoreUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := gorm.Open("sqlite3", "sessionStore.sqlite3")
		if err != nil {
			panic("データベース開ません(getUserById)")
		}
		var session TempSession
		now := time.Now()
		date := now.Add(-5 * time.Hour)
		db.Where("created_at <= ?", date).Delete(&session)
		c.Next()
	}

}

//session IDを生成するための関数群
func SecureRandom() string {
	return uuid.New().String()
}

func SecureRandomBase64() string {
	return base64.StdEncoding.EncodeToString(uuid.New().NodeID())
}

func LongSecureRandomBase64() string {
	return SecureRandomBase64() + SecureRandomBase64()
}

func MultipleSecureRandomBase64(n int) string {
	if n <= 1 {
		return SecureRandomBase64()
	}
	return SecureRandomBase64() + MultipleSecureRandomBase64(n-1)
}
