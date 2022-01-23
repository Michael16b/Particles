package particles

import (
	"math/rand"
	"project-particles/config"
	"testing"
)

// Vérification que NewSystem crée bien un système avec le nombre de particules
// initiales indiqué dans config.General.InitNumParticles
func TestNewSystemNumParticles(t *testing.T) {
	for i := 0; i < 100; i++ {
		config.General.InitNumParticles = rand.Intn(1000)
		if len(NewSystem().Content) != config.General.InitNumParticles {
			t.Fail()
		}
	}
	config.General.InitNumParticles = 0
	if len(NewSystem().Content) != 0 {
		t.Fail()
	}
}

// Vérification que Update ajoute bien le nombre de particules attendues par
// config.General.SpawnRate
func TestSpawnRate(t *testing.T) {
	config.General.SpawnRate = float64(rand.Intn(10))
	config.General.InitNumParticles = 2
	config.General.WindowSizeX = rand.Intn(1000)+250
	config.General.WindowSizeY = rand.Intn(1000)+250
	config.General.SpawnX = config.General.WindowSizeX/2
	config.General.SpawnY = config.General.WindowSizeY/2
	var content_after System = NewSystem()
	var update_value float64 = float64(rand.Intn(30))
	config.General.Life = int(update_value)+10
	var particle_dead int
	for i := 0; i < int(update_value); i++ {
		content_after.Update()
	}
	for t := 0; t < len(content_after.Content); t++ {
		if content_after.Content[t].Isdead {
			particle_dead++
		}
	}
	if len(content_after.Content) != config.General.InitNumParticles+int(config.General.SpawnRate*update_value)-particle_dead {
		t.Fail()
	}
}
