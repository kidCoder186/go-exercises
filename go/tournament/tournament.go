package tournament

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team structure
type Team struct {
	name          string
	matchesPlayed int
	matchesWon    int
	matchesDrawn  int
	matchesLost   int
	points        int
}

// NewTeam constructor
func NewTeam(s string) *Team {
	return &Team{name: s}
}

// Table structure
type Table struct {
	exist map[string]*Team
	teams []*Team
}

// NewTable constructor
func NewTable() Table {
	return Table{
		exist: map[string]*Team{},
		teams: []*Team{},
	}
}

func (t *Table) String() string {
	res := []string{}
	res = append(res, fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n",
		"Team", "MP", "W", "D", "L", "P"))
	for _, team := range t.teams {
		res = append(res, fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			team.name, team.matchesPlayed, team.matchesWon, team.matchesDrawn, team.matchesLost, team.points))
	}
	return strings.Join(res, "")
}

func (t *Table) sort() {
	sort.Slice(t.teams, func(i, j int) bool {
		return (t.teams[i].points > t.teams[j].points) ||
			(t.teams[i].points == t.teams[j].points && t.teams[i].name < t.teams[j].name)
	})
}

func (t *Table) win(team1, team2 string) {
	t.exist[team1].matchesPlayed++
	t.exist[team1].matchesWon++
	t.exist[team1].points += 3
	t.exist[team2].matchesPlayed++
	t.exist[team2].matchesLost++
}

func (t *Table) draw(team1, team2 string) {
	t.exist[team1].matchesPlayed++
	t.exist[team1].matchesDrawn++
	t.exist[team1].points++
	t.exist[team2].matchesPlayed++
	t.exist[team2].matchesDrawn++
	t.exist[team2].points++
}

func (t *Table) loss(team1, team2 string) {
	t.exist[team1].matchesPlayed++
	t.exist[team1].matchesLost++
	t.exist[team2].matchesPlayed++
	t.exist[team2].matchesWon++
	t.exist[team2].points += 3
}

func (t *Table) update(res string) error {
	fields := strings.FieldsFunc(strings.TrimSpace(res), func(r rune) bool {
		return r == ';'
	})
	if len(fields) != 3 {
		return errors.New("wrong input format")
	}
	if fields[2] == "win" || fields[2] == "draw" || fields[2] == "loss" {
		_, ok1 := t.exist[fields[0]]
		_, ok2 := t.exist[fields[1]]
		if !ok1 {
			t.exist[fields[0]] = NewTeam(fields[0])
			t.teams = append(t.teams, t.exist[fields[0]])
		}
		if !ok2 {
			t.exist[fields[1]] = NewTeam(fields[1])
			t.teams = append(t.teams, t.exist[fields[1]])
		}
		switch fields[2] {
		case "win":
			t.win(fields[0], fields[1])
		case "draw":
			t.draw(fields[0], fields[1])
		case "loss":
			t.loss(fields[0], fields[1])
		}
	} else {
		return errors.New("wrong match result format")
	}
	return nil
}

// Tally reads the data from io.Reader and writes the table result to io.Writer.
func Tally(r io.Reader, w io.Writer) error {
	buf := make([]byte, 256)
	input := []byte{}
	for {
		n, err := r.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		input = append(input, buf[:n]...)
	}
	lines := strings.FieldsFunc(string(input), func(r rune) bool {
		return r == '\n'
	})
	table := NewTable()
	for _, line := range lines {
		if line[0] == '#' {
			continue
		}
		err := table.update(line)
		if err != nil {
			return err
		}
	}
	table.sort()
	output := bytes.NewBufferString(table.String())
	_, err := w.Write(output.Bytes())
	return err
}
