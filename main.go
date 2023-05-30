package main

import (
	_driverFactory "Capstone/drivers"
	"Capstone/utils"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_threadController "Capstone/controllers/threads"
	_threadUseCase "Capstone/models/threads"

	_userController "Capstone/controllers/users"
	_userUseCase "Capstone/models/users"

	_commentController "Capstone/controllers/comments"
	_commentUseCase "Capstone/models/comments"

	_likeController "Capstone/controllers/likes"
	_likeUseCase "Capstone/models/likes"

	_followController "Capstone/controllers/follows"
	_followUseCase "Capstone/models/follows"

	_savedthreadController "Capstone/controllers/savedthreads"
	_savedthreadUseCase "Capstone/models/savedthreads"

	_dbDriver "Capstone/drivers/mysql"

	_middleware "Capstone/app/midleware"
	_routes "Capstone/app/routes"

	echo "github.com/labstack/echo/v4"
)

type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.DBConfig{
		DB_USERNAME: utils.GetConfig("DB_USERNAME"),
		DB_PASSWORD: utils.GetConfig("DB_PASSWORD"),
		DB_HOST:     utils.GetConfig("DB_HOST"),
		DB_PORT:     utils.GetConfig("DB_PORT"),
		DB_NAME:     utils.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	_dbDriver.MigrateDB(db)

	configJWT := _middleware.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	configLogger := _middleware.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	e := echo.New()

	likeRepo := _driverFactory.NewLikeRepository(db)
	likeUsecase := _likeUseCase.NewLikeUsecase(likeRepo)
	likeCtrl := _likeController.NewLikeController(likeUsecase)

	commentRepo := _driverFactory.NewCommentRepository(db)
	commentUsecase := _commentUseCase.NewCommentUsecase(commentRepo)
	commentCtrl := _commentController.NewCommentController(commentUsecase)

	threadRepo := _driverFactory.NewThreadRepository(db)
	threadUsecase := _threadUseCase.NewThreadUsecase(threadRepo)
	threadCtrl := _threadController.NewThreadController(threadUsecase)

	savedthreadRepo := _driverFactory.NewSavedthreadRepository(db)
	savedthreadUsecase := _savedthreadUseCase.NewSavedthreadUsecase(savedthreadRepo)
	savedthreadCtrl := _savedthreadController.NewSavedthreadController(savedthreadUsecase)

	followRepo := _driverFactory.NewFollowRepository(db)
	followUsecase := _followUseCase.NewFollowUsecase(followRepo)
	followCtrl := _followController.NewFollowController(followUsecase)

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUseCase(userRepo, &configJWT)
	userCtrl := _userController.NewAuthController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		AuthController:   *userCtrl,
		ThreadController: *threadCtrl,
		CommentController: *commentCtrl,
		LikeController: *likeCtrl,
		FollowController: *followCtrl,
		SavedthreadController: *savethreadCtrl,
		JWTMiddleware:    configJWT.Init(),
	}

	routesInit.RegisterRoutes(e)

	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server": func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait
}

// gracefulShutdown performs application shut down gracefully.
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
