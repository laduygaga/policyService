package app


import (
	"log"
	"strconv"
	"go.mongodb.org/mongo-driver/bson"
)

func ToDoc(v interface{}) (doc bson.M) {
    data, err := bson.Marshal(v)
    if err != nil {
        return
    }
    err = bson.Unmarshal(data, &doc)
    return
}

func ToInt(s string) int {
    i , err := strconv.ParseInt(s, 10, 32)
    if err != nil {
        log.Fatal(err)
    }
    return int(i)
}

func ToFloat(s string) float64{
    i , err := strconv.ParseFloat(s, 64)
    if err != nil {
        log.Fatal(err)
    }
    return i
}
