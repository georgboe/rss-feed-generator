package podcast_test

import (
	"fmt"

	podcast "github.com/georgboe/rss-feed-generator"
)

func ExampleNew() {
	ti, l, d := "title", "link", podcast.Description{Text: "description"}

	// instantiate a new Podcast
	p := podcast.New(ti, l, d, &pubDate, &updatedDate)
	p.AddLanguage("en-us")

	fmt.Println(p.Title, p.Link, p.Description.Text, p.Language)
	// Output:
	// title link description en-us
}

func ExamplePodcast_AddAuthor() {
	p := podcast.New("title", "link", podcast.Description{Text: "Description"}, nil, nil)

	// add the Author
	// p.AddAuthor("the name", "me@test.com")
	p.AddAuthor([]string{"the name"})

	fmt.Println(p.ManagingEditor)
	fmt.Println(p.IAuthor)
	// Output:
	// the name
}

func ExamplePodcast_AddCategory() {
	p := podcast.New("title", "link", podcast.Description{Text: "Description"}, nil, nil)

	// add the Category
	p.AddCategory("Bombay", nil)
	p.AddCategory("American", []string{"Longhair", "Shorthair"})
	p.AddCategory("Siamese", nil)

	fmt.Println(len(p.ICategories), len(p.ICategories[1].ICategories))
	// Output:
	// 3 2
}

func ExamplePodcast_AddImage() {
	p := podcast.New("title", "link", podcast.Description{Text: "Description"}, nil, nil)

	// add the Image
	p.AddImage("http://example.com/image.jpg")

	if p.Image != nil {
		fmt.Println(p.Image.URL)
	}
	// Output:
	// http://example.com/image.jpg
}

// func ExamplePodcast_AddItem() {
// 	p := podcast.New("title", "link", "description", &pubDate, &updatedDate)
// 	// p.AddAuthor("the name", "me@test.com")
// 	p.AddAuthor([]string{"the name"})
// 	p.AddImage("http://example.com/image.jpg")

// 	// create an Item
// 	item := podcast.Item{
// 		Title:       "Episode 1",
// 		Description: "Description for Episode 1",
// 		ISubtitle:   "A simple episode 1",
// 	}
// 	item.AddEnclosure(
// 		"http://example.com/1.mp3",
// 		"audio/mpeg",
// 		183,
// 	)
// 	item.AddSummary("See more at <a href=\"http://example.com\">Here</a>")

// 	// add the Item
// 	if _, err := p.AddItem(item); err != nil {
// 		fmt.Println("item validation error: " + err.Error())
// 	}

//		if len(p.Items) != 1 {
//			fmt.Println("expected 1 item in the collection")
//		}
//		pp := p.Items[0]
//		fmt.Println(
//			pp.GUID, pp.Title, pp.Link, pp.Description,
//			pp.AuthorFormatted, pp.Category, pp.Comments, pp.Source, *pp.Enclosure,
//			pp.IAuthor, pp.IDuration, pp.IExplicit, pp.IIsClosedCaptioned,
//			pp.IOrder, pp.ISubtitle, pp.ISummary)
//		// Output:
//		// http://example.com/1.mp3 Episode 1 http://example.com/1.mp3 Description for Episode 1     {{ } http://example.com/1.mp3 183 183 audio/mpeg audio/mpeg} the name     A simple episode 1 &{{ } See more at <a href="http://example.com">Here</a>}
//	}
func ExamplePodcast_AddSummary() {
	p := podcast.New("title", "link", podcast.Description{Text: "Description"}, nil, nil)

	// add a summary
	p.AddSummary(`A very cool podcast with a long summary!

See more at our website: <a href="http://example.com">example.com</a>
`)

	if p.ISummary != nil {
		fmt.Println(p.ISummary.Text)
	}
	// Output:
	// A very cool podcast with a long summary!
	//
	// See more at our website: <a href="http://example.com">example.com</a>
}

// func ExamplePodcast_Bytes() {
// 	p := podcast.New(
// 		"eduncan911 Podcasts",
// 		"http://eduncan911.com/",
// 		"An example Podcast",
// 		&pubDate, &updatedDate,
// 	)
// 	p.AddAuthor("Jane Doe", "me@janedoe.com")
// 	p.AddImage("http://janedoe.com/i.jpg")
// 	p.AddSummary(`A very cool podcast with a long summary using Bytes()!

