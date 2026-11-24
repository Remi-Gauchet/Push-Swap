# Push-Swap

Programme qui trie dans l'ordre croissant une liste de *n* nombres entiers relatifs, avec le moins d'opérations possibles.  
Répartit les valeurs sous forme de deux piles qu'il réorganise avec des opérations bien définies.  
Push-Swap est un défi d'algorithmie, celui-ci est codé en Go.

## Sommaire

- [Principe](#principe)
- [Compilation](#compilation)
- [Utilisation](#utilisation)
    - [Push-Swap](#1-push-swap)
    - [Checker](#2-checker)
- [Détail des opérations](#détail-des-opérations)
- [Stratégie de tri](#stratégie-de-tri)
- [Mesurer l'optimisation](#mesurer-loptimisation)
- [Tests](#tests)
- [Limites connues](#limites-connues)

## Principe

Soient *a* et *b* deux piles dont le sommet est à gauche.  
Les nombres relatifs donnés en arguments sont placés dans la pile *a*.  
La pile *b*, vide au départ, sert d'espace de travail.  

Avec les opérations autorisées (détaillées dans la section suivante), le programme réorganise *a* jusqu'à ce que ses nombres soient triés dans l'ordre croissant, avec *b* vide. Il affiche alors la liste des opérations utilisées, une par ligne.  
L'objectif est d'obtenir un tri correct avec le moins d'opérations possible.

* Le programme `push-swap` calcule et affiche dans le terminal une suite d'opérations.
* Le programme `checker` lit des opérations sur son entrée standard, les applique à la pile et indique si elle est triée.

## Compilation

```bash
go build -o push-swap ./cmd/push-swap
go build -o checker ./cmd/checker
go build -o comparer-tailles ./cmd/comparer/tailles
go build -o comparer-tri ./cmd/comparer/tri
```

## Utilisation

### 1. Push-Swap
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

### 2) Checker
Entrez une suite d'opération puis votre suite de nombres à trier, en suivant le format de cet exemple :

```bash
echo -e "sa\nrra\npb" | ./checker "3 2 1 0"
```

Il affiche `OK` si la pile *a* est triée et que *b* est vide, `KO` sinon.  


## Détail des opérations

| Opération | Effet | Exemple avant | Exemple après |
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

L'efficacité du tri par bloc repose sur le choix optimal de la taille du bloc en fonction de *n*.  
Une taille trop petite oblige à parcourir *a* un grand nombre de fois, une taille trop grande range mal *b* et rend le retour coûteux.  
Selon nos observations : `Taille de bloc optimale = coefficient x √n`  
Nous avons trouvé le coefficient **1,4** par tattonnement avec la fonction `./comparer-tailles`.

## Mesurer l'optimisation

Voici deux outils qui utilisent des piles tirées au hasard :  
```bash
./comparer-tri      # compare les méthodes de tri selon la taille de la pile
./comparer-tailles   # compare plusieurs tailles de bloc selon la taille de la pile
```

Pendant l'optimisation du programme, deux méthodes de tri ont été éliminées : TrierRadix et TrierCinq.  
Elles sont fonctionnelles, mais moins efficaces. Nous les gardons comme point de comparaison.

## Tests

```bash
go test -v ./piles
```


## Limites connues

- Push-Swap n'accepte que les nombres entiers relatifs. Les doublons sont refusés.
- Les nombres sont des `int` Go, donc limités aux bornes de ce type :  
de -2 147 483 648 à 2 147 483 647 sur une machine 32 bits   
de -9 223 372 036 854 775 808 à 9 223 372 036 854 775 807 sur une machine 64 bits  
Au-delà, le programme affiche `Error`.  
- La vitesse d'exécution de la fonction recherche dépend de la puissance de la machine. Elle devient beaucoup trop lente au-delà de 8 éléments. C'est pourquoi on change de méthode par la suite.
- Pour les grandes piles, notre tri par blocs est efficace mais n'est pas garanti optimal.
