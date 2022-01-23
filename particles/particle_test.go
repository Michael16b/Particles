package particles

import (
	"math"
	"math/rand"
	"project-particles/config"
	"testing"
)

// Vérification que les particules générées quand config.General.RandomSpawn
// vaut true sont bien toutes à l'intérieur de l'écran
func TestNewParticleInScreen(t *testing.T) {
	config.General.RandomSpawn = true
	config.General.InitNumParticles = 5
	var content []Particle = NewSystem().Content
	for i := 0; i < len(content); i++ {
		config.General.WindowSizeX = rand.Intn(600) + 200
		config.General.WindowSizeY = rand.Intn(600) + 400
		p := content[i]
		if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX) ||
			p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY) {
			t.Fail()
		}
	}
}

// Vérification que les particules générées quand config.General.RandomSpawn
// vaut false sont bien toutes à la position demandée, c'est-à-dire aux
// coordonnées (config.General.SpawnX, config.General.SpawnY)
func TestNewParticleAtSpawnPoint(t *testing.T) {
	config.General.RandomSpawn = false
	config.General.InitNumParticles = 5
	var content []Particle = NewSystem().Content

	for i := 0; i < len(content); i++ {
		var p Particle = content[i]
		if int(math.Round(p.PositionX)) != config.General.SpawnX ||
			int(math.Round(p.PositionY)) != config.General.SpawnY {
			t.Fail()
		}
	}
}

// Vérification que la fonction update met bien à jour la position des particules
// sur lesquelles on l'utilise
func TestParticleUpdate(t *testing.T) {
	config.General.InitNumParticles = 5
	var content []Particle = NewSystem().Content
	for i := 0; i < len(content); i++ {
		var p Particle = content[i]
		goalX := p.PositionX + p.SpeedX
		goalY := p.PositionY + p.SpeedY
		p.update()
		if p.PositionX != goalX || p.PositionY != goalY {
			t.Fail()
		}
	}
}

//Cette fonction vérifie que les différentes particules se déplacent à des vitesses aléatoires en vérifiant que les vitesses
// sont bien dans l'intervalle donné.
func TestRandomSpeed(t *testing.T) {
	config.General.InitNumParticles = 5
	var content []Particle = NewSystem().Content
	for i := 0; i < len(content); i++ {
		if content[i].SpeedX > 5 || content[i].SpeedX < -5 || content[i].SpeedY > 5 || content[i].SpeedY < -5 {
			t.Fail()
		}
	}
}

// Vérification que les particules sont mortes après n mises à jour du système.
func Test_Dead_Particule(t *testing.T) {
	for i:=0; i<100; i++ {
		config.General.InitNumParticles = 5
		var s System = NewSystem()
		config.General.Life = rand.Intn(100)
		config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX+700, config.General.WindowSizeY+400
		var content []Particle = s.Content
		for n := 0; n < config.General.Life+1; n++ {
			s.Update()
		}
		for i := 0; i < len(content); i++ {
			if content[i].Isdead != true {
				t.Fail()
			}
		}
	}
}

// Vérification que lorsque des particules sont à l'extérieur de l'écran elles sont supprimées.
func Test_Outside_the_screen(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 300
	config.General.Wall = false
	config.General.Life = 10
	config.General.Margin = 10
	config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX+config.General.Margin+10, config.General.WindowSizeY+config.General.Margin+10
	var s System = NewSystem()
	var content []Particle = s.Content
	s.Update()
	for i := 0; i < len(content); i++ {
		if content[i].Isdead != true {
			t.Fail()
		}
	}
}

// Vérification que les particules mortes sont bien mises à la fin du tableau de particules.
func Test_memory_opti(t *testing.T) {
	for i := 0; i < 100; i++ {
		config.General.InitNumParticles = rand.Intn(10) + 1
		config.General.WindowSizeX = rand.Intn(800)
		config.General.WindowSizeY = rand.Intn(800)
		config.General.Margin = 0
		var tab []int = []int{-1, 1}
		var s System = NewSystem()
		s.Content[0].PositionX, s.Content[0].PositionY = float64((config.General.WindowSizeX+rand.Intn(50))*tab[rand.Intn(len(tab))]), float64((config.General.WindowSizeY+rand.Intn(50))*tab[rand.Intn(len(tab))])
		s.Update()
		if !s.Content[len(s.Content)-1].Isdead && s.Content[len(s.Content)-1].Opacity != 0 {
			t.Fail()
		}
	}
}

