package htmlparser

import "fmt"

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

func (list LinkList) Each(function func(link Hyperlink)) LinkList {
	for _, link := range list {
		function(link)
	}
	return list
}

func (list LinkList) PrintAll() LinkList {
	return list.Each(func(link Hyperlink) { fmt.Println(link) })
}
