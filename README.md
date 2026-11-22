# Push-Swap

Programme en Go qui trie une liste d'entiers avec le moins d'instructions possible, en n'utilisant que deux piles (*a* et *b*) et un jeu d'opérations limité. Aucune dépendance hors bibliothèque standard.

Deux programmes : `push-swap` affiche les instructions qui trient la pile, `checker` vérifie qu'une suite d'instructions la trie.

## Compilation

Go, dans la version indiquée par `go.mod`.

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

## Utilisation

```bash
./push-swap "4 67 3 87 23"          # affiche 6 instructions, une par ligne
ARG="4 67 3 87 23"
./push-swap "$ARG" | ./checker "$ARG"   # affiche OK (ou KO)
./push-swap "$ARG" | wc -l              # nombre d'instructions
```

- Seuls les entiers relatifs sont acceptés, sans doublon.
- En cas d'erreur (argument invalide, doublon, instruction inconnue pour `checker`) : `Error` sur la sortie d'erreur, code de sortie 1.
- Sans argument, ou pile déjà triée : aucune sortie. Les lignes vides sont ignorées par `checker`.

## Instructions

Le sommet d'une pile est son premier élément.

| Instruction | Effet |
|---|---|
| `sa` / `sb` / `ss` | échange les deux premiers éléments de *a* / de *b* / des deux |
| `pa` / `pb` | déplace le premier élément de *b* vers *a* / de *a* vers *b* |
| `ra` / `rb` / `rr` | le premier élément devient le dernier, sur *a* / *b* / les deux |
| `rra` / `rrb` / `rrr` | le dernier élément devient le premier, sur *a* / *b* / les deux |

Une instruction impossible (par exemple `pa` avec *b* vide) est sans effet.

## Stratégie de tri

Chaque nombre est d'abord remplacé par son rang (0 à *n*-1). La méthode dépend ensuite de la taille *n* :

| Taille | Méthode |
|---|---|
| 2 | `sa` |
| 3 | analyse des 6 cas (optimal : 2 instructions au maximum) |
| 4 à 8 | recherche exhaustive : suite d'instructions la plus courte possible |
| 9 à 11 | tri par sélection |
| 12 et plus | tri par blocs, taille de bloc ≈ 1,4 × √*n* |

- **Recherche exhaustive** : parcours en largeur des situations atteignables, en mémorisant celles déjà vues. Le coût croît comme (*n*+1)!, d'où le seuil à 8 éléments (pire cas ≈ 5 s).
- **Tri par sélection** : le plus petit élément de *a* passe dans *b*, puis tout revient dans *a* avec `pa`.
- **Tri par blocs** : aller dans *b* par tranches de valeurs, puis retour dans *a* en reprenant à chaque fois le plus grand de *b*. Le coefficient 1,4 a été trouvé par mesure.

## Mesures

```bash
go run ./cmd/comparer/tri       # compare les méthodes selon la taille de la pile
go run ./cmd/comparer/tailles   # compare plusieurs tailles de bloc
```

Nombre moyen d'instructions (20 tirages) :

| n | Sélection | Blocs | Radix |
|---|---|---|---|
| 20 | 90 | 75 | 160 |
| 100 | 1 476 | 557 | 1 084 |
| 500 | 32 560 | 4 886 | 6 784 |

Le tri radix (`piles.TrierRadix`) n'est pas utilisé par `push-swap` : il sert de point de comparaison.

## Tests et structure

```bash
go test ./...    # le test de la recherche sur 8 éléments dure environ 5 s
```

```text
cmd/push-swap      programme principal
cmd/checker        vérificateur d'instructions
cmd/comparer/tri       comparaison des méthodes de tri
cmd/comparer/tailles   comparaison des tailles de bloc
piles/             structure Piles, instructions, algorithmes de tri
outils/            lecture des arguments, doublons, rangs, taille de bloc
```
