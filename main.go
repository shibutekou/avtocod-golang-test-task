package main

import (
    "app/internal/delivery"
    "app/internal/domain"
    "app/internal/repository"
    "log"
    "net/http"
    "os"
)

func main() {
    client := http.Client{}
    var token = os.Getenv("MOS_TOKEN")

    rdb := repository.NewRedisClient()
    redisRepo := repository.NewRedisRepository(rdb)

    service := domain.NewService(token, client, redisRepo)

    handler := delivery.NewHandler(service)
    mux := handler.InitRoutes()

    log.Fatal(http.ListenAndServe("localhost:8800", mux))
}
