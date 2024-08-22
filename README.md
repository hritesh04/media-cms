### Functional Requirements
- User should be able to create an account.
- User should be able to subscribe for premium or free services.
- User should be able to comment on articles.
- User should get notification for new articles or updates.
- An Author should be able to publish articles.
- Article should have statistics like views, engagement and subscriptions.

### Non-Functional Requirements
- The system should be scaleable and fault tolerant.
- The system must have minimum down-time.
- The system latency must be low.
- System should be highly consistent.

### Capacity Estimation

Assuming that the platform get around 1 million request per month and around 10k new users every month. Out of the 1 million requests 20% are of new articles and 80% are articles reads, that means approx 33k requests per day and approx 9 requests per second.

which means 200k articles are created every month which would be approx 6.6k new articles every day and approx 2 new articles every second and 800k article reads per month which would be approx 26.7k reads per day and approx 7 reads per second.

Therefore the number of request our platform handles per second is 9 (2 + 7).

#### Data Storage

###### Articles :
Assuming the average length of an article is around 1500 words, and each word averages around 6 character that would make it 9000 characters. so the size of each article will be 9000 x 1 byte/character = 9000 bytes or 8.79kbs.

Our platform crates 2 articles per second that means 2 x 8.79kb = 17.58 kbs of writes per second and 7 x 8.79 kbs = 61.5 kbs read per second.

###### Users :
we store users name,email,password,subscription,comments.
let us assume it to be 1 kb of storage space for each user.

our platform gets 5 user per minutes that means 5kb of writes per minute.

###### Total
The total Data write per second  : 17.7 kbs (approx)
The total Data read per second :  61.5kbs (approx)
