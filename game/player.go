package game

import (
	"my-game/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image    *ebiten.Image
	position Vector
	x        float64
	y        float64
	width    float64
	height   float64
}

func NewPlayer() *Player {
	image := assets.PlayerSprite

	bounds := image.Bounds()
	width := float64(bounds.Dx())
	height := float64(bounds.Dy())

	// Define a posição inicial do jogador no centro da tela.
	position := Vector{
		X: (screenWidth / 2) - (width / 2),
		Y: 500, // Ajuste conforme necessário.
	}
	return &Player{
		image:    image,
		position: position,
		x:        position.X,
		y:        position.Y,
		width:    width,
		height:   height,
	}
}

func (p *Player) Update() {
	// Controle do jogador, ajustado de acordo com teclas pressionadas
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.x -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.x += 2
	}

	// Garantir que o jogador não saia da tela
	if p.x < 0 {
		p.x = 0
	}
	if p.x+p.width > screenWidth {
		p.x = screenWidth - p.width
	}
	p.position.X = p.x // Atualiza a posição para que possa ser usada na função Draw.
}

func (p *Player) Position() (float64, float64) {
	return p.x, p.y
}

func (p *Player) Draw(screen *ebiten.Image) {
	// Opções de desenho, para definir a posição da imagem na tela
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.x, p.y) // Posiciona a imagem nas coordenadas x e y do jogador

	// Desenha a imagem do jogador na tela
	screen.DrawImage(p.image, op)
}