// Vérification que le générateur de particules change de position (en "simulant" une souris)
// en fonction de la taille de l'écran.
func Test_Generator_Displacement(t *testing.T) {
	config.General.InitNumParticles = 5
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 300
	config.General.RandomSpawn = false
	config.General.Mouse = false
	var content []Particle = NewSystem().Content
	content_before := make([]Particle, len(content))
	copy(content_before, content)
	var mx, my float64 = float64(config.General.WindowSizeX), float64(config.General.WindowSizeY)

	for i := 0; i < len(content); i++ {
		content[i].PositionX, content[i].PositionY = mx, my
	}
	if int(math.Round(content[0].PositionX)) == int(content_before[0].PositionX) ||
		int(math.Round(content[0].PositionY)) == int(content_before[0].PositionY) {
		t.Fail()
	}
}

// Vérification que les particules sont bien soumises à la gravité quand celle-ci est activée
func TestGravity(t *testing.T) {
	var update_value int = 15
	config.General.Gravity = 0.5
	config.General.Life = 16
	config.General.InitNumParticles = 1
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 300
	var s System = NewSystem()
	s.Content[0].SpeedY = 1
	s.Content[0].PositionX = float64(config.General.WindowSizeX / 2)
	s.Content[0].PositionY = float64(config.General.WindowSizeY / 2)
	for i := 0; i < update_value; i++ {
		s.Update()
	}
	if int(math.Round(s.Content[0].SpeedY)) != int(math.Round(1+(config.General.Gravity*float64(update_value)))) {
		t.Fail()
	}
}

// Vérification que l'opacité baisse, augmente ou ne change pas selon l'option sélectionnée au cours de la vie de la particule.
func Test_Opacity(t *testing.T) {
	for test := 0; test < 100; test++ {
		var update_value int = 10
		config.General.InitNumParticles = 5
		config.General.WindowSizeX = 800
		config.General.WindowSizeY = 300
		config.General.RandomSpawn = false
		config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX/2, config.General.WindowSizeY/2
		config.General.Life = update_value + 10
		config.General.Opacity = rand.Intn(2)
		var s System = NewSystem()
		var life float64 = (1 / float64(config.General.Life) * 100)
		var content []Particle = s.Content
		if config.General.Opacity == 2 {
			for i := 1; i <= update_value; i++ {
				s.Update()
				if int(math.Round(content[0].Opacity*100)) != int(life) {
					t.Fail()
				}
				life += (1 / float64(config.General.Life) * 100)
			}
		} else if config.General.Opacity == 1 {
			for i := 1; i <= update_value; i++ {
				s.Update()
				if int(math.Round(content[0].Opacity*100)) != 100-int(life) {
					t.Fail()
				}
				life += (1 / float64(config.General.Life) * 100)
			}
		} else {
			for i := 1; i <= update_value; i++ {
				s.Update()
				if int(content[0].Opacity) != 1 {
					t.Fail()
				}
			}
		}
	}
}

// Vérification que les particules changent de couleur
// en passant progressivement du blanc au bleu en fonction de la vie de la particule.
func Test_Variation_Color(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 300
	config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX/2, config.General.WindowSizeY/2
	config.General.Life = 100
	config.General.Color = true
	var particle_before Particle
	var particle Particle
	var s System = NewSystem()
	for i := 0; i < 25; i++ {
		particle_before = s.Content[0]
		s.Update()
		particle = s.Content[0]
		if int(particle_before.ColorRed*100) == int(particle.ColorRed*100) &&
			int(particle_before.ColorGreen*100) == int(particle.ColorGreen*100) {
			t.Fail()
		}
	}
}

