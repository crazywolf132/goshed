package template

import (
	"fmt"

	"github.com/crazywolf132/goshed/internal/model"
)

var templates = map[string]*model.Template{
	"basic": {
		Name:        "basic",
		Description: "A basic Go program template",
		Files: map[string]string{
			"main.go": `package main

import "fmt"

func main() {
	fmt.Println("Hello, GoShed!")
}
`,
		},
	},
	"web": {
		Name:        "web",
		Description: "A basic web server template",
		Files: map[string]string{
			"main.go": `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, GoShed!")
	})

	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
`,
		},
	},
	"cli": {
		Name:        "cli",
		Description: "A command-line application template using Cobra",
		Files: map[string]string{
			"main.go": `package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "A brief description of your CLI application",
	Long: ` + "`" + `A longer description that spans multiple lines and likely contains
examples and usage of using your application.` + "`" + `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from your CLI app!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
`,
			"cmd/root.go": `package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	// Add your flags here
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mycli.yaml)")
}
`,
		},
		Dependencies: []string{
			"github.com/spf13/cobra",
			"github.com/spf13/viper",
		},
	},
	"api": {
		Name:        "api",
		Description: "A RESTful API template using Chi router",
		Files: map[string]string{
			"main.go": `package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Message string ` + "`json:\"message\"`" + `
}

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Welcome to the API"}
		json.NewEncoder(w).Encode(response)
	})

	log.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", r)
}
`,
		},
		Dependencies: []string{
			"github.com/go-chi/chi/v5",
		},
	},
	"graphql": {
		Name:        "graphql",
		Description: "A GraphQL API template using gqlgen",
		Files: map[string]string{
			"main.go": `package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("Connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
`,
			"graph/schema.graphqls": `type Query {
  hello: String!
}
`,
			"graph/resolver.go": `package graph

type Resolver struct{}
`,
		},
		Dependencies: []string{
			"github.com/99designs/gqlgen",
		},
	},
}

// Get returns a template by name
func Get(name string) (*model.Template, error) {
	t, ok := templates[name]
	if !ok {
		return nil, fmt.Errorf("template %s not found", name)
	}
	return t, nil
}

// List returns all available templates
func List() map[string]*model.Template {
	return templates
}
