import React, {useState} from 'react'
import JSONDATA from '../MOCK_DATA.json';
import ShowFriends from './showFriends'

import '../App.css'
import 'bootstrap/dist/css/bootstrap.min.css';

const Users = () => {
  const [showFriends, setshowFriends] = useState(JSONDATA)
  return (
    <div className='users'>
        <h2 className="display-1">Number of friends: {showFriends.length} </h2>
        <ShowFriends showfriends={showFriends}/>
        <button type='button' className="btn btn-success" onClick={()=>setshowFriends(JSONDATA)}>Get all</button>
        <button type='button' className="btn btn-success" onClick={()=>setshowFriends([])}>Clear all</button>
      
    </div>
  )
}

export default Users
