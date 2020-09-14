package usecase

import "context"

func (p *products) Check(ctx context.Context, prefix string) (string, error) {
	return prefix, nil
}
