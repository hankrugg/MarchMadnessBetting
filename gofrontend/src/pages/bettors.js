import React, { useState } from 'react';
import axios from 'axios';
import {Link} from "react-router-dom";

const BettorsPage = () => {
    const [bettorData, setBettorData] = useState([]);
    const [bettorName, setBettorName] = useState('');

    const handleAddBettor = () => {
        setBettorData(prevData => [...prevData, { name: bettorName }]);
        setBettorName('');
    };

    const handleSubmit = async () => {
        try {
            // Ensure bettorData is an array of strings
            const bettorsArray = bettorData.map(bettor => bettor.name);

            // Send the array of bettors to the server
            await axios.post('http://localhost:8080/set-bettors', { bettors: bettorsArray });

            alert('Bettors data submitted successfully!');
        } catch (error) {
            console.error('Error submitting bettors data:', error);
            alert('Failed to submit bettors data.');
        }
    };


    return (
        <div>
            <h1>Bettors Page</h1>
            <div>
                <label htmlFor="bettorName">Name:</label>
                <input type="text" id="bettorName" value={bettorName} onChange={(e) => setBettorName(e.target.value)} />
            </div>
            <button onClick={handleAddBettor}>Add Bettor</button>
            <Link to="/" className="button">
            <button onClick={handleSubmit}>
                Submit
            </button>
            </Link>
            <div>
                <h2>Current Bettors</h2>
                <ul>
                    {bettorData.map((bettor, index) => (
                        <li key={index}>{bettor.name}</li>
                    ))}
                </ul>
            </div>
        </div>
    );
};

export default BettorsPage;
