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
        title: "Messages",
        path: '/messages',
        icon: <FiIcons.FiMessageSquare />,
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
