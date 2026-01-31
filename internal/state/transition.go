package state

import "errors"

// TransitionContext membawa dependency runtime
type TransitionContext struct {
	Gas GasMeter
}

// ApplyTransaction menjalankan satu transisi state
func ApplyTransaction(ctx *TransitionContext, fn func() error, gasCost uint64) error {
	if ctx == nil || ctx.Gas == nil {
		return errors.New("missing gas context")
	}

	// consume gas dulu (anti DoS)
	if err := ctx.Gas.Consume(gasCost); err != nil {
		return err
	}

	// jalankan logika state
	return fn()
}
