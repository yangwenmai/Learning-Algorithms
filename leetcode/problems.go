package leetcode

type problems []Problem

func (ps *problems) add(p Problem) {
	if len(*ps) <= p.ID {
		*ps = append(*ps, make([]Problem, p.ID-len(*ps)+1)...)
	}
	(*ps)[p.ID] = p
}

func (ps problems) accepted() problems {
	res := make([]Problem, 0, len(ps))
	for _, p := range ps {
		if p.IsAccepted {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) available() problems {
	res := make([]Problem, 0, len(ps))
	size := len(ps)
	for i := 1; i < size; i++ {
		p := ps[i]
		if p.IsAvailable() {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) completed() problems {
	var res []Problem
	for i := 0; i < len(ps); i++ {
		p := ps[i]
		if p.IsAccepted {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) favorite() problems {
	res := make([]Problem, 0, len(ps))
	size := len(ps)
	for i := 0; i < size; i++ {
		p := ps[i]
		if p.IsFavor {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) unavailable() problems {
	res := make([]Problem, 0, len(ps))
	for _, p := range ps {
		if p.HasNoGoOption {
			res = append(res, p)
		}
	}
	return res
}

func (ps problems) table() string {
	res := "|题号|题目|通过率|难度|收藏|\n"
	res += "|:-:|:-|:-: | :-: | :-: |\n"
	for _, p := range ps {
		res += p.tableLine()
	}
	return res
}

func (ps problems) list() string {
	res := ""
	for _, p := range ps {
		res += p.listLine()
	}
	return res
}
