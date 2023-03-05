package model

import "time"

// Contest has Name, ID, Kind, State.
type Contest struct {
	Name  string
	ID    string // like "abc123"
	Kind  string // like "Algorithm", "Heuristics"...
	State string // "permanent", "upcoming", "archive"
	// TODO: StateをEnumに
}

type Task struct {
	Name   string
	IdName string // "A", "90"...
	ID     string // "abc123_a", "typical90_cl"...
}

// TaskInfo is detailed information of a task.
// It is target to express [commonest task page] neither too much nor little.
type TaskInfo struct {
	ProblemStatement     string
	ProblemStatementHTML string
	Constraints          string
	ConstraintsHTML      string
	IoStyle              IoStyle
	TestCases            []*TestCase
}

// IoStyle represents input and output signature.
// Input and Output are machine readable sections.
// ~Desc is Description of ones.
type IoStyle struct {
	InputSig, OutputSig   string
	InputDesc, OutputDesc string
}

// TestCase.
// Can be used to ascertain (.Input > `user program` == .Output).
type TestCase struct{ Input, Output string }

// Language stand for
// "<option value={Language.Value} data-mime={Language.Datamime}>{Language.Text}</option>"
type Language struct {
	Value    string
	Datamime string
	Text     string
}

type Submission struct {
	ID           string
	Time         time.Time
	TaskName     string
	UserID       string
	LanguageName string
	Score        string
	CodeSize     string
	Status       string
	ExecTime     string
	Memory       string
}
