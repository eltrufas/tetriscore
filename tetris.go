package tetriscore

import (
	"math/rand"
)

type Block int

// Definir las piezas como tipo bloque
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

// Tiempo que tarda en actualizar
const Timestep float32 = 16.66

// Definición de piezas
// |
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

// J
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

// L
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

// []
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

// S
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

// T
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

// Z
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

// Arreglo con todas las piezas
var Tetrominos [7]*[4][16]int = [7]*[4][16]int{
	&IStates,
	&OStates,
	&TStates,
	&SStates,
	&ZStates,
	&JStates,
	&LStates,
}

// Colores de las piezas
var TetrominoColors [7]Block = [7]Block{
	Cyan,
	Yellow,
	Purple,
	Green,
	Red,
	Blue,
	Orange,
}

// Vector para rotación
type Vec2 struct {
	X, Y int
}

// Estructura de una pieza
type Piece struct {
	X             int
	Y             int
	State         int
	TetrominoType int
}

/*
type InputState struct {
	Left, Right, Up, Down, Space, Shift, Enter bool
}

type InputTimers struct {
	Left, Right, Up, Down, Space, Shift, Enter int
}
*/

// Representa la entrada en una actualización
type InputState uint32

// Representa el tiempo que han estado activas las entradas
type InputTimers [32]int

// Posibles inputs
const (
	Left InputState = iota
	Right
	Up
	Down
	Space
	Shift
	Zed
)

// Vectores para los wall kicks en rotaciones
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

// Definición del tetris
type Tetris struct {
	Board            [220]Block //10*22
	CurrentPiece     Piece
	Score            int
	PieceQueue       [14]int // lista de piezas revueltas
	NextIndex        int
	Gravity          float32
	DropTimer        float32 // Tiempo que falta para avanzar un pixel
	LockTime         float32 // Tiempo para inmovilizar una pieza
	lockTimerStarted bool    // Indica si se inició el contador de movilización
	lockTimer        float32 // Contador
	It               InputTimers
	FlagLoss         bool // Para saber si el jugador perdió
	Level            int
	speedUp          bool
	Tetris           bool
	Held             bool
	HoldPiece        int // Pieza que detiene el jugador
	ClearLines       int // Líneas que se han limpiado bajo la logica de tetris
  Tspin            bool // Saber si se hizo un wallkick para T-spin
  Combo            int // Cuantos combos lleva
}

// Función que regresa la siguiente pieza
func (t *Tetris) NextPiece() int {
	return t.PieceQueue[t.NextIndex]
}

// Revolver las piezas
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

// Actualizar el tablero
func (t *Tetris) Update(is InputState) {
	t.updateInputTimers(is)
	t.checkHold()
	t.applyMovement()
	t.updateLockTimer()

	for t.DropTimer >= 1. {
		t.DropTimer--
		if !t.SoftDrop(&t.CurrentPiece) {
			t.startLockTimer()
		}
	}

	var multiplier float32 = 1.0
	if t.speedUp {
		multiplier = 4.0
	}

	t.DropTimer += multiplier * t.Gravity
}

// Guardar una pieza
func (t *Tetris) checkHold() {
	if t.It[Shift] == 1 {
		if t.Held {
			return
		}

		t.Held = true

		oldHold := t.HoldPiece
		t.HoldPiece = t.CurrentPiece.TetrominoType

		t.stopLockTimer()

		if oldHold >= 0 {
			t.CurrentPiece.TetrominoType = oldHold
			t.resetPiece()
		} else {
			t.spawnNextPiece()
		}
	}
}

func shouldMove(timer int) bool {
	return timer == 1 || (timer > 5 && timer%5 == 0)
}

// Función que hace inputs
func (t *Tetris) applyMovement() {
	initPiece := t.CurrentPiece

	if shouldMove(t.It[Left]) {
		t.moveLeft()
	}

	if shouldMove(t.It[Right]) {
		t.moveRight()
	}

	t.speedUp = t.It[Down] > 0

	if t.It[Up] == 1 {
		t.rotate(Clockwise)
	}

	if t.It[Zed] == 1 {
		t.rotate(CounterClockwise)
	}

	if t.It[Space] == 1 {
		for t.SoftDrop(&t.CurrentPiece) {
      t.Score += 1
		}

		t.lockPiece()
	}

	if t.lockTimerStarted && (t.CurrentPiece.X != initPiece.X ||
		t.CurrentPiece.Y != initPiece.Y ||
		t.CurrentPiece.State != initPiece.State) {
		t.stopLockTimer()
	}
}

// Actualizar el tiempo de cada input
func (t *Tetris) updateInputTimers(is InputState) {
	var i InputState
	for i = 0; i < 32; i++ {
		var m InputState = 1 << i
		if (is & m) > 0 {
			t.It[i]++
		} else {
			t.It[i] = 0
		}
	}
}

func (t *Tetris) GhostPiece() Piece {
	gp := t.CurrentPiece
	for t.SoftDrop(&gp) {
	}

	return gp
}

const (
	Clockwise        int = 1
	CounterClockwise int = 3
)

// Rotar una pieza
func (t *Tetris) rotate(direction int) bool {
	p := &t.CurrentPiece

	cc := 0
	if direction == CounterClockwise {
		cc = 1
	}

	// Evita que las piezas se salgan del tablero utilizando wall kick
	var wallKicks *[5]Vec2
	if p.TetrominoType == 0 {
		wallKicks = &IWallKicks[2*p.State+cc]
	} else {
		wallKicks = &WallKicks[2*p.State+cc]
	}

	state := (p.State + direction) % 4

	return t.tryRotations(p, state, wallKicks)
}

