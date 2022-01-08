package particles

import (
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	s.Rate = config.General.SpawnRate + s.Reste
	//mise à jour de la position de chaque particule en fonction de la vitesse qui leur est donnée
	for i := 0; i < len(s.Content); i++ {
		s.Content[i].PositionX += s.Content[i].VitesseX
		s.Content[i].PositionY += s.Content[i].VitesseY
	}
	//génération du nombre de particules voulues (indiqué par le SpawnRate défini dans config.json)
	for s.Rate >= 1 {
		s.Add()
		s.Rate--
	}
	s.Reste = s.Rate
}
