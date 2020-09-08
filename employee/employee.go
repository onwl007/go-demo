package employee

import (
	"fmt"
)

type employee struct {
	firstName   string
	lastName    string
	totalLeave  int
	leavesTaken int
}

// 相当于构造函数
func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeave - e.leavesTaken))
}
