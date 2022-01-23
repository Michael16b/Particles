package main

// Update se charge d'appeler la fonction Update du système de particules
// g.system. Elle est appelée automatiquement exactement 60 fois par seconde par
// la bibliothèque Ebiten.
func (g *game) Update() error {

	g.system.Update()

	return nil
}
