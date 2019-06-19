package main

// import "log"

type Room struct {
	// 
	hub *Hub

	// 
	roomID string
	player1 *Client
	player2 *Client

	// Inbound messages from the clients.
	broadcast chan Message

	playerLeave chan *Client

	// close channel
	close chan bool
}

type ClientRoomMessage struct {
	Client *Client
	Message Message
}

func (room *Room) run(){
	defer func(){
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
			close(room.playerLeave)
			close(room.close)
		}
	}()
	for{
		select {
		case message := <-room.broadcast:
			if room.player1 != nil {
				room.player1.send <- message
			}
			if room.player2 != nil {
				room.player2.send <- message
			}
		case close := <-room.close:
			if close {
				break;
			}
		case client := <- room.playerLeave:
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
				room.close <- true
			}
		}

	}
}
