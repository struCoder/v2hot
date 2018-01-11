package main

import (
  "log"
  "github.com/PuerkitoBio/goquery"
  "github.com/olekukonko/tablewriter"
  "os"
)

func main() {
  basePath := "https://www.v2ex.com"
  doc, err := goquery.NewDocument(basePath)
  if err != nil {
    log.Fatal(err)
  }

  table := tablewriter.NewWriter(os.Stdout)
  // Find the review items
  doc.Find(".content .item_hot_topic_title").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    linkText := s.Find("a").Text()
    
    linkSrc, _ := s.Find("a").Attr("href")
    table.Append([]string{linkText, basePath + linkSrc})
  })

 
  table.SetHeader([]string{"topic", "link"})
  
  table.Render()
}
