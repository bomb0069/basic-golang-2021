package tasks

type Cli struct {
	counter int
}

type Member struct {
	Name string
	Age  int
}

func (cli Cli) NumberOfMember() int {
	return cli.counter
}

func (cli *Cli) AddMember(member Member) int {
	cli.counter = cli.counter + 1
	return 0
}
