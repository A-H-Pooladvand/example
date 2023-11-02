package main

import (
	"app/internal/models"
	"app/pkg/mysql"
	"context"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	client := mysql.New()

	group, _ := errgroup.WithContext(context.Background())
	group.SetLimit(10)

	usersCh := make(chan *models.User, 1000000)

	for i := 0; i < 1000; i++ {
		group.Go(func() error {
			for j := 0; j < 1000; j++ {
				u := &models.User{}

				if err := u.Fake(); err != nil {
					return err
				}

				usersCh <- u
			}

			return nil
		})

	}

	go func() {
		if err := group.Wait(); err != nil {
			panic(err)
		}

		close(usersCh)
	}()

	var users []*models.User

	for v := range usersCh {
		users = append(users, v)

		if len(users) >= 1000 {
			client.Create(&users)
			users = nil
		}
	}
}
