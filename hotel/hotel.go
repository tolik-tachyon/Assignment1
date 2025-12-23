package Hotel

import "fmt"

type Room struct {
	RoomNumber    string
	Type          string
	PricePerNight float64
	IsOccupied    bool
}

type Hotel struct {
	Rooms map[string]Room
}

func NewHotel() Hotel {
	return Hotel{Rooms: make(map[string]Room)}
}

func (h *Hotel) AddRoom(room Room) {
	h.Rooms[room.RoomNumber] = room
}

func (h *Hotel) CheckIn(roomNumber string) {
	room := h.Rooms[roomNumber]
	if room.IsOccupied {
		fmt.Println("Room already occupied")
		return
	}
	room.IsOccupied = true
	h.Rooms[roomNumber] = room
}

func (h *Hotel) CheckOut(roomNumber string) {
	room := h.Rooms[roomNumber]
	room.IsOccupied = false
	h.Rooms[roomNumber] = room
}

func (h *Hotel) ListVacantRooms() {
	fmt.Println("Vacant rooms:")
	for _, room := range h.Rooms {
		if !room.IsOccupied {
			fmt.Println(room.RoomNumber, room.Type, room.PricePerNight)
		}
	}
}

func DemoHotel() {
	fmt.Println("\n--- HOTEL DEMO ---")
	hotel := NewHotel()

	hotel.AddRoom(Room{"101", "Single", 50, false})
	hotel.AddRoom(Room{"102", "Double", 80, false})

	hotel.CheckIn("101")
	hotel.ListVacantRooms()
	hotel.CheckOut("101")
}
