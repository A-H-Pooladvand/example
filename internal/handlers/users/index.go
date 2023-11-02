package users

import (
	"app/internal/models"
	"app/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/sync/errgroup"
	"math"
	"strconv"
)

func Index(e echo.Context) error {

	// Todo:: refactor to a pipeline
	take := 10
	t := e.QueryParam("take")

	if t != "" {
		take, _ = strconv.Atoi(t)

		if take == 0 {
			take = 10
		}
	}
	// Todo:: End

	var usersCh = make(chan models.User, take)

	service := services.NewUserService()

	group, ctx := errgroup.WithContext(e.Request().Context())
	group.SetLimit(1)

	// Todo:: Refactor to dedicated func
	loop := take

	if take/1000 <= 1 {
		loop = 1
	} else {
		loop = int(math.Ceil(float64(take / 1000)))
	}
	// Todo:: End

	for i := 0; i < loop; i++ {
		group.Go(func() error {
			var users []models.User
			service.TakeUsers(ctx, i, 1000, &users)

			for _, u := range users {
				usersCh <- u
			}

			return nil
		})
	}

	go func() {
		if err := group.Wait(); err != nil {
			log.Error(err)
		}

		close(usersCh)
	}()

	var allUsers []models.User
	for v := range usersCh {
		allUsers = append(allUsers, v)
	}

	return e.JSON(200, map[string]any{
		"len":  len(allUsers),
		"data": allUsers,
	})
}
