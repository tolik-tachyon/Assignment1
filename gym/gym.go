package Gym

import "fmt"

type Member interface {
	GetDetails() string
}

type BasicMember struct {
	Name string
}

func (b BasicMember) GetDetails() string {
	return "Basic Member: " + b.Name
}

type PremiumMember struct {
	Name  string
	Level string
}

func (p PremiumMember) GetDetails() string {
	return "Premium Member: " + p.Name + " Level: " + p.Level
}

type Gym struct {
	Members map[uint64]Member
}

func NewGym() Gym {
	return Gym{Members: make(map[uint64]Member)}
}

func (g *Gym) AddMember(id uint64, m Member) {
	g.Members[id] = m
}

func (g *Gym) ListMembers() {
	for id, member := range g.Members {
		fmt.Println(id, member.GetDetails())
	}
}

func DemoGym() {
	fmt.Println("\n--- GYM DEMO ---")

	gym := NewGym()
	gym.AddMember(1, BasicMember{"Alex"})
	gym.AddMember(2, PremiumMember{"John", "Gold"})

	gym.ListMembers()
}
