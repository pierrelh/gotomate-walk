package scraping

import (
	"fmt"
	"gotomate/fiber/variable"
	"reflect"

	"github.com/gocolly/colly"
)

// AddAllowedDomain Setting allowed Domains for scraper
func AddAllowedDomain(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	scraper.AllowedDomains = append(scraper.AllowedDomains, path)
	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// AddDisallowedDomain Setting disallowed Domains for scraper
func AddDisallowedDomain(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var path string
	if pathIsVar := instructionData.FieldByName("PathIsVar").Interface().(bool); pathIsVar {
		pathVarName := instructionData.FieldByName("PathVarName").Interface().(string)
		if val := variable.SearchVariable(pathVarName).Value; val != nil {
			path = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", pathVarName)
			finished <- true
			return -1
		}
	} else {
		path = instructionData.FieldByName("Path").Interface().(string)
	}

	scraper.DisallowedDomains = append(scraper.DisallowedDomains, path)
	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// IgnoreRobotsTxt Setting ignoreRobotsTxt Domains for scraper
func IgnoreRobotsTxt(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var ignore bool
	if ignoreIsVar := instructionData.FieldByName("IgnoreIsVar").Interface().(bool); ignoreIsVar {
		ignoreVarName := instructionData.FieldByName("IgnoreVarName").Interface().(string)
		if val := variable.SearchVariable(ignoreVarName).Value; val != nil {
			ignore = val.(bool)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", ignoreVarName)
			finished <- true
			return -1
		}
	} else {
		ignore = instructionData.FieldByName("Ignore").Interface().(bool)
	}

	scraper.IgnoreRobotsTxt = ignore
	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// NewScraper Create a new scraper
func NewScraper(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	scraper := colly.NewCollector()
	variable.SetVariable(instructionData.FieldByName("Output").Interface().(string), scraper)
	finished <- true
	return -1
}

// OnFindScrapAttribute Scrap a target's attribute when the wanted element if found
func OnFindScrapAttribute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	var attribute string
	if attributeIsVar := instructionData.FieldByName("AttributeIsVar").Interface().(bool); attributeIsVar {
		attributeVarName := instructionData.FieldByName("AttributeVarName").Interface().(string)
		if val := variable.SearchVariable(attributeVarName).Value; val != nil {
			attribute = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", attributeVarName)
			finished <- true
			return -1
		}
	} else {
		attribute = instructionData.FieldByName("Attribute").Interface().(string)
	}

	var tab []string
	tabVarName := instructionData.FieldByName("TabVarName").Interface().(string)
	if val := variable.SearchVariable(tabVarName).Value; val != nil {
		tab = val.([]string)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", tabVarName)
		finished <- true
		return -1
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		value := e.Attr(attribute)
		tab = append(tab, value)
		variable.SetVariable(tabVarName, tab)
	})

	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// OnFindScrapChildAttribute Scrap a target's children attribute when the wanted element if found
func OnFindScrapChildAttribute(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	var childAttribute string
	if childAttributeIsVar := instructionData.FieldByName("ChildAttributeIsVar").Interface().(bool); childAttributeIsVar {
		childAttributeVarName := instructionData.FieldByName("ChildAttributeVarName").Interface().(string)
		if val := variable.SearchVariable(childAttributeVarName).Value; val != nil {
			childAttribute = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", childAttributeVarName)
			finished <- true
			return -1
		}
	} else {
		childAttribute = instructionData.FieldByName("ChildAttribute").Interface().(string)
	}

	var attribute string
	if attributeIsVar := instructionData.FieldByName("AttributeIsVar").Interface().(bool); attributeIsVar {
		attributeVarName := instructionData.FieldByName("AttributeVarName").Interface().(string)
		if val := variable.SearchVariable(attributeVarName).Value; val != nil {
			attribute = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", attributeVarName)
			finished <- true
			return -1
		}
	} else {
		attribute = instructionData.FieldByName("Attribute").Interface().(string)
	}

	var tab []string
	tabVarName := instructionData.FieldByName("TabVarName").Interface().(string)
	if val := variable.SearchVariable(tabVarName).Value; val != nil {
		tab = val.([]string)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", tabVarName)
		finished <- true
		return -1
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		value := e.ChildAttr(childAttribute, attribute)
		tab = append(tab, value)
		variable.SetVariable(tabVarName, tab)
	})

	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// OnFindScrapText Scrap a target's text when the wanted element if found
func OnFindScrapText(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	var tab []string
	tabVarName := instructionData.FieldByName("TabVarName").Interface().(string)
	if val := variable.SearchVariable(tabVarName).Value; val != nil {
		tab = val.([]string)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", tabVarName)
		finished <- true
		return -1
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		value := e.Text
		tab = append(tab, value)
		variable.SetVariable(tabVarName, tab)
	})
	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// OnFindScrapChildText Scrap a target's children text when the wanted element if found
func OnFindScrapChildText(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var childAttribute string
	if childAttributeIsVar := instructionData.FieldByName("ChildAttributeIsVar").Interface().(bool); childAttributeIsVar {
		childAttributeVarName := instructionData.FieldByName("ChildAttributeVarName").Interface().(string)
		if val := variable.SearchVariable(childAttributeVarName).Value; val != nil {
			childAttribute = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", childAttributeVarName)
			finished <- true
			return -1
		}
	} else {
		childAttribute = instructionData.FieldByName("ChildAttribute").Interface().(string)
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	var tab []string
	tabVarName := instructionData.FieldByName("TabVarName").Interface().(string)
	if val := variable.SearchVariable(tabVarName).Value; val != nil {
		tab = val.([]string)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", tabVarName)
		finished <- true
		return -1
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		value := e.ChildText(childAttribute)
		tab = append(tab, value)
		variable.SetVariable(tabVarName, tab)
	})
	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// OnFindVisit Visit a target when the wanted element if found
func OnFindVisit(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		t := e.Attr("href")
		scraper.Visit(t)
	})

	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// OnFindChildVisit Visit a target's children when the wanted element if found
func OnFindChildVisit(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var childAttribute string
	if childAttributeIsVar := instructionData.FieldByName("ChildAttributeIsVar").Interface().(bool); childAttributeIsVar {
		childAttributeVarName := instructionData.FieldByName("ChildAttributeVarName").Interface().(string)
		if val := variable.SearchVariable(childAttributeVarName).Value; val != nil {
			childAttribute = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", childAttributeVarName)
			finished <- true
			return -1
		}
	} else {
		childAttribute = instructionData.FieldByName("ChildAttribute").Interface().(string)
	}

	var element string
	if elementIsVar := instructionData.FieldByName("ElementIsVar").Interface().(bool); elementIsVar {
		elementVarName := instructionData.FieldByName("ElementVarName").Interface().(string)
		if val := variable.SearchVariable(elementVarName).Value; val != nil {
			element = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", elementVarName)
			finished <- true
			return -1
		}
	} else {
		element = instructionData.FieldByName("Element").Interface().(string)
	}

	scraper.OnHTML(element, func(e *colly.HTMLElement) {
		t := e.ChildAttr(childAttribute, "href")
		scraper.Visit(t)
	})

	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// ScraperEndCondition Create a new scraper
func ScraperEndCondition(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	scraper.OnResponse(func(r *colly.Response) {
		var end int64
		if endIsVar := instructionData.FieldByName("EndIsVar").Interface().(bool); endIsVar {
			endVarName := instructionData.FieldByName("EndVarName").Interface().(string)
			if val := variable.SearchVariable(endVarName).Value; val != nil {
				end = val.(int64)
			}
		} else {
			end = instructionData.FieldByName("End").Interface().(int64)
		}

		end--

		if end <= 0 {
			panic("Scrap finished")
		}
		if endIsVar := instructionData.FieldByName("EndIsVar").Interface().(bool); endIsVar {
			endVarName := instructionData.FieldByName("EndVarName").Interface().(string)
			variable.SetVariable(endVarName, end)
		} else {
			instructionData.FieldByName("End").SetInt(end)
		}

	})

	variable.SetVariable(scraperVarName, scraper)
	finished <- true
	return -1
}

// ScrapStart Visit a url to start a scraping
func ScrapStart(instructionData reflect.Value, finished chan bool) int {
	fmt.Println("FIBER INFO: Starting new process ...")

	var scraper *colly.Collector
	scraperVarName := instructionData.FieldByName("ScraperVarName").Interface().(string)
	if val := variable.SearchVariable(scraperVarName).Value; val != nil {
		scraper = val.(*colly.Collector)
	} else {
		fmt.Println("FIBER WARNING: Unable to find var ...", scraperVarName)
		finished <- true
		return -1
	}

	var url string
	if urlIsVar := instructionData.FieldByName("UrlIsVar").Interface().(bool); urlIsVar {
		urlVarName := instructionData.FieldByName("UrlVarName").Interface().(string)
		if val := variable.SearchVariable(urlVarName).Value; val != nil {
			url = val.(string)
		} else {
			fmt.Println("FIBER WARNING: Unable to find var ...", urlVarName)
			finished <- true
			return -1
		}
	} else {
		url = instructionData.FieldByName("Url").Interface().(string)
	}

	scraper.Visit(url)
	scraper.Wait()
	finished <- true
	return -1
}
