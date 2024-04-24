package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Player struct {
    id    int
    dice  []int
    score int
}

func main() {
    rand.Seed(time.Now().UnixNano())

    var numPlayers, numDice int

    // Input jumlah pemain dan jumlah dadu
    fmt.Println("Masukkan jumlah pemain:")
    fmt.Scanln(&numPlayers)

    fmt.Println("Masukkan jumlah dadu per pemain:")
    fmt.Scanln(&numDice)

    // Inisialisasi pemain
    players := make([]Player, numPlayers)
    for i := range players {
        players[i] = Player{id: i + 1, dice: make([]int, numDice)}
    }

    // Mulai permainan
    for {
        // Lemparkan dadu
        for i := range players {
            players[i].rollDice()
        }

        // Evaluasi lemparan dadu
        for i := range players {
            players[i].evaluate()
        }

        // Cek apakah ada pemain yang masih memiliki dadu
        remainingPlayers := make([]Player, 0)
        for _, player := range players {
            if len(player.dice) > 0 {
                remainingPlayers = append(remainingPlayers, player)
            }
        }

        // Jika hanya ada satu pemain yang tersisa, akhiri permainan
        if len(remainingPlayers) == 1 {
            fmt.Printf("Pemain %d menang dengan %d poin!\n", remainingPlayers[0].id, remainingPlayers[0].score)
            break
        }

        // Persiapkan lemparan dadu selanjutnya
        players = remainingPlayers
    }
}

// Method untuk melemparkan dadu
func (p *Player) rollDice() {
    for i := range p.dice {
        p.dice[i] = rand.Intn(6) + 1 // Lemparkan dadu dengan angka 1-6
    }
}

// Method untuk mengevaluasi hasil lemparan dadu
func (p *Player) evaluate() {
    for i := 0; i < len(p.dice); i++ {
        switch p.dice[i] {
        case 1:
            // Berikan dadu angka 1 kepada pemain disampingnya
            if i+1 < len(p.dice) {
                adjacentPlayer := i + 1
                p.dice = append(p.dice[:i], p.dice[i+1:]...)
                fmt.Printf("Pemain %d memberikan dadu angka 1 kepada pemain %d\n", p.id, adjacentPlayer+1)
            }
        case 6:
            // Keluarkan dadu angka 6 dan tambahkan poin
            p.dice = append(p.dice[:i], p.dice[i+1:]...)
            p.score++
            fmt.Printf("Pemain %d mendapatkan poin karena mendapat dadu angka 6\n", p.id)
        }
    }
}
