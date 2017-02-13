package minesweepergo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type Game struct {
	Id                string `json:"id"`
	Title             string `json:"title"`
	State             string `json:"state"`
	BoardView         [][]string `json:"board_view"`
	DurationSeconds   int `json:"duration_seconds"`
	ElapsedSeconds    int `json:"elapsed_seconds"`
	Score             int `json:"score"`
	ResumedTimestamp  string `json:"resumed_timestamp"`
}

func sendRequest(req *http.Request) *Game {
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
	return &game
}

func State(host string, game_id string) *Game {
	safeId := url.QueryEscape(game_id)
	url := fmt.Sprintf("http://%s/api/v1/games/state/?game_id=%s", host, safeId)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
	game := sendRequest(req)
	return game
}

func New(host string, rows int, columns int, mines int) *Game {
	type Arg struct {
		Rows    int `json:"rows"`
		Columns int `json:"columns"`
		Mines   int `json:"mines"`
	}
	data, _ := json.Marshal(Arg{Rows: rows, Columns: columns, Mines: mines})
	url := fmt.Sprintf("http://%s/api/v1/games/new/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
}

func Pause(host string, game_id string) *Game {
	type Arg struct {
		GameId string `json:"game_id"`
	}
	data, _ := json.Marshal(Arg{GameId: game_id})
	url := fmt.Sprintf("http://%s/api/v1/games/pause/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
}

func Resume(host string, game_id string) *Game {
	type Arg struct {
		GameId string `json:"game_id"`
	}
	data, _ := json.Marshal(Arg{GameId: game_id})
	url := fmt.Sprintf("http://%s/api/v1/games/resume/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
}

func MarkAsFlag(host string, game_id string, x int, y int) *Game {
	type Arg struct {
		GameId string `json:"game_id"`
		X int `json:"x"`
		Y int `json:"y"`
	}
	data, _ := json.Marshal(Arg{GameId: game_id, X: x, Y: y})
	url := fmt.Sprintf("http://%s/api/v1/games/mark_as_flag/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
}

func MarkAsQuestion(host string, game_id string, x int, y int) *Game {
	type Arg struct {
		GameId string `json:"game_id"`
		X int `json:"x"`
		Y int `json:"y"`
	}
	data, _ := json.Marshal(Arg{GameId: game_id, X: x, Y: y})
	url := fmt.Sprintf("http://%s/api/v1/games/mark_as_question/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
	return &Game{}
}

func Reveal(host string, game_id string, x int, y int) *Game {
	type Arg struct {
		GameId string `json:"game_id"`
		X int `json:"x"`
		Y int `json:"y"`
	}
	data, _ := json.Marshal(Arg{GameId: game_id, X: x, Y: y})
	url := fmt.Sprintf("http://%s/api/v1/games/reveal/", host)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatal("Error in NewRequest: ", err)
		return nil
	}
    req.Header.Add("Content-Type", "application/json")
	game := sendRequest(req)
	return game
}
