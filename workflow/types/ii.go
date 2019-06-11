package types

import "fmt"

var (
	IndexSubmitForm = "gosion-workflow-forms"
	DBFlow          = "gosion-workflow"
)

func GetSubmitFormCollection(instanceId, nodeId string) string {
	return fmt.Sprintf("submit_form_%s_%s", instanceId[:1], nodeId[:1])
}
