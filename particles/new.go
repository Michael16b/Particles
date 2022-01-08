package particles

import "project-particles/config"

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.

func NewSystem() System {
	var system System
	for i := 0; i < config.General.InitNumParticles; i++ {
		//ajout du nombre de paticules voulu (défini par InitNumParticles dans le fichier config.json) dans le système créé
		system.Add()
	}
	return system
}
