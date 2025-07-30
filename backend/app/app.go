package app

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adriel-timoteo/olist-fullstack/backend/constant"
	"github.com/adriel-timoteo/olist-fullstack/backend/db"
	"github.com/adriel-timoteo/olist-fullstack/backend/handler"
	"github.com/adriel-timoteo/olist-fullstack/backend/middleware"
	"github.com/adriel-timoteo/olist-fullstack/backend/repository"
	"github.com/adriel-timoteo/olist-fullstack/backend/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func (a *App) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}

	a.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatalf("error connect DB: %s\n", err)
	}

	a.Router = gin.Default()
	a.Router.ContextWithFallback = true
	a.Router.Use(middleware.CORSMiddleware())
	a.Router.Use(middleware.ErrorMiddleware())

	a.initRoutes()
}

func (a *App) initRoutes() {
	trx := repository.NewTransactor(a.DB)

	ur := repository.NewUserRepo()
	uuc := usecase.NewUserUsecaseImpl(ur, trx)
	uh := handler.NewUserHandler(uuc)

	pr := repository.NewProductRepo()
	puc := usecase.NewProductUsecaseImpl(pr, trx)
	ph := handler.NewProductHandler(puc)

	cr := repository.NewCustomerRepo()
	cuc := usecase.NewCustomerUsecaseImpl(cr, trx)
	ch := handler.NewCustomerHandler(cuc)

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := a.Router.Group("/api/v1")
	{
		v1.POST("/register", uh.RegisterUser)
		v1.POST("/login", uh.LoginUser)
		products := v1.Group("/products", middleware.Authenticate())
		{
			products.GET("/top-categories", ph.GetTopCategories)
			products.GET("/trends/delivered", ph.GetDeliveredTrend)
			products.GET("/trends/status", ph.GetProductStatusSnapshot)
			products.GET("/delivery/ontime-rate", ph.GetOnTimeDeliveryRate)
		}
		customers := v1.Group("/customer", middleware.Authenticate())
		{
			customers.GET("/top-cities", ch.GetTopCities)
			customers.GET("/total", ch.GetTotalUniqueCustomers)
			customers.GET("/repeat-rate", ch.GetRepeatPurchaseRate)
		}

		// FUTURE DEVELOPMENT
		// orders := v1.Group("/orders", middleware.Authenticate())
		// {
		// 	orders.GET("/kpi/revenue", oh.GetTotalRevenue)
		// 	orders.GET("/kpi/aov", oh.GetAverageOrderValue)
		// 	orders.GET("/traffic/daily", oh.GetOrdersPerDay)
		// 	orders.GET("/revenue/monthly", oh.GetRevenueByMonth)
		// }

		// marketing := v1.Group("/marketing", middleware.Authenticate())
		// {
		// 	marketing.GET("/conversion-rate", mh.GetLeadsToCustomerConversionRate)
		// 	marketing.GET("/revenue-by-channel", mh.GetRevenueByChannel)
		// }

		// reviews := v1.Group("/reviews", middleware.Authenticate())
		// {
		// 	reviews.GET("/average-score", rh.GetAverageReviewScore)
		// 	reviews.GET("/score-distribution", rh.GetReviewScoreDistribution)
		// }
	}

}

func (a *App) Run() {
	a.Init()

	defer a.DB.Close()

	srv := &http.Server{
		Addr:    os.Getenv(constant.ServerPort),
		Handler: a.Router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown: ", err)
	}

	<-ctx.Done()
	log.Println("Timeout of 5 seconds")
	log.Println("Server exiting")
}
