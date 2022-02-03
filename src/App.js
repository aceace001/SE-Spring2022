import {useState, useEffect} from 'react'
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'  //To use router we need wrap everything in our output
import Header from "./components/Header"
import Footer from "./components/Footer"
import Tasks from "./components/Tasks"
import AddTask from './components/AddTask'
import About from './components/About'

const App = () => {
  const [showAddTask, setShowAddTask] = useState(false)
  const [tasks, setTasks] = useState([
    {
      id: 1,
      text: 'Doctors Appointment',
      day: 'Feb 5th at 2: 30pm',
      reminder: true, 
    },
    {
      id: 2,
      text: 'Meeting at School',
      day: 'Fed 6th at 1:30pm',
      reminder: true,
    },
    {
      id: 3,
      text: 'Food Shopping',
      day: 'Feb 5th at 2:30pm',
      reminder: false,
    },
  ])

  //useEffect to fetch tasks
  useEffect(() => {
    const getTasks = async () => {
      const tasksFromServer = await fetchTasks()
      setTasks(tasksFromServer)
    }
    getTasks()
  },[])
  //[] shows the empty dependency array

  //Fetch Tasks
  const fetchTasks = async () => {
    const res = await fetch('http://localhost:5000/tasks')
    const data = await res.json()

    return data
  }

    //Fetch Task
    const fetchTask = async (id) => {
      const res = await fetch(`http://localhost:5000/tasks/${id}`)
      const data = await res.json()
  
      return data
    }

  //Add Task
  const addTask = async (task) => {                                                                                       //stringfy: change JSON to a string
    const res = await fetch('http://localhost:5000/tasks',{method: 'POST', headers: {'Content-type': 'application/json'}, body: JSON.stringify(task)})
    //data is the returned result
    //it is a promise make sure to do await
    const data = await res.json()

    //Take the exisitng task then add on to it
    setTasks([...tasks, data])

    // const id = Math.floor(Math.random()*10000) + 1
    // const newTask = {id, ...task}
    // setTasks([...tasks, newTask])
  }

  //Delete Task
  const deleteTask = async (id) =>{
    const res = await fetch(`http://localhost:5000/tasks/${id}`,{method: 'DELETE'})
    setTasks(tasks.filter((task)=>task.id !== id))
  }

  //Toggle Reminder
  const toggleReminder = async (id) => {
    const taskToToggle = await fetchTask(id)
    const updTask = {...taskToToggle, reminder: !taskToToggle.reminder}
    //Add header since we are sending data, set method
    const res = await fetch(`http://localhost:5000/tasks/${id}`,{method:'PUT',headers: {'Content-type': 'application/json'}, body: JSON.stringify(updTask)})
    const data = await res.json()
    setTasks(
      tasks.map((task) =>
        task.id === id ? {...task, reminder:data.reminder} : task
      )
    )
  }

  return (
    <Router>
    <div className="container">
      <Header onAdd={()=>setShowAddTask(!showAddTask)} showAdd={showAddTask} />
      <Routes>
          <Route
            path='/'
            element={
              <>
                {showAddTask && <AddTask onAdd={addTask} />}
                {tasks.length > 0 ? (
                  <Tasks
                    tasks={tasks}
                    onDelete={deleteTask}
                    onToggle={toggleReminder}
                  />
                ) : (
                  'No Tasks To Show'
                )}
              </>
            }
          />
          <Route path='/about' element={<About />} />
        </Routes>
      <Footer />
    </div>
    </Router>
  )
}




export default App;
