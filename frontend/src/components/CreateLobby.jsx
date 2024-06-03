import React, { useState } from 'react';
import '../../src/App.css';
import corner from '../images/classic.jpg'
import { useNavigate } from 'react-router-dom';

function LobbyPage() {
  const [lobbyName, setLobbyName] = useState('');
  const[displayJoinLobbyTag, setDisplayJoinLobbyTag] = useState(false);
  const navigate = useNavigate();
  

  const handleCreateLobby = () => {
   navigate('/waiting')
    
  };

  const handleJoinLobby = () => {
   setDisplayJoinLobbyTag(!displayJoinLobbyTag);
    
  };
  const handleJoinlobbyWithCode = () =>{
    navigate('/waiting')
  }

  return (
    <>
     <div className='flex-center-outer col'>
        <img className='corner-img left-top' src={corner} alt="" />
        <img className='corner-img right-btm' src={corner} alt="" />
      <h2 className='title'>Lobby</h2>
      
     <div className='flex-center col sml-cont-box fix-width-n-height align-justify-center'>
      <button className='btn-lobby' onClick={handleCreateLobby}>Create Lobby</button>
     </div>
      <div className='flex-center col sml-cont-box fix-width-n-height align-justify-center'>
      
      {displayJoinLobbyTag ? 
        <div className='flex-center'> <input className='join-lobby-tag' type="text" placeholder="Lobby Name" value={lobbyName} onChange={(e) => setLobbyName(e.target.value)} /> <div onClick={handleJoinlobbyWithCode} className="btn-go">Go</div></div>
        : <button className='btn-lobby' onClick={handleJoinLobby}>Join Lobby</button>
      }
      </div>
      </div>
    </>
  );
}

export default LobbyPage;
