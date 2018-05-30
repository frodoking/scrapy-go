package scrapy

import (
	"os"
	"strings"
	"scrapy/commands"
	"scrapy/settings"
)

func execute(argv []string, mSettings *settings.CrawlerSettings) {
	if argv == nil {
		argv = os.Args
	}

	if mSettings == nil {
		mSettings = settings.NewCrawlerSettings()
	}

	if mSettings == nil {
		mSettings = GetProjectSettings()
	}

	inproject := InsideProject()
	cmds := getCommandsDict(mSettings, inproject)

	cmdName, argv := popCommandName(argv)
	if len(cmdName) == 0 {
		panic("cmd name empty.")
		os.Exit(0)
	}

	cmdNameIncluded := false
	for _, cmd := range cmds {
		if strings.EqualFold(cmd.Name, cmdName) {
			cmdNameIncluded = true
		}
	}
	if !cmdNameIncluded {
		print("xxxxx")
		os.Exit(2)
	}

	cmd := cmds[cmdName]
	cmd.Syntax()
	cmd.LongDesc()
}

func iterCommandClasses(moduleName string) []commands.ScrapyCommand {
	return make([]commands.ScrapyCommand, 10)
}

func getCommandsFromModule(module string, inproject bool) map[string]commands.ScrapyCommand {
	d := make(map[string]commands.ScrapyCommand)
	cmds := iterCommandClasses(module)
	for _, cmd := range cmds {
		if inproject || cmd.RequiresProject {
			cmdName := ""
			d[cmdName] = cmd
		}
	}
	return d
}

func getCommandsFromEntryPoints(inproject bool) map[string]commands.ScrapyCommand {
	cmds := make(map[string]commands.ScrapyCommand)
	return cmds
}

func getCommandsDict(mSettings *settings.CrawlerSettings, inproject bool) map[string]commands.ScrapyCommand {
	cmds := getCommandsFromModule("scrapy.commands", inproject)
	cmds2 := getCommandsFromEntryPoints(inproject)
	for cmd := range cmds2 {
		cmds[cmd] = cmds2[cmd]
	}
	cmdsModuleName := mSettings.Get("COMMANDS_MODULE")
	if len(cmdsModuleName) > 0 {
		var cmdsModule = getCommandsFromModule(cmdsModuleName, inproject)
		for cmd := range cmdsModule {
			cmds[cmd] = cmds2[cmd]
		}
	}
	return cmds
}

func popCommandName(argv []string) (string, []string) {
	for i, arg := range argv {
		if strings.HasPrefix(arg, "-") {
			return arg, append(argv[:i], argv[i+1:]...)
		}
	}
	return "", argv
}
