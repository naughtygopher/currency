package currency

// UpdateWithFractional will update all the relevant values of currency based on the
// fractional unit provided.
func (c *Currency) UpdateWithFractional(frac int) {
	fus := int(c.FUShare)

	c.Main = (frac / fus)
	c.Fractional = (frac % fus)

	if c.Main < 0 {
		c.Fractional = -c.Fractional
	}
}

// Add adds the given currency with the base currency.
func (c *Currency) Add(acur Currency) error {
	if c.Code != acur.Code {
		return ErrMismatchCurrency
	}

	c.UpdateWithFractional(c.FractionalTotal() + acur.FractionalTotal())
	return nil
}

// AddInt adds main & fractional value provided to the currency
func (c *Currency) AddInt(main int, frac int) {
	if main < 0 && frac > 0 {
		frac = -frac
	}

	c.UpdateWithFractional(c.FractionalTotal() + main*int(c.FUShare) + frac)
}

// SubtractInt subtracts main & fractional value provided from the currency
func (c *Currency) SubtractInt(main int, frac int) {
	if main < 0 && frac > 0 {
		frac = -frac
	}

	c.UpdateWithFractional(c.FractionalTotal() - (main*int(c.FUShare) + frac))
}

// Subtract subtracts the given currency from the base currency.
func (c *Currency) Subtract(scur Currency) error {
	if c.Code != scur.Code {
		return ErrMismatchCurrency
	}

	c.UpdateWithFractional(c.FractionalTotal() - scur.FractionalTotal())
	return nil
}

// Percent returns a new instance of currency which is n percent of c.
func (c *Currency) Percent(n float64) *Currency {
	totalFrac := round(float64(c.FractionalTotal())*(n/100.00), c.magnitude)
	c1 := *c
	c1.UpdateWithFractional(totalFrac)
	return &c1
}

// Multiply multiplies the currency by an integer.
func (c *Currency) Multiply(by int) {
	c.UpdateWithFractional(c.FractionalTotal() * by)
}

// MultiplyFloat64 multiplies the currency by a float value.
func (c *Currency) MultiplyFloat64(by float64) {
	t := float64(c.FractionalTotal()) * by
	c.UpdateWithFractional(round(t, c.magnitude))
}

// Divide is a deprecated method which does allocations
// Deprecated: Divide is not the technical term when dealing with currency.
func (c *Currency) Divide(by int, retain bool) ([]Currency, bool) {
	return c.Allocate(by, retain)
}

// Allocate does fair allocation of the currency by the given integer and returns a list of currencies and bool.
/*
   If `retain` is set as true, the balance will not be distributed among the splits,
   instead retained inside c. It returns a list because, when the currency cannot
   be split/divided equally, then the remainder has to be distributed.
   The bool value if `true`, means the currency was split equally.
*/
func (c *Currency) Allocate(by int, retain bool) ([]Currency, bool) {
	sE := false

	ft := c.FractionalTotal()

	d := make([]Currency, by)

	c1 := *c
	c1.UpdateWithFractional(ft / by)

	balance := ft % by

	if balance == 0 {
		sE = true
	}

	for i := 0; i < by; i++ {
		d[i] = c1
		if !retain && balance > 0 {
			d[i].AddInt(0, 1)
			balance--
		}
	}

	if retain {
		c.UpdateWithFractional(balance)
	}

	return d, sE
}
