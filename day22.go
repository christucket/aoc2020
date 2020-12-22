package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	guuid "github.com/google/uuid"
)

var all_games map[string]bool

func RecordGame(cards map[int][]int, game_id guuid.UUID) bool {
	game_string := fmt.Sprintf("%v: %v-%v", game_id, cards[1], cards[2])

	if !all_games[game_string] {
		all_games[game_string] = true
		return false
	}

	return true
}

func PlayCombatGame(cards map[int][]int, subgames bool) (map[int][]int, int) {
	round := 1
	game_id := guuid.New()
	for {
		same_game := RecordGame(cards, game_id)
		if same_game {
			return cards, 1
		}

		// take top card
		player1_card := cards[1][0]
		player2_card := cards[2][0]

		// remove it from the deck
		cards[1] = cards[1][1:]
		cards[2] = cards[2][1:]

		player1_winner := player1_card > player2_card

		if subgames && len(cards[1]) >= player1_card && len(cards[2]) >= player2_card {
			// sub game
			sub_game_cards := make(map[int][]int)
			// slices dont create copies! woops
			sub_game_cards[1] = append([]int{}, cards[1][0:player1_card]...)
			sub_game_cards[2] = append([]int{}, cards[2][0:player2_card]...)

			_, sub_game_winner := PlayCombatGame(sub_game_cards, true)
			player1_winner = sub_game_winner == 1
		}

		if player1_winner {
			cards[1] = append(cards[1], player1_card)
			cards[1] = append(cards[1], player2_card)
		} else {
			cards[2] = append(cards[2], player2_card)
			cards[2] = append(cards[2], player1_card)
		}

		if round > 10099 {
			fmt.Println("ended", round, game_id)
			break
		}
		if len(cards[1]) == 0 {
			return cards, 2
		} else if len(cards[2]) == 0 {
			return cards, 1
		}
		round++
	}

	return cards, 0
}

func day22() {
	inp, _ := ioutil.ReadFile("./inputs/day22.input")

	data := GetDoubleStringInput(inp)

	player_cards := make(map[int][]int)
	all_games = make(map[string]bool)

	for _, players_hand := range data {
		cards := strings.Split(players_hand, "\r\n")
		player := cards[0]
		player = player[len(player)-2 : len(player)-1]
		player_int, _ := strconv.Atoi(player)
		for _, card := range cards[1:] {
			card_int, _ := strconv.Atoi(card)
			player_cards[player_int] = append(player_cards[player_int], card_int)
		}
	}

	cards, winner := PlayCombatGame(player_cards, true)
	fmt.Println("winner is:", winner, "with cards", cards[winner])

	score := 0
	for i, card := range cards[winner] {
		score += (len(cards[winner]) - i) * card
	}

	fmt.Println("final score is:", score)
}
