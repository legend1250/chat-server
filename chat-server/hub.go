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
	leaverRoom chan *Client
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
		leaverRoom: 		make(chan *Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
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
			
			// if room was found
			// TODO: many cases for handling
			if room, found := h.rooms[roomID]; found {
				room.player2 = client

				client.joinRoom <- roomID
				} else {
					// add player 1 first
				newRoom := Room{roomID: roomID, player1: client}
				// set client room
				client.room = roomID
				// create new room of hub
				h.rooms[roomID] = &newRoom
				// pass message to client
				client.joinRoom <- roomID
			}
		case client := <- h.leaverRoom:
			joinedRoom := client.room

			if room, found := h.rooms[joinedRoom]; found {
				if room.player1 == client {
					room.player1 = nil
				} else if room.player2 == client {
					room.player2 = nil
				}
				// if both players are nil => delete room
				if room.player1 == nil && room.player2 == nil {
					log.Println("delete room: ", joinedRoom)
					delete(h.rooms, joinedRoom)
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