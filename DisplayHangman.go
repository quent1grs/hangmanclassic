package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Displayhangman() {
	file, err := os.Open("etapes_pendu.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}
	defer file.Close()

	// Map pour stocker le numéro de l'étape et le nombre d'essais correspondant
	etapesPendu := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elements := strings.Fields(line)
		if len(elements) != 2 {
			continue
		}

		numEtape, err := strconv.Atoi(elements[0])
		if err != nil {
			continue
		}

		nbEssais, err := strconv.Atoi(elements[1])
		if err != nil {
			continue
		}

		etapesPendu[numEtape] = nbEssais
	}

	// Vérifier si des étapes ont été chargées
	if len(etapesPendu) == 0 {
		fmt.Println("Aucune étape chargée depuis le fichier.")
		return
	}

	// Exemple d'utilisation : affichage du nombre d'essais pour chaque étape
	for numEtape, nbEssais := range etapesPendu {
		fmt.Printf("Étape %d : %d essais\n", numEtape, nbEssais)
	}
}
