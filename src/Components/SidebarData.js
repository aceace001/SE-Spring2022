import React from 'react';
import * as BsIcons from 'react-icons/bs'
import * as AiIcons from "react-icons/ai"
import * as IoIcons from 'react-icons/io'
import * as FiIcons from 'react-icons/fi'

export const SidebarData = [
    {
        title: "Home",
        path: '/',
        icon: <AiIcons.AiFillHome />,
        cName: 'nav-text'
    },
    {
        title: "Settings",
        path: '/Settings',
        icon: <IoIcons.IoMdSettings />,
        cName: 'nav-text'
    },
    {
        title: "Users",
        path: '/users',
        icon: <FiIcons.FiUsers />,
        cName: 'nav-text'
    },
    {
        title: "Search",
        path: '/search',
        icon: <FiIcons.FiSearch />,
        cName: 'nav-text'
    },
    {
        title: "Messages",
        path: '/messages',
        icon: <FiIcons.FiMessageSquare />,
        cName: 'nav-text'
    },
    {
        title: "Addfriends",
        path: '/addfriends',
        icon: <FiIcons.FiUserPlus />,
        cName: 'nav-text'
    },
    {
        title: "Posts",
        path: '/posts',
        icon: <BsIcons.BsFillFilePostFill />,
        cName: 'nav-text'
    },
    {
        title: "Support",
        path: '/support',
        icon: <IoIcons.IoMdHelpCircle />,
        cName: 'nav-text'
    },

]