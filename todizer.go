package todizer

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cursorStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#DF8965"))
var checkedStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#54AEFF"))
var headerStyle = lipgloss.NewStyle().Bold(true)

func Execute(menu *Menu) error {
	prog := tea.NewProgram(menu, tea.WithAltScreen())
	if _, err := prog.Run(); err != nil {
		return err
	}
	return nil
}

var header = "TODOS"

type Menu struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	Input    io.Reader
	Output   io.Writer
}

func (self *Menu) Init() tea.Cmd {
	return nil
}

func (self *Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+[":
			self.flush()
			return self, tea.Quit
		case "up", "k", "ctrl+p":
			if self.cursor > 0 {
				self.cursor--
			}
		case "down", "j", "ctrl+n":
			if self.cursor < len(self.choices) {
				self.cursor++
			}
		case "enter", "ctrl+y", " ":
			// toggle selection
			if _, ok := self.selected[self.cursor]; ok {
				delete(self.selected, self.cursor)
			} else {
				self.selected[self.cursor] = struct{}{}
			}
		}
	}
	return self, nil
}

func (self *Menu) View() string {
	var data = []string{headerStyle.Render(strings.ToUpper(header))}
	for i, item := range self.choices {
		row := self.format(i, item, true)
		data = append(data, row)
	}
	return strings.Join(data, "\n")
}

func (self *Menu) flush() {
	if self.Output == nil {
		return
	}
	for i, row := range self.choices {
		row := self.format(i, row, false)
		row = strings.TrimPrefix(row, cursorStyle.Render(">"))
		row = strings.Trim(row, " ")
		row = fmt.Sprintf("- %s", row)
		fmt.Fprintln(self.Output, row)
	}
}

func (self *Menu) format(idx int, item string, colorized bool) string {
	// mark he position of the cursor
	cursor := ""
	if self.cursor == idx {
		cursor = cursorStyle.Render(">")
	}

	// check current item if it is selected
	checked := "[ ]"
	if _, ok := self.selected[idx]; ok {
		checked = "[x]"
		if colorized {
			checked = checkedStyle.Render(checked)
		}
	}

	return fmt.Sprintf("%s %s %s", cursor, checked, item)
}

func New(input io.Reader, ouput io.Writer) *Menu {
	var choices []string

	sc := bufio.NewScanner(input)
	for sc.Scan() {
		choices = append(choices, sc.Text())
	}

	return &Menu{
		choices:  choices,
		selected: make(map[int]struct{}),
		Input:    input,
		Output:   ouput,
	}
}
