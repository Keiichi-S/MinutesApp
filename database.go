package main

import(
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3" //DBのパッケージだが、操作はGORMで行うため、importだけして使わない
)


type Message struct{
  gorm.Model
  Message string
}
/*
テーブル名：messages -->　テーブル名は自動で複数形になる
カラム
  ・id
  ・created_at
  ・updated_at
  ・deleted_at
  ・Message (追加)
*/
/*外部からカラムを参照するときは
id → ID
created_at → CreatedAt
updated_at → UpdatedAt
deleted_at → DeletedAt
*/

//DBマイグレート
//main関数の最初でdbInit()を呼ぶことでデータベースマイグレート
func dbInit(){
  db, err := gorm.Open("sqlite3", "minutes.sqlite3") //第一引数：使用するDBのデバイス。第二引数：ファイル名
  if err != nil{
    panic("データベース開ません(dbinit)")
  }
  db.AutoMigrate(&Message{}) //ファイルがなければ、生成を行う。すでにあればマイグレート。すでにあってマイグレートされていれば何も行わない
  defer db.Close()
}

//DB追加
//追加したいメッセージは、dbInsert(message.Message)のような感じで呼べば追加される
func dbInsert(message string){
  db, err := gorm.Open("sqlite3", "minutes.sqlite3")
  if err != nil{
    panic("データベース開ません(dbInsert)")
  }
  db.Create(&Message{Message: message})
  defer db.Close()
}

//DB全取得
//dbGetAll()と呼ぶことで、データベース内の全てのMessageオブジェクトが返される
func dbGetAll() []Message{
  db, err := gorm.Open("sqlite3", "minutes.sqlite3")
  if err != nil{
    panic("データベース開ません(dbGetAll)")
  }
  var messages []Message
  db.Order("created_at desc").Find(&messages) //db.Find(&messages)で構造体Messageに対するテーブルの要素全てを取得し、それをOrder("created_at desc")で新しいものが上に来るように並び替えている
  db.Close()
  return messages
}

//DB一つ取得
//idを与えることで、該当するMessageオブジェクトが一つ返される
func dbGetOne(id int) Message{
  db, err := gorm.Open("sqlite3", "minutes.sqlite3")
  if err != nil{
    panic("データベース開ません(dbGetOne)")
  }
  var message Message
  db.First(&message, id)
  db.Close()
  return message
}

//DB更新
//idとmessageを与えることで、該当するidのMessageオブジェクトのMessageが更新される
func dbUpdate(id int, update_message string){
  db, err := gorm.Open("sqlite3", "minutes.sqlite3")
  if err != nil{
    panic("データベース開ません(dgUpdate)")
  }
  var message Message
  db.First(&message, id)
  message.Message = update_message
  db.Save(&message)
  db.Close()
}

//DB削除
//指定したidのMessageオブジェクトが削除される
func dbDelete(id int){
  db, err := gorm.Open("sqlite3", "minutes.sqlite3")
  if err != nil{
    panic("データベース開ません(dbDelete)")
  }
  var message Message
  db.First(&message, id)
  db.Delete(&message)
  db.Close()
}