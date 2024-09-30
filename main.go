package main

import (
	"my-game/game"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

// Responsavel por atualizar a logica do jogo
func (g *Game) Update() error {
	return nil
}

// Responsavel por desenhar objetos na tela
func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {

	g := game.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)

	}

}
