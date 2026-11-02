# Push-Swap
Programme d'organisation et de tri d'une liste de nombre.  
Défi d'algorithmie.  


## Objectif et fonctionnement
Push-Swap trie dans l'ordre croissant une liste de *n* nombres entiers relatifs, entrés en arguments dans le terminal.  
Pour faire cela, il répartit les valeurs sous la forme de deux piles *a*=[]int et *b*=[]int dont le sommet de la pile se situe à gauche.  
À travers plusieurs opérations bien définies (swap, rotate et push), le programme réorganise la liste de nombres, et vous rend la renvoie triée, accompagné du nombre d'opérations réalisées.  

La finalité de cet exercice est d'obtenir le bon résultat en faisant le moins d'opérations possibles.


## Exemple d'utilisation
Tapez :
```bash
go run . 5 87 3 26 7
```

Résultat attendu :
```bash
opérations : [rra rra rra pb rra rra pb sa rra pa pa]
nombre d'opérations : 11
```

## Règles d'utilisation
Ce programme ne gère que les nombres entiers relatifs.  
Les doublons ne sont pas autorisés.  



## Détail des opérations  
|    Opération    |              Effet           |  Exemple | Résultat |
|----------------|------------------------------|------------|----------|
|     `sa`     | Échange les deux premiers éléments de *a* | a=[1 2 3 4 5]  |   a=[2 1 3 4 5]   |
|     `sb`    | Échange les deux premiers éléments de *b* | b=[1 2 3 4 5] |   b=[2 1 3 4 5]  |
|     `pa`    | Prend le premier de *b* et le pose sur *a* | a=[1 2 3] b=[4 5] |   a=[4 1 2 3] b=[5] |
|    `pb` | Prend le premier de *a* et le pose sur *b* | a=[1 2 3] b=[4 5] | a=[2 3] b=[1 4 5] |
|   `ra` | Décale *a* vers le haut : le premier devient le dernier | a=[1 2 3] | a=[2 3 1] |
| `rb`   | Décale *b* vers le haut : le premier devient le dernier | b=[1 2 3] | b=[2 3 1] |
| `rra` | Décale *a* vers le bas : le dernier devient le premier | a=[1 2 3] | a=[3 1 2] |
| `rrb`| Décale *b* vers le bas : le dernier devient le premier | b=[1 2 3] | b=[3 1 2] |


## Mesurer l'optimisation
Pour comparer la méthode de tri par selection avec la méthode de tri par bloc, tapez :
```bash
go run . comparer
```
