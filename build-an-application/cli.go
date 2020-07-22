package poker

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

type CLI struct {
	Game GameController
	in   *bufio.Scanner
	out  io.Writer
}

func NewCLI(in io.Reader, out io.Writer, game GameController) *CLI {
	return &CLI{Game: game, in: bufio.NewScanner(in), out: out}
}

const PlayerPrompt = "Please enter the number of players: "
const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"
const BadWinnerInputErrMsg = "Bad received for winner input"

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, err := strconv.Atoi(cli.readLine())

	if err != nil {
		fmt.Fprint(cli.out, BadPlayerInputErrMsg)
		return
	}

	cli.Game.Start(numberOfPlayers, ioutil.Discard)

	userInput := cli.readLine()
	winner, err := extractWinner(userInput)
	if err != nil {
		fmt.Fprintf(cli.out, BadWinnerInputErrMsg)
	}

	cli.Game.Finish(winner)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) (string, error) {
	winner := strings.TrimSuffix(userInput, " wins")
	if winner == userInput {
		return "", fmt.Errorf("wrong winning command got %v", userInput)
	}
	return winner, nil
}
