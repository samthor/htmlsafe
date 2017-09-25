package htmlsafe

type state struct {
	stack  []string
	skipAt int
}

func (s *state) popTo(tag string) []string {
	var idx int
	if tag != "" {
		for idx = len(s.stack) - 1; idx >= 0; idx-- {
			if s.stack[idx] == tag {
				break
			}
			idx--
		}
		if idx < 0 {
			return nil
		}
	}

	if s.skipAt >= idx {
		// pop to the skipAt point
		s.stack = s.stack[:s.skipAt]
		s.skipAt = -1
	} else if s.skipAt >= 0 {
		// don't render anything
		s.stack = s.stack[:idx]
		return nil
	}

	// render closing
	out := s.stack[idx:]
	s.stack = s.stack[:idx]
	return out
}

func (s *state) push(tag string) bool {
	s.stack = append(s.stack, tag)
	if s.skipAt >= 0 {
		return true
	}

	// TODO: call to policy
	if tag == "script" || tag == "style" || tag == "link" {
		s.skipAt = len(s.stack) - 1
		return true
	}
	return false
}

func (s *state) topTag() string {
	l := len(s.stack)
	if l == 0 {
		return ""
	}
	return s.stack[l-1]
}

func (s *state) skipText() bool {
	return s.skipAt >= 0
}

func (s *state) skipNode(tag string) bool {
	return s.skipAt >= 0
}
