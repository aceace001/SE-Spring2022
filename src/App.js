import React from 'react'
import './App.css'
import Taskbar from './Components/Taskbar'
import {BrowserRouter as Router, Switch, Route} from 'react-router-dom'
import Home from './pages/Home'
import Report from './pages/Report'
import Product from './pages/Product'


function App() {
  return (
    <>
    <Router>
      <Taskbar />
      <Switch>
        <Route path='/' exact component={Home}/> //We should have exact to make the page render
        <Route path='/reports' component={Report}/>
        <Route path='/products' component={Product}/>
      </Switch>
    </Router>
    
    </>
  );
}

export default App;
