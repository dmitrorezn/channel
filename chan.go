package channel

func SelectN[T any](ch ...chan T) chan T {
	result := make(chan T)
	if len(ch) == 0 {
		close(result)

		return result
	}

	switch len(ch) {
	case 1:
		result = ch[0]

		return result
	case 2:
		go func() {
			var v T
			select {
			case v = <-ch[0]:
			case v = <-ch[1]:
			}
			select {
			case result <- v:
			default:
			}
		}()

		return result
	}
	go func() {
		var v T
		select {
		case v = <-SelectN(ch[:2]...):
		case v = <-SelectN(ch[2:]...):
		}
		select {
		case result <- v:
		default:
		}
	}()

	return result
}
