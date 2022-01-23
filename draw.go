package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"project-particles/assets"
	"project-particles/config"
)

// Draw se charge d'afficher à l'écran l'état actuel du système de particules
// g.system. Elle est appelée automatiquement environ 60 fois par seconde par
// la bibliothèque Ebiten.
func (g *game) Draw(screen *ebiten.Image) {

	for _, p := range g.system.Content {
		if !(p.PositionX > float64(config.General.WindowSizeX+config.General.Margin) || p.PositionX < 0-float64(config.General.Margin) || p.PositionY > float64(config.General.WindowSizeY+config.General.Margin) || p.PositionY < 0-float64(config.General.Margin)) || p.Life > config.General.Life {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorM.Scale(p.ColorRed, p.ColorGreen, p.ColorBlue, p.Opacity)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	if config.General.Debug {
		ebitenutil.DebugPrint(screen, fmt.Sprint(ebiten.CurrentTPS()))
	}

}
