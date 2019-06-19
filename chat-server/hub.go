// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"
	"log"
	// "github.com/rs/xid"
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

	// boardcast room
	broadcastRoom chan *ClientRoomMessage
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		// rooms of namepsace
		rooms: 					make(map[string]*Room),
		registerRoom: 	make(chan *Client),
		leaveRoom: 			make(chan *Client),
		broadcastRoom:	make(chan *ClientRoomMessage),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			// log.Printf("register client %p\n", client)
			h.clients[client] = true
		case client := <-h.unregister:
			// log.Printf("unregister client %p\n", client)
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			// this was block, how to fix???
			// h.leaveRoom <-client

			if room, ok := h.rooms[client.room]; ok {
				room.playerLeave <- client
				client.leaveRoom <- room.roomID
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
		case client := <- h.registerRoom:
			// TODO: temporary hardcode
			roomID := "hardcode_room01"
			// roomID := xid.New().String()
			log.Println("register room")
			// if room was found
			// TODO: many cases for handling
			room, found := h.rooms[roomID] 
				if found {
					if room.player2 == nil {
						room.player2 = client
					} else if room.player1 == nil {
						room.player1 = client
					}
					// set room client
					client.joinRoom <- roomID
				} else {
					// add player 1 first
					newRoom := Room{
						hub: h, 
						roomID: roomID, 
						player1: client,
						player2: nil,
						broadcast: make(chan Message), 
						playerLeave: make(chan *Client, 2), 
						close: make(chan bool),
					}
						go newRoom.run()
						// create new room of hub
						h.rooms[roomID] = &newRoom
						// pass message to client
						// set room client
						client.joinRoom <- roomID
				}
			
		case c := <- h.leaveRoom:
			log.Println("leave room: ",c)
			if room, ok := h.rooms[c.room]; ok {
				room.playerLeave <- c
			}
			c.leaveRoom <- c.room

		case clientMessage := <- h.broadcastRoom:
			roomID := clientMessage.Client.room
			message := clientMessage.Message
			if room, found := h.rooms[roomID]; found {
				if room.player1 == clientMessage.Client || room.player2 == clientMessage.Client {
					room.broadcast <- message
				}
			}
		}

	}
}

func (h *Hub) loggingRooms() {
	go func(){
		for{
			for k := range h.rooms {
				log.Println("roomId: ", k)
			time.Sleep(time.Second * 1)

			}
			log.Println("end of rooms")
			time.Sleep(time.Second * 1)
		}
	}()
}