package spider

import (
	"errors"
	"fmt"
)

const (
	uninit = iota
	inited
	started
	end
)

type Spider struct {
	start_urls        []string
	start_task_count  int
	finish_task_count int
	status            int
}

type Spider_info Spider

func (s *Spider) init() (err error) {
	if s.status != uninit {
		err_string := fmt.Sprintf("Spider is already inited!, Spider status is %+v", s.status)
		return errors.New(err_string)
	}

	s.start_urls = []string{}
	s.status = inited
}

func (s *Spider) set_start_urls(err error) {

}
