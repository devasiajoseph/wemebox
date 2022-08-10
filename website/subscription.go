package website

type Subscription struct {
	SubsID   int `db:"subs_id" json:"subs_id"`
	DomainID int `db:"domain_id" json:"domain_id"`
	Credits  int `db:"credits" json:"credits"`
}

func AddCredits(dID int, credits int) error {

}
