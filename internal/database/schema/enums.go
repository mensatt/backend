package schema

type TagPriority string

const (
	Hide   TagPriority = "HIDE"
	Low    TagPriority = "LOW"
	Medium TagPriority = "MEDIUM"
	High   TagPriority = "HIGH"
)

func (TagPriority) Values() (kinds []string) {
	for _, s := range []TagPriority{Hide, Low, Medium, High} {
		kinds = append(kinds, string(s))
	}
	return
}
