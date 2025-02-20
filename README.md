# March Madness Betting  

## Overview  
This is a Go web application that provides an interface for managing and tracking bets on NCAA Women's Basketball games. Users can input a betting grid, load daily game data, and see real-time updates on the current winners.  

## Features  
- Automatically loads daily NCAA Women's Basketball games  
- Allows users to add and manage bettors  
- Provides a bettor grid for tracking bets  
- Displays real-time game updates and highlights current winners  

## Prerequisites  
Before running the project, ensure you have the following installed:  
- [Go](https://golang.org/dl/)  
- [Node.js and npm](https://nodejs.org/)  

## Setup Instructions  

### 1. Install Frontend Dependencies  
Navigate to the frontend directory and install the required packages:  
```sh
cd gofrontend  
npm install  
```

### 2. Start the Frontend  
Run the following command to launch the frontend:  
```sh
npm start  
```
This will start the React app and serve it at `http://localhost:3000/`.  

### 3. Start the Backend  
In a separate terminal, navigate to the root directory where the Go files are located and start the backend:  
```sh
go run .  
```

### 4. Access the Web App  
Once both the frontend and backend are running, open your browser and go to:  
```
http://localhost:3000
```

## How to Use  
1. **Load Daily Games:** The app will automatically retrieve the latest game data.  
2. **Add Bettors:** Click "Set Custom Bettor Data" to enter bettor names individually. Once all bettors are added, click "Submit."  
3. **Select a Game:** Choose a game from the dropdown menu. The app will populate the bettor grid with relevant data.  
4. **Track Results:** Click "Refresh" to update game data and see the current winner highlighted. You can refresh anytime to get live updates.  

## Future Improvements  
- Add user authentication for personalized betting history  
- Implement an AI-powered betting suggestion system  
- Enhance UI for a better user experience  
  

---
