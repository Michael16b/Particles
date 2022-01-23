package particles

import("project-particles/config")

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// Dans sa version actuelle, cette fonction renvoie le nombre de particules indiqué par InitNumParticles,
// celles-ci ayant les caractéristiques définies dans config.json.

func NewSystem() System {
	var system System
	system.Pos = len(system.Content)-1
	for i:=0; i<config.General.InitNumParticles; i++ {
		system.add()
	}
	return system
}
