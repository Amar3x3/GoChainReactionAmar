import React, { useState, useEffect } from 'react';
import corner from '../images/classic.jpg'
import { useNavigate } from 'react-router-dom';

function WaitingPage() {
  const [players, setPlayers] = useState([
    {name: 'dummy1'},
    {name: 'dummy2'},
    {name: 'dummy3'},
    {name: 'dummy4'}
  ]);
  const navigate = useNavigate();
  

  useEffect(() => {
   
  }, []);

  const handleStartGame = () => {
    navigate('/game')
    
  };

  return (
    <div className='flex-center-outer col'>
        <img className='corner-img left-top' src={corner} alt="" />
        <img className='corner-img right-btm' src={corner} alt="" />
        <h2 className='title'>Waiting for Players</h2>
      
        {players.map((player, index) => (
          <div className='player-wait-card' key={index}>{player.name}</div>
        ))}
      
      <button className='btn-lobby' onClick={handleStartGame}>Start Game</button>
      </div>
  );
}

export default WaitingPage;
