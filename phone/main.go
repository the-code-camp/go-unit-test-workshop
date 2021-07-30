package main

/**
* This example shows how a phone state goes between different transitions based on the events
* For this example we just have two events "Pressing Home Button" and "Providing PIN" Pin can be correct or incorrect
* Use case
  1. A new phone with battery should have status as DISPLAY OFF
  2. A new phone without battery should have status as SWITCHED OFF
  3. A new phone with battery and Home button Pressed should have status as LOCKED
  4. A new phone with battery and correct PIN should have status as UNLOCKED
  5. A new phone with battery and incorrect PIN should have status as SWITCHED ON AND LOCKED
  6. A new phone with battery and 3 Unlock attempts with incorrect PIN should have status as PHONE IS BLOCKED
*/

const (
	SwitchedOff = iota
	DisplayOff
	SwitchedOnAndLocked
	Blocked
	Locked
	Unlocked
)

var statusMessages map[int]string = map[int]string{
	DisplayOff:          "DISPLAY OFF",
	SwitchedOff:         "SWITCHED OFF",
	Locked:              "LOCKED",
	Unlocked:            "UNLOCKED",
	SwitchedOnAndLocked: "SWITCHED ON AND LOCKED",
	Blocked:             "PHONE IS BLOCKED",
}

type Phone struct {
	Pin                string
	IsBatteryAvailable bool
	state              int
	unlockCounter      int
}

func (p *Phone) PressHome() {
	if p.IsBatteryAvailable {
		p.state = SwitchedOnAndLocked
	}
}

func (p *Phone) GetStatus() string {
	if p.IsBatteryAvailable == true {
		return statusMessages[p.state]
	} else {
		return statusMessages[SwitchedOff]
	}
}

func (p *Phone) UnLock(s string) {
	if p.IsBatteryAvailable {
		if p.unlockCounter > 2 {
			p.state = Blocked
		} else if p.Pin == s {
			p.state = Unlocked
			p.unlockCounter = 0
		} else {
			p.state = Locked
			p.unlockCounter++
		}
	} else {
		p.state = SwitchedOff
	}
}

func NewPhone(pin string, batteryStatus bool) Phone {
	return Phone{
		Pin:                pin,
		IsBatteryAvailable: batteryStatus,
		state:              DisplayOff,
	}
}
