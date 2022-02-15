import React from 'react'
import './App.css'
import Taskbar from './Components/Taskbar'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import Home from './pages/Home'
import Report from './pages/Report'
import Product from './pages/Product'


function App() {
  return (
    <>
    <Router>
      <Taskbar />
      <Routes>
        <Route path='/' element = {<Home />}/> 
        <Route path='/reports' element={<Report />}/>
        <Route path='/products' element={<Product />}/>
      </Routes>
    </Router>
    
    </>
  );
}

export default App;
