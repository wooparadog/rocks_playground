//
// main.go
// Copyright (C) 2016 WooParadog <guohaochuan@gmail.com>
//
// Distributed under terms of the MIT license.
//

package main

import (
	"fmt"
	"log"

	"github.com/bsm/redeo"
	"github.com/tecbot/gorocksdb"
)

var db *gorocksdb.DB
var writeOptions *gorocksdb.WriteOptions
var readOptions *gorocksdb.ReadOptions

func main() {
	srv := redeo.NewServer(&redeo.Config{Addr: "localhost:9736"})
	srv.HandleFunc("get", GetHandler)
	srv.HandleFunc("set", SetHandler)

	log.Printf("Listening on tcp://%s", srv.Addr())
	log.Fatal(srv.ListenAndServe())
}

func SetHandler(out *redeo.Responder, req *redeo.Request) error {
	fmt.Println("%v", req.Args)
	return db.Put(writeOptions, []byte(req.Args[0]), []byte(req.Args[1]))
}

func GetHandler(out *redeo.Responder, req *redeo.Request) error {
	value, err := db.Get(readOptions, []byte(req.Args[0]))
	if err != nil {
		return err
	}
	out.WriteInlineString(string(value.Data()))
	return nil
}

func init() {
	fmt.Println("Initing DB")
	var err error
	writeOptions = gorocksdb.NewDefaultWriteOptions()
	readOptions = gorocksdb.NewDefaultReadOptions()

	options := gorocksdb.NewDefaultOptions()
	options.SetCreateIfMissing(true)
	db, err = gorocksdb.OpenDb(options, "this.db")
	if err != nil {
		panic(err)
	}
}
