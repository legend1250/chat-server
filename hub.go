// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"
	"time"

	"github.com/rs/xid"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	// Room create from clients
	rooms map[string]*Room

	// Client register room
	registerRoom chan *Client

	// Client leave room
	leaveRoom chan *Client

	// Client join a room with code
	joinRoom chan *ClientRoomMessage

	// Client join a room quickly
	joinRoomQuickly chan *ClientRoomMessage

	// Client move next
	movePlayer chan *ClientRoomMessage

	// broadcast room
	broadcastRoom chan *ClientRoomMessage
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		// rooms of namespace
		rooms: make(map[string]*Room),
		// register (create) a room
		registerRoom: make(chan *Client, 1024),
		// leave / quit a room
		leaveRoom: make(chan *Client, 1024),
		// join a room with code
		joinRoom: make(chan *ClientRoomMessage, 1024),
		// join a room quickly
		joinRoomQuickly: make(chan *ClientRoomMessage, 1024),

		// player move next
		movePlayer: make(chan *ClientRoomMessage, 1024),

		// broadcast entire room
		broadcastRoom: make(chan *ClientRoomMessage, 1024),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			log.Printf("register client %p\n", client)
			h.clients[client] = true
		case client := <-h.unregister:
			log.Printf("unregister client %p\n", client)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			if room, ok := h.rooms[client.room]; ok {
				room.playerLeave <- client
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case client := <-h.registerRoom:
			// random roomID
			roomID := xid.New().String()
			// create new room
			newRoom := Room{
				hub:         h,
				roomID:      roomID,
				player1:     client,
				broadcast:   make(chan Message, 1024),
				playerJoin:  make(chan *Client, 2),
				playerLeave: make(chan *Client, 2),
			}
			go newRoom.run()
			// create new room of hub
			h.rooms[roomID] = &newRoom
			// pass message to client
			// set room client
			client.joinRoom <- roomID

		case c := <-h.leaveRoom:
			if room, ok := h.rooms[c.room]; ok {
				room.playerLeave <- c
			}
			c.leaveRoom <- c.room

		case clientMessage := <-h.joinRoom:
			// get roomID from sent message
			requestRoomID := clientMessage.Message.RoomID
			// found room in hub
			if room, ok := h.rooms[requestRoomID]; ok {
				if clientMessage.Client.room == "" {
					room.playerJoin <- clientMessage.Client
				}
			} else {
				// join room error code = 8
				notFoundRoomMsg := Message{Type: 8, Message: "Room is not exist"}
				clientMessage.Client.send <- notFoundRoomMsg
			}

		case clientMessage := <-h.broadcastRoom:
			c := clientMessage.Client
			if valid, room := h.isClientBelongToRoom(c); valid {
				room.broadcast <- clientMessage.Message
			}

		case clientMessage := <-h.joinRoomQuickly:
			c := clientMessage.Client
			// TODO: check whether client will join multiple of rooms
			findRoomChannel := make(chan string)
			go func() {
				var isFound bool = false
				var roomID string
				for _, room := range h.rooms {
					//if room.player1 != nil && room.player2 == nil {
					//	c.room = id
					//	room.player2 = c
					//} else if room.player1 == nil && room.player2 != nil {
					//	c.room = id
					//	room.player1 = c
					//} else {
					//	msg := Message{Type: 10, Message: "No rooms are available for now"}
					//	c.send <- msg
					//}
					if !isFound {
						if (room.player1 != nil && room.player2 == nil) || (room.player1 == nil && room.player2 != nil) {
							// room.playerJoin <- c
							roomID = room.roomID
							isFound = true
							break
						}
					}
				}
				if isFound {
					findRoomChannel <- roomID
				} else {
					findRoomChannel <- ""
				}
			}()
			// receiver round found
			roomIDFound := <-findRoomChannel
			close(findRoomChannel)
			// send player to room
			if roomIDFound != "" {
				h.rooms[roomIDFound].playerJoin <- c
			} else {
				// send message to room
				message := Message{Type: 10, Message: "There is no room available"}
				c.send <- message
			}
		}
	}
}

func (h *Hub) loggingRooms() {
	go func() {
		for {
			for k, v := range h.rooms {
				log.Printf("roomId: %v player1 %p player2 %p", k, v.player1, v.player2)
				time.Sleep(time.Second * 1)
			}
			log.Println("end of rooms")
			time.Sleep(time.Second * 1)
		}
	}()
}

func (h *Hub) isClientBelongToRoom(c *Client) (valid bool, room *Room) {
	roomID := c.room
	if room, found := h.rooms[roomID]; found {
		if room.player1 == c || room.player2 == c {
			return true, room
		}
		return false, nil
	}

	return false, nil
}
