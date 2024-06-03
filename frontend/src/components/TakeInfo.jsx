import React, { useState } from 'react';
import corner from '../images/classic.jpg'
import { useNavigate } from 'react-router-dom';

const TakeInfo = () => {
  const [playerName, setPlayerName] = useState('');
  const [playerTag, setPlayerTag] = useState('');
  const navigate = useNavigate();


  const handleStartGame = () => {
 navigate('/lobby')

  };

  return (
    <>
      <div className='flex-center-outer col'>
        <img className='corner-img left-top' src={corner} alt="" />
        <img className='corner-img right-btm' src={corner} alt="" />
        
      <h2 className='title'>Chain Reaction</h2>
        <div className='flex-center col sml-cont-box'>
          
          <input className='input-tags' type="text" placeholder="Name" value={playerName} onChange={(e) => setPlayerName(e.target.value)} />
          <input className='input-tags' type="text" placeholder="Player Tag" value={playerTag} onChange={(e) => setPlayerTag(e.target.value)} />
          <button className='btn' onClick={handleStartGame}>Start Game</button>
          
        </div>
      </div>
    </>
  );
}
export default TakeInfo;


