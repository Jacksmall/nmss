package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Jacksmall/gospike/local"
	"github.com/Jacksmall/gospike/remote"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/sync/errgroup"
)

var (
	remoteSpike remote.RemoteSpikeKeys
	redisPool   *redis.Pool
	done        chan int
)

func init() {
	remoteSpike = remote.RemoteSpikeKeys{
		SpikeOrderHashKey:  "ticket_hash_key",
		TotalInventoryKey:  "ticket_total_nums",
		QuantityOfOrderKey: "ticket_sold_nums",
	}
	redisPool = remote.NewPool()
	done = make(chan int, 1)
	done <- 1
}

var (
	g errgroup.Group
)

func router3001() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery())
	localSpike := local.LocalSpike{
		LocalInStock:     100,
		LocalSalesVolume: 0,
	}

	e.GET("/buy/ticket", func(ctx *gin.Context) {
		redisConn := redisPool.Get()
		LogMsg := ""

		<-done

		if localSpike.LocalDeductStock() && remoteSpike.RemoteDeductionStock(redisConn) {
			LogMsg = LogMsg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "抢票成功",
				"data": nil,
			})
		} else {
			LogMsg = LogMsg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "已售罄",
				"data": nil,
			})
		}

		done <- 1

		writeLog(LogMsg, "./stat.log", "3001")
	})
	return e
}

func router3002() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery())
	localSpike := local.LocalSpike{
		LocalInStock:     100,
		LocalSalesVolume: 0,
	}

	e.GET("/buy/ticket", func(ctx *gin.Context) {
		redisConn := redisPool.Get()
		LogMsg := ""

		<-done

		if localSpike.LocalDeductStock() && remoteSpike.RemoteDeductionStock(redisConn) {
			LogMsg = LogMsg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "抢票成功",
				"data": nil,
			})
		} else {
			LogMsg = LogMsg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "已售罄",
				"data": nil,
			})
		}

		done <- 1

		writeLog(LogMsg, "./stat.log", "3002")
	})
	return e
}
func router3003() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery())
	localSpike := local.LocalSpike{
		LocalInStock:     100,
		LocalSalesVolume: 0,
	}

	e.GET("/buy/ticket", func(ctx *gin.Context) {
		redisConn := redisPool.Get()
		LogMsg := ""

		<-done

		if localSpike.LocalDeductStock() && remoteSpike.RemoteDeductionStock(redisConn) {
			LogMsg = LogMsg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "抢票成功",
				"data": nil,
			})
		} else {
			LogMsg = LogMsg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "已售罄",
				"data": nil,
			})
		}

		done <- 1

		writeLog(LogMsg, "./stat.log", "3003")
	})
	return e
}

func router3004() http.Handler {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery())
	localSpike := local.LocalSpike{
		LocalInStock:     100,
		LocalSalesVolume: 0,
	}

	e.GET("/buy/ticket", func(ctx *gin.Context) {
		redisConn := redisPool.Get()
		LogMsg := ""

		<-done

		if localSpike.LocalDeductStock() && remoteSpike.RemoteDeductionStock(redisConn) {
			LogMsg = LogMsg + "result:1,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "抢票成功",
				"data": nil,
			})
		} else {
			LogMsg = LogMsg + "result:0,localSales:" + strconv.FormatInt(localSpike.LocalSalesVolume, 10)
			ctx.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "已售罄",
				"data": nil,
			})
		}

		done <- 1

		writeLog(LogMsg, "./stat.log", "3004")
	})
	return e
}

func main() {
	server3001 := &http.Server{
		Addr:         ":3001",
		Handler:      router3001(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server3002 := &http.Server{
		Addr:         ":3002",
		Handler:      router3002(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server3003 := &http.Server{
		Addr:         ":3003",
		Handler:      router3003(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server3004 := &http.Server{
		Addr:         ":3004",
		Handler:      router3004(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server3001.ListenAndServe()
	})
	g.Go(func() error {
		return server3002.ListenAndServe()
	})
	g.Go(func() error {
		return server3003.ListenAndServe()
	})
	g.Go(func() error {
		return server3004.ListenAndServe()
	})
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}

func writeLog(msg string, logPath string, port string) {
	fd, _ := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	defer fd.Close()

	content := strings.Join([]string{msg, "\r\n"}, ", port:"+port)
	buf := []byte(content)
	fd.Write(buf)
}
