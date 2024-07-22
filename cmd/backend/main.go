package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/server"
	"backend/internal/service"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "net/http/pprof"

	_ "github.com/lib/pq"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

func InitServer(handler *handler.Handler) *server.Server {
	s := server.New(
		handler.InitRoutes(),
		server.WithTimeout(10*time.Second),
	)
	return s
}

func InitDB() (*gorm.DB, error) {
	db, err := repository.NewConnectionPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
		Password: "qwerty",
	})

	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	logrus.SetFormatter(new(logrus.JSONFormatter))
	db, err := InitDB()
	if err != nil {
		logrus.Fatal("К базе данных не подключиться ", err, db)
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handler := handler.NewHandler(service)

	serv := InitServer(handler)

	go func() {
		if err := serv.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("Сервер не запустился ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Fatal("Сервер не закрылся ", err)
	}

	// Завершаем работу, закрывая очередь и дожидаясь завершения воркеров
	close(service.QueryQueue)
	service.QueueWG.Wait()
}

// // Создание файла для записи профиля CPU
// cpuProfile, err := os.Create("cpu.pprof")
// if err != nil {
// 	logrus.Fatal("could not create CPU profile: ", err)
// }
// defer cpuProfile.Close()

// if err := pprof.StartCPUProfile(cpuProfile); err != nil {
// 	logrus.Fatal("could not start CPU profile: ", err)
// }
// defer pprof.StopCPUProfile()

// // Создание файла для записи профиля кучи
// heapProfile, err := os.Create("heap.pprof")
// if err != nil {
// 	logrus.Fatal("could not create heap profile: ", err)
// }
// defer heapProfile.Close()

// // Запись профиля кучи в файл при завершении
// defer pprof.WriteHeapProfile(heapProfile)
