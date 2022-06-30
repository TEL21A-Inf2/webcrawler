package htmlparser

type LinkList []Hyperlink

func (list LinkList) Filter(pred func(link Hyperlink) bool) LinkList {
	result := make(LinkList, 0)
	for _, link := range list {
		if pred(link) {
			result = append(result, link)
		}
	}
	return result
}
