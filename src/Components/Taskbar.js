import React,{useState}from 'react'
import * as FaIcons from "react-icons/fa"
import * as AiIcons from "react-icons/ai"
import {Link} from 'react-router-dom'
import {SidebarData} from './SidebarData'
import './Taskbar.css'
import {IconContext} from 'react-icons'
// IconContext allow you to modify the top layer of the icons


function Taskbar() {
  const [sidebar, setSidebar] = useState(false)
  //Here is like toggling: switching between true and false
  const showSidebar = () => setSidebar(!sidebar)
  return (
    <>
    <IconContext.Provider value={{color: '#fff' }}>
        <div className="taskbar">
            <Link to='#' className="menu-bars">
                <FaIcons.FaBars onClick={showSidebar}/>
            </Link>

        </div>
        <nav className={sidebar ? 'nav-menu active' : 'nav-menu'}>
            <ul className='nav-menu-items' onClick={showSidebar}>
                <li className='navbar-toggle'>
                    <Link to="#" className='menu-bars'>
                        <AiIcons.AiOutlineClose />
                    </Link>
                </li>
                {SidebarData.map((item,index) =>{
                    return(
                        <li key={index} className={item.Cname}>
                            <Link to={item.path}>
                                {item.icon}
                                <span>{item.title}</span>
                            </Link>
                        </li>
                    )
                })}
            </ul>
        </nav>
        </IconContext.Provider>

      
    </>
  )
}

export default Taskbar
