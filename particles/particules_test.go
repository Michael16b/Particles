package particles

import (
	"project-particles/config"
	"testing"
)

/*
Cette fonction teste si le nombre initial de particules choisi est bien le même que le nombre de particules créées dans le système
et donc que la fonction add créée bien les particules.
*/
func TestNumParticles(t *testing.T) {
	config.General.InitNumParticles = 5
	var system System = NewSystem()
	if len(system.Content) != config.General.InitNumParticles {
		t.Fail()
	}
}

/*
Ces fonctions testent si les particules créées dans le système apparaissent bien au point d'apparition souhaité.
*/
func TestRandomSpawnFalse1(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 100
	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		if content[i].PositionX-float64(config.General.SpawnX) > 0.0001 || content[i].PositionY-float64(config.General.SpawnY) > 0.0001 {
			t.Fail()
		}
	}
}

func TestRandomSpawnFalse2(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.SpawnX = 5476
	config.General.SpawnY = 7586
	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		if content[i].PositionX-float64(config.General.SpawnX) > 0.0001 || content[i].PositionY-float64(config.General.SpawnY) > 0.0001 {
			t.Fail()
		}
	}
}

/*
Ces fonctions testent si les particules apparaissent bien de manière aléatoire en vérifiant
que les particules apparaissent dans l'écran.
*/
func TestRandomSpawnTrue1(t *testing.T) {
	config.General.RandomSpawn = true
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		if content[i].PositionX < 0 ||
			content[i].PositionY < 0 ||
			content[i].PositionX-10*content[i].ScaleX > float64(config.General.WindowSizeX) ||
			content[i].PositionY-10*content[i].ScaleY > float64(config.General.WindowSizeY) {
			t.Fail()
		}
	}
}

func TestRandomSpawnTrue2(t *testing.T) {
	config.General.RandomSpawn = true
	config.General.WindowSizeX = 1920
	config.General.WindowSizeY = 1080
	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		if content[i].PositionX < 0 ||
			content[i].PositionY < 0 ||
			content[i].PositionX-10*content[i].ScaleX > float64(config.General.WindowSizeX) ||
			content[i].PositionY-10*content[i].ScaleY > float64(config.General.WindowSizeY) {
			t.Fail()
		}
	}
}

/*
Cette fonction teste si les différentes particules se déplacent à des vitesses aléatoires en vérifiant que les vitesses
sont bien dans l'intervalle donné.
*/
func TestRandomSpeed(t *testing.T) {

	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		if content[i].VitesseX > 5 || content[i].VitesseX < -5 || content[i].VitesseY > 5 || content[i].VitesseY < -5 {
			t.Fail()
		}
	}
}

/*
Ces fonctions vérifient s'il y a bien le bon nombre de particules créées à la fin de la dernière mise à jour souhaitée.
*/
func TestSpawnRate1(t *testing.T) {
	config.General.SpawnRate = 0.5
	config.General.InitNumParticles = 10
	content_after := NewSystem()
	var update_value float64 = 10
	// Boucle de mises à jour du système
	for i := 0; i < int(update_value); i++ {
		content_after.Update()
	}
	if len(content_after.Content) != config.General.InitNumParticles+int(config.General.SpawnRate*update_value) {
		t.Fail()
	}
}

func TestSpawnRate2(t *testing.T) {
	config.General.SpawnRate = 1.3
	config.General.InitNumParticles = 10
	content_after := NewSystem()
	var update_value float64 = 10
	// Boucle de mises à jour du système
	for i := 0; i < int(update_value); i++ {
		content_after.Update()
	}
	if len(content_after.Content) != config.General.InitNumParticles+int(config.General.SpawnRate*update_value) {
		t.Fail()
	}
}

func TestSpawnRate3(t *testing.T) {
	config.General.SpawnRate = 73
	config.General.InitNumParticles = 10
	content_after := NewSystem()
	var update_value float64 = 10
	// Boucle de mises à jour du système
	for i := 0; i < int(update_value); i++ {
		content_after.Update()
	}
	if len(content_after.Content) != config.General.InitNumParticles+int(config.General.SpawnRate*update_value) {
		t.Fail()
	}
}
