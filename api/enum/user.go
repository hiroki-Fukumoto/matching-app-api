package enum

type Sex string

var SEX = struct {
	MALE   Sex
	FEMALE Sex
	OTHER  Sex
}{
	MALE:   "male",
	FEMALE: "female",
	OTHER:  "other",
}
