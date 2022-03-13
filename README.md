# Go-Vote Description
Voting App backend built with Go.

## Dependencies
1. Go version 1.17.8
2. MySQL version 5.7.35-38

## The Business Process
1. Admin make a voting by their own preferences, the voting status is now "draft"
2. Admin adds voters and/or specified the number of ballots to be generated
3. Admin finalize the voting preferences and voters, the voting status is now "ready"
4. Once the voting status is "ready", the system generates the ballot tokens
5. Admin opens the voting, the voting status is now "opened"
6. Voters vote
7. Admin can cancel the voting. if this happened, the voting status is become "canceled"
8. System close the vote when it meets the deadline or when all voters already voted, the voting status is now "finished"
9. System shows the voting results in realtime

## The Business Rules
1. Admin can set the voting preferences like voting deadline, voters, voting visibility, voters visibility
2. Admin can add the voters and generate ballot tokens when the voting has not been finalized yet
3. Admin can remove voters when the voting has not been finalized yet
4. Admin can set the number of ballots to be generated when the voting has not been finalized yet
5. Admin can shorten the voting deadline when the voting has not been finalized yet
6. Admin can extend the voting deadline when the voting has not meet the 90% votes
7. Admin can change the voting visibility from private to public, but cannot otherwise
8. Admin can change the voters visibility from known to anonymous anytime
9. Admin can change the voters visibility from anonymous to known when the voting has no ballot tokens generated and when the voting has not been finalized yet
10. 

## How To
### How to make a voting
1. create an account at `/user/register`
2. make a voting using the account at `/voting/create`
3. you will get the voting id as the return

### How to add or remove voters
1. login with your account at `/auth/login`
2. to add voters, you can:
    - invite the other users by using their email at `/voters/invite`
    - or if you set the voting's voters to be anonymous, you can generate ballots tokens to be used by other users anonymously at `/voters/make_ballots`
3. to remove voters, you can go to `/voters/uninvite`. you can only remove the voters who hasn't voted yet.

### How to vote
you have 2 ways to vote:
1. login to your account at `/auth/login`
2. login and go to `/voting/vote` and pass the voting id

or you can:
1. go to `/voting/vote` and pass the voting id followed by the ballot token

### How to get voting details
1. to get the voting's summary, go to `/voting/summary` and pass the voting id
2. if you set the voting's settings for voters to be known, you can get the voters email at `/voting/voters` using the voting id

### How to update a voting
1. 
2. to adjust the voting deadline, you can go to `/voting/adjust_deadline`
3. to close the voting immediately, go to `/voting/close`
4. 
