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

type OccurrenceStatus string

const (
	Confirmed        OccurrenceStatus = "CONFIRMED"
	Approved         OccurrenceStatus = "APPROVED"
	AwaitingApproval OccurrenceStatus = "AWAITING_APPROVAL"
	Updated          OccurrenceStatus = "UPDATED"
	PendingDeletion  OccurrenceStatus = "PENDING_DELETION"
)

func (OccurrenceStatus) Values() (kinds []string) {
	for _, s := range []OccurrenceStatus{Confirmed, Approved, AwaitingApproval, Updated, PendingDeletion} {
		kinds = append(kinds, string(s))
	}
	return
}
