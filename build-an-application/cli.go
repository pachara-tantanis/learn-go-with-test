package poker

import (
	"bufio"
	"fmt"
	"io"
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

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayers, _ := strconv.Atoi(cli.readLine())
	cli.Game.Start(numberOfPlayers)

	userInput := cli.readLine()
	cli.Game.Finish(extractWinner(userInput))
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}
