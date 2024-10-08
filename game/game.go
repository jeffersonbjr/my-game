package game

import (
	"math/rand"
	"my-game/assets"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player  *Player
	meteors []Meteor
	lives   int
	score   int
}

type Meteor struct {
	x     float64
	y     float64
	speed float64
	image *ebiten.Image
}

// Função que inicializa um novo jogo
func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	return &Game{
		player:  NewPlayer(),
		meteors: []Meteor{},
		lives:   3,
		score:   0,
	}
}

// Responsavel por atualizar a logica do jogo
func (g *Game) Update() error {
	g.player.Update()
	g.updateMeteors()
	return nil
}

// Responsavel por desenhar objetos na tela
func (g *Game) updateMeteors() {
	for i := 0; i < len(g.meteors); i++ {
		meteor := &g.meteors[i]
		meteor.y += meteor.speed

		if meteor.y > screenHeight {
			g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
			i--
			continue
		}

		if g.checkCollision(meteor) {
			g.lives--
			g.meteors = append(g.meteors[:i], g.meteors[i+1:]...)
			i--
		}
	}

	if rand.Float64() < 0.01 {
		g.spawnMeteor()
	}
}

func (g *Game) checkCollision(meteor *Meteor) bool {
	playerX, playerY := g.player.Position()

	playerWidth := g.player.width
	playerHeight := g.player.height
	meteorWidth := 64.0
	meteorHeight := 64.0

	if meteor.x < playerX+playerWidth && meteor.x+meteorWidth > playerX &&
		meteor.y < playerY+playerHeight && meteor.y+meteorHeight > playerY {
		return true
	}
	return false
}

func (g *Game) spawnMeteor() {
	newMeteor := Meteor{
		x:     rand.Float64() * screenWidth,
		y:     0,
		speed: 2 + rand.Float64()*2,
		image: assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))], // Escolhe um meteoro aleatoriamente
	}
	g.meteors = append(g.meteors, newMeteor)
}

// Desenha todos os meteoros na tela
func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, meteor := range g.meteors {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(meteor.x, meteor.y)
		screen.DrawImage(meteor.image, op) // Desenha a imagem do meteoro
	}
	ebitenutil.DebugPrint(screen, "Lives: "+strconv.Itoa(g.lives))
	ebitenutil.DebugPrintAt(screen, "Score: "+strconv.Itoa(g.score), 10, 20)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
