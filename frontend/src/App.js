import { BrowserRouter as Router, Route,Routes} from 'react-router-dom';
import TakeInfo from '../src/components/TakeInfo';
import LobbyPage from '../src/components/CreateLobby';
import WaitingPage from '../src/components/WaitingPage';
import GamePage from '../src/components/GamePage';
import React, { useState, useEffect } from 'react';

function App() {
  const [matrix, setMatrix] = useState({ Board: [], Color: [] });
  useEffect(() => {
    setMatrix({
      Board: [
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0]
      ],
      Color: [
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
        ['', '', '', '', '', ''],
      ]
    });
  }, []);
  return (
    <Router>
      <div className="App">
      <Routes>
          <Route path="/" exact element={<TakeInfo/>} />
          <Route path="/lobby" element={<LobbyPage/>} />
          <Route path="/waiting" element={<WaitingPage/>} />
          <Route path="/game" element={<GamePage matrix={matrix}/>} />
          </Routes>
      </div>
    </Router>
  );
}

export default App;
