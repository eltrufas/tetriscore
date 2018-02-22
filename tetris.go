package tetriscore

import (
	"fmt"
	"math/rand"
)

type Block int

const (
	Empty Block = iota
	Cyan
	Yellow
	Purple
	Green
	Red
	Blue
	Orange
)

const Timestep float32 = 16.66

var IStates [4][16]int = [4][16]int{
	{
		0, 0, 0, 0,
		1, 1, 1, 1,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 1, 0,
		0, 0, 1, 0,
		0, 0, 1, 0,
		0, 0, 1, 0,
	},
	{
		0, 0, 0, 0,
		0, 0, 0, 0,
		1, 1, 1, 1,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		0, 1, 0, 0,
		0, 1, 0, 0,
		0, 1, 0, 0,
	},
}

var JStates [4][16]int = [4][16]int{
	{
		1, 0, 0, 0,
		1, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 1, 0,
		0, 1, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 0, 0,
		1, 1, 1, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		0, 1, 0, 0,
		1, 1, 0, 0,
		0, 0, 0, 0,
	},
}

var LStates [4][16]int = [4][16]int{
	{
		0, 0, 1, 0,
		1, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		0, 1, 0, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 0, 0,
		1, 1, 1, 0,
		1, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		1, 1, 0, 0,
		0, 1, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
}

var OStates [4][16]int = [4][16]int{
	{
		0, 1, 1, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 1, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 1, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 1, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
}

var SStates [4][16]int = [4][16]int{
	{
		0, 1, 1, 0,
		1, 1, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		0, 1, 1, 0,
		0, 0, 1, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 0, 0,
		0, 1, 1, 0,
		1, 1, 0, 0,
		0, 0, 0, 0,
	},
	{
		1, 0, 0, 0,
		1, 1, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
}

var TStates [4][16]int = [4][16]int{
	{
		0, 1, 0, 0,
		1, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		0, 1, 1, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 0, 0,
		1, 1, 1, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		1, 1, 0, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
}

var ZStates [4][16]int = [4][16]int{
	{
		1, 1, 0, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 1, 0,
		0, 1, 1, 0,
		0, 1, 0, 0,
		0, 0, 0, 0,
	},
	{
		0, 0, 0, 0,
		1, 1, 0, 0,
		0, 1, 1, 0,
		0, 0, 0, 0,
	},
	{
		0, 1, 0, 0,
		1, 1, 0, 0,
		1, 0, 0, 0,
		0, 0, 0, 0,
	},
}

var Tetrominos [7]*[4][16]int = [7]*[4][16]int{
	&IStates,
	&JStates,
	&LStates,
	&OStates,
	&SStates,
	&TStates,
	&ZStates,
}

var TetrominoColors [7]Block = [7]Block{
	Cyan,
	Yellow,
	Purple,
	Green,
	Red,
	Blue,
	Orange,
}

type Vec2 struct {
	X, Y int
}

type Piece struct {
	X             int
	Y             int
	State         int
	TetrominoType int
}

type InputState struct {
	Left, Right, Up, Down, Space bool
}

type InputTimers struct {
	Left, Right, Up, Down, Space int
}

var WallKicks [8][5]Vec2 = [8][5]Vec2{
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: -1, Y: 1},
		Vec2{X: 0, Y: -2},
		Vec2{X: -1, Y: -2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: 1, Y: -1},
		Vec2{X: 0, Y: 2},
		Vec2{X: 1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: 1, Y: -1},
		Vec2{X: 0, Y: 2},
		Vec2{X: 1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: -1, Y: 1},
		Vec2{X: 0, Y: -2},
		Vec2{X: -1, Y: -2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: 1, Y: 1},
		Vec2{X: 0, Y: -2},
		Vec2{X: 1, Y: -2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: -1, Y: -1},
		Vec2{X: 0, Y: 2},
		Vec2{X: -1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: -1, Y: -1},
		Vec2{X: 0, Y: 2},
		Vec2{X: -1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: 1, Y: 1},
		Vec2{X: 0, Y: -2},
		Vec2{X: 1, Y: -2},
	},
}

var IWallKicks [8][5]Vec2 = [8][5]Vec2{
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -2, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: -2, Y: -1},
		Vec2{X: 1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 2, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: 2, Y: 1},
		Vec2{X: -1, Y: -2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: 2, Y: 0},
		Vec2{X: -1, Y: 2},
		Vec2{X: 2, Y: -1},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: -2, Y: 0},
		Vec2{X: 1, Y: -2},
		Vec2{X: -2, Y: 1},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 2, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: 2, Y: 1},
		Vec2{X: -1, Y: -2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -2, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: -2, Y: -1},
		Vec2{X: 1, Y: 2},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: 1, Y: 0},
		Vec2{X: -2, Y: 0},
		Vec2{X: 1, Y: -2},
		Vec2{X: -2, Y: 1},
	},
	{
		Vec2{X: 0, Y: 0},
		Vec2{X: -1, Y: 0},
		Vec2{X: 2, Y: 0},
		Vec2{X: -1, Y: 2},
		Vec2{X: 2, Y: -1},
	},
}

