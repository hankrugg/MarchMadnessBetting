import './App.css';
import {BrowserRouter, Routes, Route, Navigate} from 'react-router-dom';
import Home from "./pages/home"
import BettorsPage from "./pages/bettors";


function App() {
  return (
      <BrowserRouter>
        <Routes>
            <Route path="*"  element={<Home/>}/>
            <Route path="/set-bettors"  element={<BettorsPage/>}/>
        </Routes>
      </BrowserRouter>
  );
}

export default App;
