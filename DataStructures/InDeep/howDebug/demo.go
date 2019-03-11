/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : tao.lv
#   Email         : 547617251lrtz@gmail.com
#   File Name     : demo.go
#   Last Modified : 2018-05-10 13:26
#   Describe      :
#
# ====================================================*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync/atomic"
)

var visitors int64

var rxOptionalID = regexp.MustCompile(`^\d*$`)

func handleHi(w http.ResponseWriter, r *http.Request) {
	if !rxOptionalID.MatchString(r.FormValue("color")) {
		http.Error(w, "Optional color is invalid", http.StatusBadRequest)
		return
	}

	visitNum := atomic.AddInt64(&visitors, 1)
	fmt.Fprintf(w, "<html><h1 stype='color: \"%s\"'>Welcome!</h1>You are visitor number %d!", r.FormValue("color"), visitNum)
}

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
