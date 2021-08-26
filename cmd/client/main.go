package main

import (
	"context"
	"fmt"
	"github.com/inancgumus/screen"
	"github.com/pushkinvladislav/link_shortening/api/shorter"
	"github.com/pushkinvladislav/link_shortening/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	logger.InitLogger()
	logger.Logger.Info("Starting client...")

	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := shorter.NewShorterServiceClient(conn)

	var comand string
	var url string

	for comand != "exit" {
		screen.Clear()
		fmt.Printf("Cервис генерации сокращенных ссылок\n")
		fmt.Print("Доступные команды:\n")
		fmt.Print("Создать shortURl: Create\n")
		fmt.Print("Вернуть longURl:  Get\n")
		fmt.Print("Выход:            exit\n")
		fmt.Scanf("%v/$v\n", &comand, &url)

		switch comand {
		case "Create":
			screen.Clear()
			fmt.Printf("Введите longURL\n")
			fmt.Scanf("%v\n", &url)

			if url != "" {
				res, err := client.Create(context.Background(), &shorter.CreateRequest{LongURL: url})
				if err != nil {
					logger.Logger.Error(err)
				}
				fmt.Printf("\nВаша ссылка:  %s\n", res.GetShortURL())
			} else {
				fmt.Printf("Вы ничего не ввели\n")
			}
			fmt.Printf("\nНажмите Enter, чтобы вернуться в главное меню")
			fmt.Scanf("\n")
			screen.Clear()

		case "Get":
			screen.Clear()
			fmt.Printf("Введите shortURL без localhost:8080/\n")
			fmt.Scanf("%v\n", &url)

			if url != "" {
				res, err := client.Get(context.Background(), &shorter.GetRequest{ShortURL: url})
				if err != nil {
					logger.Logger.Error(err)
				}
				if res.LongURL == "" {
					fmt.Printf("Короткая ссылка для этого URL не создавалась\n")
				} else {
					fmt.Printf("Ваша ссылка: %s\n", res.GetLongURL())
				}

			} else {
				fmt.Printf("Вы ничего не ввели\n")
			}
			fmt.Printf("Нажмите Enter, чтобы вернуться в главное меню")
			fmt.Scanf("\n")
			screen.Clear()

		case "exit":

		default:
			fmt.Printf("Команда введена неверно, попробуйте еще раз\n")
		}

	}
}
