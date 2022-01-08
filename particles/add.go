package particles

import (
	"math/rand"
	"project-particles/config"
)

func (s *System) Add() {
	var posX, posY float64
	//Si la génération aléatoire des particules est activée,
	//alors les particules sont générées aléatoirement dans l'espace défini par la taille de la fenêtre
	if config.General.RandomSpawn {
		posX = float64(rand.Intn(config.General.WindowSizeX))
		posY = float64(rand.Intn(config.General.WindowSizeY))
	} else {
		//Sinon, elles sont générées à un point fixe défini par les valeurs de SpawnX et SpawnY dans config.json
		posX = float64(config.General.SpawnX)
		posY = float64(config.General.SpawnY)
	}
	//création d'une particule avec une structure de type Particle
	var p Particle = Particle{
		PositionX: posX,
		PositionY: posY,
		ScaleX:    1, ScaleY: 1,
		ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
		Opacity:  1,
		VitesseX: rand.Float64()*config.General.Vitesse - (config.General.Vitesse)/2,
		VitesseY: rand.Float64()*config.General.Vitesse - (config.General.Vitesse)/2,
	}
	//ajout de la particule au système
	s.Content = append(s.Content, p)
}
