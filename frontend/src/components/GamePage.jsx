import React, { useEffect } from 'react';
import corner from '../images/classic.jpg';
function GamePage({ matrix }) {
  useEffect(() => {
   
  }, []);

  return (
    <div className='flex-center-outer'>
      <img className='corner-img left-top' src={corner} alt="" />
        <img className='corner-img right-btm' src={corner} alt="" />
      <div>
      <h2>Game Board</h2>
      <table>
        <tbody>
          {matrix?.Board.map((row, i) => (
            <tr key={i}>
              {row.map((cell, j) => (
                <td className='cell' key={j} style={{ backgroundColor: matrix.Color[i][j] }}>{cell}</td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
    </div>
  );
}

export default GamePage;
