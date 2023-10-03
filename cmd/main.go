package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gigaflex-co/ppt_backend/app/commodity"
	"github.com/gigaflex-co/ppt_backend/app/encdec"
	"github.com/gigaflex-co/ppt_backend/app/land_status"
	"github.com/gigaflex-co/ppt_backend/app/menu"
	"github.com/gigaflex-co/ppt_backend/app/report_category"
	"github.com/gigaflex-co/ppt_backend/app/role"
	"github.com/gigaflex-co/ppt_backend/app/service"
	"github.com/gigaflex-co/ppt_backend/app/user"
	"github.com/gigaflex-co/ppt_backend/app/ws"
	"github.com/gigaflex-co/ppt_backend/config"
	"github.com/gigaflex-co/ppt_backend/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	config, err := config.LoadConfig("./.")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	psql, err := sql.Open(config.PSQLDBDriver, config.PSQLDBSource)
	if err != nil {
		log.Fatal("cannot connect to PostgreSQL database:", err)
	}
	defer psql.Close()

	// ds := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.MYSQLDBUsername, config.MYSQLDBPassword, config.MYSQLDBHost, config.MYSQLDBPort, config.MYSQLDBName)
	// mysql, err := sql.Open("mysql", ds)
	// if err != nil {
	// 	log.Fatal("cannot connect to MYSQL database:", err.Error())
	// 	return
	// }
	// defer mysql.Close()

	// err = mysql.Ping()
	// if err != nil {
	// 	log.Fatal("cannot connect to MYSQL database:", err.Error())
	// 	return
	// }

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.RedisDBAddress,
		Password: config.RedisDBPassword,
		DB:       config.RedisDBIndex,
	})

	err = rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatal("cannot connect to Redis database:", err)
		return
	}

	err = rdb.Set(context.Background(), "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Error setting value in Redis:", err)
		return
	}

	_, err = rdb.Get(context.Background(), "key").Result()
	if err != nil {
		fmt.Println("Error getting value from Redis:", err)
		return
	}

	escfg := elasticsearch.Config{
		Addresses: []string{config.ElasticDBAddress},
		Username:  config.ElasticDBUser,
		Password:  config.ElasticDBPassword,
	}

	es, err := elasticsearch.NewClient(escfg)
	if err != nil {
		log.Fatal("cannot connect to Elastic database:", err)
		return
	}

	util.NewKafkaClient(config)
	go GinServer(config, psql, rdb, es)
	SocketServer(config)
}

func CORSMiddleware(config config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		allowHeaders := "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"
		allowMethods := "POST, GET, PUT, DELETE, OPTIONS"

		c.Header("Access-Control-Allow-Origin", config.AllowOrigin)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", allowHeaders)
		c.Header("Access-Control-Allow-Methods", allowMethods)
		c.Header("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	}
}

func GinServer(config config.Config, db *sql.DB, rdb *redis.Client, edb *elasticsearch.Client) {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware(config))

	userRepo := user.NewRepository(db, rdb, edb)
	userUsecase := user.NewUsecase(userRepo, rdb, config)
	user.NewHandler(router, userUsecase, rdb)

	roleRepo := role.NewRepository(db)
	roleUsecase := role.NewUsecase(roleRepo)
	role.NewHandler(router, roleUsecase, rdb)

	serviceRepo := service.NewRepository(db)
	serviceUsecase := service.NewUsecase(serviceRepo)
	service.NewHandler(router, serviceUsecase, rdb)

	menuRepo := menu.NewRepository(db)
	menuUsecase := menu.NewUsecase(menuRepo)
	menu.NewHandler(router, menuUsecase, rdb)

	landStatusRepo := land_status.NewRepository(db)
	landStatusUsecase := land_status.NewUsecase(landStatusRepo)
	land_status.NewHandler(router, landStatusUsecase, rdb)

	commodityRepo := commodity.NewRepository(db)
	commodityUsecase := commodity.NewUsecase(commodityRepo)
	commodity.NewHandler(router, commodityUsecase, rdb)

	reportCategoryRepo := report_category.NewRepository(db)
	reportCategoryUsecase := report_category.NewUsecase(reportCategoryRepo)
	report_category.NewHandler(router, reportCategoryUsecase, rdb)

	encdec.NewHandler(router, rdb)

	router.Run(config.HTTPServerAddress)
}

func SocketServer(config config.Config) {
	wsRouter := gin.Default()
	wsRouter.Use(CORSMiddleware(config))

	ws.NewHandler(wsRouter)
	wsRouter.Run(config.WebSocketServerAddress)
}
