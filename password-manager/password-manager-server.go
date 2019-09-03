package main

import (
	context "context"

	"github.com/lpegoraro/password-manager/remote"
)

type DefaultPasswordManagerServer struct {
}

func (*DefaultPasswordManagerServer) Add(ctx context.Context, req *remote.AddPasswordReq) (*remote.PasswordValue, error) {
	var configuration = PasswordConfiguration{}
	if req.OverrideConfiguration != nil {
		configuration = PasswordConfiguration{
			Method:  req.OverrideConfiguration.Method,
			Seed:    req.OverrideConfiguration.Seed,
			Factor:  req.OverrideConfiguration.Factor,
			Storage: req.OverrideConfiguration.Storage,
			Output:  req.OverrideConfiguration.Output,
		}
	}
	password := AddPassword(req.Tag, req.Username, configuration)
	return &remote.PasswordValue{
		Password: password,
	}, nil

}

func (*DefaultPasswordManagerServer) Get(ctx context.Context, req *remote.GetPasswordReq) (*remote.PasswordValue, error) {
	password := GetPassword(req.Tag, req.Username)
	return &remote.PasswordValue{
		Password: password,
	}, nil
}