type Tetris struct {
	Board            [220]Block //10*22
	CurrentPiece     Piece
	Score            int
	PieceQueue       [14]int
	NextIndex        int
	Gravity          float32
	DropTimer        float32
	LockTime         float32
	lockTimerStarted bool
	lockTimer        float32
	ToClear          [20]bool
	It               InputTimers
	FlagLoss		 bool
}

func (t *Tetris) NextPiece() int {
	return t.PieceQueue[t.NextIndex]
}

func (t *Tetris) ShuffleQueue() {
	offset := 0
	if t.NextIndex < 7 {
		offset = 7
	}

	for i := offset + 6; i > offset; i-- {
		j := offset + rand.Intn(7)

		t.PieceQueue[i], t.PieceQueue[j] = t.PieceQueue[j], t.PieceQueue[i]
	}
}

func (t *Tetris) Update(is InputState) {
	t.updateInputTimers(is)
	t.applyMovement()
	t.updateLockTimer()

	for t.DropTimer >= 1. {
		t.DropTimer--
		if !t.SoftDrop() {
			t.startLockTimer()
		}
	}

	t.DropTimer += t.Gravity
}

func (t *Tetris) applyMovement() {
	initPiece := t.CurrentPiece

	if t.It.Left == 1 || t.It.Left > 15 {
		t.moveLeft()
	}

	if t.It.Right == 1 || t.It.Right > 15 {
		t.moveRight()
	}

	if t.It.Down == 1 || t.It.Down > 15 {
		t.SoftDrop()
	}

	if t.It.Up == 1 {
		t.rotate(Clockwise)
	}

	if t.It.Space == 1 {
		for t.SoftDrop() {
		}

		t.lockPiece()
		t.spawnNextPiece()
	}

	if t.lockTimerStarted && (t.CurrentPiece.X != initPiece.X ||
		t.CurrentPiece.Y != initPiece.Y ||
		t.CurrentPiece.State != initPiece.State) {
		t.stopLockTimer()
	}
}

func (t *Tetris) updateInputTimers(is InputState) {
	if is.Left {
		t.It.Left++
	} else {
		t.It.Left = 0
	}

	if is.Right {
		t.It.Right++
	} else {
		t.It.Right = 0
	}

	if is.Down {
		t.It.Down++
	} else {
		t.It.Down = 0
	}

	if is.Up {
		t.It.Up++
	} else {
		t.It.Up = 0
	}

	if is.Space {
		t.It.Space++
	} else {
		t.It.Space = 0
	}
}

const (
	Clockwise        int = 1
	CounterClockwise int = 3
)

func (t *Tetris) rotate(direction int) bool {
	p := &t.CurrentPiece

	cc := 0
	if direction == CounterClockwise {
		cc = 1
	}

	var wallKicks *[5]Vec2
	if p.TetrominoType == 0 {
		wallKicks = &IWallKicks[2*p.State+cc]
	} else {
		wallKicks = &WallKicks[2*p.State+cc]
	}

	state := (p.State + direction) % 4

	fmt.Println(state)

	return t.tryRotations(p, state, wallKicks)
}

