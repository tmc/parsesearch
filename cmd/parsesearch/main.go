package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sheki/parsesearch"
)

func main() {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	var (
		whkey = os.Getenv("PARSE_WEBHOOK_KEY")
		mkey  = os.Getenv("PARSE_MASTER_KEY")
		appid = os.Getenv("PARSE_APPLICATION_ID")
	)
	if whkey == "" || mkey == "" || appid == "" {
		log.Fatalln("Must provide PARSE_WEBHOOK_KEY, PARSE_MASTER_KEY, and PARSE_APPLICATION_ID")
	}
	i, err := parsesearch.NewIndexer(whkey, mkey, appid)
	if err != nil {
		log.Fatalln("error creating Indexer:", err)
	}
	err = i.RegisterHooks("/test/")
	if err != nil {
		log.Fatalln("error creating Indexer:", err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/search", i.Search)
	mux.HandleFunc("/index", i.Index)
	http.ListenAndServe(":"+port, mux)
}
