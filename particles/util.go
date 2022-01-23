package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"github.com/hajimehoshi/ebiten/v2"
)

// Cette fonction renvoie la position en x, y d'une particule en fonction des extensions activées dans le config.json.
func getPosition() (posX, posY float64) {
	var mx, my int
	if config.General.RandomSpawn {
		posX = float64(config.General.WindowSizeX)
		posY = float64(config.General.WindowSizeY)
	} else {
		posX = float64(config.General.SpawnX)
		posY = float64(config.General.SpawnY)
	}
	if config.General.Mouse {
		mx, my = ebiten.CursorPosition()
		posX = float64(mx)
		posY = float64(my)
	}
	if config.General.SpawnCircle {
		var rayon float64 = config.General.CircleRadius
		var t []float64 = []float64{1, -1}
		var cx float64 = (rand.Float64() * (rayon * 2)) - rayon
		var cy float64 = t[rand.Intn(2)] * math.Sqrt(rayon*rayon-cx*cx)
		posX += cx
		posY += cy
	}
	return posX, posY
}

// Cette fonction renvoie une vitesse aléatoire en x, y d'une particule.
func getSpeed() (vitX, vitY float64) {
	vitX = rand.Float64()*config.General.Speed - config.General.Speed/2
	vitY = rand.Float64()*config.General.Speed - config.General.Speed/2
	return vitX, vitY
}

// Cette fonction renvoie la couleur d'une particule en fonction de l'activation du randomColor (ou blanche par défaut).
func getColor() (Rouge, Vert, Bleu float64) {
	if config.General.RandomColor {
		Rouge = rand.Float64()
		Vert = rand.Float64()
		Bleu = rand.Float64()
	} else {
		Rouge = 1
		Vert = 1
		Bleu = 1
	}
	return Rouge, Vert, Bleu
}

// Cette fonction renvoie l'opacité d'une particule en fonction de la valeur donnée à config.General.Opacity (par défaut elle est opaque).
func getOpacity() (op int) {
	if config.General.Opacity == 2 {
		op = 0
	} else {
		op = 1
	}
	return op
}

// Cette fonction met à jour les différentes caractéristiques d'une particule à son appel.
func (p Particle) update() Particle {
	p.PositionX += p.SpeedX
	p.PositionY += p.SpeedY
	p.SpeedY += config.General.Gravity
	p.Rotation += config.General.Rotation
	p = update_opacity(p)
	p = update_scale(p)
	if config.General.Color {
		p = update_color(p)
	}
	p.Life++
	return p
}

// Cette fonction fait évoluer l'opacité de la particule en fonction de sa vie.
func update_opacity(p Particle) Particle {
	if config.General.Opacity == 1 {
		p.Opacity -= 1 / float64(config.General.Life)
	} else if config.General.Opacity == 2 {
		p.Opacity += 1 / float64(config.General.Life)
	}
	return p
}

// Cette fonction fait évoluer la taille de la particule en fonction de config.General.ScaleVariation.
func update_scale(p Particle) Particle {
	p.ScaleX = p.ScaleX * config.General.ScaleVariation
	p.ScaleY = p.ScaleY * config.General.ScaleVariation
	return p
}

// Cette fonction fait évoluer la couleur de la particule en fonction de sa vie.
func update_color(p Particle) Particle {
	p.ColorGreen -= 1 / float64(config.General.Life)
	p.ColorRed -= 1 / float64(config.General.Life)
	return p
}

// Cette fonction fait "rebondir" les particules quand celles-ci touchent un ou plusieurs bords de l'écran.
func (p Particle) wall() Particle {
	if p.PositionX+p.ScaleX*10 >= float64(config.General.WindowSizeX) || p.PositionX <= 0 {
		p.SpeedX *= -1
	}
	if p.PositionY+p.ScaleY*10 >= float64(config.General.WindowSizeY) || p.PositionY <= 0 {
		p.SpeedY *= -1
	}
	return p
}
