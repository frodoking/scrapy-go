package src

// Indicates a missing configuration situation
type NotConfigured struct {
	err error
	msg string
}

// Indicates a decision was made not to process a request
type IgnoreRequest struct {
	err error
	msg string
}

// Request the spider not to be closed yet
type DontCloseSpider struct {
	err error
	msg string
}

// Raise this from callbacks to request the spider to be closed
type CloseSpider struct {
	err error
	msg string
}

// Drop item from the item pipeline
type DropItem struct {
	err error
	msg string
}

// Indicates a feature or method is not supported
type NotSupported struct {
	err error
	msg string
}

// To indicate a command-line usage error
type UsageError struct {
	err error
	msg string
}

// Warning category for deprecated features, since the default
type ScrapyDeprecationWarning struct {
	err error
	msg string
}

// Error raised in case of a failing contract
type ContractFail struct {
	err error
	msg string
}
