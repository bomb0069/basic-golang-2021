package tasks

type Cli struct {
	counter int
}

type Member struct {
	Id   int
	Name string
	Age  int
}

func (cli Cli) NumberOfMember() int {
	return cli.counter
}

func (cli *Cli) AddMember(member Member) int {
	cli.counter = cli.counter + 1
	return cli.counter
}

func (cli *Cli) GetMember(id int) Member {
	return Member{Id: 1, Name: "Karan", Age: 15}
}
