package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.ResetUser(ctx)
	if err != nil {
		return fmt.Errorf("something went wrong while reseting the users table: %w", err)
	}

	fmt.Println("successfully resetted table!")

	return nil
}
