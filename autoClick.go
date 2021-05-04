package main

import (
	"container/list"
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type coordenada struct {
	x int
	y int
}

func temporizador(s int) {
	for {
		if s <= 0 {
			break
		} else {
			fmt.Println(s)
			time.Sleep(1 * time.Second)
			s--
		}
	}
}

func main() {

	var listaClicks list.List
	var xy coordenada
	fmt.Println("Situa tu click donde quieras, y para guardar presiona a ")
	robotgo.EventHook(hook.KeyDown, []string{"a"}, func(e hook.Event) {
		x, y := robotgo.GetMousePos()
		xy.x = x
		xy.y = y
		fmt.Println(xy)
		listaClicks.PushBack(xy)
	})

	fmt.Println("--- Para terminar de agregar a la secuencia y empezar los clickeos presiona enter ---")
	robotgo.EventHook(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		for {
			cnt:=1
			for e := listaClicks.Front(); e != nil; e = e.Next() {
				if coord, ok := e.Value.(coordenada); ok {
					if cnt == 3 || cnt == 8 || cnt == 13 || cnt == 18{
						robotgo.KeyTap("esc")
						time.Sleep(2 * time.Second)
					}
						robotgo.MoveMouse(coord.x, coord.y)
						time.Sleep(3 * time.Second)
						robotgo.Click()

					
					cnt++
				}
			}
			temporizador(90 * 60)
		}
	})

	robotgo.EventHook(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		fmt.Println("Adios")
		robotgo.EventEnd()
	})

	s := robotgo.EventStart()
	<-robotgo.EventProcess(s)
}
