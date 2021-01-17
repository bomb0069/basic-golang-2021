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

func Test_add_first_member_and_should_get_1_for_member_id(t *testing.T) {
	cli := tasks.Cli{}

	memberId := cli.AddMember(tasks.Member{Name: "Karan", Age: 15})

	if memberId != 1 {
		t.Errorf("Member Id of new member should be %v but %v", 1, memberId)
	}
}

func Test_add_Karan_18_to_first_member_and_GetMember_at_1_should_found_Karan(t *testing.T) {
	cli := tasks.Cli{}

	cli.AddMember(tasks.Member{Name: "Karan", Age: 15})

	member := cli.GetMember(1)

	if member.Id != 1 {
		t.Errorf("Member Id of new member should be %v but %v", 1, member.Id)
	}
	if member.Name != "Karan" {
		t.Errorf("Member Name of new member should be %v but %v", "Karan", member.Name)
	}
	if member.Age != 15 {
		t.Errorf("Member Age of new member should be %v but %v", 15, member.Age)
	}
}