// Vérification que les particules rebondissent (leur vitesse est inversée) sur les bords de l'écran.
func Test_Wall(t *testing.T) {
	for i := 0; i < 100; i++ {
		config.General.Life = 10
		config.General.InitNumParticles = 1
		config.General.WindowSizeX = rand.Intn(800) + 10
		config.General.WindowSizeY = rand.Intn(800) + 10
		config.General.Wall = true
		var tabx []int = []int{config.General.WindowSizeX - 5, 5}
		var taby []int = []int{config.General.WindowSizeY - 5, 5}
		config.General.SpawnX = tabx[rand.Intn(2)]
		config.General.SpawnY = taby[rand.Intn(2)]
		var s System = NewSystem()
		if config.General.SpawnY == config.General.WindowSizeY-5 {
			s.Content[0].SpeedY = 6
		} else {
			s.Content[0].SpeedY = -6
		}
		if config.General.SpawnX == config.General.WindowSizeX-5 {
			s.Content[0].SpeedX = 6
		} else {
			s.Content[0].SpeedX = -6
		}
		var SpeedY_before, SpeedX_before float64 = s.Content[0].SpeedY, s.Content[0].SpeedX
		s.Update()
		var p Particle = s.Content[0]
		if ((SpeedY_before < 0 && p.SpeedY < 0) && (SpeedX_before < 0 && p.SpeedX < 0)) ||
			((SpeedY_before < 0 && p.SpeedY < 0) && (SpeedX_before > 0 && p.SpeedX > 0)) ||
			((SpeedY_before > 0 && p.SpeedY > 0) && (SpeedX_before < 0 && p.SpeedX < 0)) ||
			((SpeedY_before > 0 && p.SpeedY > 0) && (SpeedX_before > 0 && p.SpeedX > 0)) {
			t.Fail()
		}
	}
}

// Vérification que les particules tournent lorsqu'on leur applique une rotation.
func Test_Rotation(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 300
	config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX/2, config.General.WindowSizeY/2
	config.General.Life = 100
	config.General.Rotation = 4
	var particle_before Particle
	var s System = NewSystem()
	for i := 0; i < 25; i++ {
		particle_before = s.Content[0]
		s.Update()
		if int(particle_before.Rotation) == int(s.Content[0].Rotation) {
			t.Fail()
		}
	}
}

// Vérification que les particules grossissent ou rapetissent en fonction de config.General.ScaleVariation.
func Test_Scale(t *testing.T) {
	config.General.Life = 10
	config.General.InitNumParticles = 1
	for i := 0; i < 100; i++ {
		config.General.WindowSizeX, config.General.WindowSizeY = 100, 100
		config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX/2, config.General.WindowSizeY/2
		config.General.ScaleVariation = 2
		var s System = NewSystem()
		s.Content[0].ScaleX = rand.NormFloat64()
		s.Content[0].ScaleY = rand.NormFloat64()
		var p Particle = s.Content[0]
		s.Update()
		if int(s.Content[0].ScaleX*1000) != int(p.ScaleX*config.General.ScaleVariation*1000) ||
			int(s.Content[0].ScaleY*1000) != int(p.ScaleY*config.General.ScaleVariation*1000) {
			t.Fail()
		}
	}
}

// Vérification que les particules sont générées sur un cercle de rayon config.General.CircleRadius.
func Test_Circle(t *testing.T) {
	config.General.SpawnCircle = true
	config.General.InitNumParticles = 1
	for i := 0; i < 100; i++ {
		config.General.CircleRadius = rand.Float64() + float64(rand.Intn(100))
		config.General.WindowSizeX, config.General.WindowSizeY = 1000, 1000
		config.General.SpawnX, config.General.SpawnY = config.General.WindowSizeX/2, config.General.WindowSizeY/2
		var s System = NewSystem()
		var CircleRadiusTest float64 = math.Sqrt(math.Pow(float64(config.General.SpawnX)-s.Content[0].PositionX, 2) + math.Pow(float64(config.General.SpawnY)-s.Content[0].PositionY, 2))
		if int(math.Round(CircleRadiusTest)) != int(math.Round(config.General.CircleRadius)) {
			t.Fail()
		}
	}
}
