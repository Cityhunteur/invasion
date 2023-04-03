package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"strconv"
	"syscall"

	"github.com/cityhunteur/invasion/invasion"
)

var (
	aliensFlag = flag.String("aliens", "10", "number of aliens")
	mapFlag    = flag.String("map", "testdata/example.map", "path to file containing map of the world")
)

func main() {
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			fmt.Println("program panic", r, string(debug.Stack()))
			os.Exit(1)
		}
	}()

	err := realMain(ctx)
	done()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func realMain(_ context.Context) error {
	nAliens, err := strconv.Atoi(*aliensFlag)
	if err != nil {
		return errors.New("--aliens must specify a number. Default is 10")
	}

	if !mapExist(*mapFlag) {
		return errors.New("--map must provide a file. Default is 'testdata/example.map'")
	}

	fmt.Printf("Simulating invasion with %d aliens using map %s\n", nAliens, *mapFlag)

	f, err := os.Open(*mapFlag)
	if err != nil {
		return errors.New("--aliens must specify a number. Default is 10")
	}

	w, err := invasion.NewWorld(f)
	if err != nil {
		return fmt.Errorf("failed to create new world: %w", err)
	}

	s := invasion.NewSimulation(w, nAliens)
	if err != nil {
		return fmt.Errorf("failed to create simulation: %w", err)
	}

	s.Start()

	return nil
}

func mapExist(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}
