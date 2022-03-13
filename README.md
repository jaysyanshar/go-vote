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
2. Admin can add the voters and generate ballot tokens when the voting has not been finalized (status "ready") yet
3. Admin can remove voters when the voting has not been finalized yet
4. Admin can set the number of ballots to be generated when the voting has not been finalized yet
5. Admin can shorten the voting deadline when the voting has not been finalized yet
6. Admin can extend the voting deadline when the voting has not meet the 90% votes
7. Admin can change the voting visibility from private to public, but cannot otherwise
8. Admin can change the voters visibility from known to anonymous anytime
9. Admin can change the voters visibility from anonymous to known when the voting has no ballot tokens and when the voting has not been finalized yet
10. Voters can vote the voting only when the voting status is "opened"
11. Voters can only vote when they're invited or has ballots for the voting
12. Users can only see the votings only when they're the one that created it or when they're invited to vote for it
13. The anonymous users can only see the votings when they have the ballot tokens of the voting

## How To
### How to make a voting
1. create an account at `/user/register` and/or login at `/auth/login`
2. make a voting using the account at `/voting/create`
3. you will get the voting id as the return

### How to add or remove voters
1. login with your account at `/auth/login`
2. to add/remove voters, you can go to `/voters/specify` and specify the user email, allow email domain, and/or set the number of ballots to be generated

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
1. to adjust the voting deadline, you can go to `/voting/adjust_deadline`
2. to close the voting immediately, go to `/voting/close`
3. 
