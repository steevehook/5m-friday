package main

import "fmt"

type person struct {
	name string
}

type people []*person

func (p people) toPersonMap() personMap {
	peopleMap := make(personMap)
	for _, v := range p {
		peopleMap[v.name] = v
	}
	return peopleMap
}

func (p people) search(name string) *person {
	for _, v := range p {
		if v.name == name {
			return v
		}
	}
	return &person{}
}

func (p *people) add(name string) {
	var exists bool
	for _, v := range *p {
		if v.name == name {
			exists = true
			return
		}
	}
	if !exists {
		*p = append(*p, &person{name: name})
	}
}

func (p people) update(oldVal, newVal string) {
	for _, v := range p {
		if v.name == oldVal {
			v.name = newVal
		}
	}
}

func (p *people) delete(name string) {
	res := people{}
	for _, v := range *p {
		if v.name != name {
			res = append(res, v)
		}
	}
	*p = res
}

type personMap map[string]*person

func (pMap personMap) search(name string) *person {
	p, ok := pMap[name]
	if ok {
		return p
	}
	return &person{}
}

func (pMap personMap) add(name string) {
	pMap[name] = &person{name: name}
}

func (pMap personMap) update(oldVal, newVal string) {
	delete(pMap, oldVal)
	pMap[newVal] = &person{name: newVal}
}

func (pMap personMap) delete(name string) {
	delete(pMap, name)
}

func main() {
	p := people{
		{"John"},
		{"Steve"},
		{"Jane"},
		{"Jane"},
		{"Mike"},
	}
	fmt.Printf("SLICE:\n")
	steve := p.search("Steve")
	fmt.Println(steve.name)
	fmt.Println("len before:", len(p))

	p.add("Chris")
	fmt.Println("len after:", len(p))

	p.update(steve.name, "Stephen")
	fmt.Println(steve.name)

	fmt.Println("len before:", len(p)) // 6
	p.delete("Jane")
	fmt.Println("len after:", len(p)) // 4 (2 Jane(s) were removed)

	fmt.Printf("\nMAP:\n")
	pMap := p.toPersonMap()
	stephen := pMap.search("Stephen")
	fmt.Println(stephen.name)

	pMap.add("Sophie")
	fmt.Println(pMap.search("Sophie").name)

	pMap.update("Stephen", "Steve")
	fmt.Println(pMap.search("Steve"))

	fmt.Println("len before:", len(pMap))
	pMap.delete("Mike")
	fmt.Println("len after:", len(pMap))
}
