package cmd

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/spf13/cobra"
)

var err error
var output string

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var RootCmd = &cobra.Command{
	Use:   "envtpl",
	Short: "Render go templates from environment variables",
	Long:  `Render go templates from environment variables.`,
	Run: func(cmd *cobra.Command, args []string) {
		// load template; if an argument is not specified, default to stdin
		var t *template.Template
		if len(args) > 0 {
			t, err = parseFiles(args...)
			checkError(err)
		} else {
			bytes, err := ioutil.ReadAll(os.Stdin)
			checkError(err)
			t, err = parse(string(bytes))
			checkError(err)
		}

		// get environment variables to supply to the template
		env := readEnv()

		// get writer for rendered output; if an output file is not
		// specified, default to stdout
		var w io.Writer
		if len(output) > 0 {
			f, err := os.Create(output)
			checkError(err)
			defer f.Close()
			w = io.Writer(f)
		} else {
			w = os.Stdout
		}

		// render the template
		err := t.Execute(w, env)
		checkError(err)
	},
}

func Execute() {
	err := RootCmd.Execute()
	checkError(err)
}

func init() {
	RootCmd.Flags().StringVarP(&output, "output", "o", "", "The rendered output file")
}

func parse(s string) (*template.Template, error) {
	return template.New("").Funcs(sprig.TxtFuncMap()).Funcs(customFuncMap()).Parse(s)
}

func parseFiles(files ...string) (*template.Template, error) {
	return template.New(filepath.Base(files[0])).Funcs(sprig.TxtFuncMap()).Funcs(customFuncMap()).ParseFiles(files...)
}

func readEnv() (env map[string]string) {
	env = make(map[string]string)
	for _, setting := range os.Environ() {
		pair := strings.SplitN(setting, "=", 2)
		env[pair[0]] = pair[1]
	}
	return
}

// returns key, value for all environment variables starting with prefix
func environment(prefix string) map[string]string {
	env := make(map[string]string)
	for _, setting := range os.Environ() {
		pair := strings.SplitN(setting, "=", 2)
		if strings.HasPrefix(pair[0], prefix) {
			env[pair[0]] = pair[1]
		}
	}
	return env
}

func customFuncMap() template.FuncMap {
	var functionMap = map[string]interface{}{"environment": environment}
	return template.FuncMap(functionMap)
}
