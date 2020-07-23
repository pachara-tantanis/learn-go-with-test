package poker_test

import (
	"../../build-an-application"
	"bytes"
	"io"
	"strings"
	"testing"
)

var dummyPlayerStore = &poker.StubPlayerStore{}
var dummyStdOut = &bytes.Buffer{}

type GameSpy struct {
	StartCalled  bool
	StartedWith  int
	BlindAlert   []byte
	FinishedWith string
}

func (g *GameSpy) Start(numberOfPlayers int, out io.Writer) {
	g.StartCalled = true
	g.StartedWith = numberOfPlayers
	out.Write(g.BlindAlert)
}

func (g *GameSpy) Finish(winner string) {
	g.FinishedWith = winner
}

func TestCLI(t *testing.T) {
	t.Run("it prompts the user to enter the number of players and starts the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7", "I wins")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt)

		assertGameStartedWith(t, game, 7)
	})

	t.Run("it prints an error when a non numeric value is entered and does not start the game", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := strings.NewReader("Pies\n")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertGameNotStarted(t, game)
		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadPlayerInputErrMsg)
	})

	t.Run("it prints an error when input wrong winner pattern", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		in := userSends("7", "something")
		game := &GameSpy{}

		cli := poker.NewCLI(in, stdout, game)
		cli.PlayPoker()

		assertMessagesSentToUser(t, stdout, poker.PlayerPrompt, poker.BadWinnerInputErrMsg)
	})

	t.Run("start game with 3 players and finish game with 'Chris' as winner", func(t *testing.T) {
		game := &GameSpy{}

		out := &bytes.Buffer{}
		in := userSends("3", "Chris wins")

		poker.NewCLI(in, out, game).PlayPoker()

		assertMessagesSentToUser(t, out, poker.PlayerPrompt)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, "Chris")
	})
}

func assertFinishCalledWith(t *testing.T, game *GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("wanted %v but got %d", winner, game.StartedWith)
	}
}

func assertGameStartedWith(t *testing.T, game *GameSpy, count int) {
	t.Helper()
	if game.StartedWith != count {
		t.Errorf("wanted Start called with %v but got %d", count, game.StartedWith)
	}
}

func userSends(messages ...string) *strings.Reader {
	in := strings.Join(messages, "\n")
	return strings.NewReader(in)
}

func assertGameNotStarted(t *testing.T, game *GameSpy) {
	t.Helper()
	if game.StartCalled {
		t.Errorf("game should not have started")
	}
}

func assertMessagesSentToUser(t *testing.T, stdout *bytes.Buffer, messages ...string) {
	t.Helper()
	want := strings.Join(messages, "")
	got := stdout.String()
	if got != want {
		t.Errorf("got %q sent to stdout but expected %+v", got, messages)
	}
}
