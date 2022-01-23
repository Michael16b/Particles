package particles

// System définit un système de particules.
type System struct {
	Content []Particle
	Rate    float64
	Pos     int
}

// Particle définit une particule.
// Elle possède une position, une rotation, une taille, une couleur, une
// opacité, une vitesse, une vie et une indication sur l'état de sa vie.
type Particle struct {
	PositionX, PositionY            float64
	Rotation                        float64
	ScaleX, ScaleY                  float64
	ColorRed, ColorGreen, ColorBlue float64
	Opacity                         float64
	SpeedX                          float64
	SpeedY                          float64
	Life                            int
	Isdead                          bool
}
