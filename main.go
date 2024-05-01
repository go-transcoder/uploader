package main

import (
	"fmt"
	events2 "github.com/go-transcoder/uploader/internal/application/events"
	"github.com/go-transcoder/uploader/internal/application/interfaces"
	"github.com/go-transcoder/uploader/internal/application/services"
	postgres2 "github.com/go-transcoder/uploader/internal/infrastructure/db/postgres"
	"github.com/go-transcoder/uploader/internal/infrastructure/events/kafka"
	"github.com/go-transcoder/uploader/internal/interface/api/rest"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	// we load the .env.test only if we are working locally
	PROJECT_ENV := os.Getenv("PROJECT_ENV")

	if PROJECT_ENV != "prod" {
		godotenv.Load(".env.test")
	}

	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	DBNAME := os.Getenv("DBNAME")
	DBUSER := os.Getenv("DBUSER")
	DBPASS := os.Getenv("DBPASS")
	SSLMODE := os.Getenv("SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", DBHOST, DBUSER, DBPASS, DBNAME, DBPORT, SSLMODE)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	unitOfWork := postgres2.NewUnityOfWork(gormDB)
	videoService := services.NewVideoService(unitOfWork)

	e := echo.New()
	rest.NewVideoController(e, videoService)

	// Start a new goroutine which will listen to the queue
	KAFKAHOST := os.Getenv("KAFKAHOST")
	KAFKATOPIC := os.Getenv("KAFKATOPIC")

	// Create the events consumer here and pass them as an slice to the eventListener func
	videoTranscodedEventConsumer := kafka.NewEventConsumerService(KAFKAHOST, KAFKATOPIC)
	videoTranscodedEventHandler := events2.NewVideoTranscodedEventHandler(videoTranscodedEventConsumer, unitOfWork.GetVideosRepo())

	eventHandlers := []interfaces.EventHandlers{
		videoTranscodedEventHandler,
	}

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	exitChan := make(chan bool, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-sigChan
		fmt.Println("Received SIGTERM. Exiting...")
		exitChan <- true
	}()

	eventListenerChannel := eventListener(eventHandlers)
	webServerChannel := startWebServer(e)

	wg.Add(1)
	go func() {
		defer func() {
			err := videoTranscodedEventConsumer.Close()
			if err != nil {
				log.Fatalf("Failed to close the kafka listener: %v", err)
			}
			fmt.Println("Kafka closed")
			wg.Done()
		}()

		for {
			select {
			case <-exitChan:
				fmt.Println("Exiting the event processing routine...")
				return
			case webserver := <-webServerChannel:
				fmt.Println(webserver)
			case eventListener := <-eventListenerChannel:
				fmt.Println(eventListener)
			}
		}
	}()
	wg.Wait()
	fmt.Println("Exited")
}

func eventListener(eventHandlers []interfaces.EventHandlers) <-chan string {
	c := make(chan string)

	go func() {
		for {
			for _, handler := range eventHandlers {
				err := handler.Process()

				if err != nil {
					fmt.Printf("Error getting the event message: %v", err)
					c <- fmt.Sprintf("Error getting the event message: %v", err)
				}

			}
		}
	}()

	return c
}

func startWebServer(e *echo.Echo) <-chan string {
	port := os.Getenv("PORT")

	c := make(chan string)

	go func() {
		if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
			fmt.Printf("Error starting the web server: %v", err)
			c <- fmt.Sprintf("Error starting the web server: %v", err)
		}
	}()

	return c
}