func (t *Tetris) tryRotations(p *Piece, state int, table *[5]Vec2) bool {
	oldX := p.X
	oldY := p.Y
	oldState := p.State

	p.State = state
	for _, trans := range table {
		p.X += trans.X
		p.Y += trans.Y

		if !t.Collides(*p) {
			return true
		}

		p.X = oldX
		p.Y = oldY
	}

	p.State = oldState

	return false
}

func (t *Tetris) startLockTimer() {
	if !t.lockTimerStarted {
		t.lockTimerStarted = true
		t.lockTimer = 0
	}
}

func (t *Tetris) stopLockTimer() {
	t.lockTimerStarted = false
	t.lockTimer = 0
}

func (t *Tetris) updateLockTimer() {
	if t.lockTimerStarted {
		fmt.Println(t.lockTimer)
		t.lockTimer += Timestep
		if t.lockTimer >= t.LockTime {
			t.lockPiece()
		}
	}
}

func (t *Tetris) spawnNextPiece() {
	t.CurrentPiece.TetrominoType = t.PieceQueue[t.NextIndex]
	t.NextIndex = (t.NextIndex + 1) % 14

	t.CurrentPiece.X = 3
	t.CurrentPiece.Y = 0
	t.CurrentPiece.State = 0
}

func (t *Tetris) lockPiece() {
	p := t.CurrentPiece
	mask := Tetrominos[p.TetrominoType][p.State]
	color := TetrominoColors[p.TetrominoType]

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4
			x := p.X + i%4

			j := x + y*10

			t.Board[j] = color
		}
	}
	t.lockTimerStarted = false

	t.cleanLine()
	t.checkLoss()
	t.spawnNextPiece()
}

func (t *Tetris) cleanLine(){
	flag := false
	for i := 0; i < 220; i += 10 {
		flag = true
		for j := 0; j < 10; j++ {
			if t.Board[i+j] == 0 {
				flag = false
			}
		}
		if flag == true {
			for j := i+9; j >= 10; j-- {
				t.Board[j] = t.Board[j-10]
			}
		}
	}
}

func (t *Tetris) checkLoss(){
	for i := 0; i < 20; i++{
		if t.Board[i] != Empty {
			t.FlagLoss = true
			return
		}
	}
	t.FlagLoss = false
}

func (t *Tetris) score(){

}

func (t *Tetris) Collides(p Piece) bool {
	mask := Tetrominos[p.TetrominoType][p.State]

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4
			x := p.X + i%4

			if x < 0 || x >= 10 || y < 0 || y >= 22 {
				return true
			}

			j := x + y*10
			if t.Board[j] != Empty {
				return true
			}
		}
	}

	return false
}

func (t *Tetris) moveRight() {
	t.CurrentPiece.X++

	if t.Collides(t.CurrentPiece) {
		t.CurrentPiece.X--
	}
}

func (t *Tetris) moveLeft() {
	t.CurrentPiece.X--

	if t.Collides(t.CurrentPiece) {
		t.CurrentPiece.X++
	}
}

func (t *Tetris) SoftDrop() bool {
	t.CurrentPiece.Y++

	if t.Collides(t.CurrentPiece) {
		t.CurrentPiece.Y--
		return false
	}

	return true
}

func CreateTetris() *Tetris {
	var t Tetris

	// Llena la cola de piezas con 0...7 dos veces
	for i := 0; i < 14; i++ {
		t.PieceQueue[i] = i % 7
	}

	for i := 6; i > 0; i-- {
		j := rand.Intn(7)
		k := 7 + rand.Intn(7)

		t.PieceQueue[i], t.PieceQueue[j] = t.PieceQueue[j], t.PieceQueue[i]
		t.PieceQueue[i+7], t.PieceQueue[k] = t.PieceQueue[j], t.PieceQueue[i+7]
	}

	t.NextIndex = 0
	t.Gravity = 0.1
	t.LockTime = 500

	t.spawnNextPiece()

	return &t
}
