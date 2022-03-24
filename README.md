# Go-Vote Description
Voting App backend built with Go. (still on progress)

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
8. Admin can change the voters visibility from known to unknown anytime
9. Admin can change the voters visibility from unknown to known when the voting has no ballot tokens and when the voting has not been finalized yet
10. Voters can vote the voting only when the voting status is "opened"
11. Voters can only vote when they're invited or has ballots for the voting
12. Users can only see the votings only when they're the one that created it or when they're invited to vote for it, based on the user id and/or ip address
13. The anonymous users can only see the votings when they have the ballot tokens of the voting

## Scope and Limitation
1. This app cannot make a vote at global visibility (global visibility: all users can vote whether they're invited or not)

## How To
### How to manage an account
1. Create an account at `POST /user/register`
2. Get an account details at `GET /user/profile/:user_id`
3. Update your account at `PUT /user/update/:user_id` (N/A)
4. Login at `POST /auth/login`
5. Refresh session at `PUT /auth/refresh`
6. Logout at `PUT /auth/logout` (N/A)

### How to make a voting
1. login to your account
2. make a voting using the account at `POST /voting/create` (N/A)
3. you will get the voting id as the return

### How to specify voters
1. login with your account
2. to add voters by email, go to `POST /voting/voters/invite` (N/A)
3. to remove voters by email, go to `POST /voting/voters/uninvite` (N/A)
4. to allow organization to vote, you can use `POST /voting/voters/allow_organization` (N/A)
5. to disallow organization, use `POST /voting/voters/disallow_organization` (N/A)
6. to set the ballots to be generated, use `POST /voting/voters/set_ballots` (N/A)
7. to get the ballot tokens, finalize the voting and then use `GET /voting/voters/get_ballots/:voting_id` (N/A)

### How to update voting status
1. login with your account
2. to finalize voting preferences, use `PUT /voting/finalize/:voting_id` (N/A)
3. to open the voting, use `PUT /voting/open/:voting_id` (N/A)
4. to cancel voting, use `PUT /voting/cancel/:voting_id` (N/A)

### How to vote
you have 2 ways to vote:
1. login to your account
2. go to `POST /voting/vote` and pass the voting id on its body (N/A)

or you can:
1. go to `POST /voting/vote` anonymously and pass the voting id and ballot token on its body (N/A)

### How to get voting details
1. to get the voting summary, go to
   - `GET /voting/summary/:voting_id` or (N/A)
   - `GET /voting/summary/by_ballot/:ballot_id` (N/A)
2. if you set the voter visibility to be known, you can get the voters email at `GET /voting/voters/:voting_id` (N/A)
3. to get the list of voting and filter it, go to `POST /voting/filter` (N/A)

### How to update a voting
1. to update the voting preferences, you can go to `PUT /voting/update/:voting_id` (N/A)
