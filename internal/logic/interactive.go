// Package logic provides ...
package logic

import (
	"github.com/glepnir/jarvis/internal/render"
	"github.com/glepnir/jarvis/internal/render/dein"
	"github.com/glepnir/jarvis/internal/vim"
	"github.com/glepnir/jarvis/pkg/cli"
)

func PluginManage() render.Render {
	message := "What is plugin manage do you use?"
	options := []string{"dein", "vim-plug"}
	pm := cli.SingleSelectTemplate(message, options)
	if pm == "dein" {
		return new(dein.Dein)
	}
	return nil
}

func LeaderKey() string {
	message := "What is your Leader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

func LocalLeaderKey() string {
	message := "What is your LocalLeader Key?"
	options := []string{"Space", "Comma(,)", "Semicolon(;)"}
	return cli.SingleSelectTemplate(message, options)
}

func Colorscheme() []string {
	questionname := "Colorscheme Question"
	message := "Choose your favorite colorscheme"
	pagesize := 19
	options := make([]string, 0)
	for k, _ := range vim.ColorschemeMap {
		options = append(options, k)
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func DashboardPlugin() bool {
	message := "Do you want use dashboard-nvim a better StartScreenPlugin?"
	return cli.ConfirmTemplate(message)
}

func BufferLinePlugin() bool {
	message := "Do you want use vim-buffet as your bufferline?"
	return cli.ConfirmTemplate(message)
}

func SpacelinePlugin() bool {
	message := "Do you want use spaceline.vim a light and beautiful statusline?"
	return cli.ConfirmTemplate(message)
}

func ExplorerPlugin() string {
	message := "What is your explorer plugin?"
	options := []string{"defx.nvim", "nerdtree", "coc-explorer"}
	return cli.SingleSelectTemplate(message, options)
}

func DatabasePlugin() bool {
	message := "Do you need database plugins?"
	return cli.ConfirmTemplate(message)
}

func FuzzyFindPlugin() bool {
	message := "Do you want use fuzzy find plugin vim-clap?"
	return cli.ConfirmTemplate(message)
}

func EditorConfigPlugin() bool {
	message := "Do you want use editorconfig to control program style(like indent,whitespace etc)"
	return cli.ConfirmTemplate(message)
}

func IndentLinePlugin() string {
	message := "Choose  your favorite indenline plugin?"
	options := []string{"Yggdroot/indentLine", "nathanaelkane/vim-indent-guides"}
	return cli.SingleSelectTemplate(message, options)
}

func CommentPlugin() bool {
	message := "Do you want to use Caw.vim as comment plugin?"
	return cli.ConfirmTemplate(message)
}

func ViewSymbolsPlugin() bool {
	message := "Do you want to use vista.vim to view tags and LSP symbols in sidebar"
	return cli.ConfirmTemplate(message)
}

func GentagsPlugin() bool {
	message := "Do you want to use vim-gutentags to gen tags"
	return cli.ConfirmTemplate(message)
}

func QuickRunPlugin() bool {
	message := "Do you want to use vim-quickrun to fast run program in vim?"
	return cli.ConfirmTemplate(message)
}

func DataTypeFile() []string {
	questionname := "Data filetype"
	message := "Which Data filetype you need?"
	pagesize := 10
	options := make([]string, 0)
	for k, _ := range vim.DataFileMap {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func EnhancePlugin() []string {
	questionname := "Enhance question"
	message := "Choose the enhace plugins that you need "
	pagesize := 10
	options := make([]string, 0)
	for k, _ := range vim.EnhancePluginMap {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func SandWichPlugin() bool {
	message := "Do you want use vim-sandwich more useful than vim-surround?"
	return cli.ConfirmTemplate(message)
}

func VersionControlPlugin() []string {
	questionname := "Version Control plugin"
	message := "Choose the version control plugins that you need"
	pagesize := 10

	options := make([]string, 0)
	for k, _ := range vim.VersionControlMap {
		options = append(options, k)
	}

	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}

func LanguageServerProtocol() []string {
	questionname := "LanguageQuestion"
	message := "What Languages do you write"
	pagesize := 19
	options := make([]string, 0)
	for k, _ := range vim.LanguagesPluginMap {
		options = append(options, k)
	}
	return cli.MultiSelectTemplate(questionname, message, options, pagesize)
}
