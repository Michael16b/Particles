package config

// Config définit les champs qu'on peut trouver dans un fichier de config.
// Dans le fichier les champs doivent porter le même nom que dans le type si
// dessous, y compris les majuscules. Tous les champs doivent obligatoirement
// commencer par des majuscules, sinon il ne sera pas possible de récupérer
// leurs valeurs depuis le fichier de config.
type Config struct {
	WindowTitle              string
	WindowSizeX, WindowSizeY int
	ParticleImage            string
	Debug                    bool
	InitNumParticles         int
	RandomSpawn              bool
	SpawnX, SpawnY           int
	SpawnRate                float64
	Gravity                  float64
	Margin                   int
	Life                     int
	Rotation                 float64
	Circle                   bool
	RandomColor              bool
	Speed                    float64
	ScaleVariation           float64
	SpawnScale               float64
	Wall                     bool
	Opacity                  int
	Mouse                    bool
	SpawnCircle              bool
	CircleRadius             float64
	Color                    bool
}

var General Config
