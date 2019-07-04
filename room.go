package main

import "log"

type Room struct {
	//
	hub *Hub

	//
	roomID  string
	player1 *Client
	player2 *Client

	// Inbound messages from the clients.
	broadcast chan Message

	playerJoin  chan *Client
	playerLeave chan *Client

	game *Game
}

type ClientRoomMessage struct {
	Client  *Client
	Message Message
}

func (room *Room) run() {
	defer func() {
		// remove reference of hub
		room.hub = nil
		// cleanup player
		room.player1 = nil
		room.player2 = nil

		// user already leave before closing room
		// room.player1.leaveRoom <- roomID
		// room.player2.leaveRoom <- roomID

		// if both players are nil => delete room
		if room.player1 == nil && room.player2 == nil {
			delete(room.hub.rooms, room.roomID)
			close(room.broadcast)
			close(room.playerJoin)
			close(room.playerLeave)
		}
	}()
	for {
		select {
		case message := <-room.broadcast:
			if room.player1 != nil {
				room.player1.send <- message
			}
			if room.player2 != nil {
				room.player2.send <- message
			}

		case client := <-room.playerJoin:
			// if join fail -> send message type = 8
			if room.player1 != nil && room.player2 != nil {
				fullCapacityRoomMsg := Message{Type: 8, Message: "Room is full of capacity"}
				client.send <- fullCapacityRoomMsg
			} else {
				// if join success -> send message type = ??
				log.Println("room available")
				if room.player1 == nil {
					room.player1 = client
				} else if room.player2 == nil {
					room.player2 = client
				}
				client.joinRoom <- room.roomID
				// TODO: should broadcast msg new user has just joined to channel
				// implementation
			}
		case client := <-room.playerLeave:
			if room.player1 == client {
				room.player1 = nil
			} else if room.player2 == client {
				room.player2 = nil
			}
			// if both players are nil => remove room from hub and close room itself
			if room.player1 == nil && room.player2 == nil {
				// remove from hub
				delete(room.hub.rooms, room.roomID)
				// close room itself
				break
			}
		}

	}
}
