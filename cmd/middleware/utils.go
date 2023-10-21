package middleware

type OptionsStruct struct {
	Operator  string `bson:"operator"`
	Action    string `bson:"action"`
	DryRun    string `bson:"dryRun"`
	Timestamp string `bson:"timestamp"`
}

func DataGenerator(o OptionsStruct) OptionsStruct {
	var res OptionsStruct

	res.Operator = o.Operator
	res.Action = o.Action
	if o.DryRun == "" {
		res.DryRun = "false"
	} else {
		res.DryRun = o.DryRun
	}

	return res
}
