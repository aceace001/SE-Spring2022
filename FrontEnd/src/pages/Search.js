import React, {useState} from 'react'
import JSONDATA from '../MOCK_DATA.json';
import "../App.css";

//map is like the for loop, use filter can traverse faster then the map

function Search() {
  const [searchTerm, setsearchTerm] = useState('')
  return (
    <div className='search'>
        <h1>Search for users</h1>
        <input className='searchterm' type='text' placeholder='Search'onChange={event => {setsearchTerm(event.target.value)}}/>
        {JSONDATA.filter((val) =>{
          if (searchTerm === ''){
            return val
          } else if (val.first_name.toLowerCase().includes(searchTerm.toLowerCase())){
            return val
          }
        }).map((val,key) => {
          return(
            <div className='user' key={key}>
              <p>{val.first_name}</p>
            </div>
          );
        })}
      
    </div>
  )
}

export default Search
