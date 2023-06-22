package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"gitlab.com/xsysproject/ppt_backend/config"
	"gitlab.com/xsysproject/ppt_backend/internal/encdec"
	"gitlab.com/xsysproject/ppt_backend/internal/role"
	"gitlab.com/xsysproject/ppt_backend/internal/service"
	"gitlab.com/xsysproject/ppt_backend/internal/user"
)

func main() {
	config, err := config.LoadConfig("./.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	pqconn, err := sql.Open(config.PSQLDBDriver, config.PSQLDBSource)
	if err != nil {
		log.Fatal("cannot connect to PostgreSQL database:", err)
	}

	rdconn := redis.NewClient(&redis.Options{
		Addr:     config.RedisDBAddress,
		Password: config.RedisDBPassword,
		DB:       config.RedisDBIndex,
	})

	rErr := rdconn.Ping(context.Background()).Err()
	if rErr != nil {
		log.Fatal("cannot connect to Redis database:", rErr)
	}

	escfg := elasticsearch.Config{
		Addresses: []string{config.ElasticDBAddress},
		Username:  config.ElasticDBUser,
		Password:  config.ElasticDBPassword,
	}

	esconn, esRrr := elasticsearch.NewClient(escfg)
	if esRrr != nil {
		log.Fatal("cannot connect to Elastic database:", esRrr)
	}

	// res, sserr := esconn.Info()
	// if sserr != nil {
	// 	log.Fatalf("Error getting response: %s %s", err)
	// }

	// fmt.Println(config, "____________", pqconn, "____________", rdconn, "____________", esconn, "____________", res)

	GinServer(config, pqconn, rdconn, esconn)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE")
		c.Header("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func GinServer(config config.Config, db *sql.DB, rdb *redis.Client, edb *elasticsearch.Client) {
	// gin.SetMode(gin.ReleaseMode) // production uncomment this
	router := gin.Default()
	router.Use(CORSMiddleware())

	userRepo := user.NewRepository(db, rdb, edb)
	userUseCase := user.NewUseCase(userRepo)
	user.NewHandler(router, userUseCase)

	roleRepo := role.NewRepository(db)
	roleUseCase := role.NewUseCase(roleRepo)
	role.NewHandler(router, roleUseCase)

	serviceRepo := service.NewRepository(db)
	serviceUseCase := service.NewUseCase(serviceRepo)
	service.NewHandler(router, serviceUseCase)

	encdec.NewHandler(router)

	router.Run(config.HTTPServerAddress)
}
