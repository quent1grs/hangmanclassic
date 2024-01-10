package hangmanclassic

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func RandomWord() string {
	filename := os.Args[1]
	fmt.Println(filename)
	file, err := os.Open("assets/" + filename)
	if err != nil {
		log.Fatalf("Error When Opening File : %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	var randomword []string
	for fileScanner.Scan() {
		randomword = append(randomword, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error While Reading File : %s", err)

	}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(randomword))
	chosenWord := randomword[randomIndex]
	defer file.Close()
	return chosenWord
}
