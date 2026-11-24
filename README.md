# Push-Swap

Programme qui trie dans l'ordre croissant une liste de *n* nombres entiers relatifs, avec le moins d'opérations possibles.  
Répartit les valeurs sous forme de deux piles, et utilise une liste d'opérations bien définies.  
Push-Swap est un défi d'algorithmie, celui-ci est codé en Go.

## Sommaire

- [Principe](#principe)
- [Prérequis et compilation](#prérequis-et-compilation)
- [Utilisation](#utilisation)
- [Les instructions](#les-instructions)
- [Stratégie de tri](#stratégie-de-tri)
- [Mesurer l'optimisation](#mesurer-loptimisation)
- [Tests](#tests)
- [Structure du projet](#structure-du-projet)
- [Limites connues](#limites-connues)

## Principe

Les nombres donnés en arguments sont placés dans une pile *a*. Une seconde pile *b*, vide au départ, sert d'espace de travail.  
Le sommet de chaque pile est son premier élément (à gauche).

Avec les opérations autorisées (swap, rotate et push), le programme réorganise *a* jusqu'à ce que ses nombres soient triés dans l'ordre croissant, avec *b* vide. Il affiche alors la liste des instructions utilisées, une par ligne. L'objectif est d'obtenir un tri correct avec le moins d'instructions possible.

Le projet contient deux programmes :

- `push-swap` calcule et affiche les instructions qui trient la pile ;
- `checker` lit des instructions sur son entrée standard, les applique à la pile et indique si elle est triée.

## Compilation

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
```

## Utilisation

### 1) push-swap
Entrez en argument la pile à trier (fonctionne avec ou sans guillemets).  
Exemple :

```bash
./push-swap "4 67 3 87 23"
```

Résultat (6 instructions) :

```text
pb
sa
pb
rrr
pa
pa
```

### 2) checker

Le `checker` reçoit la pile en argument et les instructions sur l'entrée standard, avec une instruction par ligne.  
Exemple :

```bash
echo -e "sa\nrra\npb" | ./checker "3 2 1 0"
```

Il affiche `OK` si la pile *a* est triée et que *b* est vide, `KO` sinon.  


## Les instructions

| Instruction | Effet | Exemple avant | Exemple après |
|---|---|---|---|
| `sa` | Échange les deux premiers éléments de *a* | a=[1 2 3 4 5] | a=[2 1 3 4 5] |
| `sb` | Échange les deux premiers éléments de *b* | b=[1 2 3 4 5] | b=[2 1 3 4 5] |
| `ss` | `sa` et `sb` en même temps | a=[1 2 3] b=[4 5 6] | a=[2 1 3] b=[5 4 6] |
| `pa` | Prend le premier de *b* et le pose sur *a* | a=[1 2 3] b=[4 5] | a=[4 1 2 3] b=[5] |
| `pb` | Prend le premier de *a* et le pose sur *b* | a=[1 2 3] b=[4 5] | a=[2 3] b=[1 4 5] |
| `ra` | Décale *a* vers le haut : le premier devient le dernier | a=[1 2 3] | a=[2 3 1] |
| `rb` | Décale *b* vers le haut : le premier devient le dernier | b=[1 2 3] | b=[2 3 1] |
| `rr` | `ra` et `rb` en même temps | a=[1 2 3] b=[4 5 6] | a=[2 3 1] b=[5 6 4] |
| `rra` | Décale *a* vers le bas : le dernier devient le premier | a=[1 2 3] | a=[3 1 2] |
| `rrb` | Décale *b* vers le bas : le dernier devient le premier | b=[1 2 3] | b=[3 1 2] |
| `rrr` | `rra` et `rrb` en même temps | a=[1 2 3] b=[4 5 6] | a=[3 1 2] b=[6 4 5] |


## Stratégie de tri

Avant de trier, chaque nombre est remplacé par son **rang** (de 0 pour le plus petit jusqu'à *n*-1 pour le plus grand)  
Puis, la méthode utilisée dépend du nombre *n* d'éléments :

| Taille n | Méthode |
|---|---|
| 2 | `sa` |
| 3 | analyse des 6 cas possibles |
| 4 à 8 | fonction recherche |
| 9 à 11 | tri par sélection |
| 12 et plus | tri par blocs |

### Fonction Recherche

Le programme explore les situations atteignables en 1 instruction, puis en 2, puis en 3, et ainsi de suite, en mémorisant celles déjà rencontrées pour ne pas les refaire.  
La première situation triée trouvée est la suite d'instructions la plus courte possible.

Nous avons fixé le seuil à 8 éléments maximum pour limiter le temps d'exécution.  
Passer à 9 éléments multiplierait ce temps d'exécution par 10 environ.

### Tri par sélection

À chaque étape, le plus petit élément de *a* est ramené au sommet par le chemin le plus court (`ra` ou `rra`) puis envoyé dans *b*.  
Une fois *a* vide, tous les éléments sont remis dans *a* avec `pa`, dans le bon ordre.

### Tri par blocs

1. **Aller** : les éléments de *a* sont envoyés dans *b* par tranches de valeurs, des plus petites aux plus grandes.
2. **Retour** : on remonte vers *a* en ramenant à chaque fois le plus grand élément de *b*, par le chemin le plus court (`rb` ou `rrb`).

La taille des tranches est calculée par `outils.TailleDeBloc` : environ 1,4 × √*n*, arrondie à l'entier le plus proche.  
Une taille trop petite oblige à parcourir *a* un grand nombre de fois, une taille trop grande range mal *b* et rend le retour coûteux.  
Le meilleur compromis évolue comme √*n*, et le coefficient 1,4 a été trouvé par mesure avec `cmd/comparer/tailles`.

## Mesurer l'optimisation

Voici deux outils qui utilisent des piles tirées au hasard :

```bash
go run ./cmd/comparer/tri       # compare les méthodes de tri selon la taille de la pile
go run ./cmd/comparer/tailles   # compare plusieurs tailles de bloc selon la taille de la pile
```

Pendant l'optimisation du programme, deux méthodes de tri ont été éliminées : TrierRadix et TrierCinq.  
Elles sont fonctionnelles, mais moins efficaces. Nous les gardons comme point de comparaison.

## Tests

```bash
go test ./...
```

Le test de la fonction recherche sur 8 éléments dure environ 5 secondes.  
Pour l'afficher en détail : `go test -v ./piles`.

## Limites connues

- Push-Swap n'accepte que les nombres entiers relatifs. Les doublons sont refusés.
- Les nombres sont des `int` Go, donc limités aux bornes de ce type :  
de -2 147 483 648 à 2 147 483 647 sur une machine 32 bits   
de -9 223 372 036 854 775 808 à 9 223 372 036 854 775 807 sur une machine 64 bits  
Au-delà, le programme affiche `Error`.  
- La vitesse d'exécution de la fonction recherche dépend de la puissance de la machine. Elle devient beaucoup trop lente au-delà de 8 éléments. C'est pourquoi on change de méthode par la suite.
- Pour les grandes piles, notre tri par blocs est efficace mais n'est pas garanti optimal.
