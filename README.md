# GlobalWebIndex Engineering Challenge

## Introduction

This challenge is designed to give you the opportunity to demonstrate your abilities as a software engineer and specifically your knowledge of the Go language.

On the surface the challenge is trivial to solve, however you should choose to add features or capabilities which you feel demonstrate your skills and knowledge the best. For example, you could choose to optimise for performance and concurrency, you could choose to add a robust security layer or ensure your application is highly available. Or all of these.

Of course, usually we would choose to solve any given requirement with the simplest possible solution, however that is not the spirit of this challenge.

## Challenge

Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

A working server application with functional API is required, along with a clear readme.md. Useful and passing tests would be also be viewed favourably

It is appreciated, though not required, if a Dockerfile is included.

## Submission

Just create a fork from the current repo and send it to us!

Good luck, potential colleague!

## Docker Setup

### Build the Docker Image
To build the Docker image, navigate to the root directory of the project and run the following command:
```sh
docker build -t gwi-app .
```

### Run the Docker Container
```sh
docker run -p 8080:8080 gwi-app
```

This will start the web server and make it accessible at [http://localhost:8080](http://localhost:8080).

## Unit Tests

The following methods in the `InMemoryUserDataService` have unit tests implemented:

### GetUser
- `GetUser` retrieves a user by ID.
- Tests:
    - Successful retrieval of a user.
    - Retrieval of a non-existent user.

### AddUser
- `AddUser` adds a new user.
- Tests:
    - Successful addition of a user.

### AddAssetToFavourites
- `AddAssetToFavourites` adds an asset to the user's list of favourites.
- Tests:
    - Successful addition of an asset to the user's favourites.
    - Attempt to add a duplicate asset to the user's favourites.

### RemoveAssetFromFavourites
- `RemoveAssetFromFavourites` removes an asset from the user's list of favourites.
- Tests:
    - Successful removal of an asset from the user's favourites.
    - Attempt to remove a non-existent asset from the user's favourites.

### EditAssetDescription
- `EditAssetDescription` edits the description of an asset in the user's list of favourites.
- Tests:
    - Successful editing of an asset's description.
    - Attempt to edit the description of a non-existent asset.


## Unit Tests for Handlers

The following methods in the `handlers` package have unit tests implemented:

### MockUserDataService
- `MockUserDataService` is a mock implementation of `UserDataService` for testing purposes.
- Mocked Methods:
    - `GetUser`
    - `AddUser`
    - `AddAssetToFavourites`
    - `RemoveAssetFromFavourites`
    - `EditAssetDescription`

### GetUserFavourites
- `GetUserFavourites` handles the retrieval of a user's favourite assets.
- Tests:
    - Successful response with the user's favourite assets.
    - Response when the user is not found.
