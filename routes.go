package main

import (
	"log"
	"net/http"

	"github.com/battlesnakeio/starter-snake-go/api"
)

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}
	dump(decoded)

	respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	dump(decoded)

	/*
	directions := []string{
		"up",
		"down",
		"left",
		"right",
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	
	respond(res, api.MoveResponse{
		Move: directions[r.Intn(4)],
	})
	*/
	/*
	boardWidth := decoded.Board.Width
	myX := decoded.You.Body.X
	if (boardWidth - myX) > (myX) {
		respond(res, api.MoveResponse{
			Move: "right",
		}
	} else {
		respond(res, api.MoveResponse{
			Move: "left",
		}
	}
	*/
	respond(res, api.MoveResponse{
			Move: "right",
	}
}

func End(res http.ResponseWriter, req *http.Request) {
	return
}

func Ping(res http.ResponseWriter, req *http.Request) {
	return
}
