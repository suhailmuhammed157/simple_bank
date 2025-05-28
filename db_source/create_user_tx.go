package db_source

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreateUser func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *Store) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}
		err = arg.AfterCreateUser(result.User)
		return err
	})

	return result, err

}
