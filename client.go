package minesweepergo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Game struct {
	Id                bool   `json:"id"`
	Title             string `json:"title"`
	State             string `json:"state"`
	BoardView         string `json:"board_view"`
	DurationSeconds   string `json:"duration_seconds"`
	ElapsedSeconds    string `json:"elapsed_seconds"`
	Score             string `json:"score"`
	ResumedTimestamp  string `json:"resumed_timestamp"`
	Player            string `json:"player"`
}

func Sum(i, j int) int {
    return i + j
}

func State(game_id string) *Game {
	safeId := url.QueryEscape(game_id)
	url := fmt.Sprintf("http://127.0.0.1:7000/api/v1/games/state/?game_id=%s", safeId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error in Do: ", err)
		return nil
	}

	defer resp.Body.Close()
	var game Game
	if err := json.NewDecoder(resp.Body).Decode(&game); err != nil {
		log.Println(err)
	}

	fmt.Println("Id: ", game.Id)
	return &game
}

func New(rows int, columns int, mines int) *Game {
	return &Game{}
}

func Pause(game_id string) *Game {
	return &Game{}
}

func Resume(game_id string) *Game {
	return &Game{}
}

func MarkAsFlag(game_id string, x int, y int) *Game {
	return &Game{}
}

func MarkAsQuestion(game_id string, x int, y int) *Game {
	return &Game{}
}

func Reveal(game_id string, x int, y int) *Game {
	return &Game{}
}
