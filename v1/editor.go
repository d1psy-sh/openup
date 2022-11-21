package v1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type EditorListModel struct {
	list list.Model
}

func EditorChoiceUpdate(m model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	m.Editor.Focus()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
                        m.state = ItemList
			return m, nil
		case "enter":
			index, _ := strconv.Atoi(m.Editor.Value())
                        editor = editors[index]
			data.Editor = editors[index]
			saveData(data)
			data = data.GetData()
			m.state = ItemList
			return m, nil
		case "cntl+l":
			m.state = ItemList
			return m, nil
		}
	case tea.WindowSizeMsg:
		h, v := dockstyle.GetFrameSize()
		m.ListModel.list.SetSize(msg.Width-h, msg.Height-v)
	}

	m.Editor, cmd = m.Editor.Update(msg)

	return m, cmd
}

func renderEditor(m model) string {

	var b strings.Builder

	b.WriteString("Pick your editor of choice.\n")
	for i := 0; i < len(editors); i++ {
                if editors[i] == editor {
                        fmt.Fprintf(&b, accentColor2Text.Render("\n"+fmt.Sprintf("%v",i) +"."+editor+"\n"))
        }else{
		fmt.Fprintf(&b, "\n%v.%s\n", i, editors[i])
        }
	}


	fmt.Fprintf(&b, "\n\n%s\n", m.Editor.View())

	return b.String()
}
