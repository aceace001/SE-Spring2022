import React, {useEffect, useState} from 'react';

import './App.css'
import Taskbar from './Components/Taskbar'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Users from './pages/Users'
import Settings from './pages/Settings'
import Search from './pages/Search'
import Addfriends from './pages/Addfriends'
import Posts from './pages/Posts'
import Support from './pages/Support'
import Messages from './pages/Messages'
import Login from './pages/Login'
import Register from './pages/Register'
import HomePost from './pages/HomePost'



function App() {
  return (
    <>
    
    <Router>
      <Taskbar />
      <Routes>
        <Route path='/' element={<Home />} />
        <Route path='/login' element={<Login />} />
        <Route path='/HomePost' element={<HomePost />} />
        <Route path='/register' element={<Register />} />
        <Route path='/settings' element={<Settings />}/>
        <Route path='/users' element={<Users />}/>
        <Route path='/messages' element={<Messages />}/>
        <Route path='/search' element={<Search />}/>
        <Route path='/addfriends' element={<Addfriends />}/>
        <Route path='/posts' element={<Posts />}/>
        <Route path='/support' element={<Support />}/>
      </Routes>
    </Router>
    
    </>
  );
}

export default App;
