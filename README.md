# go-vote
Voting App backend built with Go.

## dependencies
1. Go version 1.17.8
2. MySQL version 5.7.35-38

## how to
### how to make a vote
1. create an account at `/user/register`
2. make a vote using the account at `/vote/create`
3. you will get the vote id as the return

### how to add voters
1. login with your account at `/auth/login`
2. to add voters, you can:
    - invite the other users by using their email at `/voters/invite`
    - or if you set the vote's voters to be anonymous, you can generate ballots tokens to be used by other users anonymously at `/voters/make_ballots`

### how to vote
you have 2 ways to vote:
1. login to your account at `/auth/login`
2. login and go to `/vote/commit` and pass the vote id
or you can:
1. go to `/vote/commit` and pass the vote id followed by the ballot token

