package chat

import "testing"

func Test_Command_Types(t *testing.T) {
	tests := []struct {
		desc           string
		command        Command
		expectedFormat string
	}{
		{
			desc:           "MessageCommand",
			command:        MSG{To: "Foo", Payload: []byte("hello")},
			expectedFormat: "MSG Foo 5\nhello",
		},
	}

	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			if string(tc.command.Format()) != tc.expectedFormat {
				t.Error("command format not equal")
			}
		})
	}
}
