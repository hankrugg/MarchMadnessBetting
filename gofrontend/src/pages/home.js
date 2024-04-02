import React, { useState, useEffect } from 'react';
import axios from 'axios';
import '../styles/styles.css';
import { Link } from "react-router-dom";

const Home = () => {
    const [games, setGames] = useState([]);
    const [selectedGameId, setSelectedGameId] = useState('');
    const [selectedGame, setSelectedGame] = useState(null);
    const [loading, setLoading] = useState(true);
    const [bettorData, setBettorData] = useState([]);
    const [winner, setWinner] = useState({});

    useEffect(() => {
        fetchData();
    }, []);

    useEffect(() => {
        if (selectedGameId) {
            const game = games.find(game => game.game.gameID === selectedGameId);
            setSelectedGame(game);
        }
    }, [selectedGameId, games]);

    const fetchData = async () => {
        try {
            const response = await axios.get('http://localhost:8080/');
            const gamesData = response.data.games;
            setGames(gamesData);
            setLoading(false);
            console.log('Game data fetched:', gamesData);
            console.log("GameID", selectedGameId)
            const winner = await axios.post('http://localhost:8080/refresh', {gameID : selectedGameId});
            const winningData = winner.data;
            setWinner(winningData);
            console.log('Winning data fetched:', winningData);
        } catch (error) {
            console.error('Error fetching game data:', error);
            setLoading(false);
        }
    };

    const fetchBettorData = async () => {
        try {
            const bettorResponse = await axios.get('http://localhost:8080/bettors');
            const bettors = bettorResponse.data.bettors;
            setBettorData(bettors);

        } catch (error) {
            console.error('Error fetching bettor data:', error);
        }
    };

    const handleGameSelect = (event) => {
        const gameId = event.target.value;
        setSelectedGameId(gameId);
        handleRefresh();
    };

    const handleRefresh = async () => {
        setLoading(true); // Set loading state to true while fetching data
        fetchData(); // Fetch game data
        fetchBettorData();
    };

    console.log("bettor data", bettorData)

    return (
        <div className="container">
            <header className="header">
                <h1 className="title">Game Details</h1>
                <button className="button" onClick={handleRefresh}>Refresh</button>
                <Link to="/set-bettors" className="button">Set Custom Bettor Data</Link>
            </header>
            {loading ? (
                <p>Loading...</p>
            ) : (
                <div>
                    <select value={selectedGameId} onChange={handleGameSelect}>
                        <option value="">Select a Game</option>
                        {games.map(game => (
                            <option key={game.game.gameID} value={game.game.gameID}>
                                {game.game.title}
                            </option>
                        ))}
                    </select>
                    <br />
                    {selectedGame ? (
                        <div>
                            <div>
                                <h2>Selected Game Details</h2>
                                <div className="game-details">
                                    <div>
                                        <p><strong>Title:</strong> {selectedGame.game.title}</p>
                                        <p><strong>Start Time:</strong> {selectedGame.game.startTime}</p>
                                    </div>
                                    <div>
                                        <p><strong>Time:</strong> {selectedGame.game.contestClock}</p>
                                        <p><strong>Period:</strong> {selectedGame.game.currentPeriod}</p>
                                    </div>
                                    <div>
                                        <p><strong>Home:</strong> {selectedGame.game.home.names.short}</p>
                                        <p><strong>Score:</strong> {selectedGame.game.home.score}</p>
                                    </div>
                                    <div>
                                        <p><strong>Away:</strong> {selectedGame.game.away.names.short}</p>
                                        <p><strong>Score:</strong> {selectedGame.game.away.score}</p>
                                    </div>
                                </div>
                            </div>

                            <h2>Bettors</h2>
                            {bettorData && (
                                <table className="table">
                                    <thead>
                                    <tr>
                                    <th>Spot</th>
                                        {bettorData.map((_, index) => (
                                            <th key={index}>{index} {selectedGame.game.away.names.char6}</th>
                                        ))}
                                    </tr>
                                    </thead>
                                    <tbody>
                                    {bettorData.map((row, rowIndex) => (
                                        <tr key={rowIndex}>
                                            <td>{rowIndex} {selectedGame.game.home.names.char6}</td>
                                            {row.map((player, colIndex) => (
                                                <td
                                                    key={colIndex}
                                                    className={winner.home_score_digit === rowIndex && winner.away_score_digit === colIndex ? 'winner' : ''}
                                                    style={{backgroundColor: winner.home_score_digit === rowIndex && winner.away_score_digit === colIndex ? 'green' : 'inherit'}}
                                                >
                                                    {player}
                                                </td>
                                            ))}
                                        </tr>
                                    ))}
                                    </tbody>
                                </table>
                            )}
                        </div>
                    ) : (
                        <p>Please select a game.</p>
                    )}
                </div>
            )}
        </div>
    );
};

export default Home;
