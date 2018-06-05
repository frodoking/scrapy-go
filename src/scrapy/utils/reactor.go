package utils

/**
	Schedule a function to be called in the next reactor loop, but only if
    it hasn't been already scheduled since the last time it ran.
 */
type CallLaterOnce struct {
	function func()
}

func (c *CallLaterOnce) Schedule() {

}

func (c *CallLaterOnce) Cancel() {

}
