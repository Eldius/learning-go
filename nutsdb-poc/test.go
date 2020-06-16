package main

import (
	"encoding/json"
	"github.com/xujiajun/nutsdb"
	"log"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	bucket001 := "game001"
	key := []byte("192.168.100.1")
	val := map[string]string{
		"opt0": "value0",
		"opt1": "value1",
		"opt2": "value2",
	}
	if err := db.Update(
		func(tx *nutsdb.Tx) error {
			b, err := json.Marshal(val)
			if err != nil {
				log.Println("error:", err)
				return err
			}
			if err := tx.Put(bucket001, key, b, 0); err != nil {
				return err
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			//key := []byte("name1")
			if e, err := tx.Get(bucket001, key); err != nil {
				return err
			} else {
				log.Println(string(e.Value)) // "val1-modify"
				var result map[string]string
				json.Unmarshal(e.Value, &result)
				log.Println(result) // "val1-modify"
			}
			return nil
		}); err != nil {
		log.Println(err)
	}

	if err := db.View(
		func(tx *nutsdb.Tx) error {
			prefix := []byte("192")
			// Constrain 100 entries returned
			entries, _ := tx.PrefixScan(bucket001, prefix, 25)

			for _, e := range entries {
				log.Println(string(e.Key), string(e.Value))
				var result map[string]string
				json.Unmarshal(e.Value, &result)
				log.Println(result) // "val1-modify"
			}
			return nil
		}); err != nil {
		log.Fatal(err)
	}

}
