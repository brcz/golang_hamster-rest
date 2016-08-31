
package main

// Model is a sample basic data model.
type Model struct {
	get    func(filename string) interface{}
	getList func() interface{}
	add    func(items ...interface{})
	//put    func(id string, item interface{})
	delete func(filename string)
}

func proxyModel() Model {
    
	var model Model
	var records = Videos //[]VideoRecord{}
	
    search := func(id string) int {
		for i := range records {
			if records[i].ID == id {
				return i
			}
		}
		return -1
	}
	
    model.get = func(id string) interface{} {
		if i := search(id); i > -1 {
			return records[i]
		}
		return nil
	}
	
    model.getList = func() interface{} {
		return records
	}
    
   
	model.add = func(items ...interface{}) {
		for i := range items {
			records = append(records, items[i].(VideoRecord))
		}
	}
    
     /*
	model.put = func(id string, item interface{}) {
		if i := search(id); i > -1 {
			records[i] = item.(VideoRecord)
		}
	}
     */
    
	model.delete = func(id string) {
		if i := search(id); i > -1 {
			part := append(records[:i])
			if i < len(records)-1 {
				part = append(part, records[i+1:]...)
			}
			records = part
		}
	}
    
	return model
}
