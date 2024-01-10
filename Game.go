package hangmanclassic

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Game(mot string) {
	longueur := len(mot)
	essaisRestants := 10
	lettresRestantes := longueur

	rand.Seed(time.Now().UnixNano())
	indices := make([]int, longueur/2-1)

	for i := 0; i < longueur/2-1; i++ {
		indice := rand.Intn(longueur)
		indices[i] = indice
	}

	motPartiel := make([]rune, longueur)
	for i := range mot {
		afficher := false
		for _, index := range indices {
			if i == index {
				afficher = true
				lettresRestantes--
				break
			}
		}
		if afficher {
			motPartiel[i] = rune(mot[i])
		} else {
			motPartiel[i] = '_'
		}
	}

	reader := bufio.NewReader(os.Stdin)

	for essaisRestants > 0 && lettresRestantes > 0 {
		fmt.Printf("Mot à deviner : %s\n", string(motPartiel))
		fmt.Printf("Essais restants : %d\n", essaisRestants)
		fmt.Print("Entrez une lettre ou devinez le mot : ")

		tentative, _ := reader.ReadString('\n')
		tentative = strings.TrimSpace(tentative)

		if tentative == "" {
			fmt.Println("Veuillez entrer une lettre ou deviner le mot.")
			continue
		}

		if len(tentative) == 1 {
			lettreRune := []rune(tentative)[0]
			trouve := false
			for i, c := range mot {
				if c == lettreRune && motPartiel[i] == '_' {
					motPartiel[i] = lettreRune
					trouve = true
					lettresRestantes--
				}
			}

			if !trouve {
				essaisRestants--
				fmt.Println("Lettre incorrecte.")
			} else {
				fmt.Println("Bonne lettre !")
			}
		} else if tentative == mot {
			fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", mot)
			return
		} else {
			essaisRestants -= 2
			fmt.Println("Tentative incorrecte pour deviner le mot.")
		}
	}

	if lettresRestantes == 0 {
		fmt.Printf("Félicitations, vous avez deviné le mot : %s\n", mot)
	} else {
		fmt.Printf("Dommage, le mot était : %s\n", mot)
	}
}
