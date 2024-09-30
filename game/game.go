package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	player *Player
}

func NewGame() *Game {
	player := NewPlayer()
	return &Game{
		player: player,
	}
}

// Responsavel por atualizar a logica do jogo
func (g *Game) Update() error {
	return nil
}

// Responsavel por desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {

	g.player.Draw(screen)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
