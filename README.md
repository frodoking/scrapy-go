# scrapy-go
rewrite scrapy with golang.

[scrapy architecture overview](https://doc.scrapy.org/en/latest/topics/architecture.html)

![scrapy architecture overview](https://doc.scrapy.org/en/latest/_images/scrapy_architecture_02.png)


The data flow in Scrapy is controlled by the execution engine, and goes like this:

> * The Engine gets the initial Requests to crawl from the Spider.
> * The Engine schedules the Requests in the Scheduler and asks for the next Requests to crawl.
> * The Scheduler returns the next Requests to the Engine.
> * The Engine sends the Requests to the Downloader, passing through the Downloader Middlewares (see process_request()).
> * Once the page finishes downloading the Downloader generates a Response (with that page) and sends it to the Engine, passing through > * the Downloader Middlewares (see process_response()).
> * The Engine receives the Response from the Downloader and sends it to the Spider for processing, passing through the Spider Middleware > * (see process_spider_input()).
> * The Spider processes the Response and returns scraped items and new Requests (to follow) to the Engine, passing through the Spider > * Middleware (see process_spider_output()).
> * The Engine sends processed items to Item Pipelines, then send processed Requests to the Scheduler and asks for possible next Requests > * to crawl.
> * The process repeats (from step 1) until there are no more requests from the Scheduler.


Related Projects

* [scrapy-parsel-go](https://github.com/frodoking/scrapy-parsel-go)
* [logrus](https://github.com/Sirupsen/logrus) Structured, pluggable logging for Go.