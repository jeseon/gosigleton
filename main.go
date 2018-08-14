package main

import (
    "net/http"
    "time"
    "log"
    "os"
    "sync"
)

type singleton struct {
    t time.Time
}

var (
    logger *log.Logger
    instance *singleton
    once sync.Once
)

func GetInstance() *singleton {
    once.Do(func() {
        instance = &singleton{t: time.Now()}
    })
    return instance
}

func init() {
    logger = log.New(os.Stdout, "", 0)
    instance = GetInstance()
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        go PrintStuff(r)
    })
    http.ListenAndServe(":5000", nil)
}

func PrintStuff(_ *http.Request) {
    time.Sleep(10 * time.Second)

    logger.Println(instance.t)
}