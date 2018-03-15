package alarms

import "testing"

const checkMark = "\u2713"
const ballotX = "\u2717"

func TestParse(t *testing.T) {

	var cases = []struct {
		input    string
		expected Alarm
	}{
		{"MemoryUtilization_production_aircall-web-42_aircall-web-42_i-00012c8d739386c8e", Alarm{
			Name:        "MemoryUtilization_production_aircall-web-42_aircall-web-42_i-00012c8d739386c8e",
			Project:     "aircall-web-42",
			Component:   "web",
			Environment: "production",
		}},
		{"MemoryUtilization_staging_aircall-strongbox_sidekiq_i-01fe9da63c0c79cb4", Alarm{
			Name:        "MemoryUtilization_staging_aircall-strongbox_sidekiq_i-01fe9da63c0c79cb4",
			Project:     "aircall-strongbox",
			Component:   "sidekiq",
			Environment: "staging",
		}},
		{"InstanceHealth_production_aircall-livecall_livecall_i-0107b66a0a538125a", Alarm{
			Name:        "InstanceHealth_production_aircall-livecall_livecall_i-0107b66a0a538125a",
			Project:     "aircall-livecall",
			Component:   "livecall",
			Environment: "production",
		}},
		{"InstanceHealth_production_aircall-strongbox_strongbox_i-07e51cfefd2efe7f3", Alarm{
			Name:        "InstanceHealth_production_aircall-strongbox_strongbox_i-07e51cfefd2efe7f3",
			Project:     "aircall-strongbox",
			Component:   "strongbox",
			Environment: "production",
		}},
	}

	for _, tc := range cases {
		t.Logf(tc.input)
		{
			actual, _ := Parse(tc.input)
			if !alarmMatches(actual, tc.expected) {
				t.Error(
					"expected", tc.expected,
					"got", actual,
					"%s", ballotX,
				)
			} else {
				t.Log(checkMark)
			}
		}
	}
}

func alarmMatches(a, b Alarm) bool {
	if a.Name != b.Name {
		return false
	}

	if a.Component != b.Component {
		return false
	}

	if a.Environment != b.Environment {
		return false
	}

	if a.Project != b.Project {
		return false
	}
	return true
}
