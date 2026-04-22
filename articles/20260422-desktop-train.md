---
title: "Weekend Projects: j'ai créé un desktop pet"
url: project_desktop_train
desc: ≈3min · Après avoir vu un post sur internet à propos de Desktop Goose, j'ai eu envie de créer mon propre desktop pet, mais en forme de train!
date: 2026-04-22
tags:
  - go
  - ebitengine
  - trains
  - weekend-project
---
Après avoir vu un post sur internet à propos de [Desktop Goose](https://samperson.itch.io/desktop-goose) j'ai eu envie de créer mon propre desktop pet, mais en forme de train !

[↳ Repo Github du projet](https://github.com/nrissot/desktop_train) <br>
[↳ Dernière release des exécutables](https://github.com/nrissot/desktop_train/releases/latest) <br>
[↳ Page ITCH-IO des assets](https://kaawan.itch.io/tiny-trains)

## C'est quoi un desktop pet
Les desktop pet sont une variation des virtual pet, littéralement, des animaux de compagnie virtuels (tamagotchi), mais vivant sur le bureau de votre ordinateur.

Selon [Wikipedia](https://en.wikipedia.org/wiki/Virtual_pet#Software-based) le premier desktop pet était un programme appelé Neko, dans lequel un petit chat s'amuse a chasser votre curseur.

## Concept du projet
Plutôt que d'avoir un véritable animal de compagnie virtuel qui interagirait avec le curseur par exemple, je préférerait quelque de plus discret et oubliable, qui soit non-disruptif, et qui ne bloque pas de fonctionnalité des outils que j'utilise (VS Code)

Il est donc impératif par exemple que les zones d'outils ne soit pas bloquées par le desktop pet.

En regardant un tram passer en allant à la fac j'ai eu l'idée de faire un petit train, circulant en bas de l'écran de l'ordinateur.

Avec le concept en place j'ai pu dresser le cahier des charges de ce projet:

### Cahier des charges
- ne doit pas bloquer les clics
- ne doit pas perturber les actions de l'utilisateur
- ne doit pas être trop distrayant (l'objectif c'est que je puisse le garder sur mon écran en travaillant après tout)

## Programmation
Pour programmer ce projet, j'ai décidé d'utiliser mon language préféré: Go (<3), avec la librairie de développement de jeux [Ebitengine](https://ebitengine.org/) (que j'avais déjà utilisé pour mon projet [starmap](https://nrissot.github.io/misc/starmap) sur lequel j'écrirai peut être un article un jour).

### Setup de la fenêtre

Après environ dix minutes de recherche dans la doc, j'ai fini par trouver toute les options qu'ils me fallait:

```go
ebiten.SetWindowMousePassthrough(true) // permet de cliquer à travers la fenêtre
ebiten.SetWindowDecorated(false)       // retire les bordure de la fenêtre
ebiten.SetWindowFloating(true)         // met toujours la fenetre au dessus des autres

// précise que la fenêtre est transparente par défaut.
ebiten.RunGameWithOptions(&g, &ebiten.RunGameOptions{ScreenTransparent: true})
```

Et voila, la partie "dure" est faite, la fenêtre est donc toujours affichée au dessus des autres, complètement transparente, et ne bloque pas les clics.

Maintenant il ne reste plus qu'a afficher le petit train

### Logique
Avec Ebitengine, notre application/jeu est construit autour de 2 méthodes sur un struct représentant l'état de notre jeu:

`func (g *Game) Update() error { ... }` dans laquelle on effectue les mises a jour d'état, les calculs etc. Dans un jeu "normal" c'est dans cette méthode que l'on gérerait les inputs du joueurs, les calculs de physiques, etc etc. Ici, on se contente de soustraire 1 à la position X qui marque l'endroit a partir du quel dessiner le train, et de décider si le train est sorti de l’écran, et donc si on doit le remplacer par un nouveau train. C'est aussi dans cette fonction qu'on compte le nombre de trains pour l'écrire dans les statistiques.

`func (g *Game) Draw(screen *ebiten.Image) { ... }` dans laquelle on dessine ce qui doit apparaître sur la fenêtre. Pour des raisons d’efficacité, on redessine toute la fenêtre toute les frames (notre jeu tourne à 60 FPS, donc on dessine 60 fois par secondes). On redessine tout toutes les frames, car, dans la majorité des cas, c'est plus rapide de calculer les nouveaux pixels plutôt que de décider cas par cas si il faut modifier les existants. Pour ne pas bloquer visuellement des parties de l'écran, si la souris est autour de la ligne de chemin de fer, celle ci devient translucide pour qu'on puisse lire ce qu'il y'a en dessous.

### Assets
J'ai ensuite dessiner plusieurs assets de train et de rails. J'ai utilise le programme [LibreSprite](https://libresprite.github.io/#!/), et la palette de couleurs [PICO-8 Secret Palette](https://lospec.com/palette-list/pico-8-secret-palette). J'ai ensuite dessiné trois styles de trains avec des variantes de wagons et de locomotives, ainsi que des tags que je peut appliquer par dessus pour plus de variations:

![Exemple d'un train construit à partir des assets que j'ai dessiné](/static/assets/images/train-example.png)

Tout ces assets sont placé sur une seule image que l'on appelle la spritesheet (ou tilesheet) qui est ensuite découpée par le moteur de jeu en tuiles individuelles qui sont affiché par la méthode `Draw` 

## Conclusion
Ce projet s'est déroulé sans problème, c’était un petit projet sur un weekend (weekend+ si on compte les deux trois finitions faite après), je me suis un peu plus familiarisé avec Ebitengine et j'ai refait du pixel art, ce que je n'avais pas fait depuis longtemps.

Depuis j'ai mes petits trains qui défilent sur le bas de mon écran, y compris pendant que je suis en train d'écrire ces lignes

![Capture d'écran d'un train passant sur mon bureau](/static/assets/images/desktop-train-screenshot.png)

## Remarques:
Pour utiliser ce projet, vous pouvez soit 
1. Compiler l'exécutable à partir du code source vous même en suivant les instructions
2. Télécharger l'exécutable adapté à votre système dans [la dernière release](https://github.com/nrissot/desktop_train/releases/latest)