package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gitlab.com/xsysproject/ppt_backend/config"
	"gitlab.com/xsysproject/ppt_backend/internal/role"
)

func main() {
	config, err := config.LoadConfig("./.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	GinServer(config, conn)
}

func GinServer(config config.Config, db *sql.DB) {
	router := gin.Default()

	roleRepo := role.NewRepository(db)
	roleUseCase := role.NewUseCase(roleRepo)
	role.NewHandler(router, roleUseCase)

	router.Run(config.HTTPServerAddress)
}
