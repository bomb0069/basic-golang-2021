package tasks_test

import (
	"tasks"
	"testing"
)

func Test_number_of_members_for_init_cli_should_0(t *testing.T) {
	cli := tasks.Cli{}

	if cli.NumberOfMember() != 0 {
		t.Errorf("Number of Member for Init Cli should be %v but %v", 0, cli.NumberOfMember())
	}
}

func Test_number_of_member_after_add_new_member_should_be_1(t *testing.T) {
	cli := tasks.Cli{}

	cli.AddMember(tasks.Member{Name: "Karan", Age: 15})
	if cli.NumberOfMember() != 1 {
		t.Errorf("Number of Member after add new member should be %v but %v", 1, cli.NumberOfMember())
	}
}
