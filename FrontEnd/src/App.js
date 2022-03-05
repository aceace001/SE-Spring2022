import React, {useEffect, useState} from 'react';
import './App.css';
import Home from "./page/Home";
import Login from "./page/Login";
import Register from "./page/Register";
import TaskBar from "./TaskBar/TaskBar";
import TaskBar1 from "./TaskBar/TaskBar1";
import HomePost from "./page/HomePost";
import {BrowserRouter, Route, Routes} from "react-router-dom";

function App() {
  const [name, setName] = useState('');
  return (
      <div className="App">
        <BrowserRouter>
          <TaskBar name={name} setName={setName}/>
          <div className={"home"}>
              <Routes>
                  <Route path="/" element={<Home/>}/>
              </Routes>

          </div>
          <main className="form-signin">
              <Routes>
                  <Route path="/Login" element={<Login/>}/>
                  <Route path="/Register" element={<Register/>}/>
              </Routes>
          </main>
          <div className={"homepost"}>
              <Routes>
                  <Route path="/HomePost" element={<HomePost/>}/>
              </Routes>
          </div>
        </BrowserRouter>
      </div>
  );
}

export default App;

