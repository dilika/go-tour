# Mes predictions sur Fonctions (Avant lecture, sans regarder le code)


## Example1 – basic-multiple-return
  –Predictions: Une fonction go peut retourner N valeurs, N > 1, parmi ces valeurs on peut retourner l'erreur
  –Raison: la raison qui me viennent a l'esprit est que go ne gere pas les exceptions dont il permet de renvoyé les erreur a la place


## Example2 — Identifiant blanc/vide
  –Predictions: on peut decider de retourner un valeurs qui qu'on veut ignorer ou ne pas traiter
  –Raison: Franchement j'ignore la raison pour cela


## Example3 – Ré-déclaration
  –Predictions: Avec la Ré-déclaration j'attend par declarer une variable avec un type specifique et ensuite lui affecter une autre valeur
  –Raison: Dans le cadre d'un fonction Franchement je n'en sais rien


## Example4 – Fonctions Anonymes/Closures
  –Predictions: Je crois bien qu'une fonction anonyme est une fonction qui n'a pas de nom comme les fonction ordinaire du coup on peut les passer en parametre a une fonction ordinaire ou les retourner comme valeur dans une fonction
  –Raison: Pour ce qui de la raison, je ne saurai le dire


## Example5 –Récupérer des Panics
  –Predictions: Les panics sont des fonction qu'on fait appelle lorsqu'il y a une erreur
  –Raison: J'avoue que je ne vois pas pour le moment la raison.

## Raffinement obligatoire
 ### Raffinement A 
   `name, age := "Bob", 30
   name, city := "Bob", "NYC"
   fmt.Println(name, age, city)`
  

Oui le code compile, la sortie exacte est Bob, 30, NYC
La regle := permet une affectation direct avec assignation de type implicite
La règle de la ré-affectation dit que si il existe au mois une nouvelle variable LHS on utilise :=, le cas name, city, name est deja declarer et city ne l'est pas sinon on utilise =

### Raffinement B
   `code
    var fns []func()
    for i:= 0; i < 3: i++ {
       fns = append(fns, func(){ fmt.Println(i)})
   }
    for _, f := range fns {f()}
   `
Prédit : 1 2 3  ← FAUX (vérifié contre go1.26 : `go run` donne 0 1 2)
Le pourquoi (corrigé) : la boucle commence à i := 0 (donc 0, pas 1).
  - go1.22+ : `i` est une variable fraîche par itération → chaque closure capture SA propre `i` → 0 1 2.
  - pre-1.22 : `i` était une seule variable partagée, valant 3 à la fin → toutes les closures affichaient 3 3 3.
Mon erreur : j'ai confondu l'index de départ (0, pas 1) et le comportement post-1.22.

### Raffinement C
    2 cas ou panic est approprié
    un index hors limite dans le fonctionnement interne du programme qu'on controle, ou un nil receiver qui ne devrait jamais arriver 
    `code
     if user == nil {
        panic("user should never be nil here!")
}
    Une situation irrecuperable qui corromperait l'etat du programme
    echec critique lors de l'initialisation d'un composant essentiel au programme(base de donnée, configuration obligatoire) ou continuer serati dangereux
`
    2 cas ou erreur est approprié
    Erreur attendu et irrecuperable:
    Erreur de reseau, fichier non trouvé, erreur de validation d'un utilisateur etc..
    `code
    err, file := os.Open("config.json")
    if err != nill {
      fmt.Errorf("Impossible de lire le fichier de configuration", err)
    }
    `
    Erreur que l'appellant peut recuperer et gerer en sa maniere
    Lorsqu'un utilisateur n'existe pas, l'appellant peut de creer un nouveau ou afficher un message d'erreur a l'utilisateur


     Le cas de recover 
      Recover fonctionnemnt uniquement si :
     il est appélé directemnt dans une fonction defer 
     il est dans le goroutine que celle qui panic
     il est appéleé avant que la fonction differé ne retourne


## Confrontation (après lecture)

### Example 1
- Prédit : Une fonction go peut retourner N> 1 value et parmi ces values on peut retourner l'erreur
- Chapitre dit : Go permet aux fonction de renvoyé plusieurs valeurs, ceci est utilse si vous voulez envoyé plus d'un resultat
- Surprise : Rien just une reformulation, l'idée general est comprise

### Example 2
- Prédit :  on peut decider de retourner un valeurs qui qu'on veut ignorer ou ne pas traiter
- Chapitre dit : cet example l'illustre bien :
`code
// Update the user Name. Don't care about the update stats.
	if _, err := updateUser(&u); err != nil {
		fmt.Println(err)
		return
	}
` ou decide d'ignorer la deuxieme valeur de retour et traiter uniquement l'erreur
- Surprise : Pas de surprise sauf que je l'ai affirmé sans vraiment le mettre en pratique et je comprend mieux avec la pratique

### Example 3
- Prédit : Avec la Ré-déclaration j'attend par declarer une variable avec un type specifique et ensuite lui affecter une autre valeur
- Chapitre dit : de memoire, le chapitre dit pour une re-declaration d'une variable avec := il faut une deuxieme variable nouvelle avec laquelle declarer sinon utiliser = a la place
- Surprise : J'etais completement a coté de la plaque et je comprend mieux

### Example 4
- Prédit: Je crois bien qu'une fonction anonyme est une fonction qui n'a pas de nom comme les fonction ordinaire du coup on peut les passer en parametre a une fonction ordinaire ou les retourner comme valeur dans une fonction
- Chapitre dit: En plus de ce que j'ai dit, les fonctions peut etre des types auquel cas on peut les passer en parametre a une fonction ou les retourné comme valeur de retour d'une fonciton.
- Surprise: Je ne m'attendais pas a ce que les fonctions soit considerer comme type.
Direct n  0 (la sortie 1 parceque la fonction est déclaré et appélé, c'est une fonction anonyme evoqué aussitot )
variable  0 (cette fonction n'est appélé que par 	f() c'est une variable)
variable  3  ( parce que on a affecté la valeur 3 et evoqué la fonction anonyme qui affiche variable avec l'appelle de la variable f())
Defer 2:  3  (Les defer sont appélé apres l'excution de toutes les fonctions sans defer, et et le premier defer appélé se retrouve avec first in last out )
Defer 1 :   (Les defer sont appélé apres l'excution de toutes les fonctions sans defer, et et le premier defer appélé se retrouve avec last in first out )


# Example 5 
- Prédit: Les panics sont des fonctions qu'on fait appelle lorsqu'il y a une erreur
- Chapitre dit: Le panic est utiliser pour déclenché une erreur d'execution et derouler la stact
- Surprise: déclenché une erreur d'execution est different de "fait appelle lorsqu'il y a une erreur"




