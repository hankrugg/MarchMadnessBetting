# Go Web Project

This is a Go web project that provides an interface that allows the user to provide a grid for NCAA Womens Basketball games. 

## Using the Program
The app will automatically load the daily games, but before you can bet, you must load the bettors. Click on Set Custom Bettor Data to add bettors.
This opens the bettors page where you will add name individually and click add bettor. When all you bettors have been added, click submit.

Next, select a game from the drop down menu. This will load game data and fill in the bettor grid. Finally, click refresh and you will get up to date game data as
well as a highlighted grid of the current winner. Click refresh any time you want to see who the current winner is!


## Prerequisites

Before running the project, ensure you have the following dependencies installed:

- Go programming language
- Node.js and npm

## Setup Instructions

Navigate to the gofrontend directory and run

```
npm install
```

This will install the needed packages for the frontend to run. Once this has completed, run

```
npm start
```
to start the frontend.

In a different command line window, navigate to the root directory where the .go files are located. Run 
```
go run .
```
to start the backend.

Find the web app at 

```
localhost:3000
```

and have fun betting!
