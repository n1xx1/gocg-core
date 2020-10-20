package main

import (
	"io/ioutil"
	"log"
	"ocgcore/database"
	"ocgcore/server"
	"path/filepath"
	"regexp"
)

func scriptReader() func(path string) []byte {
	cardScriptRegex := regexp.MustCompile(`c\d+\.lua`)

	return func(path string) []byte {
		if cardScriptRegex.MatchString(path) {
			path = filepath.Join("official", path)
		}

		contents, err := ioutil.ReadFile(filepath.Join("script", path))
		if err != nil {
			log.Println("script reader error: ", err)
			return nil
		}
		return contents
	}
}

func main() {
	db := database.NewDatabase()
	if err := db.Load("cards.cdb"); err != nil {
		log.Fatal(err)
	}
	if err := db.Load("release.cdb"); err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(server.Config{
		Address:      "0.0.0.0:8080",
		ScriptReader: scriptReader(),
		Database:     db,
	})

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
