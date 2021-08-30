package main

import "unsafe"

var USERMEMORY = (*[58975]uint8)(unsafe.Pointer(uintptr(0x19a0)))

var ship = NewShip(SCREEN_SIZE, SCREEN_SIZE)

var land = NewLand(SCREEN_SIZE, SCREEN_SIZE)

func UNUSED(x ...interface{}) {}

//go:export update
func update() {

	PALETTE[0] = 0x2a3abd
	PALETTE[1] = 0xcdc23c
	PALETTE[2] = 0x335715
	PALETTE[3] = 0x7c3f58
	*DRAW_COLORS = 2

	var gamepad = *GAMEPAD1
	ship.setupFrame()
	if gamepad&BUTTON_RIGHT != 0 {
		ship.moveRight(SCREEN_SIZE, SCREEN_SIZE, &land)
	}
	if gamepad&BUTTON_LEFT != 0 {
		ship.moveLeft(SCREEN_SIZE, SCREEN_SIZE, &land)
	}
	if gamepad&BUTTON_UP != 0 {
		ship.moveUp(SCREEN_SIZE, SCREEN_SIZE, &land)
	}
	if gamepad&BUTTON_DOWN != 0 {
		ship.moveDown(SCREEN_SIZE, SCREEN_SIZE)
	}
	land.render(SCREEN_SIZE, SCREEN_SIZE)
	ship.render()

}
