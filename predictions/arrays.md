# Mes predictons sur le chapitre des arrays( Sans regarder le code )

## Example1: Déclarer, initier et iterer
 - Predicitons: Je n'ai aucune predictions a faire ici car c'est une question de syntaxe
 - Raison: pas de raison non plus puisqu'aucune preditions n'a été faite

## Example2: Tableau de different type
 - Predicitons: Venant de l'univers Java je dirai qu'un tableau doit contenir des elements de meme type, je pense que la meme chose doit etre possible ici
 - Raison: Un tableau doit contenir des elements de meme type


## Example3: Allocation de memoire contigue
 - Predicitons: L'allocation de memoire contigue est comme c'est fait avec les struct, c'est a dire on alloue l'addresse memoire contigue au precedent champ declaré dans le struct
 - Raison: ça permet une bonne gestion de la memoire et facilite le netoyage de la memoire


## Example4: Mecanique de parcourt
 - Predicitons: La mecanique de parcourt d'un arrays pour moi qui vient de l'univers java doit commencé par i le length du tableau -1
 - Raison: Puisque depuis le langagee machine l'adresse d'un element du tableau est obtenu en ajoutant l'index a l'adresse du debut du tableau et en commençant par 1 resulterait gaspillage de l'espace memoire et un gachis lors de la recuperation de l'adresse de l'element en enlevant le 1 qu'on avait ajouter



# Confrontation apres lecture:
## Example1: Déclarer, initier et iterer
   - Predictons: Je n'ai aucune predictions a faire ici car c'est une question de syntaxe
   - Chapitre dit:
      - Déclaration: 
        `code 
        var strings [5]string
      `
      - Initialisation: 
       `code 
        strings := [5]string{"Apple", "Orange", "Tomato", "Pinapple", "Mango"}
      `
      - Iteration: The chapter say that we two kind of iteration with arrays which are : semantic value iteration and semantic pointer iteration
         - Semanctic value iteration:
          `code 
          for i, fruit := range strings {
                fmt.Println(i, fruit)
         }
        `
         - Semanctic pointer iteration:
          `code 
          for i := range strings {
                fmt.Println(i, strings[i])
         }
         `
    - Surprises: Je n'ai pas été surpris sur la syntaxe et l'initialisation, par contre l'iteration par value et la copie effectué a chaque iteration m'a vraiment surpris et le cout n'est pas enorme

## Example2: Tableau de type different
  - Predicitons: Venant de l'univers Java je dirai qu'un tableau doit contenir des elements de meme type, je pense que la meme chose doit etre possible ici
  - Chapitre dit: La longueur du tableau fait partie des information sur le type donc go considere deux arrays de meme type avec des longueur different comme des type type different:
     - Exmample:
     `code
      var fiveArray [5]int
      var fourArray [4]int

      fiveArray = fourArray // go le rejete en disant que les type [5]int est different du type [4]int
    `
  - Surprises: J'etais completement a coté de la plaque, j'ai confondu l'affectation de variable avec les types de contenu que peut avoir un array en go

## Example3: Allocation de memoire contigue
  - Predicitons: L'allocation de memoire contigue est comme c'est fait avec les struct, c'est a dire on alloue l'addresse memoire contigue au precedent champ declaré dans le struct
  - Chapitre dit: Chaque élément est exactement 16 bytes après le précédent. Le CPU adore ça — il peut prédire où sera le prochain élément avant même d'en avoir besoin.&v est identique à chaque itération (preuve que v est une seule variable copiée), tandis que &friends[i] avance de 16 octets. Cette opposition (&v constant vs &arr[i] qui avance) est la signature visuelle de la sémantique par valeur.
  - Surprises: L'allocation de memoire de memoire contigue s'avere etre plus optimal pour un cpu.


## Exmample4: Mecanique de parcourt
   - Predicitons:  La mecanique de parcourt d'un arrays pour moi qui vient de l'univers java doit commencé par i le length du tableau -1
   - Chapitre dit: enseigne la visibilité des mutations pendant l'iteration, avec trois forme.
`code
friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

// Forme 1 — pointer semantics
for i := range friends {
    friends[1] = "Jack"
    if i == 1 { fmt.Printf("Aft[%s]\n", friends[1]) }
}  // → Aft[Jack]   (on lit l'original → mutation visible)

// Forme 2 — value semantics
for i, v := range friends {
    friends[1] = "Jack"
    if i == 1 { fmt.Printf("v[%s]\n", v) }
}  // → v[Betty]   (v = copie d'avant la mutation)

// Forme 3 — DANGER, "DON'T DO THIS"
for i, v := range &friends {
    friends[1] = "Jack"
    if i == 1 { fmt.Printf("v[%s]\n", v) }
}  // → v[Jack]   (itération par pointeur → v reflète la mutation)
`

 - Surprises: surpris d'apprendre d'apprndre les trois forme de l'iteration et comment l'iteration y s'effectue je n'en avais aucune idée
   
