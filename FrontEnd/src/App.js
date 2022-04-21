import React, {useEffect, useState} from 'react';

import './App.css'
import Taskbar from './Components/Taskbar'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Users from './pages/Users'
import Submit from './pages/Submit'
import Settings from './pages/Settings'

import Addfriends from './pages/Addfriends'
import Posts from './pages/Posts'
import Support from './pages/SupportAdmin'
import Messages from './pages/Messages'
import Login from './pages/Login'
import Register from './pages/Register'
import HomePost from './pages/HomePost'
import HomePost1 from './pages/HomePost1'




function App() {
  return (
    <>
    <div className="c" id="container">
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
        
        <Route path='/addfriends' element={<Addfriends />}/>
        <Route path='/HomePost1' element={<HomePost1 />} />
        <Route path='/support' element={<Support />}/>
        <Route path='/posts' element={<Submit />}/>
      </Routes>
    </Router>
    </div>
    </>
  );
}

export default App;
