package gobelt

type Generator interface {
  Generate(args []string) error
}