// See more at our website: <a href="http://example.com">example.com</a>
// `)
// 	p.AddLanguage("en-us")

// 	for i := int64(5); i < 7; i++ {
// 		n := strconv.FormatInt(i, 10)
// 		d := pubDate.AddDate(0, 0, int(i+3))

// 		item := podcast.Item{
// 			Title:       "Episode " + n,
// 			Link:        "http://example.com/" + n + ".mp3",
// 			Description: "Description for Episode " + n,
// 			PubDate:     &d,
// 		}
// 		if _, err := p.AddItem(item); err != nil {
// 			fmt.Println(item.Title, ": error", err.Error())
// 			break
// 		}
// 	}

// 	// call Podcast.Bytes() to return a byte array
// 	os.Stdout.Write(p.Bytes())

// 	// Output:
// 	// <?xml version="1.0" encoding="UTF-8"?>
// 	// <rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
// 	//   <channel>
// 	//     <title>eduncan911 Podcasts</title>
// 	//     <link>http://eduncan911.com/</link>
// 	//     <description>An example Podcast</description>
// 	//     <generator>go podcast v1.3.1 (github.com/eduncan911/podcast)</generator>
// 	//     <language>en-us</language>
// 	//     <lastBuildDate>Mon, 06 Feb 2017 08:21:52 +0000</lastBuildDate>
// 	//     <managingEditor>me@janedoe.com (Jane Doe)</managingEditor>
// 	//     <pubDate>Sat, 04 Feb 2017 08:21:52 +0000</pubDate>
// 	//     <image>
// 	//       <url>http://janedoe.com/i.jpg</url>
// 	//       <title>eduncan911 Podcasts</title>
// 	//       <link>http://eduncan911.com/</link>
// 	//     </image>
// 	//     <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
// 	//     <itunes:summary><![CDATA[A very cool podcast with a long summary using Bytes()!
// 	//
// 	// See more at our website: <a href="http://example.com">example.com</a>
// 	// ]]></itunes:summary>
// 	//     <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
// 	//     <item>
// 	//       <guid>http://example.com/5.mp3</guid>
// 	//       <title>Episode 5</title>
// 	//       <link>http://example.com/5.mp3</link>
// 	//       <description>Description for Episode 5</description>
// 	//       <pubDate>Sun, 12 Feb 2017 08:21:52 +0000</pubDate>
// 	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
// 	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
// 	//     </item>
// 	//     <item>
// 	//       <guid>http://example.com/6.mp3</guid>
// 	//       <title>Episode 6</title>
// 	//       <link>http://example.com/6.mp3</link>
// 	//       <description>Description for Episode 6</description>
// 	//       <pubDate>Mon, 13 Feb 2017 08:21:52 +0000</pubDate>
// 	//       <itunes:author>me@janedoe.com (Jane Doe)</itunes:author>
// 	//       <itunes:image href="http://janedoe.com/i.jpg"></itunes:image>
// 	//     </item>
// 	//   </channel>
// 	// </rss>
// }

// func ExampleItem_AddPubDate() {
// 	p := podcast.New("title", "link", "description", nil, nil)
// 	i := podcast.Item{
// 		Title:       "item title",
// 		Description: "item desc",
// 		Link:        "item link",
// 	}
// 	d := pubDate.AddDate(0, 0, -11)

// 	// add the pub date
// 	i.AddPubDate(&d)

// 	// before adding
// 	if i.PubDate != nil {
// 		fmt.Println(i.PubDateFormatted, *i.PubDate)
// 	}

// 	// this should not override with Podcast.PubDate
// 	if _, err := p.AddItem(i); err != nil {
// 		fmt.Println(err)
// 	}

// 	// after adding item
// 	fmt.Println(i.PubDateFormatted, *i.PubDate)
// 	// Output:
// 	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
// 	// Tue, 24 Jan 2017 08:21:52 +0000 2017-01-24 08:21:52 +0000 UTC
// }

func ExampleItem_AddDuration() {
	i := podcast.Item{
		Title:       "item title",
		Description: &podcast.Description{Text: "item desc"},
		Link:        "item link",
	}
	d := int64(533)

	// add the Duration in Seconds
	i.AddDuration(d)

	fmt.Println(i.IDuration)
	// Output:
	// 533
}
