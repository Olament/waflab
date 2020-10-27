package object

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Testcase struct {
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Title        string `xorm:"varchar(100)" json:"title"`
	Method       string `xorm:"varchar(100)" json:"method"`
	UserAgent    string `xorm:"varchar(1000)" json:"userAgent"`
	QueryStrings []Pair `xorm:"mediumtext" json:"queryStrings"`
	Status       int    `json:"status"`
	TrueStatus   int    `json:"trueStatus"`
	Response     string `xorm:"mediumtext" json:"response"`
}

func GetTestcases() []*Testcase {
	testcases := []*Testcase{}
	err := ormManager.engine.Desc("created_time").Find(&testcases)
	if err != nil {
		panic(err)
	}

	return testcases
}

func GetFilteredTestcases(testsetId string) []*Testcase {
	testset := GetTestset(testsetId)
	m := map[string]int{}
	for _, item := range testset.Testcases {
		m[item.Name] = 1
	}

	testcases := GetTestcases()
	res := []*Testcase{}
	for _, testcase := range testcases {
		if _, ok := m[testcase.Name]; ok {
			res = append(res, testcase)
		}
	}
	return res
}

func GetTestcase(id string) *Testcase {
	s := Testcase{Name: id}
	existed, err := ormManager.engine.Get(&s)
	if err != nil {
		panic(err)
	}

	if existed {
		return &s
	} else {
		return nil
	}
}

func UpdateTestcase(id string, testcase *Testcase) bool {
	if GetTestcase(id) == nil {
		return false
	}

	_, err := ormManager.engine.Id(id).AllCols().Update(testcase)
	if err != nil {
		panic(err)
	}

	//return affected != 0
	return true
}

func AddTestcase(testcase *Testcase) bool {
	affected, err := ormManager.engine.Insert(testcase)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func DeleteTestcase(testcase *Testcase) bool {
	affected, err := ormManager.engine.Id(testcase.Name).Delete(&Testcase{})
	if err != nil {
		panic(err)
	}

	return affected != 0
}