// Intenta realizar las rotaciones posibles
func (t *Tetris) tryRotations(p *Piece, state int, table *[5]Vec2) bool {
	oldX := p.X
	oldY := p.Y
	oldState := p.State

	p.State = state
	for i, trans := range table {
		p.X += trans.X
		p.Y += trans.Y
    if i != 0 && p.TetrominoType == 2{
      t.Tspin = true
    } else {
      t.Tspin = false
    }

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
		t.lockTimer += Timestep
		if t.lockTimer >= t.LockTime {
			t.lockPiece()
		}
	}
}

// Inicializa la siguiente pieza
func (t *Tetris) spawnNextPiece() {
	t.CurrentPiece.TetrominoType = t.PieceQueue[t.NextIndex]
	t.NextIndex = (t.NextIndex + 1) % 14
	if t.NextIndex%7 == 0 {
		t.ShuffleQueue()
	}

	t.resetPiece()
	if t.Collides(t.CurrentPiece) {
		t.FlagLoss = true
	}
}

func (t *Tetris) resetPiece() {
	t.CurrentPiece.X = 3
	t.CurrentPiece.Y = 0
	t.CurrentPiece.State = 0

}

// Sella la pieza en la estructura
func (t *Tetris) lockPiece() {
	p := t.CurrentPiece
	mask := Tetrominos[p.TetrominoType][p.State]
	color := TetrominoColors[p.TetrominoType]

	t.FlagLoss = true

	for i := 0; i < 16; i++ {
		if mask[i] == 1 {
			y := p.Y + i/4
			x := p.X + i%4

			j := x + y*10

			if j >= 20 {
				t.FlagLoss = false
			}

			t.Board[j] = color
		}
	}
	t.lockTimerStarted = false
	t.Held = false

	t.cleanLine()
	t.spawnNextPiece()
}

// Limpia las líneas cuando se completan
func (t *Tetris) cleanLine() {
	var flag bool
	lines := 0
	for i := 0; i < 220; i += 10 {
		flag = true
		for j := 0; j < 10; j++ {
			if t.Board[i+j] == 0 {
				flag = false
			}
		}
		if flag == true {
			lines++
			for j := i + 9; j >= 10; j-- {
				t.Board[j] = t.Board[j-10]
			}
		}
	}
  if lines == 0 {
    t.Combo = 0
  } else if t.Combo < 20 {
    t.Combo = lines
    if t.Combo > 20 {
      t.Combo = 20
    }
  }
	t.score(lines)
	t.upgradeLevel(lines)
}

// Actualiza el nível en el que está el jugador
func (t *Tetris) upgradeLevel(lines int) {
	switch lines {
	case 1:
    if t.Tspin {
      t.ClearLines += 8
    } else {
		  t.ClearLines++
    }
		break
	case 2:
    if t.Tspin {
      t.ClearLines += 12
    } else {
      t.ClearLines += 3
    }
		break
	case 3:
    if t.Tspin {
      t.ClearLines += 16
    } else {
      t.ClearLines += 5
  	}
		break
	case 4:
		t.ClearLines += 8
		break
	}
	// Los níveles aumentan level * 5 líneas completadas, nível max 15
	if t.ClearLines >= 5*t.Level && t.Level <= 15 {
		t.Level++
		// Aumenta la velocidad de caída de las piezas
		t.Gravity *= 1.08
	}
}

// Actualiza el puntaje del jugador
func (t *Tetris) score(lines int){
  t.Score += 50 * t.Combo * t.Level
  switch lines {
    case 0:
      if t.Tspin {
        t.Score += t.Level * 400
      } else {
        t.Score += 0
      }
      break
  	// Una línea vale 100
  	case 1:
      if t.Tspin {
        t.Score += t.Level * 800 + t.Level * 400
      } else {
  		  t.Score += t.Level * 100
      }
      break
  	// 2 líneas valen 300
  	case 2:
      if t.Tspin {
        t.Score += t.Level * 1200 + t.Level * 400
      } else {
        t.Score += t.Level * 300
      }
      break
  		// 3 líneas valen 500
  	case 3:
      if t.Tspin {
        t.Score += t.Level * 1600 + t.Level * 400
      } else {
  		  t.Score += t.Level * 500
      }
      break
  		// 4 líneas valen 800
  	case 4:
  		if t.Tetris == false {
  			t.Score += t.Level * 800
  		} else {
  			// Si la jugada anterior se limpiaron 4 líneas y esta jugada también
  			// la puntuación se multiplica por 1200
  			t.Score += t.Level * 1200
  		}
	}
}

// Revisa si la pieza actual puede inmovilizarse en ese lugar
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
	} else {
    t.Tspin = false
  }
}

func (t *Tetris) moveLeft() {
	t.CurrentPiece.X--

	if t.Collides(t.CurrentPiece) {
		t.CurrentPiece.X++
	} else {
    t.Tspin = false
  }
}

// Caída normal
func (t *Tetris) SoftDrop(p *Piece) bool {
	p.Y++

	if t.Collides(*p) {
		p.Y--
		return false
	}
  t.Tspin = false
  t.Score += 1
	return true
}

// Crea el juego
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
	// Valores de inicialización
	t.NextIndex = 0
	t.Gravity = 0.07
	t.LockTime = 500
	t.Level = 1
	t.HoldPiece = -1

	t.spawnNextPiece()

	return &t
}
