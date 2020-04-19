package erratum

// Use opens a resource, calls func Frob() in the resource and handles errors are occurred.
func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	res, err = o()
	if err != nil {
		switch err.(type) {
		case TransientError:
			return Use(o, input)
		}
		return
	}
	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			if val, ok := r.(FrobError); ok {
				res.Defrob(val.defrobTag)
			}
			err = r.(error)
		}
	}()
	res.Frob(input)
	return
}
