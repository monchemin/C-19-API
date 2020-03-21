package main

import (
	"c19/connector/pgsql"
)

func main() {
	pg, err := pgsql.OpenWithConfig(pgsql.Config{
		HostName:     "localhost",
		UserName:     "nyemo",
		Password:     "",
		Port:         5432,
		DataBaseName: "clinickp",
	})
	if err != nil{
		panic(err)
	}
	defer pg.Close()
}
