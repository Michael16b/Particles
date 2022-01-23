package particles

import (
	"project-particles/config"
)

// Update met à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) et otimise la mémoire utilisée par le système de particules
// à chaque pas de temps. Elle est appelée exactement 60 fois par seconde (de manière régulière)
// par la fonction principale du projet.

func (s *System) Update() {
	s.Rate += config.General.SpawnRate
	for i := 0; i < len(s.Content); i++ {
		if !s.Content[i].Isdead {
			if s.Content[i].Life <= config.General.Life {
				s.Content[i] = s.Content[i].update()
			}
			if config.General.Wall {
				s.Content[i] = s.Content[i].wall()
			}

			if s.Content[i].PositionX > float64(config.General.WindowSizeX+config.General.Margin) ||
				s.Content[i].PositionX < 0-float64(config.General.Margin) ||
				s.Content[i].PositionY > float64(config.General.WindowSizeY+config.General.Margin) ||
				s.Content[i].PositionY < 0-float64(config.General.Margin) ||
				s.Content[i].Life > config.General.Life {
				s.Content[i].Isdead = true
			}
			if s.Content[i].Isdead {
				s.Content[i], s.Content[s.Pos] = s.Content[s.Pos], s.Content[i]
				s.Content[s.Pos].Opacity = 0
				s.Pos--
				i--
			}
		}
	}
	for s.Rate >= 1 {
		s.add()
		s.Rate--
	}
}
