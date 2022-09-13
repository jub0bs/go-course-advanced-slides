package auth

import "context"

type tokenCtxKey struct{} // HL

func WithToken(ctx context.Context, token string) context.Context { // HL
	return context.WithValue(ctx, tokenCtxKey{}, token)
}

func Token(ctx context.Context) (string, bool) { // HL
	v := ctx.Value(tokenCtxKey{})
	s, ok := v.(string)
	return s, ok
}
