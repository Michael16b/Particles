package particles

import (
	"project-particles/config"
)

// La fonction add ajoute une particule à un système s.
func (s *System) add() {
	var posX, posY float64 = getPosition()
	var vitX, vitY float64 = getSpeed()
	var Rouge, Vert, Bleu float64 = getColor()
	var op int = getOpacity()

	var p Particle = Particle{
		PositionX: posX,
		PositionY: posY,
		ScaleX:    config.General.SpawnScale, ScaleY: config.General.SpawnScale,
		ColorRed: Rouge, ColorGreen: Vert, ColorBlue: Bleu,
		Opacity: float64(op),
		SpeedX:  vitX,
		SpeedY:  vitY,
		Life:    0,
	}

	if s.Pos == len(s.Content)-1 {
		s.Content = append(s.Content, p)
		s.Pos++
	} else {
		s.Content[s.Pos+1] = p
		s.Pos++
	}
}
