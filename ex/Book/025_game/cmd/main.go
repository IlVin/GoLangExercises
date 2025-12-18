package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {
	val := reflect.ValueOf(object)

	fn := val.MethodByName("ReceiveSpell")
	if fn.IsValid() {
		fn.Call([]reflect.Value{reflect.ValueOf(spell)})
	} else {
		field := val.Elem().FieldByName(spell.Char())
		if field.IsValid() {
			if field.CanSet() {
				field.SetInt(field.Int() + int64(spell.Value()))
			}
		}
	}

	// реализуйте эту функцию.
}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}

func ReadTextFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", fmt.Errorf(`не удалось прочитать файл: %w`, err)
	}
	fmt.Println(string(data))
	return string(data), err
}

func main() {

	data, err := ReadTextFile(`./nothing.txt`)
	fmt.Printf("DATA: %s, ERR: %v\n", data, err)
	fmt.Printf("DATA: %s, ERR: %v\n", data, errors.Unwrap(err))
	fmt.Printf("DATA: %s, ERR: %v\n", data, errors.Unwrap(errors.Unwrap(err)))

	player := &Player{
		name:   "Player_1",
		health: 100,
	}

	enemies := []interface{}{
		&Zombie{Health: 1000},
		&Zombie{Health: 1000},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Orc{Health: 500},
		&Daemon{Health: 1000},
		&Daemon{Health: 1000},
		&Wall{Durability: 100},
	}

	fmt.Println(player)
	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	fmt.Println(player)
	CastToAll(newSpell("heal", "Health", 190), append(enemies, player))

	fmt.Println(player)
}
