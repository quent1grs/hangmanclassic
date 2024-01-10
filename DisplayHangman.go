package hangmanclassic

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

	if len(etapesPendu) == 0 {
		fmt.Println("Aucune étape chargée depuis le fichier.")
		return
	}

	for numEtape, nbEssais := range etapesPendu {
		fmt.Printf("Étape %d : %d essais\n", numEtape, nbEssais)
	}
}
