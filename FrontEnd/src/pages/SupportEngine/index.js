import React, { Component } from 'react';
import { useRef, useEffect, useState } from 'react'
import Avatar from './Avatar'
import SupportWindow from './SupportWindow';

const SupportEngine = () => {
    const [visible, setVisible] = useState(false)
    const wrapperRef = useRef(null);
    useOutsideAlerter(wrapperRef);

    function useOutsideAlerter(ref){
        useEffect(() =>{
            function handleClickOutside(event){
                if (ref.current && !ref.current.contains(event.target)){
                    setVisible(false)
                }
            }
            document.addEventListener("mousedown",handleClickOutside);
            return () => {
                //when you refresh or move the other pages, we just removeEventListener
                //In this case we can reduce the memroy leak
                document.removeEventListener("mousedown",handleClickOutside);
            }
        },[ref]);
    }

    return(
        <div ref={wrapperRef}>
        
            <SupportWindow
                visible={visible}
            />
            <Avatar 
                onClick={() => setVisible(true)}
                style={{position: 'fixed', bottom: '24px', right: '24px'}}
            />
        </div>
    )
}

export default SupportEngine;
